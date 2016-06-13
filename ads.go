package mgtvPmpProto

const (
	IMPRESSION_TYPE_ZERO  = 0
	IMPRESSION_TYPE_ONE   = 1
	IMPRESSION_TYPE_TWO   = 2
	IMPRESSION_TYPE_THREE = 3
	IMPRESSION_TYPE_FOUR  = 4
	AdTypePicture         = 1
	AdTypeVideo           = 2
	AdTypeFlv             = 3
)

type IURL struct {
	Event int //0-开始, 1-四分之一, 2-二分之一, 3-四分之三, 4-结束
	Url   string
}

type ADS struct {
	//千次展现金额，计划的最高竞标价格, 单位为分
	Price int

	//物料地址，需要http://或者https://协议头
	AdUrl string

	//广告的目标跳转地址，需要http://或者https://协议头
	ClickThroughUrl string `json:"click_through_url"`

	//展示监播地址
	IUrl []IURL

	//点击监播地址
	CUrl []string

	//视频广告ID
	AdId string

	//视频广告的标题
	Title string

	//视频广告时长，单位为秒
	Duration int

	//视频类型 1-图片，2-视频，3-flv
	Ctype int

	//创意宽度
	Width int

	//创意高度
	Height int
}
