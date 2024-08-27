package mqt

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"time"
)

var messageRecHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("Clenit Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
}

// 其实consumer既可以收到消息，也可以发送消息
// 作为互联网硬件收集器，采集的环境信息数据（温度、湿度等）发送到broker
// 作为互联网硬件执行器，可以接受broker的消息（执行指令信息，如显示文字、声音等），并根据消息执行硬件行为

func ConsumerPoint() {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%d", broker, port))
	opts.SetClientID("go_mqtt_consumer")
	opts.SetUsername(userName)
	opts.SetPassword(passwd)
	opts.SetKeepAlive(8 * time.Second)
	opts.SetDefaultPublishHandler(messageRecHandler)
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	sub(client, false)
	time.Sleep(30 * time.Second)
	client.Disconnect(250)
}
