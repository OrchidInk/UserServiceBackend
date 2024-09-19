package models

type CategoryEn struct {
	CategoryEnID   int32  `json:"category_en"`
	CategoryNameEn string `json:"category_name_en"`
}

type CategoryMn struct {
	CategoryMnID   int32  `json:"category_mn"`
	CategoryNameMn string `json:"category_name_mn"`
}

type UpdateCategoryEn struct {
	CategoryNameEn string `json:"category_name_en" validate:"required"`
}

type UpdateCategoryMn struct {
	CategoryNameMn string `json:"category_name_mn" validate:"required"`
}

type CreateCategoryEn struct {
	CategoryNameEn string `json:"category_name_en" validate:"required"`
}

type CreateCategoryMn struct {
	CategoryNameMn string `json:"category_name_mn" validate:"required"`
}
