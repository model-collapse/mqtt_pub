package main

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"flag"
	"net/url"
	"log"
	"fmt"
)

var clientName = flag.String("name", "ws-client", "name of client")
var host = flag.String("host", "hivemq", "path to host")
var port = flag.String("port", "1883", "port")
var protocol = flag.String("protocol", "tcp", "protocol")
var topic = flag.String("topic", "test", "topic")
var message = flag.String("message", "haha", "message")
var path = flag.String("path", "mqtt", "path of ws")

func toURL(u string) (ret *url.URL) {
	ret, err := url.Parse(u)
	if err != nil {
		log.Fatalf("Error in parsing url, %s, %v", u, err)
	}

	return
}

func main() {
	flag.Parse()
	urlp := fmt.Sprintf("%s://%s:%s", *protocol, *host, *port)
	if *protocol == "ws" {
		urlp = fmt.Sprintf("%s://%s:%s/%s", *protocol, *host, *port, *path)
	}

	cl := mqtt.NewClient(&mqtt.ClientOptions{
		Servers: []*url.URL{toURL(urlp)},
		ClientID: *clientName,
	})

	if token := cl.Connect(); token.Wait() && token.Error() != nil {
		log.Fatal(token.Error())
	}

	log.Printf("Connected")
	token := cl.Publish(*topic, 0, false, *message)
	token.Wait()
	log.Printf("Command published!")
}