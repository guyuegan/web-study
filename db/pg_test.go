package db_test

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"testing"
)

/*
从上面的代码我们可以看到，PostgreSQL 是通过 $1, $2 这种方式来指定要传递的参数，而不是 MySQL 中的 ?，
另外在 sql.Open 中的 dsn 信息的格式也与 MySQL 的驱动中的 dsn 格式不一样，所以在使用时请注意它们的差异。

还有 pg 不支持 LastInsertId 函数，因为 PostgreSQL 内部没有实现类似 MySQL 的自增 ID 返回，其他的代码几乎是一模一样。

 */
func Test2(t *testing.T)  {
	db, err := sql.Open("postgres", "user=postgres password=123456 dbname=postgres sslmode=disable")
	checkErr(err)

	defer db.Close()

// 插入数据
	stmt, err := db.Prepare("insert into userinfo (username, department, created) values ($1, $2, $3) returning uid")
	checkErr(err)
	res, err := stmt.Exec("pg-gyg", "baas-java", "2020-09-11")
	checkErr(err)

	// pg 不支持这个函数，因为他没有类似 MySQL 的自增 ID
	//lastId, err := res.LastInsertId()
	//CheckErr(err)
	//fmt.Println(lastId)
	var lastId int
	err = db.QueryRow("insert into userinfo(username, department, created) values ($1, $2, $3) returning uid;",
		"pg-gyg2", "baas-java", "2020-09-11").Scan(&lastId)
	checkErr(err)
	fmt.Println("lastId=", lastId)

//	更新数据
	stmt, err = db.Prepare("update userinfo set username=$1 where uid=$2")
	checkErr(err)
	res, err = stmt.Exec("gygUpdate", lastId)
	checkErr(err)

	affected, err := res.RowsAffected()
	checkErr(err)
	fmt.Println(affected)

//	查询数据
	rows, err := db.Query("select * from userinfo")
	checkErr(err)

	for rows.Next() {
		var uid int
		var username string
		var department string
		var created string
		err := rows.Scan(&uid, &username, &department, &created)
		checkErr(err)
		fmt.Println(uid, username, department, created)
	}

//	删除数据
	stmt, err = db.Prepare("delete from userinfo where uid=$1")
	checkErr(err)

	res, err = stmt.Exec(1)
	checkErr(err)

	affected, err = res.RowsAffected()
	checkErr(err)
	fmt.Println(affected)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

