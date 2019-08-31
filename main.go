package main

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
	"github.com/rs/xid"
)

type TestData struct {
	ID   string `dynamo:"ID"`
	Name string `dynamo:"Name"`
}

func main() {
	log.Println("New Session")
	dynamoSession, err := session.NewSession(&aws.Config{
		Region:      aws.String("ap-northeast-1"),
		Endpoint:    aws.String("http://dynamodb:8000"),
		Credentials: credentials.NewEnvCredentials(),
	})
	if err != nil {
		panic(err)
	}

	db := dynamo.New(dynamoSession)
	table := db.Table("TestTable")

	// put item
	guid := xid.New()
	log.Println("Put", guid.String())
	w := TestData{ID: guid.String(), Name: "hello"}
	err = table.Put(w).Run()
	if err != nil {
		panic(err)
	}

	// get item
	log.Println("Get")
	var ret []TestData
	err = table.Scan().All(&ret)
	if err != nil {
		panic(err)
	}

	log.Printf("%#v\n", ret)

	// delete item
	log.Println("Delete", guid.String())
	err = table.Delete("ID", guid.String()).Run()
	if err != nil {
		panic(err)
	}

}
