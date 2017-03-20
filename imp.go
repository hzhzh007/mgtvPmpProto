package mgtvPmpProto

const (
	LocationFrontAd  = 1
	LocationMidAd    = 2
	LocationLastAd   = 3
	LocationPauseAd  = 4
	LocationFloatAd  = 5
	LocationBufferAd = 6
)

type Imp struct {
	//广告位ID，广告位的唯一标识
	SpaceId string `json:"space_id"`

	//广告位宽度
	Width int `json:"width"`

	//广告位高度
	Height int `json:"height"`

	//流量最低竞标价格, DSP出价需要不小于这个值, 单位为分
	MinCpmPrice int `json:"min_cpm_price"`

	//播放器ID
	PlayerId int `json:"player_id"`

	//该广告位支持的素材类型 1-图片，2-视频，3-flv
	Ctype []int

	//视频类广告播放时长，单位为秒
	PlayTime int `json:"playtime"`

	//视频位置类型 1-前贴片，2-中贴片，3-后贴片，4-暂停，5-浮层，6-缓冲（APP无）
	Location int

	//视频广告位序 同类型下该广告位所处位序，从1开始，例：前贴第三个广告位，其位序位3
	Order int

	//Preferred Deal相关参数
	Pmp Deals `json:"pmp"`
	
	//视频广告最短播放时长，单位为秒
	MinPlayTime int `json:min_playtime`

	//视频广告最大播放时长，单位为秒
	MaxPlayTime int `json:max_playtime`
}
