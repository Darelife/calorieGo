package cache

const (
	foodBarcodePrefix = "food:barcode:"
)

func FoodBarcodeKey(barcode string) string {
	return foodBarcodePrefix + barcode
}

