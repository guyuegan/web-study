package db_test

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"testing"
	"time"
)

func Test(t *testing.T)  {
	/*
	user@unix(/path/to/socket)/dbname?charset=utf8
	user:password@tcp(localhost:5555)/dbname?charset=utf8
	user:password@/dbname
	user:password@tcp([de:ad:be:ef::ca:fe]:80)/dbname
	 */
	db, err := sql.Open("mysql", "root:123456@tcp(localhost:3306)/go?charset=utf8")
	checkErr(err)

	defer db.Close()

// 插入数据
	stmt, err := db.Prepare("insert userinfo set username=?,department=?,created=?")
	checkErr(err)
	res, err := stmt.Exec("gyg", "baas-java", "2020-09-09")
	checkErr(err)

	lastId, err := res.LastInsertId()
	checkErr(err)
	fmt.Println(lastId)

//	更新数据
	stmt, err = db.Prepare("update userinfo set username=? where uid=?")
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
		var created time.Time
		err := rows.Scan(&uid, &username, &department, &created)
		checkErr(err)
		fmt.Println(uid, username, department, created)
	}

//	删除数据
//	stmt, err = db.Prepare("delete from userinfo where uid=?")
//	CheckErr(err)
//
//	res, err = stmt.Exec(lastId)
//	CheckErr(err)
//
//	affected, err = res.RowsAffected()
//	CheckErr(err)
//	fmt.Println(affected)
}

