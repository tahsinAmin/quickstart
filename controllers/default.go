package controllers

import (
	"encoding/json"
	"fmt"

	"github.com/beego/beego/v2/client/httplib"
	beego "github.com/beego/beego/v2/server/web"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {

	type Categories struct {
		Name string `json:"name"`
		Id   int    `json:"id"`
	}

	type Breeds struct {
		Name string `json:"name"`
		Id   string `json:"id"`
	}

	req := httplib.Get("https://api.thecatapi.com/v1/categories")
	str, err := req.String()
	if err != nil {
		fmt.Println(err.Error())
	}

	var categories []Categories
	json.Unmarshal([]byte(str), &categories)

	req = httplib.Get("https://api.thecatapi.com/v1/breeds")
	str, err = req.String()
	if err != nil {
		fmt.Println(err.Error())
	}
	// fmt.Println(str, "\n---")

	var breeds []Breeds
	json.Unmarshal([]byte(str), &breeds)

	// for i := 0; i < len(breeds); i++ {
	// 	fmt.Println(breeds[i])
	// }

	this.Data["Website"] = "beego.me"
	this.Data["Email"] = "muhammadtahsin@gmail.com"
	this.Data["Fruits"] = []string{"Apple", "Banana", "Carrot", "Date", "Eggplant", "Grape"}
	this.Data["Categories"] = categories
	this.Data["Breeds"] = breeds
	this.Data["Title"] = ""
	this.TplName = "index.html"
}

func (this *MainController) Post() {
	// order := this.GetString("order")
	// type := this.GetString("type")
	title := this.GetString("title")
	fmt.Println(title)
	this.Data["Title"] = title
	this.TplName = "index.html"
}
