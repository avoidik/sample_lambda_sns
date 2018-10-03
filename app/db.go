package main

import (
	"strconv"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

const tableName = "Items"

var db = dynamodb.New(session.New(), aws.NewConfig().WithRegion("eu-west-1"))

func (e *eventData) createOrUpdateDB() error {
	input := &dynamodb.PutItemInput{
		TableName: aws.String(tableName),
		Item: map[string]*dynamodb.AttributeValue{
			"id": {
				N: aws.String(strconv.Itoa(e.ID)),
			},
			"data": {
				S: aws.String(e.Data),
			},
		},
	}
	_, err := db.PutItem(input)
	return err
}

func (e *eventData) deleteDB() error {
	input := &dynamodb.DeleteItemInput{
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				N: aws.String(strconv.Itoa(e.ID)),
			},
		},
	}
	_, err := db.DeleteItem(input)
	return err
}
