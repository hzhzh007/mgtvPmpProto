package mgtvPmpProto

/*
{
    "version": 3,
    "bid": "fa8a8a16d5cd46159765688d79337806",
    "request_type": 0,
    "imp": [
        {
            "space_id": "148",
            "width": 0,
            "height": 0,
            "min_cpm_price": 1000,
            "player_id": 4620,
            "location": 1,
            "ctype": [
                2
            ],
            "playtime": 15,
            "order": 1
        }
    ],
    "device": {
        "idfa": "C8CAE811-4A1F-4901-BBF6-D3D4646A789E",
        "openudid": "14779b66e5fa0e6e8561781e10e45cb8a5da773e",
        "os": "ios 9.200000",
        "brand": "apple",
        "model": "iPad",
        "sw": 768,
        "sh": 1024,
        "ip": "175.10.84.44",
        "connectiontype": 0,
        "type": 33,
        "version": "4.2.4",
        "screen_orientation": 2
    },
    "video": {
        "video_id": 1819244,
        "video_name": "天天向上20151023期：孙杨谈母亲动情落泪",
        "collection_id": 18,
        "collection_name": "天天向上",
        "item_ids": "20,22,15,21",
        "item_names": "脱口秀,真人秀,搞笑,娱乐节目",
        "area_id": 2015,
        "area_name": "内地",
        "duration": 15,
        "type": 1
    }
}

*/
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
