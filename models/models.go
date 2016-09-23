package models

import (
	"time"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"path"
)

const (
	DB_NAME = "test"

)

//分类
//tag orm 表示只有orm 会去读这个tag
type Category struct {
	Id              int64                   //id
	Title           string                  //标题
	Created         time.Time `orm:"index"` //创建时间
	Views           int64     `orm:"index"` //浏览次数
	TopicTime       time.Time `orm:"index"` // 文章发表时间
	TopicCount      int64                   //分类中有多少篇文章
	TopicLastUserId int64                   //最后操作分类用户的id
}

//文章
type Topic struct {
	Id              int64                     //id
	Uid             int64                     //用户id, 谁写的
	Title           string                    //标题
	Content         string `orm:"size(5000)"` //内容，设置长度5000个字节
	Attachment      string                    //附件
	Created         time.Time `orm:"index"`   //创建时间
	Update          time.Time `orm:"index"`   //更新时间
	Views           int64     `orm:"index"`   //浏览次数
	Author          string                    //作者名
	ReplyTime       time.Time `orm:"index"`   //回复时间
	ReplyCount      int64                     //评论的个数
	ReplyLastUserId int64                     //最后评论用户id
}

func init() {
	orm.RegisterModel(new(Category), new(Topic)) //使用这个来指定建表
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:root@/test?charset=utf8", 10) //10 最大连接数
}

//另外一种注册
func RegisterDB()  {
	//不存在文件夹时创建文件夹和文件
	os.MkdirAll(path.Dir(DB_NAME), os.ModePerm)
	os.Create(DB_NAME)


	//beego 注册model
	orm.RegisterModel(new(Category), new(Topic))
}
