package main

import (
	"testing"
	"sync"
)

func TestWriteTofile(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1)

	go WriteTofile(&wg)
	//go WriteTofile(&wg)
	//go WriteTofile(&wg)
	//go WriteTofile(&wg)

	wg.Wait()
}
