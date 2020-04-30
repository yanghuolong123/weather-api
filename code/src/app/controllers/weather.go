package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"time"
//	"fmt"
	"app/constant"
	"net/http"
	"io/ioutil"
)

// operation about weather
type WeatherController struct {
	beego.Controller
}

// @Title Get
// @Description get weather by date
// @Param       date	path    string  false "The key for date"
// @Success 200 {string} date format 
// @Failure 403 :date is  error
// @router /:date [get]
// @router / [get]
func (w *WeatherController) Get() {
	d := w.GetString(":date")
	if d == "" {
		d = time.Now().Format(constant.DateFormat)
	}
	city := w.GetString("city")
	if city == "" {
		city = "西安"
	}

	weatherUrl := "https://www.toutiao.com/stream/widget/local_weather/data/?city="+city
	res,err :=http.Get(weatherUrl)
	if err != nil {
		w.Data["json"] = "error"
	}
	defer res.Body.Close()
	body,err := ioutil.ReadAll(res.Body)
	weather := make(map[string]interface{})//string(body)

	json.Unmarshal(body,&weather)

	m := map[string]interface{}{
		"date":d,
		"city":city,
		"weather":weather,
	}

	//fmt.Println(m)
	w.Data["json"]= m
	w.ServeJSON()
}
