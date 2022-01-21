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
	req := httplib.Get("https://api.thecatapi.com/v1/categories")
	str, err := req.String()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(str, "\n---")
	var categories []Categories
	json.Unmarshal([]byte(str), &categories)
	fmt.Println(categories)

	this.Data["Website"] = "beego.me"
	this.Data["Email"] = "muhammadtahsin@gmail.com"
	this.Data["Fruits"] = []string{"Apple", "Banana", "Carrot", "Date", "Eggplant", "Grape"}
	// this.data = get()
	this.TplName = "index.tpl"
}
