package main

import (
	"SmartEyes/Library"
	"fmt"
	"reflect"

	//"bytes"
	//"encoding/binary"
	//"fmt"
	"sync"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var Json []byte

func ModbusRead(wg *sync.WaitGroup, mbc Library.MBClient) interface{} {

	var add uint16
	var leng uint16
	leng = 10
	add = 0
	for {
		Regdata, err := mbc.ReadHoldReg(1, 0, 127)
		if err != nil {
			fmt.Print(err)
		}
		Library.Parser(Regdata)

		fmt.Println(reflect.TypeOf(Regdata))
		println("Modbus Read")
		Json = Library.JsonMaker(add, leng, b)
		time.Sleep(time.Second)

	}
	return Json
}

func MQTTPublish(wg *sync.WaitGroup, mbc Library.MBClient, client mqtt.Client, topic string) {

	for {

		Library.MQTTPublish(client, topic, Json)
		println("data Publish")
		time.Sleep(time.Second)
	}
}

func main() {
	var wg sync.WaitGroup

	//Modbus Client Creat and TCP access
	mbc := Library.NewClient("192.168.0.77", 502)
	mbc.Open()

	//MQTT Client Creat and MQTT Broker access
	//uri := "tcp://broker.hivemq.com:1883"
	//topic := "test/topic12/1"
	//client := Library.Connect("enitt", uri)

	println("@@SmartEyes Start@@")
	wg.Add(1)
	go ModbusRead(&wg, *mbc)

	//println("@@MQTTPublish Start@@")
	//wg.Add(2)
	//go MQTTPublish(&wg, *mbc, client, topic)

	wg.Wait()
}
