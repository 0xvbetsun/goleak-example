package main

import (
	"bytes"
	"encoding/json"
	"io"
	"os"
	"sync"
	"testing"
)

func Test_main(t *testing.T) {
	out := captureOutput(func() {
		main()
	})
	var arr []int
	err := json.Unmarshal([]byte(out), &arr)
	if err != nil {
		t.Fatal(err)
	}
	// Be careful this is a flaky test !!
	if len(arr) == 0 {
		t.Errorf("expected len(arr) > 0, got %d", len(arr))
	}
}

func captureOutput(f func()) string {
	r, w, err := os.Pipe()
	if err != nil {
		panic(err)
	}
	stdout := os.Stdout
	stderr := os.Stderr
	defer func() {
		os.Stdout = stdout
		os.Stderr = stderr
	}()
	os.Stdout = w
	os.Stderr = w
	out := make(chan string)
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func() {
		var buf bytes.Buffer
		wg.Done()
		io.Copy(&buf, r)
		out <- buf.String()
	}()
	wg.Wait()
	f()
	w.Close()
	return <-out
}
