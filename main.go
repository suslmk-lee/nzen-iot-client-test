package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"nzen-iot-client-test/common"
)

var MQTTBroker string
var MQTTClientID string
var MQTTTopic string

func init() {
	MQTTBroker = common.ConfInfo["mqtt.broker.url"]
	MQTTClientID = common.ConfInfo["mqtt.producer.client.id"]
	MQTTTopic = common.ConfInfo["mqtt.topic"]
}

// AccelerometerData 구조체 정의
type AccelerometerData struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
	Z float64 `json:"z"`
}

func main() {
	// MQTT 클라이언트 옵션 설정
	opts := mqtt.NewClientOptions().AddBroker(MQTTBroker).SetClientID(MQTTClientID)

	// MQTT 클라이언트 생성 및 연결
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatalf("Error connecting to MQTT broker: %v", token.Error())
	}

	// 주기적으로 가속도 데이터를 생성하여 퍼블리시
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		// 가속도 데이터 생성
		data := AccelerometerData{
			X: rand.Float64() * 10,
			Y: rand.Float64() * 10,
			Z: rand.Float64() * 10,
		}

		// JSON 직렬화
		payload, err := json.Marshal(data)
		if err != nil {
			log.Printf("Error marshalling JSON: %v", err)
			continue
		}

		// MQTT 토픽에 퍼블리시
		token := client.Publish(MQTTTopic, 0, false, payload)
		token.Wait()
		if token.Error() != nil {
			log.Printf("Error publishing message: %v", token.Error())
		} else {
			fmt.Printf("Published accelerometer data: X=%.2f, Y=%.2f, Z=%.2f\n", data.X, data.Y, data.Z)
		}
	}

	// 프로그램 종료를 위한 신호 처리 설정
	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, syscall.SIGINT, syscall.SIGTERM)
	<-sigc

	// MQTT 클라이언트 종료
	client.Disconnect(250)
	log.Println("Program terminated")
}
