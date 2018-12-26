package http

import (
	"fmt"
	"io/ioutil"
	"net/http"

	simplejson "github.com/bitly/go-simplejson"
	. "github.com/kcwebapply/go-sample/config"
)

func GetTranslate(word string) string {
	var config Config = GetConfig()
	url := config.Http.HOST + "?key=" + config.Http.APIKey + "&large_area=Z011&format=json"
	req, _ := http.NewRequest("GET", url, nil)
	client := new(http.Client)
	fmt.Println("http request ")
	resp, _ := client.Do(req)
	fmt.Println(resp)
	byteArray, _ := ioutil.ReadAll(resp.Body)
	js, _ := simplejson.NewJson(byteArray)
	fmt.Println("responseBody:", js.Get("results").Get("shop").GetIndex(0).Get("name_kana"))
	return "ok"
}
