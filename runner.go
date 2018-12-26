package main

import (
	"fmt"

	httpClient "github.com/kcwebapply/go-sample/http"
)

func main() {
	fmt.Println("test")
	httpClient.GetTranslate("sex")
}
