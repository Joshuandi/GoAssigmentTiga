package handler

import (
	"GoAssigmentTiga/model"
	"encoding/json"
	"io/ioutil"

	"net/http"
	"text/template"
)

const HtmlPath = "web/web.html"
const JsonPath = "web/weather.json"

type AutoReloadInterface interface {
	AutoReload(w http.ResponseWriter, r *http.Request)
}

type AutoReload struct {
	status *model.Web
}

func NewAutoReload(status *model.Web) AutoReloadInterface {
	return &AutoReload{status: status}
}

func (a *AutoReload) AutoReload(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/html")
	// a.status.StatusWater = "aman"
	// a.status.StatusWind = "aman"
	// a.status.Water = rand.Intn(100)
	// a.status.Wind = rand.Intn(100)
	// if a.status.Water < 5 || a.status.Wind < 6 {
	// 	a.status.StatusWater = "aman"
	// 	a.status.StatusWind = "aman"
	// }
	// if a.status.Water >= 6 && a.status.Water <= 8 || a.status.Wind >= 7 && a.status.Wind <= 15 {
	// 	a.status.StatusWater = "Siaga"
	// 	a.status.StatusWind = "Siaga"
	// }
	// if a.status.Water > 8 || a.status.Wind > 15 {
	// 	a.status.StatusWater = "Siaga"
	// 	a.status.StatusWind = "Siaga"
	// }

	file, _ := ioutil.ReadFile(JsonPath)
	json.Unmarshal(file, a.status)
	templates, _ := template.ParseFiles(HtmlPath)
	context := model.Web{
		StatusWater: a.status.StatusWater,
		StatusWind:  a.status.StatusWind,
		Water:       a.status.Water,
		Wind:        a.status.Wind,
	}
	templates.Execute(w, context)

}
