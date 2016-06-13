package mgtvPmpProto

type Device struct {
	//地理位置-经度
	Len string `json:"len"`

	//地理位置-纬度
	Lon string `json:"lon"`

	//IMEI的sha1值
	Imei string

	//iOS的idfa
	Idfa string

	//Android ID
	Anid string

	//Mac地址的sha1值
	Mac string

	//设备操作系统
	Os string `json:"os"`

	//设备品牌
	Brand string

	//设备型号
	Model string

	//屏幕分辨率宽度，单位为像素
	Sw int

	// 屏幕分辨率高度，单位为像素
	Sh int

	//访问者的IP地址
	Ip string

	//访问者的代理浏览器类型
	Ua string

	//运营商，0-未知，1-中国移动，2-中国联通，3-中国电信，4-互联网电视
	ConnectionType int `json:"connectiontype"`

	//网络环境，0-wifi，1-mobile，2-no network
	Carrier int

	//设备平台类型，1-PC,21-安卓H5平板，22-安卓H5手机，23-苹果H5平板，
	//24-苹果H5手机，31-安卓APP平板，32-安卓APP手机，33-苹果APP平板，
	//34-苹果APP手机，41-小米手机APP_SDK,42-百度视频手机APP_SDK,100-OTT
	Type int `json:"type"`

	//客户端版本
	Version string

	//横竖屏状态，0-未知，1-竖屏，2-横屏
	ScreenOrientation int `json:"screen_orientation"`

	//IOS设备的OpenUDID，完全删除全部带有OpenUDID SDK包的App或者升级iOS系统后会变化
	Openudid string `json:"openudid"`

	//iOS 6.0以前(无任何限制)，iOS 6.0~7.0之间(限制了UDID)，iOS7开始禁掉了UDID
	Odin string

	//The Open Device Identification Number (ODIN) is a number designed for mobile app developers and service providers to uniquely identify their users’ devices in a convenient and interoperable manner.
	Udid string

	//Windows Phone用户终端的DUID经过md5后的值
	Duid string
}
