package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	pb "github.com/heraldSim/grpcExample/protos"
	"io/ioutil"
	"log"
)

func main()  {
	book := &pb.AddressBook{}
	p := pb.Person{
		Id:    1234,
		Name:  "John Doe",
		Email: "jdoe@example.com",
		Phones: []*pb.Person_PhoneNumber{
			{Number: "555-4321", Type: pb.Person_HOME},
		},
	}
	book.People = append(book.People, &p)

	// Write the new address book back to disk.
	out, err := proto.Marshal(book)
	if err != nil {
		log.Fatalln("Failed to encode address book:", err)
	}
	if err := ioutil.WriteFile("./myAddress", out, 0644); err != nil {
		log.Fatalln("Failed to write address book:", err)
	}

	// Read the existing address book.
	in, err := ioutil.ReadFile("./myAddress")
	if err != nil {
		log.Fatalln("Error reading file:", err)
	}

	bookFromDisk := &pb.AddressBook{}
	if err := proto.Unmarshal(in, bookFromDisk); err != nil {
		log.Fatalln("Failed to parse address book:", err)
	}

	fmt.Println("Get info as string")
	fmt.Println(bookFromDisk.String())
}