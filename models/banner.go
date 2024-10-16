package models

type BannerInfo struct {
	BannerId        int32  `json:"bannerId"`
	BannerImagePath string `json:"bannerImagePath"`
	BannerImageUrl  string `json:"bannerImageUrl"`
}

type CreateBannerRequest struct {
	BannerId        int32  `json:"bannerId"`
	BannerImagePath string `json:"bannerImagePath" validate:"required"`
	BannerImageUrl  string `json:"bannerImageUrl" validate:"required"`
}

type UpdateBannerRequest struct {
	BannerId        int32  `json:"bannerId" validate:"required"`
	BannerImagePath string `json:"bannerImagePath" validate:"required"`
	BannerImageUrl  string `json:"bannerImageUrl" validate:"required"`
}
