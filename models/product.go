package models

type CreateProductEnRequest struct {
	ProductNameEn         string   `json:"productNameEN" validate:"required"`
	SCategoryEnID         int32    `json:"sCategoryEnId" validate:"required"`
	PriceEn               string   `json:"priceEn" validate:"required"`
	StockQuantity         int32    `json:"stockQuantity" validate:"required"`
	ImagesPathEn          []string `json:"imagesPath" `
	DescriptionEn         string   `json:"descriptionEn" validate:"required"`
	BrandEn               string   `json:"brandEn"`
	ManufacturedCountryEn string   `json:"manufacturedCountryEn"`
	// Instead of a single ColorId and SizeId, use arrays:
	ColorIds         []int32 `json:"colorIds" validate:"required"` // client sends an array e.g. [16,17,18]
	SizeIds          []int32 `json:"sizeIds" validate:"required"`  // e.g. [1,2,5]
	PenOutputEn      string  `json:"penOutputEn"`
	FeaturesEn       string  `json:"featuresEn"`
	MaterialEn       string  `json:"materialEn"`
	StapleSizeEn     string  `json:"stapleSizeEn"`
	CapacityEn       string  `json:"capacityEn"`
	WeightEn         string  `json:"weightEn"`
	ThicknessEn      string  `json:"thinknessEn"`
	PackagingEn      string  `json:"packagingEn"`
	UsageEn          string  `json:"usageEn"`
	InstructionsEn   string  `json:"instructionsEn"`
	ProductCodeEn    string  `json:"productCodeEn"`
	CostPriceEn      string  `json:"costPriceEn"`
	RetailPriceEn    string  `json:"retailPriceEn"`
	WarehouseStockEn int32   `json:"warehouseStockEn"`
}

type CreateProductMnRequest struct {
	ProductNameMn         string   `json:"productNameMn" validate:"required"`
	SCategoryMnID         int32    `json:"sCategoryMnId" validate:"required"`
	PriceMn               string   `json:"priceMn" validate:"required"`
	StockQuantity         int32    `json:"stockQuantity" validate:"required"`
	ImagesPathMn          []string `json:"imagesPath" validate:"required"`
	DescriptionMn         string   `json:"descriptionMn" validate:"required"`
	BrandMn               string   `json:"brandMn"`
	ManufacturedCountryMn string   `json:"manufacturedCountryMn"`

	// Arrays now
	ColorIds []int32 `json:"colorIds"`
	SizeIds  []int32 `json:"sizeIds"`

	PenOutputMn      string `json:"penOutputMn"`
	FeaturesMn       string `json:"featuresMn"`
	MaterialMn       string `json:"materialMn"`
	StapleSizeMn     string `json:"stapleSizeMn"`
	CapacityMn       string `json:"capacityMn"`
	WeightMn         string `json:"weightMn"`
	ThicknessMn      string `json:"thinknessMn"`
	PackagingMn      string `json:"packagingMn"`
	UsageMn          string `json:"usageMn"`
	InstructionsMn   string `json:"instructionsMn"`
	ProductCodeMn    string `json:"productCodeMn"`
	CostPriceMn      string `json:"costPriceMn"`
	RetailPriceMn    string `json:"retailPriceMn"`
	WarehouseStockMn int32  `json:"warehouseStockMn"`
}

type UpdateProductEnPriceRequest struct {
	ProductEnID int32  `json:"productEnId" validate:"required"`
	PriceEn     string `json:"priceEn" validate:"required"`
}

type UpdateSProductEn struct {
	ProductEnID   int32 `json:"productEnId" validate:"required"`
	StockQuantity int32 `json:"stockQuantity" validate:"required"`
}

type UpdateSProductMn struct {
	ProductMnID   int32 `json:"productMnId" validate:"required"`
	StockQuantity int32 `json:"stockQuantity" validate:"required"`
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

type DetailMn struct {
	DetailMnID  int32  `json:"detailMnId"`
	ChoiceName  string `json:"choiceName"`
	ChoiceValue string `json:"choiceValue"`
}
