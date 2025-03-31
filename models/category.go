package models

import "database/sql"

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
	SubCategoryIDEn   int32         `json:"subCategoryIdEn"`
	SubCategoryNameEN string        `json:"subCategoryNameEn"`
	CategoryEnID      int32         `json:"categoryEnId"`
	SCategories       []SCategoryEn `json:"sCategories,omitempty"`
}

// models/s_category_en.go
type SCategoryEn struct {
	SCategoryIdEn   int32       `json:"sCategoryIdEn"`
	SCategoryNameEn string      `json:"sCategoryNameEn"`
	SubCategoryIDEn int32       `json:"subCategoryIdEn"`
	Products        []ProductEn `json:"products"`
}

type SubCategoryMn struct {
	SubCategoryIDMn   int32         `json:"SubCategoryIDMn"`
	SubCategoryNameMn string        `json:"SubCategoryNameMN"`
	CategoryMnID      int32         `json:"categoryMnId"`
	SCategories       []SCategoryMn `json:"sCategories,omitempty"`
}

type SCategoryMn struct {
	SCategoryIdMn   int32       `json:"sCategoryIdMn"`
	SCategoryNameMn string      `json:"sCategoryNameMn"`
	SubCategoryIDMn int32       `json:"subCategoryIDMn"`
	Products        []ProductMn `json:"products"`
}

type UpdateSubCategoryEn struct {
	SubCategoryNameEn string `json:"subCategoryNameEn"`
	SubCategoryEnId   int32  `json:"subCategoryENID"`
}

type UpdateSCategoryEn struct {
	SCategoryNameEn string `json:"sCategoryNameEn"`
	SCategoryIdEn   int32  `json:"sCategoryIdEn"`
}

type UpdateSubCategoryMn struct {
	SubCategoryNameMn string `json:"subCategoryNameMn"`
	SubCategoryMnID   int32  `json:"subCategoryMNID"`
}

type UpdateSCategoryMn struct {
	SCategoryName string `json:"sCategoryNameMn"`
	SCategoryIdMn int32  `json:"sCategoryIdMn"`
}

type CategoryWithSubCategoriesAndProductsEn struct {
	CategoryEnID    int32                       `json:"categoryEnId"`
	CategoryNameEn  string                      `json:"categoryNameEn"`
	SubcategoriesEn []SubCategoryWithProductsEn `json:"subcategories"`
}

// models/sub_category_with_products_en.go
type SubCategoryWithProductsEn struct {
	SubCategoryIDEn   int32         `json:"subCategoryIdEn"`
	SubCategoryNameEn string        `json:"subCategoryNameEn"`
	SCategories       []SCategoryEn `json:"sCategories"`
}
type ProductEn struct {
	ProductEnID   int32          `json:"productEnId"`
	ProductNameEn string         `json:"productNameEn"`
	PriceEn       string         `json:"priceEn"`
	StockQuantity int32          `json:"stockQuantity"`
	ImagesPathEn  sql.NullString `json:"imagesPathEn"`
}

type SubCategoryWithProductsMn struct {
	SubCategoryIDMn   int32         `json:"SubCategoryIDMn"`
	SubCategoryNameMn string        `json:"subCategoryNameMn"`
	SCategories       []SCategoryMn `json:"sCategories"`
}

type CategoryWithSubCategoriesAndProductsMn struct {
	CategoryMnID    int32                       `json:"categoryMnId"`
	CategoryNameMn  string                      `json:"categoryNameMn"`
	SubCategoriesMn []SubCategoryWithProductsMn `json:"subCategoriesMn"`
}

type ProductMn struct {
	ProductMnID   int32          `json:"productMnID"`
	ProductNameMn string         `json:"productNameMn"`
	PriceMn       string         `json:"priceMn"`
	StockQuantity int32          `json:"stockQuantity"`
	ImagesPathMn  sql.NullString `json:"imagesPathMn"`
}

type UpdateSubCatogoryWithCategoryEn struct {
	CategoryEnID int32 `json:"categoryEnId"`
}
