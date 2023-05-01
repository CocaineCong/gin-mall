package consts

const (
	EmailOperationBinding = iota + 1
	EmailOperationNoBinding
	EmailOperationUpdatePassword
)

var EmailOperationMap = map[uint]string{
	EmailOperationBinding:        "您正在绑定邮箱, 请点击链接确定身份 %s",
	EmailOperationNoBinding:      "您正在解邦邮箱, 请点击链接确定身份 %s",
	EmailOperationUpdatePassword: "您正在修改密码, 请点击链接校验身份 %s",
}
