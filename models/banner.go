package models

type BannerInfo struct {
	BannerId       int32  `json:"bannerId"`
	BannerImageUrl string `json:"bannerImageUrl"`
}

type CreateBannerRequest struct {
	BannerId       int32  `json:"bannerId"`
	BannerImageUrl string `json:"bannerImageUrl" validate:"required"`
}

type UpdateBannerRequest struct {
	BannerId       int32  `json:"bannerId" validate:"required"`
	BannerImageUrl string `json:"bannerImageUrl" validate:"required"`
}
