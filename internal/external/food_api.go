package external

import "context"

type FoodData struct {
	Barcode  string
	Calories int
}

type FoodAPI interface {
	GetByBarcode(ctx context.Context, barcode string) (*FoodData, error)
}

