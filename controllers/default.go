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
		// Id   int    `json:"id"`
	}
	req := httplib.Get("https://api.thecatapi.com/v1/categories")
	str, err := req.String()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(str, "\n---")
	var categories []Categories
	json.Unmarshal([]byte(str), &categories)

	// var listCategories = []string

	// for x := 0; x < len(categories); x++ {
	// 	listCategories = append(listCategories,categories[x].Name)
	// 	// fmt.Println("The value for x is", )
	// }
	// fmt.Println(listCategories)

	this.Data["Website"] = "beego.me"
	this.Data["Email"] = "muhammadtahsin@gmail.com"
	this.Data["Fruits"] = []string{"Apple", "Banana", "Carrot", "Date", "Eggplant", "Grape"}
	this.Data["Categories"] = categories
	this.TplName = "index.tpl"
}
