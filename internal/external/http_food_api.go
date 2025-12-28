package external

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"time"
)

var ErrNotFound = errors.New("food not found")

type HTTPFoodAPI struct {
	client  *http.Client
	baseURL string
}

func NewHTTPFoodAPI(baseURL string) *HTTPFoodAPI {
	return &HTTPFoodAPI{
		client: &http.Client{
			Timeout: 3 * time.Second,
		},
		baseURL: baseURL,
	}
}

func (api *HTTPFoodAPI) GetByBarcode(
	ctx context.Context,
	barcode string,
) (*FoodData, error) {

	var lastErr error

	for attempt := 1; attempt <= 3; attempt++ {
		data, err := api.fetch(ctx, barcode)
		if err == nil {
			return data, nil
		}

		if err == ErrNotFound {
			return nil, err
		}

		lastErr = err
		time.Sleep(time.Duration(attempt) * 200 * time.Millisecond)
	}

	return nil, lastErr
}

func (api *HTTPFoodAPI) fetch(
	ctx context.Context,
	barcode string,
) (*FoodData, error) {

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		api.baseURL+"/product/"+barcode+".json",
		nil,
	)
	if err != nil {
		return nil, err
	}

	resp, err := api.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("external api error")
	}

	var parsed struct {
		Status  int `json:"status"`
		Product struct {
			Nutriments struct {
				Calories float64 `json:"energy-kcal_100g"`
			} `json:"nutriments"`
		} `json:"product"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&parsed); err != nil {
		return nil, err
	}

	if parsed.Status == 0 {
		return nil, ErrNotFound
	}

	return &FoodData{
		Barcode:  barcode,
		Calories: int(parsed.Product.Nutriments.Calories),
	}, nil
}

