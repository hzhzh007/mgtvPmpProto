package mgtvPmpProto

const (
	ImpressionTypeZero  = 0
	ImpressionTypeOne   = 1
	ImpressionTypeTwo   = 2
	ImpressionTypeThree = 3
	ImpressionTypeFour  = 4
	AdTypePicture       = 1
	AdTypeVideo         = 2
	AdTypeFLV           = 3
)

type IURL struct {
	//0-开始, 1-四分之一, 2-二分之一, 3-四分之三, 4-结束
	Event int
	//DSP曝光监测地址必须放该数组第一位，并且该监测地址中包要含价格宏（当次广告成交价）
	//例：http://www.xxx.com/jzh_show?c=%%SETTLE_PRICE%%，如果DSP还有其他曝光监测，其监测链接中不可在设置该价格宏
	Url string
}

type ADS struct {
	//千次展现金额，计划的最高竞标价格, 单位为分
	Price int `json:"price"`

	//物料地址，需要http://或者https://协议头
	AdUrl string `json:"ad_url"`

	//广告的目标跳转地址，需要http://或者https://协议头
	ClickThroughUrl string `json:"click_through_url"`

	//展示监播地址
	IUrl []IURL `json:"iurl"`

	//点击监播地址
	CUrl []string `json:"curl"`

	//视频广告ID
	AdId string `json:"adid"`

	//视频广告的标题
	Title string `json:"title"`

	//视频广告时长，单位为秒
	Duration int `json:"duration"`

	//视频类型 1-图片，2-视频，3-flv
	Ctype int `json:"ctype"`

	//创意宽度
	Width int `json:"width"`

	//创意高度
	Height int `json:"height"`

	//DSP参加的deal id
	DealId int `json:"dealid,omitempty"`
}
