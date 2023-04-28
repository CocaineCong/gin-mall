package consts

const (
	OrderTypeUnPaid = iota + 1
	OrderTypePaid
)

var OrderTypeMap = map[int]string{
	OrderTypeUnPaid: "未支付",
	OrderTypePaid:   "已支付",
}
