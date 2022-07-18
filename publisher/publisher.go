package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"testing"

	"github.com/nats-io/stan.go"
	"github.com/vladjong/L0/internal/app/model"
)

const (
	clusterID = "prod"
	clientID  = "simple-pub"
)

func createUniqueModel(o *model.Order, i int) {
	o.TrackNumber += strconv.Itoa(i)
	o.OrderId += strconv.Itoa(i)
	o.Payment.Transaction = o.OrderId
	o.Items[0].TrackNumber = o.TrackNumber
	o.Items[0].Rid = strconv.Itoa(i)
}

func main() {
	log.Println("Connecting publisher")
	sc, err := stan.Connect(clusterID, clientID)
	if err != nil {
		log.Fatal(err)
	}
	defer sc.Close()
	log.Println("Enter the number of model additions:")
	var countAgr int = 1
	fmt.Fscan(os.Stdin, &countAgr)
	log.Println("Publishing start")
	for i := 0; i < countAgr; i++ {
		o := model.TestOrder(&testing.T{})
		if i > 0 {
			createUniqueModel(o, i)
		}
		out, err := json.MarshalIndent(o, "", "")
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}
		err = sc.Publish("test", out)
		if err != nil {
			log.Println(err)
		}
		fmt.Println(string(out))
		log.Println("Done")

	}
	log.Println("Publisher is done")
}
