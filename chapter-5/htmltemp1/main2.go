package main

import (
	"html/template"
	"log"
	"net/http"
	//"strconv"
	"time"
	"github.com/gorilla/mux"
)

type Note struct {
	Title       string
	Description string
	CreatedOn   time.Time
}

type URLData struct {
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

	renderTemplate1(w, "index", "base", urlsData)
}

func viewUrl(w http.ResponseWriter, r *http.Request) {
	//Read value from route variable
	vars := mux.Vars(r)
	k := vars["url"]
	// Remove from Store
	for i := 0;i<len(urlsData);i++ {
		if len(urlsData[i].Name) > 0 && urlsData[i].Url == k {
			println("name=",urlsData[i].Name,"url=",urlsData[i].Url)
			urlsData[i].Times++

			println(urlsData[i].Url,"访问次数:",urlsData[i].Times)
		}

	}
	println("url = ",k)
	http.Redirect(w, r, "http://"+k, 302)
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
