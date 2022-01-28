package controllers

import (
	"encoding/json"
	"fmt"
	"sort"

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

var cats []Cats

const baseUrl = "https://api.thecatapi.com/v1/images/search?"

var s string

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

	req = httplib.Get(baseUrl + "limit=9&page=1")
	str, err = req.String()
	if err != nil {
		fmt.Println(err.Error())
	}

	json.Unmarshal([]byte(str), &cats)

	this.Data["Categories"] = categories
	this.Data["Breeds"] = breeds
	this.Data["Cats"] = cats
	this.TplName = "index.html"
}

func (this *MainController) AjaxTest() {
	this.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Origin", "*")
	this.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")

	types := this.GetString("types")
	category_id := this.GetString("categories")
	breed_id := this.GetString("breed_id")
	limit := this.GetString("limit")

	fmt.Println("Type:", types, "\nCategories:", category_id, "\nBreeds:", breed_id)

	urlConcat := "limit=" + limit + "&mime_types=" + types + "&order="

	if category_id != "" {
		urlConcat += "&category_ids=" + category_id
	}
	if breed_id != "" {
		urlConcat += "&breed_id=" + breed_id
	}

	url := baseUrl + urlConcat
	fmt.Println(url)

	req := httplib.Get(url)

	str, err := req.String()
	if err != nil {
		fmt.Println(err.Error())
	}

	json.Unmarshal([]byte(str), &cats)
	fmt.Println(cats)

	s = ""

	if len(cats) > 0 {
		for _, cat := range cats {
			s += fmt.Sprintf(`<img loading="lazy" alt="Loading..." class="bg-gray-200 bg-cover bg-center rounded-lg h-64 sm:h-80" src="%s" data-src="%s"/>`, cat.Url, cat.Url)
		}
	} else {
		s += `<div></div><div class="text-center text-black font-semibold text-2xl">No Image(s) Found</div><div></div>`
	}

	this.Data["json"] = map[string]interface{}{"name": s}
	this.ServeJSON()
	// // req.Header.Add("Content-Type", "application/json")
	// // req.Header.Add("x-api-key", "176a2a9c-6fc8-4d74-af01-4b07c0034e5e")
}

func (this *MainController) SortData() {
	this.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Origin", "*")
	this.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
	order := this.GetString("order")

	if order == "DESC" {
		sort.Slice(cats, func(i, j int) bool {
			return cats[i].Id > cats[j].Id
		})
	}
	if order == "ASC" {
		sort.Slice(cats, func(i, j int) bool {
			return cats[i].Id < cats[j].Id
		})
	}
	s = ""
	for _, cat := range cats {
		s += fmt.Sprintf(`<img loading="lazy" alt="Loading..." class="bg-gray-200 bg-cover bg-center rounded-lg h-64 sm:h-80" src="%s" data-src="%s"/>`, cat.Url, cat.Url)
	}
	this.Data["json"] = map[string]interface{}{"name": s}
	this.ServeJSON()
}

func (this *MainController) NextPage() {
	this.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Origin", "*")
	this.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")

	types := this.GetString("types")
	category_id := this.GetString("categories")
	breed_id := this.GetString("breed_id")
	limit := this.GetString("limit")
	pageNum := this.GetString("page")

	fmt.Println("Type:", types, "\nCategories:", category_id, "\nBreeds:", breed_id, "\nPage Number:", pageNum)

	urlConcat := "limit=" + limit + "&mime_types=" + types + "&page=" + pageNum

	if category_id != "" {
		urlConcat += "&category_ids=" + category_id
	}
	if breed_id != "" {
		urlConcat += "&breed_id=" + breed_id
	}

	url := baseUrl + urlConcat
	fmt.Println(url)

	req := httplib.Get(url)

	str, err := req.String()
	if err != nil {
		fmt.Println(err.Error())
	}

	json.Unmarshal([]byte(str), &cats)

	s = ""

	if len(cats) > 0 {
		for _, cat := range cats {
			s += fmt.Sprintf(`<img loading="lazy" alt="Loading..." class="bg-gray-200 bg-cover bg-center rounded-lg h-64 sm:h-80" src="%s" data-src="%s"/>`, cat.Url, cat.Url)
		}
	} else {
		s += `<div></div><div class="text-center text-black font-semibold text-2xl">No Image(s) Found</div><div></div>`
	}

	this.Data["json"] = map[string]interface{}{"name": s}
	this.ServeJSON()
	// // req.Header.Add("Content-Type", "application/json")
	// // req.Header.Add("x-api-key", "176a2a9c-6fc8-4d74-af01-4b07c0034e5e")
}
