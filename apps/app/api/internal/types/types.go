// Code generated by goctl. DO NOT EDIT.
package types

type Request struct {
	Name string `path:"name,options=you|me"`
}

type Response struct {
	Message string `json:"message"`
}

type QrCodeResponse struct {
	QrCodeUrl string `json:"qrCodeUrl"`
	QrCodeImg string `json:"qrCodeImg"`
	QrCodeKey string `json:"qrCodeKey"`
}

type BiliLoginRequest struct {
	QrCodeKey string `form:"qrCodeKey"`
}

type BiliLoginResponse struct {
	Code       int    `json:"code"`
	Dedeuserid string `json:"dedeuserid"`
	Sessdata   string `json:"sessdata"`
	BiliJct    string `json:"biliJct"`
}

type BiliUserRequest struct {
	Dedeuserid string `header:"dedeuserid"`
	Sessdata   string `header:"sessdata"`
	BiliJct    string `header:"biliJct"`
}

type BiliUserResponse struct {
	Id            int64   `json:"id"`
	Dedeuserid    string  `json:"dedeuserid"`
	Username      string  `json:"username"`
	Coins         float64 `json:"coins"`
	CurrentExp    int64   `json:"currentExp"`
	NextExp       int64   `json:"nextExp"`
	IsLogin       bool    `json:"isLogin"`
	UpgradeDays   int64   `json:"upgradeDays"`
	Level         int64   `json:"level"`
	Medals        string  `json:"medals"`
	VipStatus     int64   `json:"vipStatus"`
	VipType       int64   `json:"vipType"`
	VipLabelTheme string  `json:"vipLabelTheme"`
	Sign          string  `json:"sign"`
	LastRunTime   int64   `json:"lastRunTime"`
	CreateTime    int64   `json:"createTime"`
}

type BiliUserListRequest struct {
	Page int64 `form:"page"`
	Size int64 `form:"size"`
}

type BiliUserListResponse struct {
	Total int64              `json:"total"`
	Users []BiliUserResponse `json:"users"`
}

type BiliTaskConfigRequest struct {
	Dedeuserid string `header:"dedeuserid"`
	Sessdata   string `header:"sessdata"`
	BiliJct    string `header:"biliJct"`
}

type BiliTaskConfigResponse struct {
	Id                 int64  `json:"id"`
	DonateCoins        int64  `json:"donateCoins"`
	ReserveCoins       int64  `json:"reserveCoins"`
	AutoCharge         bool   `json:"autoCharge"`
	DonateGift         bool   `json:"donateGift"`
	AutoChargeTarget   string `json:"autoChargeTarget"`
	DonateGiftTarget   string `json:"donateGiftTarget"`
	DevicePlatform     string `json:"devicePlatform"`
	DonateCoinStrategy int64  `json:"donateCoinStrategy"`
	UserAgent          string `json:"userAgent"`
	SkipTask           bool   `json:"skipTask"`
	FollowDeveloper    bool   `json:"followDeveloper"`
	CreateTime         int64  `json:"createTime"`
	UpdateTime         int64  `json:"updateTime"`
}

type BiliTaskConfigAddRequest struct {
	Dedeuserid         string `header:"dedeuserid"`
	Sessdata           string `header:"sessdata"`
	BiliJct            string `header:"biliJct"`
	DonateCoins        int64  `json:"donateCoins"`
	ReserveCoins       int64  `json:"reserveCoins"`
	AutoCharge         bool   `json:"autoCharge"`
	DonateGift         bool   `json:"donateGift"`
	AutoChargeTarget   string `json:"autoChargeTarget"`
	DonateGiftTarget   string `json:"donateGiftTarget"`
	DevicePlatform     string `json:"devicePlatform"`
	DonateCoinStrategy int64  `json:"donateCoinStrategy"`
	UserAgent          string `json:"userAgent"`
	SkipTask           bool   `json:"skipTask"`
	FollowDeveloper    bool   `json:"followDeveloper"`
}

type BiliTaskConfigEditRequest struct {
	Id                 int64  `json:"id"`
	Dedeuserid         string `header:"dedeuserid"`
	Sessdata           string `header:"sessdata"`
	BiliJct            string `header:"biliJct"`
	DonateCoins        int64  `json:"donateCoins"`
	ReserveCoins       int64  `json:"reserveCoins"`
	AutoCharge         bool   `json:"autoCharge"`
	DonateGift         bool   `json:"donateGift"`
	AutoChargeTarget   string `json:"autoChargeTarget"`
	DonateGiftTarget   string `json:"donateGiftTarget"`
	DevicePlatform     string `json:"devicePlatform"`
	DonateCoinStrategy int64  `json:"donateCoinStrategy"`
	UserAgent          string `json:"userAgent"`
	SkipTask           bool   `json:"skipTask"`
	FollowDeveloper    bool   `json:"followDeveloper"`
}

type PushConfigRequest struct {
	Dedeuserid string `header:"dedeuserid"`
}

type PushConfigResponse struct {
	Id     int64                  `json:"id"`
	UserId string                 `json:"userId"`
	Config map[string]interface{} `json:"config"`
}

type PushConfigAddRequest struct {
	UserId string `json:"userId"`
	Config string `json:"config"`
}

type PushConfigEditRequest struct {
	Id     int64                  `path:"cid"`
	UserId string                 `json:"userId"`
	Config map[string]interface{} `json:"config"`
}
