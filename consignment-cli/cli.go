package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	pb "github.com/austinsilver/ship/consignment-service/proto/consignment"
	"golang.org/x/net/context"
	microclient "github.com/micro/go-micro/client"
	"github.com/micro/go-micro/cmd"
)

const (
	defaultFilename = "consignment.json"
)

func parseFile(file string) (*pb.Consignment, error){
	var consignment *pb.Consignment
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(data, &consignment)
	return consignment, err
}

func main(){
	//Setup up a connection
	cmd.Init()

	client := pb.NewShippingServiceClient("go.micro.srv.consignment", microclient.DefaultClient)

	//Contact the server and print out the response
	file := defaultFilename
	if len(os.Args)> 1{
		file = os.Args[1]
	}

	consignment, err := parseFile(file)

	if err != nil {
		log.Fatalf("Cloud not parse file: %v", err)
	}

	r, err := client.CreateConsignment(context.TODO(), consignment)
	if err != nil {
		log.Fatalf("Could not create %v", err)
	}
	log.Printf("Created: %t", r.Created)

	getAll, err := client.GetConsignments(context.Background(), &pb.GetRequest{})
	if err != nil {
		log.Fatalf("Could not lists consignments: %v", err)
	}
	for _, v := range getAll.Consignments{
		log.Println(v)
	}
}