package main

import (
	"fmt"
	"time"

	"github.com/gfonseca/rival/rival"
)

func main() {
	wp := rival.MakeWorkerPool(5)
	w := rival.MakeWorker(func(raw interface{}, index int) { fmt.Println(raw) })

	wp.Submit(w, "Hello")
	time.Sleep(3 * time.Second)
}
