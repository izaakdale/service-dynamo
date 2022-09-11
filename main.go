package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/izaakdale/service-dynamo/dao"
	"github.com/izaakdale/service-dynamo/handlers"
)

func main() {

	if "" == os.Getenv("DYNAMO_REGION") || "" == os.Getenv("DYNAMO_TABLE") {
		log.Fatal("Env error")
	}

	session, err := session.NewSession(aws.NewConfig().WithRegion(os.Getenv("DYNAMO_REGION")))
	if err != nil {
		log.Fatal(err)
	}

	dynamo := dynamodb.New(session)

	client := dao.Must(dao.New(&dao.ClientOptions{
		TableName: os.Getenv("DYNAMO_TABLE"),
		DB:        dynamo,
	}))

	opts := handlers.ServiceOptions{
		DBClient: *client,
	}
	service := handlers.New(opts)

	srv := &http.Server{
		Handler:      service.WithRoutes().Router,
		Addr:         "localhost:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
