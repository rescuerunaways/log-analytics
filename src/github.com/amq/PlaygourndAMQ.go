package amq

import (
	"fmt"
	"github.com/go-stomp/stomp"
)

var serverAddr = "localhost:61613"
var messageCount = 10
var queueName = "queue/queue"
var stop = make(chan bool)

func sendMessages() {

	defer func() {
		stop <- true
	}()

	conn, _ := stomp.Dial("tcp", serverAddr)

	for i := 1; i <= messageCount; i++ {
		event := fmt.Sprintf("Message #%d", i)
		conn.Send(queueName, "text/plain",
			[]byte(event))

	}
	println("sender finished")
}

func recvMessages(subscribed chan bool) {

	defer func() {
		stop <- true
	}()

	conn, _ := stomp.Dial("tcp", serverAddr)
	sub, _ := conn.Subscribe(queueName, stomp.AckAuto)

	close(subscribed)

	for i := 1; i <= messageCount; i++ {
		msg := <-sub.C
		expectedText := fmt.Sprintf("Message #%d", i)
		actualText := string(msg.Body)

		println("Expected:", expectedText)
		println("Actual:", actualText)

		if expectedText != actualText {
			println("Expected:", expectedText)
			println("Actual:", actualText)
		}
	}
	println("receiver finished")

}






