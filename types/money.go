package types

type ShowMoneyServiceReq struct {
	Key string `json:"key" form:"key"`
}

type MoneyResp struct {
	UserID    uint   `json:"user_id" form:"user_id"`
	UserName  string `json:"user_name" form:"user_name"`
	UserMoney string `json:"user_money" form:"user_money"`
}

// func BuildMoney(item *model.User, key string) Money {
//
// }
