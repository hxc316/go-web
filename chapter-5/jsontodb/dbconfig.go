package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type db struct {
	DbType string  `json:"dbType"`
	Tables []table `json:"tables"`
}

type table struct {
	Name    string    `json:"name"`
	Columns []columns `json:"columns"`
}

type columns struct {
	Name    string `json:"name"`
	ColType string `json:"col_type"`
}

func loadConfig(file string) (*db, error) {
	fmt.Println("开始加载配置文件：", file)
	data, error := ioutil.ReadFile(file)
	if error != nil {
		println("加载配置文件错误， errror = ", error.Error())
		return nil, error
	}
	config := &db{}
	err := json.Unmarshal(data, config)
	if config.DbType == ""{
		config.DbType = "none"
	}
	if err == nil {
		print("加载的数据: dbType = ", config.DbType,"  | tables size = " ,len(config.Tables), " | tables : ")
		for i := 0; i < len(config.Tables); i++ {
			print("table", i, " = ", config.Tables[i].Name)
		}
	}
	return config, err

}


func main() {
	println("---------------")
	loadConfig("S:\\go2018\\src\\github.com\\hxc316\\go-web\\chapter-5\\jsontodb\\config.json")
}
