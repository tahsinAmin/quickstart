package controllers

import (
	"encoding/json"
	"fmt"

	"github.com/beego/beego/v2/client/httplib"
	beego "github.com/beego/beego/v2/server/web"
)

type Categories struct {
	Name       string `json:"name"`
	Id         int    `json:"id"`
	IsSelected bool
}

type Breeds struct {
	Name string `json:"name"`
	Id   string `json:"id"`
}

type MainController struct {
	beego.Controller
}

func (this *MainController) Show() {

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

	this.Data["Website"] = "beego.me"
	this.Data["Email"] = "muhammadtahsin@gmail.com"
	this.Data["Fruits"] = []string{"Apple", "Banana", "Carrot", "Date", "Eggplant", "Grape"}
	this.Data["Categories"] = categories
	this.Data["Breeds"] = breeds
	this.Data["Title"] = ""
	this.TplName = "index.html"
}

func (this *MainController) DoShow() {
	this.TplName = "index.html"
	this.Ctx.Request.ParseForm()
	title := this.GetString("title")
	catagory_id, err := this.GetInt("catagory_id")

	if err != nil {
		this.Ctx.WriteString("get param id fail")
		return
	}
	fmt.Println(title, catagory_id)
	this.Data["Title"] = title
	return
}
