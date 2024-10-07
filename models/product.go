package models

type CreateProductEnRequest struct {
	ProductNameEn   string `json:"productNameEN" validate:"required"`
	SubCategoryEnID int32  `json:"subCategoryEnId" validate:"required"`
	PriceEn         string `json:"priceEn" validate:"required"`
	StockQuantity   int32  `json:"stockQuantity" validate:"required"`
	ImagesPathEn    string `json:"imagesPathEn" validate:"required"`
}

type CreateProductMnRequest struct {
	ProductNameMn   string `json:"productNameMN" validate:"required"`
	SubCategoryMnID int32  `json:"subCategoryMnID" validate:"required"`
	PriceMn         string `json:"priceMn" validate:"required"`
	StockQuantity   int32  `json:"stockQuantity" validate:"required"`
	ImagesPathMn    string `json:"imagesPathMn" validate:"required"`
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
