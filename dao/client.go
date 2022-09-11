package dao

import (
	"errors"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type DynamoDb interface {
	QueryWithContext(aws.Context, *dynamodb.QueryInput, ...request.Option) (*dynamodb.QueryOutput, error)
	BatchWriteItemWithContext(aws.Context, *dynamodb.BatchWriteItemInput, ...request.Option) (*dynamodb.BatchWriteItemOutput, error)
	UpdateItemWithContext(aws.Context, *dynamodb.UpdateItemInput, ...request.Option) (*dynamodb.UpdateItemOutput, error)
	DeleteItemWithContext(aws.Context, *dynamodb.DeleteItemInput, ...request.Option) (*dynamodb.DeleteItemOutput, error)
}

type Client struct {
	tableName string
	marshal   func(in interface{}) (map[string]*dynamodb.AttributeValue, error)
	unmarshal func(m map[string]*dynamodb.AttributeValue, out interface{}) error
	db        DynamoDb
}

type ClientOptions struct {
	TableName string
	DB        DynamoDb
	Marshal   func(in interface{}) (map[string]*dynamodb.AttributeValue, error)
	Unmarshal func(m map[string]*dynamodb.AttributeValue, out interface{}) error
}

func New(opts *ClientOptions) (*Client, error) {

	if opts.TableName == "" {
		opts.TableName = "serivce-dynamo"
	}

	if opts.DB == nil {
		return nil, errors.New("DB missing")
	}

	if opts.Marshal == nil {
		opts.Marshal = dynamodbattribute.MarshalMap
	}
	if opts.Unmarshal == nil {
		opts.Unmarshal = dynamodbattribute.UnmarshalMap
	}

	return &Client{
		tableName: opts.TableName,
		db:        opts.DB,
		marshal:   opts.Marshal,
		unmarshal: opts.Unmarshal,
	}, nil
}

func Must(client *Client, err error) *Client {
	if err != nil {
		panic(err)
	}
	return client
}
