package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

func main()  {
	db, err := sql.Open("sqlite3", "S:\\go2018\\src\\github.com\\hxc316\\go-web\\chapter-5\\htmltemp1\\urldb")
	checkErr(err)

	//插入数据
	stmt, err := db.Prepare("INSERT INTO urls(url, name,count) values(?,?,?)")
	checkErr(err)

	res, err := stmt.Exec("aa", "ss", 1)
	checkErr(err)
	id, err := res.LastInsertId()
	fmt.Println("id=",id)

	checkErr(err)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}


