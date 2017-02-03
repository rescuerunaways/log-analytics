package main

import (
	"os"
	"fmt"
	"bufio"
	"sync"
)

var (
	path = "result.txt"
	f *os.File
)

func WriteTofile(wg *sync.WaitGroup) {
	f, _ = os.OpenFile(path, os.O_RDWR | os.O_APPEND | os.O_CREATE, 0666)
	defer f.Close()

	writer := bufio.NewWriter(f)
	defer writer.Flush();

	for i := 0; i < 10; i++ {
		fmt.Fprintln(writer, "event",i)
		fmt.Fprintln(writer, "trash",i)
		fmt.Fprintln(writer, "trash",i)
		fmt.Fprintln(writer, "trash",i)
		fmt.Fprintln(writer, "trash",i)
	}
	wg.Done()
}