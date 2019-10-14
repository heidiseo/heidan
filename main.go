package main

import (
	"fmt"
	"os"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	//"github.com/aws/aws-sdk-go/service/s3"
)


func main() {
	createMyDBTable()


}

func createMyDBTable(){
	sess, err := session.NewSession(&aws.Config{Region: aws.String("eu-west-2")})
	svc := dynamodb.New(sess)

	input := &dynamodb.CreateTableInput{
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("Name"),
				AttributeType: aws.String("NString"),
			},
			{
				AttributeName: aws.String("Location"),
				AttributeType: aws.String("LString"),
			},
			{
				AttributeName: aws.String("EstimatedDuration"),
				AttributeType: aws.String("EString"),
			},
			{
				AttributeName: aws.String("Price"),
				AttributeType: aws.String("Float"),
			},
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("Name"),
				KeyType:       aws.String("HASH"),
			},
			{
				AttributeName: aws.String("Location"),
				KeyType:       aws.String("HASH"),
			},
			{
				AttributeName: aws.String("EstimatedDuration"),
				KeyType:       aws.String("HASH"),
			},
			{
				AttributeName: aws.String("Price"),
				KeyType:       aws.String("HASH"),
			},
		},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(1) ,
			WriteCapacityUnits: aws.Int64(1),
		},
		TableName: aws.String("activities") ,
	}
	ss, err := svc.CreateTable(input)
	if err != nil {
		fmt.Println("Got error calling CreateTable:")
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println("Created the table", "activities")
	fmt.Println(ss)
}
