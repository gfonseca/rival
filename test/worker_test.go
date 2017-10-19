package main

import (
	"fmt"
	"testing"

	"github.com/gfonseca/rival"
)

type wp2Submit struct {
	input  []rival.Worker
	output error
}

func TestWorkerPollSubmit(t *testing.T) {
	wp := rival.MakeWorkerPool(1)
	wp.Submit(rival.MakeWorker(func(raw interface{}, i int) {}), "Hello")

	if wp.Running != 1 {
		t.Error(
			"For input: ", 1,
			"expected:", 1,
			"got:", wp.Running,
		)
	}

	err := wp.Submit(rival.MakeWorker(func(raw interface{}, i int) { fmt.Println("D") }), "Hello")
	if err == nil {
		t.Error(
			"For input: ", "*Worker",
			"expected:", "error",
			"got:", err,
		)
	}
}
