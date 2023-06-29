package model

import (
	"fmt"
	"testing"

	conf "github.com/CocaineCong/gin-mall/config"
)

func TestMain(m *testing.M) {
	re := conf.ConfigReader{FileName: "../../../config/locales/config.yaml"}
	conf.InitConfigForTest(&re)
	fmt.Println("Write tests on values: ", conf.Config)
	m.Run()
}

func TestUserModelEncryptMoney(t *testing.T) {
	key := "123456" // 6位支付密码
	u := User{
		UserName: "FanOne",
		Money:    "10000",
	}
	t.Logf("u before encrypt money:%s", u.Money)
	money, err := u.EncryptMoney(key)
	u.Money = money
	if err != nil {
		fmt.Println("err EncryptMoney", err)
	}
	t.Logf("u after encrypt money:%s", u.Money)
	m, err := u.DecryptMoney(key)
	if err != nil {
		fmt.Println("err EncryptMoney", err)
	}
	t.Logf("u after encrypt money:%f", m)
}
