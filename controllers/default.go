package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/beego/beego/v2/client/httplib"
	beego "github.com/beego/beego/v2/server/web"
)

type Cats struct {
	Breeds     []string `json:"breeds"`
	Categories []string `json:"categories"`
	Id         string   `json:"id"`
	Url        string   `json:"url"`
}

type Categories struct {
	Name string `json:"name"`
	Id   int    `json:"id"`
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

func manipulate(w http.ResponseWriter, r *http.Request) {
	// this.TplName = "index.html"
	fmt.Println("method:", r.Method)
	if r.Method == "GET" {
		r.ParseForm()
		fmt.Println(r.Form)
	}
}

func (this *MainController) Manipulate() {
	// this.TplName = "index.html"
	// this.Ctx.Request.ParseForm()

	// categories := this.GetString("category")
	// fmt.Println("Hello", categories)

}

func (c *MainController) AjaxTest() {

	c.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Origin", "*")
	c.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
	// order := c.GetString("order")
	// types := c.GetString("types")
	// breeds := c.GetString("breeds")
	category_id := c.GetString("categories")

	// // fmt.Println("Order:", order, "\nType:", types, "\nCategories:", categories, "\nBreeds:", breeds)

	req := httplib.Get("https://api.thecatapi.com/v1/images/search?limit=10&category_ids=" + category_id)
	str, err := req.String()
	if err != nil {
		fmt.Println(err.Error())
	}

	var cats []Cats
	json.Unmarshal([]byte(str), &cats)

	// // req.Header.Add("Content-Type", "application/json")
	// // req.Header.Add("x-api-key", "176a2a9c-6fc8-4d74-af01-4b07c0034e5e")

	// fmt.Println("Hello", category_id)
	// fmt.Println("\n", cats)
	var s string

	for _, cat := range cats {
		s += fmt.Sprintf(`<div class="bg-cover bg-center h-80 w-80 rounded-lg" style="background-image: url(%s)"></div>`, cat.Url)

		// `<img class="h-80 w-80 rounded-lg" src="%s">`
	}

	c.Data["json"] = map[string]interface{}{"name": s}
	c.ServeJSON()
}
