package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type new_con struct {
	sql_type string
	sql_addr string
}

func (c *new_con) newDb(sql_type string, sql_addr string) *sql.DB {
	c.sql_type, c.sql_addr = sql_type, sql_addr
	db, err := sql.Open(c.sql_type, c.sql_addr)
	checkErr(err)
	return db
}

type new_ins struct {
	con        *new_con
	operate    string
	table_name string
	para       []string
}

type para struct{
	username map
}

func (i *new_ins) newIns(db *sql.DB,operate string,table_name string ,para []string) {
	stmt, err := db.Prepare("INSERT INTO userinfo SET username=?,department=?,created=?")
	checkErr(err)
	res, err := stmt.Exec("lee", "development", nowTime())
	checkErr(err)
}


func main() {

	db := &new_con{}
	d := db.newDb("mysql", "lee:lee1@tcp(192.168.0.114:3306)/test?charset=utf8")
	stmt, err := d.Prepare("INSERT INTO userinfo SET username=?,department=?,created=?")
	checkErr(err)

	res, err := stmt.Exec("lee", "development", nowTime())
	checkErr(err)
	id, err := res.LastInsertId()
	checkErr(err)
	fmt.Println(id)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func nowTime() string {
	timeUnix := time.Now().Unix() //已知的时间戳
	formatTimeStr := time.Unix(timeUnix, 0).Format("2006-01-02")
	fmt.Println(formatTimeStr) //打印结果：2017-04-11 13:30:39
	return formatTimeStr
}
