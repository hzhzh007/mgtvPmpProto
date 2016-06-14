package mgtvPmpProto

type Deal struct {
	//符合条件的deal id
	Id string `json:"id"`

	//竞价的方式，目前都是1，即第一竞价法。最高的deal获得竞价成功，取最高出价作为最终胜出价。注：只有当多个Deal同时响应时才互相之间竞价。这个字段和OpenRTB相同。
	At int `json:"at"`

	//该deal对应的结算价格
	Price int `json:"price"`
}
