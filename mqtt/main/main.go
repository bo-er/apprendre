package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func main() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	topic1 := "a/b/c"
	// topic2 := "sys/9c565ade-cf03-4432-bf4f-6c9095bac396/f51f6816-2db5-4e95-9e55-fea7d8908910/heartbeat"
	go testPublishEvery10Seconds("sws1", topic1)

	// go testPublishEvery10Seconds("sws1", topic2)
	<-sigs
}

func testPublishEvery10Seconds(clientID, topic string) {
	clientOptions := mqtt.NewClientOptions()
	clientOptions.AddBroker("tcp://192.168.1.118:1883")
	clientOptions.SetClientID(clientID)
	// clientOptions.SetUsername("admin")
	// clientOptions.SetPassword("ilovesws")
	clientOptions.SetProtocolVersion(4)
	clientOptions.SetKeepAlive(time.Duration(30) * time.Second)
	clientOptions.SetCleanSession(true)
	var f mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
		fmt.Printf("TOPIC: %s\n", msg.Topic())
		fmt.Printf("MSG: %s\n", msg.Payload())
	}

	clientOptions.SetDefaultPublishHandler(f)

	c := mqtt.NewClient(clientOptions)

	if token := c.Connect(); token.Wait() && token.Error() != nil {

		log.Fatalf("Error on Client.Connect(): %v", token.Error())
	}

	// topic := "sys/79271c53-0841-4f6f-b28c-8ec7bd18ad8a/00ba7208-b38a-4956-88b1-cf2c41bf8e30/heartbeat"

	// payload := bytes.NewBufferString("hello world")
	// count := 0
	// for {
	// 	if pubToken := c.Publish(topic, 2, false, strconv.Itoa(count)); pubToken.Wait() && pubToken.Error() != nil {
	// 		log.Fatalf("Error on Publish message: %v", pubToken.Error())
	// 	}
	// 	count++
	// 	time.Sleep(200 * time.Second)
	// }

	//c.Disconnect(250)
	// if pubToken := c.Publish("sys/093003a4-f58a-40c1-a0ff-0b8891152cdf/368f6200-f87e-4cce-8924-a915f28788e2/heartbeat", 0, false, "i am back!"); pubToken.Wait() && pubToken.Error() != nil {
	// 	log.Fatalf("Error on Publish message: %v", pubToken.Error())
	// }
}
