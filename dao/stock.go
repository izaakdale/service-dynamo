package dao

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type StockIndex struct {
	Ticker    string `json:"ticker,omitempty" dynamodbav:"ticker"`
	Attribute string `json:"attribute,omitempty" dynamodbav:"attribute"`

	// attribute - latest
	Timestamp string `json:"timestamp,omitempty"`
	Price     string `json:"price,omitempty"`

	// attribute - daily
	OpenPrice  float64 `json:"open_price,omitempty"`
	HighPrice  float64 `json:"high_price,omitempty"`
	LowPrice   float64 `json:"low_price,omitempty"`
	ClosePrice float64 `json:"close_price,omitempty"`
}

func (c *Client) StoreStockPrice(si StockIndex) {

	log.Println("Hello from DB client")

	testMap, err := dynamodbattribute.MarshalMap(si)
	if err != nil {
		log.Fatal(err)
	}

	writeRequest := []*dynamodb.WriteRequest{}
	writeRequest = append(writeRequest, &dynamodb.WriteRequest{
		PutRequest: &dynamodb.PutRequest{
			Item: testMap,
		},
	})

	_, err = c.db.BatchWriteItemWithContext(context.Background(), &dynamodb.BatchWriteItemInput{
		RequestItems: map[string][]*dynamodb.WriteRequest{
			c.tableName: writeRequest,
		},
	})
	if err != nil {
		log.Fatal(err)
	}
}
