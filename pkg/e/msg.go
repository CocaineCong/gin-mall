package e

var MsgFlags = map[int]string{
	SUCCESS:               "ok",
	UpdatePasswordSuccess: "修改密码成功",
	NotExistInentifier:    "该第三方账号未绑定",
	ERROR:                 "fail",
	InvalidParams:         "请求参数错误",

	ErrorExistNick:          "已存在该昵称",
	ErrorExistUser:          "已存在该用户名",
	ErrorNotExistUser:       "该用户不存在",
	ErrorNotCompare:         "账号密码错误",
	ErrorNotComparePassword: "两次密码输入不一致",
	ErrorFailEncryption:     "加密失败",
	ErrorNotExistProduct:    "该商品不存在",
	ErrorNotExistAddress:    "该收获地址不存在",
	ErrorExistFavorite:      "已收藏该商品",

	ERROR_BOSS_CHECK_TOKEN_FAIL:       "商家的Token鉴权失败",
	ERROR_BOSS_CHECK_TOKEN_TIMEOUT:    "商家TOken已超时",
	ERROR_BOSS_TOKEN:                  "商家的Token生成失败",
	ERROR_BOSS:                        "商家Token错误",
	ERROR_BOSS_INSUFFICIENT_AUTHORITY: "商家权限不足",
	ERROR_BOSS_PRODUCT:                "商家读文件错误",

	ErrorAuthCheckTokenFail:        "Token鉴权失败",
	ErrorAuthCheckTokenTimeout:     "TOken已超时",
	ErrorAuthToken:                 "Token生成失败",
	ErrorAuth:                      "Token错误",
	ErrorAuthInsufficientAuthority: "权限不足",
	ErrorReadFile:                  "读文件失败",
	ErrorSendEmail:                 "发送邮件失败",
	ErrorCallApi:                   "调用接口失败",
	ErrorUnmarshalJson:             "解码JSON失败",

	ErrorAdminFindUser: "管理员查询用户失败",

	ErrorDatabase: "数据库操作出错,请重试",

	ErrorOss: "OSS配置错误",
}

// GetMsg 获取状态码对应信息
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[ERROR]
}
