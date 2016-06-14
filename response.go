package mgtvPmpProto

const (
	ResponseBidOK = 200
	ResponseNoAd  = 204
)

/*
example
{
    "version": 3,
    "bid": "84687db331a64d0a9ef2fe6694defcf4",
    "err_code": 200,
    "ads": [
        {
            "price": 1100,
            "ad_url": "http://dsp.kss.ksyun.com/dsp/20151225/1451038491021.mp4",
            "click_through_url": "http://www.clickurl.com",
            "iurl": [
                {
                    "url": "http://www.showurl.com/jzh_show?c=%%SETTLE_PRICE%%",
                    "event": 0
                }
            ],
            "curl": [
                "http://www.click_stc.com/jzh_click"
            ],
            "adid": "10851_1450337808192",
            "title": "mg_test12",
            "duration": 15,
            "ctype": 2,
            "width": 0,
            "height": 0
        }
    ]
}
*/
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
