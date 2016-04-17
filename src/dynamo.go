package main

import (
    "fmt"
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/credentials"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/dynamodb"
)

func main() {
    cred := credentials.NewStaticCredentials("foo", "bar", "foobar")
    sess := session.New(&aws.Config{Region: aws.String("us-west-2"),
                                    Credentials: cred,
                                    Endpoint: aws.String("localhost:8000"),
                                    DisableSSL: aws.Bool(true)})
    svc := dynamodb.New(sess)
    params := &dynamodb.CreateTableInput{
      TableName: aws.String("TableName"),
      AttributeDefinitions: []*dynamodb.AttributeDefinition{
        {
          AttributeName: aws.String("Id"),
          AttributeType: aws.String("S"),
        },
      },
      KeySchema: []*dynamodb.KeySchemaElement{
        {
          AttributeName: aws.String("Id"),
          KeyType: aws.String("HASH"), 
        },
      },
      ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
        ReadCapacityUnits: aws.Int64(1),
        WriteCapacityUnits: aws.Int64(1),
      },
    }
    resp, err := svc.CreateTable(params)
    if err != nil {
      fmt.Println(err.Error())
      return
    }
    fmt.Println(resp)
    return
}
