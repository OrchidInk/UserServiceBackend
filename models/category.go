package models

type CategoryEn struct {
	CategoryEnID   int32           `json:"categoryEnId"`
	CategoryNameEn string          `json:"categoryNameEn"`
	SubCategories  []SubCategoryEn `json:"subcategories"`
}

type CategoryMn struct {
	CategoryMnID   int32           `json:"categoryMnId"`
	CategoryNameMn string          `json:"categoryNameMn"`
	SubCategories  []SubCategoryMn `json:"subcategories"`
}

type UpdateCategoryEn struct {
	CategoryEnId   int32  `json:"categoryEnId"`
	CategoryNameEn string `json:"categoryNameEn" validate:"required"`
}

type UpdateCategoryMn struct {
	CategoryMnId   int32  `json:"categoryMnId"`
	CategoryNameMn string `json:"categoryNameMn" validate:"required"`
}

type CreateCategoryEn struct {
	CategoryNameEn string `json:"categoryNameEn" validate:"required"`
}

type CreateCategoryMn struct {
	CategoryNameMn string `json:"categoryNameMn" validate:"required"`
}

type SubCategoryEn struct {
	SubCategoryIDEn   int32  `json:"SubCategoryIDEn"`
	SubCategoryNameEN string `json:"SubCategoryNameEN"`
	CategoryEnID      int32  `json:"categoryEnId"`
}

type SubCategoryMn struct {
	SubCategoryIDMn   int32  `json:"SubCategoryIDMn"`
	SubCategoryNameMn string `json:"SubCategoryNameMN"`
	CategoryMnID      int32  `json:"categoryMnId"`
}

type UpdateSubCategoryEn struct {
	SubCategoryNameEn string `json:"subCategoryNameEn"`
	SubCategoryEnId   int32  `json:"subCategoryENID"`
}

type UpdateSubCategoryMn struct {
	SubCategoryNameMn string `json:"subCategoryNameMn"`
	SubCategoryMnID   int32  `json:"subCategoryMNID"`
}

type SubCategoryWithProductsEn struct {
	SubCategoryIDEn   int32       `json:"SubCategoryIDEn"`
	SubCategoryNameEn string      `json:"subCategoryNameEn"`
	Products          []ProductEn `json:"products"`
}

type CategoryWithSubCategoriesAndProductsEn struct {
	CategoryEnID    int32                       `json:"categoryEnId"`
	CategoryNameEn  string                      `json:"categoryNameEn"`
	SubcategoriesEn []SubCategoryWithProductsEn `json:"subCategoriesEn"`
}

type ProductEn struct {
	ProductEnID   int32  `json:"productEnId"`
	ProductNameEn string `json:"productNameEn"`
	PriceEn       string `json:"priceEn"`
	StockQuantity int32  `json:"stockQuantity"`
	ImagesPathEn  string `json:"imagesPathEn"`
}

type SubCategoryWithProductsMn struct {
	SubCategoryIDMn   int32       `json:"SubCategoryIDMn"`
	SubCategoryNameMn string      `json:"subCategoryNameMn"`
	Products          []ProductMn `json:"products"`
}

type CategoryWithSubCategoriesAndProductsMn struct {
	CategoryMnID    int32                       `json:"cateogoryMnId"`
	CategoryNameMn  string                      `json:"categoryNameMn"`
	SubCategoriesMn []SubCategoryWithProductsMn `json:"subCategoriesMn'`
}

type ProductMn struct {
	ProductMnID   int32  `json:"productMnID"`
	ProductNameMn string `json:"productNameMn"`
	PriceMn       string `json:"priceMn"`
	StockQuantity int32  `json:"stockQuantity"`
	ImagesPathMn  string `json:"imagesPathMn"`
}

type UpdateSubCatogoryWithCategoryEn struct {
	CategoryEnID int32 `json:"categoryEnId"`
}
