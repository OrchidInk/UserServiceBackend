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
	SubCategoryIdEN   int32  `json:"subcategoryIDEN"`
	SubCategoryNameEN string `json:"SubCategoryNameEN"`
	CategoryEnID      int32  `json:"categoryEnId"`
}

type SubCategoryMn struct {
	SubCategoryIdMn   int32  `json:"subCategoryIDMn"`
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

type DeleteSubCategoryEn struct {
	SubCategoryEnID int32 `json:"subCategoryENID"`
}

type DeleteSubCategoryMn struct {
	SubCategoryMnID int32 `json:"subCategoryMNID"`
}
