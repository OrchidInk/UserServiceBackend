package models

type CategoryEn struct {
	CategoryEnID   int32           `json:"category_en_id"`
	CategoryNameEn string          `json:"category_name_en"`
	SubCategories  []SubCategoryEn `json:"subcategories"`
}

type CategoryMn struct {
	CategoryMnID   int32           `json:"category_mn_id"`
	CategoryNameMn string          `json:"category_name_mn"`
	SubCategories  []SubCategoryMn `json:"subcategories"`
}

type UpdateCategoryEn struct {
	CategoryEnId   int32  `json:"category_en_id"`
	CategoryNameEn string `json:"category_name_en" validate:"required"`
}

type UpdateCategoryMn struct {
	CategoryMnId   int32  `json:"category_mn_id"`
	CategoryNameMn string `json:"category_name_mn" validate:"required"`
}

type CreateCategoryEn struct {
	CategoryNameEn string `json:"category_name_en" validate:"required"`
}

type CreateCategoryMn struct {
	CategoryNameMn string `json:"category_name_mn" validate:"required"`
}

type SubCategoryEn struct {
	SubCategoryIdEN   int32  `json:"subcategoryIDEN"`
	SubCategoryNameEN string `json:"SubCategoryNameEN"`
	CategoryEnID      int32  `json:"category_en_id"`
}

type SubCategoryMn struct {
	SubCategoryIdMn   int32  `json:"sub_category_ID_mn"`
	SubCategoryNameMn string `json:"sub_category_name_mN"`
	CategoryMnID      int32  `json:"category_mn_id"`
}

type UpdateSubCategoryEn struct {
	SubCategoryNameEn string `json:"Sub_Category_NameEn"`
	SubCategoryEnId   int32  `json:"Sub_Category_EN_ID"`
}

type UpdateSubCategoryMn struct {
	SubCategoryNameMn string `json:"Sub_Category_NameMn"`
	SubCategoryMnID   int32  `json:"Sub_Category_MN_ID"`
}

type DeleteSubCategoryEn struct {
	SubCategoryEnID int32 `json:"Sub_Category_EN_ID"`
}

type DeleteSubCategoryMn struct {
	SubCategoryMnID int32 `json:"Sub_Category_MN_ID"`
}
