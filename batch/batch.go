package batch

import (
	"fmt"
	"time"

	"github.com/carlescere/scheduler"
)

func RunBatch() {

	job := func() {
		t := time.Now()
		fmt.Println("バッチ処理:", t.UTC())
	}

	scheduler.Every(1).Seconds().Run(job)
	forever := make(chan bool)
	<-forever
}
