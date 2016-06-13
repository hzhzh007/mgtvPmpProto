package mgtvPmpProto

type Request struct {
	//协议版本
	Version int

	//ADX生成的唯一的竞价ID
	Bid string

	//如果request_type = 1，那么这是一个测试要求。DSP 需要反回正常的BidResponse，但是结果是不显示给用户，DSP不会收到一个响应
	//如果request_type = 2，那么这个要求的目的是衡量网络延迟，即心跳测试。DSP需要返回一个空的BidResponse
	RequestType int `json:"request_type"`

	//Imp对象数组，但仅包含一个Imp对象，描述广告位信息
	Device Device `json:"device"`

	//Device对象，描述设备信息
	Imp []Imp `json:"imp"`

	//Video对象，描述视频信息
	Video Video `json:"video"`
}
