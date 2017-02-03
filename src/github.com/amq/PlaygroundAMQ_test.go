package amq

import (
	"testing"

)

func TestPlaygroundAMQ(t *testing.T) {


	subscribed := make(chan bool)
	go recvMessages(subscribed)

	<-subscribed

	go sendMessages()

	<-stop
	<-stop

}


