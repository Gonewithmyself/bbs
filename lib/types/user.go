package types

import "time"

type User struct {
	Role       int8
	Id         int
	Account    string
	Name       string
	Password   string
	LoginCount int
	LastTime   time.Time
	LastIp     string
	Created    time.Time
	Updated    time.Time
}

func (m *User) TableName() string {
	return TableName("user")
}

type RegForm struct {
	Account string `form:"account"`
	Pass    string `form:"pass"`
	Repass  string `form:"repass"`
	Name    string `form:"username"`
	Vercode string `form:"vercode"`
}
