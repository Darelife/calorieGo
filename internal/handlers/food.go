package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/darelife/calorieGo/internal/cache"
	"github.com/darelife/calorieGo/internal/external"
	"go.uber.org/zap"
)

type FoodHandler struct {
	logger *zap.Logger
	cache  *cache.RedisCache
	api    external.FoodAPI
}

func NewFoodHandler(
	logger *zap.Logger,
	cache *cache.RedisCache,
	api external.FoodAPI,
) *FoodHandler {
	return &FoodHandler{
		logger: logger,
		cache:  cache,
		api:    api,
	}
}

type barcodeRequest struct {
	Barcode string `json:"barcode"`
}

type barcodeResponse struct {
	Barcode  string `json:"barcode"`
	Calories int    `json:"calories"`
	Source   string `json:"source"`
}

func (h *FoodHandler) GetByBarcode(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req barcodeRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	cacheKey := cache.FoodBarcodeKey(req.Barcode)

	// 1️⃣ Try Redis cache
	if val, err := h.cache.Get(ctx, cacheKey); err == nil {
		h.logger.Info("cache hit", zap.String("barcode", req.Barcode))
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(val))
		return
	}

	h.logger.Info("cache miss", zap.String("barcode", req.Barcode))

	// 2️⃣ Call external API
	data, err := h.api.GetByBarcode(ctx, req.Barcode)
	if err != nil {
		if err == external.ErrNotFound {
			http.Error(w, "food not found", http.StatusNotFound)
			return
		}

		h.logger.Error("external api failed",
			zap.String("barcode", req.Barcode),
			zap.Error(err),
		)
		http.Error(w, "external api error", http.StatusBadGateway)
		return
	}

	resp := barcodeResponse{
		Barcode:  data.Barcode,
		Calories: data.Calories,
		Source:   "external_api",
	}

	payload, _ := json.Marshal(resp)

	// 3️⃣ Store in Redis (24h TTL)
	_ = h.cache.Set(ctx, cacheKey, string(payload), 24*time.Hour)

	w.Header().Set("Content-Type", "application/json")
	w.Write(payload)
}

