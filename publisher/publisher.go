package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/nats-io/stan.go"
	"github.com/vladjong/L0/internal/app/model"
)

const (
	clusterID = "prod"
	clientID  = "simple-pub"
)

func main() {
	fmt.Println("Connecting")
	sc, err := stan.Connect(clusterID, clientID)
	if err != nil {
		log.Fatal(err)
	}
	defer sc.Close()
	fmt.Println("Publishing")
	o := model.TestOrder(&testing.T{})

	out, err := json.MarshalIndent(o, "", "")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	//for {
	err = sc.Publish("test", out)
	if err != nil {
		log.Println(err)
		//break
	}
	fmt.Println(string(out))
	fmt.Println("has been send")
	//time.Sleep(20 * time.Second)
	//}

}
