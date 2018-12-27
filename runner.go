package main

import (
	mq "github.com/kcwebapply/go-sample/mq"
)

func main() {
	//httpClient.GetTranslate("sex")
	mq.PublishMessage()
}
