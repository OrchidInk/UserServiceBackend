package models

import "time"

type CreateProductEnRequest struct {
	ProductNameEn   string `json:"productNameEN" validate:"required"`
	SubCategoryEnID int32  `json:"subCategoryEnId" validate:"required"`
	PriceEn         string `json:"priceEn" validate:"required"`
	StockQuantity   int32  `json:"stockQuantity" validate:"required"`
	ImagesPathEn    string `json:"imagesPathEn" validate:"required"`
}

type CreateProductMnRequest struct {
	ProductNameMn   string `json:"productNameMN" validate:"required"`
	SubCategoryMnID int32  `json:"subCategoryMnId" validate:"required"`
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

type ProductWithDetailsEn struct {
	ProductEnID     int32      `json:"productEnId"`
	ProductNameEn   string     `json:"productNameEn"`
	SubCategoryIDEn int32      `json:"subCategoryIdEn"`
	PriceEn         string     `json:"priceEn"`
	StockQuantity   int32      `json:"stockQuantity"`
	ImagesPathEn    string     `json:"imagesPathEn"`
	CreatedAt       time.Time  `json:"createdAt"`
	UpdatedAt       time.Time  `json:"updatedAt"`
	Details         []DetailEn `json:"details"`
}

type ProductWithDetailsMn struct {
	ProductMnID     int32      `json:"productMnId"`
	ProductNameMn   string     `json:"productNameMn"`
	SubCategoryIDMn int32      `json:"subCategoryIdMn"`
	PriceMn         string     `json:"priceMn"`
	StockQuantity   int32      `json:" stockQuantity"`
	ImagesPathMn    string     `json:"imagesPathMn"`
	CreatedAt       time.Time  `json:"createdAt"`
	UpdatedAt       time.Time  `json:"updatedAt"`
	Details         []DetailMn `json:"details"`
}

type DetailMn struct {
	DetailMnID  int32  `json:"detailMnId"`
	ChoiceName  string `json:"choiceName"`
	ChoiceValue string `json:"choiceValue"`
}
