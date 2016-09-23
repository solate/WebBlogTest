package models

import (
	"testing"
	"github.com/astaxie/beego/orm"
)

func TestOrm(t *testing.T) {

	orm.Debug = true  //debug 模式

	//自动建表,
	//第二个参数 true,每次删掉重建, false 只建立一次
	//第三个参数 是否大于相关信息 true
	orm.RunSyncdb("default", false, true)


}
