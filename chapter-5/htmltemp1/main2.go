package main

import (
	"html/template"
	"log"
	"net/http"
	//"strconv"
	"time"
	"github.com/gorilla/mux"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"fmt"
)

type Note struct {
	Title       string
	Description string
	CreatedOn   time.Time
}

type URLData1 struct {
	Id	int
	Url	string
	Name	string
	Times	int
}

type URLDataList struct {
	Id	int
	URLData
}
//View Model for edit
type EditNote struct {
	Note
	Id string
	//data []URLData
}

//Store for the Notes collection

var urlsData = make([]URLData,10)

var templates map[string]*template.Template

//Compile view templates S:\workspace2017_IJ\xc\go-web\chapter-5\htmltemp\templates\add.html
func init() {
	if templates == nil {
		templates = make(map[string]*template.Template)
	}
	path := "chapter-5/htmltemp1/"
	templates["index"] = template.Must(template.ParseFiles(path+"templates/index.html", path+"templates/base.html"))
	templates["add"] = template.Must(template.ParseFiles(path+"templates/add.html", path+"templates/base.html"))
	templates["edit"] = template.Must(template.ParseFiles(path+"templates/edit.html", path+"templates/base.html"))
	urlsData[0] = URLData{Id:1,Url:"www.baidu.com",Name:"百度",Times:0}
	urlsData[1] = URLData{Id:2,Url:"www.jianshu.com",Name:"简书",Times:0}
}

//Render templates for the given name, template definition and data object
func renderTemplate1(w http.ResponseWriter, name string, template string, viewModel interface{}) {
	// Ensure the template exists in the map.
	tmpl, ok := templates[name]
	if !ok {
		http.Error(w, "The template does not exist.", http.StatusInternalServerError)
	}
	//var ss = []string{"11","22","33"}
	//d1 := URLData{Id:1,Url:"www.baidu.com",Name:"百度",Times:0}
	err := tmpl.ExecuteTemplate(w, template, viewModel)

	//println("name=",viewModel[0]["name"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

//Handler for "/" which render the index page
func getNotes1(w http.ResponseWriter, r *http.Request) {
	db := &DB1{}
	urlsData1 := db.Query1()
	renderTemplate1(w, "index", "base", urlsData1)
}

func viewUrl(w http.ResponseWriter, r *http.Request) {
	//Read value from route variable
	vars := mux.Vars(r)
	k := vars["url"]
	// Remove from Store
	db := &DB1{}
	urlsDatas1 := db.Query1()
	for i := 0;i<len(urlsDatas1);i++ {
		if len(urlsDatas1[i].Name) > 0 && urlsDatas1[i].Url == k {
			println("name=",urlsDatas1[i].Name,"url=",urlsDatas1[i].Url)
			urlsDatas1[i].Times++
			db := &DB1{}
			db.update1(urlsDatas1[i])
			println(urlsDatas1[i].Url,"访问次数增加到:",urlsDatas1[i].Times)
		}

	}
	println("url = ",k)
	http.Redirect(w, r, "http://"+k, 302)
}


type DB1 struct {

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

func (*DB1)  add1(){
	db, err := sql.Open("sqlite3", "S:\\go2018\\src\\github.com\\hxc316\\go-web\\chapter-5\\htmltemp1\\urldb")
	checkErr1(err)

	//插入数据
	stmt, err := db.Prepare("INSERT INTO urls(url, name,count) values(?,?,?)")
	checkErr1(err)

	res, err := stmt.Exec("aa", "ss", 1)
	checkErr1(err)
	id, err := res.LastInsertId()
	fmt.Println("id=",id)

	checkErr1(err)
}

func (*DB1)  update1(url URLData1){
	db, err := sql.Open("sqlite3", "S:\\go2018\\src\\github.com\\hxc316\\go-web\\chapter-5\\htmltemp1\\urldb")
	checkErr1(err)

	//插入数据
	stmt, err := db.Prepare("update urls set count = ? where id = ?")
	checkErr1(err)

	res, err := stmt.Exec(url.Times, url.Id)
	affect, err := res.RowsAffected()

	println("更新影响数据:",affect,"行")
	checkErr1(err)
}

func (*DB1) Query1()  (mm [2]URLData1){
	db, err := sql.Open("sqlite3", "S:\\go2018\\src\\github.com\\hxc316\\go-web\\chapter-5\\htmltemp1\\urldb")
	//defer db.Close()
	checkErr1(err)
	//查询数据
	rows, err := db.Query("SELECT t.* FROM urls t")
	checkErr1(err)

	i := 0
	for rows.Next(){
		var id int
		var url string
		var name string
		var count int
		err = rows.Scan(&id,&url,&name,&count)
		mm[i] = URLData1{id,url,name,count}
		//mm[i] = Data{1,"qq","qq",1}
		//i++
		checkErr1(err)
		println("url = ",url," | name = ",name," | count = ",count)
	}
	return mm
}
//Entry point of the program
func main() {

	r := mux.NewRouter().StrictSlash(false)
	fs := http.FileServer(http.Dir("public"))
	r.Handle("/public/", fs)
	r.HandleFunc("/", getNotes1)
	r.HandleFunc("/view/{url}", viewUrl)

	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}
	log.Println("Listening...")
	server.ListenAndServe()
}

func checkErr1(err error) {
	if err != nil {
		panic(err)
	}
}
