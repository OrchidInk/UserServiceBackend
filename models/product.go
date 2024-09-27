package models

type CreateProductEnRequest struct {
	CategoryEnID  int32  `json:"category_en_id" validate:"required"`
	PriceEn       string `json:"price_en" validate:"required"`
	StockQuantity int32  `json:"stock_quantity" validate:"required"`
	ImagesPathEn  string `json:"images_path_en" validate:"required"`
}

type UpdateProductEnPriceRequest struct {
	ProductEnID int32  `json:"product_en_id" validate:"required"`
	PriceEn     string `json:"price_en" validate:"required"`
}

type UpdateProductEnStockRequest struct {
	ProductEnID   int32 `json:"product_en_id" validate:"required"`
	StockQuantity int32 `json:"stock_quantity" validate:"required"`
}

type UpdateProductEnImagePathRequest struct {
	ProductEnID  int32  `json:"product_en_id" validate:"required"`
	ImagesPathEn string `json:"images_path_en"`
}
