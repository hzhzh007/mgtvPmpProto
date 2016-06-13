package mgtvPmpProto

const (
	ResponseBidOk = 200
	ResponseNoAd  = 204
)

type Response struct {
	//协议版本号
	Version int

	//Request传过来的bid
	Bid string

	//响应状态 200:OK, 204:no ad
	ErrCode int `json:"err_code"`
	//广告对象，描述广告信息
	Ads []ADS ``
}
