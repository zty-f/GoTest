package main

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type KafkaConfig struct {
	Topic            string `json:"topic"`
	GroupId          string `json:"group.id"`
	BootstrapServers string `json:"bootstrap.servers"`
	SecurityProtocol string `json:"security.protocol"`
	SaslMechanism    string `json:"sasl.mechanism"`
	SaslUsername     string `json:"sasl.username"`
	SaslPassword     string `json:"sasl.password"`
}

// config should be a pointer to structure, if not, panic
func loadJsonConfig() *KafkaConfig {
	var config = &KafkaConfig{
		Topic:            "mall_growth_center_task_test",
		GroupId:          "mall-task-center-task-test-group",
		BootstrapServers: "alikafka-pre-cn-pe33gz17l001-1.alikafka.aliyuncs.com:9093,alikafka-pre-cn-pe33gz17l001-2.alikafka.aliyuncs.com:9093,alikafka-pre-cn-pe33gz17l001-3.alikafka.aliyuncs.com:9093",
		SecurityProtocol: "sasl_ssl",
		SaslMechanism:    "PLAIN",
		SaslUsername:     "alikafka_pre-cn-pe33gz17l001",
		SaslPassword:     "3WiuaCTERiIcCM490sDIpghxxzXw83hG",
	}
	return config
}

func doInitConsumer(cfg *KafkaConfig) *kafka.Consumer {
	fmt.Print("init kafka consumer, it may take a few seconds to init the connection\n")
	//common arguments
	var kafkaconf = &kafka.ConfigMap{
		"api.version.request":       "true",
		"auto.offset.reset":         "latest",
		"heartbeat.interval.ms":     3000,
		"session.timeout.ms":        30000,
		"max.poll.interval.ms":      120000,
		"fetch.max.bytes":           1024000,
		"max.partition.fetch.bytes": 256000}
	kafkaconf.SetKey("bootstrap.servers", cfg.BootstrapServers)
	kafkaconf.SetKey("group.id", cfg.GroupId)

	switch cfg.SecurityProtocol {
	case "plaintext":
		kafkaconf.SetKey("security.protocol", "plaintext")
	case "sasl_ssl":
		kafkaconf.SetKey("security.protocol", "sasl_ssl")
		kafkaconf.SetKey("ssl.ca.location", "/Users/xwx/go/src/test/test10/conf/ca-cert")
		kafkaconf.SetKey("sasl.username", cfg.SaslUsername)
		kafkaconf.SetKey("sasl.password", cfg.SaslPassword)
		kafkaconf.SetKey("sasl.mechanism", cfg.SaslMechanism)
	case "sasl_plaintext":
		kafkaconf.SetKey("security.protocol", "sasl_plaintext")
		kafkaconf.SetKey("sasl.username", cfg.SaslUsername)
		kafkaconf.SetKey("sasl.password", cfg.SaslPassword)
		kafkaconf.SetKey("sasl.mechanism", cfg.SaslMechanism)

	default:
		panic(kafka.NewError(kafka.ErrUnknownProtocol, "unknown protocol", true))
	}

	consumer, err := kafka.NewConsumer(kafkaconf)
	if err != nil {
		panic(err)
	}
	fmt.Print("init kafka consumer success\n")
	return consumer
}

func main() {

	// Choose the correct protocol
	// 9092 for PLAINTEXT
	// 9093 for SASL_SSL, need to provide sasl.username and sasl.password
	// 9094 for SASL_PLAINTEXT, need to provide sasl.username and sasl.password
	cfg := loadJsonConfig()
	consumer := doInitConsumer(cfg)

	consumer.SubscribeTopics([]string{cfg.Topic}, nil)

	for {
		msg, err := consumer.ReadMessage(-1)
		if err == nil {
			fmt.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
		} else {
			// The client will automatically try to recover from all errors.
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
		}
	}

	consumer.Close()
}
