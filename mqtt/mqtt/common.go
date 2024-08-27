package mqt

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

// https://mqttx.app/zh/docs/get-started  以下使用emqx官方公共mqtt服务器测试
var broker = "broker.emqx.io"
var port = 1883
var userName = "emqx"
var passwd = "123456"
var topic = "topic/zty"

func sub(client mqtt.Client, producer bool) {
	token := client.Subscribe(topic, 1, nil)
	token.Wait()
	if producer {
		fmt.Printf("Producer subscribed to topic %s\n", topic)
	} else {
		fmt.Printf("Consumer subscribed to topic %s\n", topic)
	}
}
