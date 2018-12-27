package main

import (
	httpClient "github.com/kcwebapply/go-sample/http"
	mq "github.com/kcwebapply/go-sample/mq"
)

func main() {
	httpClient.GetTranslate("sex")
	mq.Work()
}
