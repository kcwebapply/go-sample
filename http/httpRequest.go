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
	req.Header.Set("Accept", "application/json")

	client := new(http.Client)
	resp, _ := client.Do(req)

	byteArray, _ := ioutil.ReadAll(resp.Body)
	js, _ := simplejson.NewJson(byteArray)

	// get loop size
	var array, _ = js.Get("results").Get("shop").Array()
	arraysize := len(array)

	for i := 0; i < arraysize; i++ {
		fmt.Println(fmt.Sprintf("店名%d,%s", i, js.Get("results").Get("shop").GetIndex(i).Get("name_kana")))
	}
	return "ok"
}
