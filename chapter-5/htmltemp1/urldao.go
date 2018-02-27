package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

type DB struct {
	
}

type Aq struct {
	Id string
}

type URLData struct {
	Id	int
	Url	string
	Name	string
	Times	int
}

func (*DB)  add(){
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

func (*DB)  update(url URLData){
	db, err := sql.Open("sqlite3", "S:\\go2018\\src\\github.com\\hxc316\\go-web\\chapter-5\\htmltemp1\\urldb")
	checkErr(err)

	//插入数据
	stmt, err := db.Prepare("update urls set count = ? where id = ?")
	checkErr(err)

	res, err := stmt.Exec(url.Times, url.Id)
	affect, err := res.RowsAffected()

	println("更新影响数据:",affect,"行")
	checkErr(err)
}

func (*DB) Query()  (mm [2]URLData){
	db, err := sql.Open("sqlite3", "S:\\go2018\\src\\github.com\\hxc316\\go-web\\chapter-5\\htmltemp1\\urldb")
	defer db.Close()
	checkErr(err)
	//查询数据
	rows, err := db.Query("SELECT t.* FROM urls t")
	checkErr(err)

	i := 0
	for rows.Next(){
		var id int
		var url string
		var name string
		var count int
		err = rows.Scan(&id,&url,&name,&count)
		mm[i] = URLData{id,url,name,count}
		//mm[i] = Data{1,"qq","qq",1}
		//i++
		checkErr(err)
		println("url = ",url," | name = ",name," | count = ",count)
	}
	return mm
}


func main()  {
	db := &DB{}

	mm := db.Query()
	for i :=0;i<len(mm);i++{
		println("接收返回的数据 :  url = ",mm[i].Url," | name = ",mm[i].Name," | count = ",mm[i].Times)
	}

	//mm[0].Count = mm[0].Count +1
	//db.update(mm[0])

	//d := Data{1,"qq","cc",1}
	//println(d.Name)

}



func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}


