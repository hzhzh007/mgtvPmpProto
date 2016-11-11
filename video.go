package mgtvPmpProto

const (
	//片源，1-正片
	VideoNormalVideoType = 1

	//片源，2-短片
	VideoShortVideoType = 1
)

type Video struct {
	//视频ID
	VideoId int `json:"video_id"`

	//视频名称
	VideoName string `json:"video_name"`

	//视频所属合集ID
	CollectionId int `json:"collection_id"`

	//视频所属合集名称
	CollectionName string `json:"collection_name"`
	
	//频道id
	ChannelId int `json:"channel_id"`

	//视频分类ID列表
	ItemIds string `json:"item_ids"`

	//视频分类名称
	ItemNames string `json:"item_names"`

	//视频所属地区ID
	AreaId int `json:"area_id"`

	//视频所属地区名称
	AreaName string `json:"area_name"`

	//出品年代
	Year int

	//视频时长，单位为秒
	Duration int `json:"duration"`

	//Grades string `json:"grades"`

	//片源，1-正片，2-短片
	Type int `json:"type"`
}
