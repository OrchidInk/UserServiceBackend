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
	ProductEnID int32  `json:"productEnId" validate:"required"`
	PriceEn     string `json:"priceEn" validate:"required"`
}

type UpdateProductEnStockRequest struct {
	ProductEnID   int32 `json:"productEnId" validate:"required"`
	StockQuantity int32 `json:"stockQuantity" validate:"required"`
}

type UpdateProductEnImagePathRequest struct {
	ProductEnID  int32  `json:"productEnId" validate:"required"`
	ImagesPathEn string `json:"imagesPathEn"`
}

type PurchaseProductEnRequest struct {
	ProductEnID       int32 `json:"productEnId" validate:"required"`
	QuantityPurchased int32 `json:"quantityPurchased" validate:"required"`
}

type PurchaseProductMnRequest struct {
	ProductMnID       int32 `json:"productMnId" validate:"required"`
	QuantityPurchased int32 `json:"quantityPurchased" validate:"required"`
}
