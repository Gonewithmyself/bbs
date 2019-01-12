package types

import "github.com/astaxie/beego"

func TableName(str string) string {
	return beego.AppConfig.String("dbprefix") + str
}
