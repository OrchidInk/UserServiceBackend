package models

type CategoryEn struct {
	CategoryEnID     int32  `json:"category_en_id"`
	CategoryNameEn   string `json:"category_name_en"`
	ParentCategoryID int32  `json:"parent_category_id"`
}

type CategoryMn struct {
	CategoryMnID     int32  `json:"category_mn_id"`
	CategoryNameMn   string `json:"category_name_mn"`
	ParentCategoryID int32  `json:"parent_category_id"`
}

type UpdateCategoryEn struct {
	CategoryNameEn   string `json:"category_name_en" validate:"required"`
	ParentCategoryID int32  `json:"parent_category_id"`
}

type UpdateCategoryMn struct {
	CategoryNameMn   string `json:"category_name_mn" validate:"required"`
	ParentCategoryID int32  `json:"parent_category_id"`
}

type CreateCategoryEn struct {
	CategoryNameEn   string `json:"category_name_en" validate:"required"`
	ParentCategoryID int32  `json:"parent_category_id"`
}

type CreateCategoryMn struct {
	CategoryNameMn   string `json:"category_name_mn" validate:"required"`
	ParentCategoryID int32  `json:"parent_category_id"`
}
