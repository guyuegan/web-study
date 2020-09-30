package db_test

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
	"testing"
)

/*
// 导入驱动
// _ "github.com/go-sql-driver/mysql"

// 注册驱动
orm.RegisterDriver("mysql", orm.DR_MySQL)

// 设置默认数据库
//mysql用户：root ，密码：zxxx ， 数据库名称：test ， 数据库别名：default
 orm.RegisterDataBase("default", "mysql", "root:zxxx@/test?charset=utf8")

 */
func init() {
	// 注册驱动
	orm.RegisterDriver("postgres", orm.DRPostgres)
	// PostgresQL用户：postgres ，密码：123456 ， 数据库名称：postgres ， 数据库别名：default
	orm.RegisterDataBase("default", "postgres", "user=postgres password=123456 dbname=postgres" +
		" host=127.0.0.1 port=5432 sslmode=disable")

	// 注册定义的model
	orm.RegisterModel(new(User))
	// RegisterModel 也可以同时注册多个 model
	// orm.RegisterModel(new(User), new(Profile), new(Post))

	// 创建 table
	orm.RunSyncdb("default", false, true)
}

type User struct {
	Id int
	Name string `orm:"size(100)"`
}

func Test3(t *testing.T)  {
	//打开到数据库的链接，然后创建一个 beego om 对象
	om := orm.NewOrm()
	user := User{Name: "gyg-om"}

//	插入
	id, err := om.Insert(&user)
	fmt.Printf("Id: %d, err: %v\n", id, err)

//	更新
	user.Name = "gyg-om-update"
	num, err := om.Update(&user)
	fmt.Printf("Num: %d, err: %v\n", num, err)

//	读取
	u := User{Id: user.Id}
	err = om.Read(&u)
	fmt.Printf("err: %v\n", err)

//	删除
//	num, err = om.Delete(&u)
//	fmt.Printf("Num: %d, err: %v\n", num, err)
}
