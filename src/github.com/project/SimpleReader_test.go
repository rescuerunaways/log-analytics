package main

import (
	"testing"

)

func TestReadFile(t *testing.T) {

	done := make(chan bool)
	go openAndRead(done)
	<-done

//	time.Sleep(1000 * time.Millisecond)

}