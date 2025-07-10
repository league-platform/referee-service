package handlers

import (
    "encoding/json"
    "fmt"
    "net/http"
    "time"

    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/dynamodb"
    "github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type Referee struct {
    ID        string    `json:"id"`
    Name      string    `json:"name"`
    Category  string    `json:"category"`
    CreatedAt time.Time `json:"createdAt"`
}

func CreateReferee(w http.ResponseWriter, r *http.Request) {
    var referee Referee
    _ = json.NewDecoder(r.Body).Decode(&referee)
    referee.CreatedAt = time.Now()

    sess := session.Must(session.NewSession(&aws.Config{
        Region: aws.String("us-east-1"),
    }))
    svc := dynamodb.New(sess)

    av, _ := dynamodbattribute.MarshalMap(referee)
    input := &dynamodb.PutItemInput{
        TableName: aws.String("Referees"),
        Item:      av,
    }

    _, err := svc.PutItem(input)
    if err != nil {
        http.Error(w, "Failed to save referee", http.StatusInternalServerError)
        return
    }

    fmt.Printf("EVENT: referee.created -> %+v\n", referee)
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(referee)
}

func GetReferees(w http.ResponseWriter, r *http.Request) {
    sess := session.Must(session.NewSession(&aws.Config{
        Region: aws.String("us-east-1"),
    }))
    svc := dynamodb.New(sess)

    input := &dynamodb.ScanInput{
        TableName: aws.String("Referees"),
    }

    result, err := svc.Scan(input)
    if err != nil {
        http.Error(w, "Failed to retrieve referees", http.StatusInternalServerError)
        return
    }

    var referees []Referee
    _ = dynamodbattribute.UnmarshalListOfMaps(result.Items, &referees)
    json.NewEncoder(w).Encode(referees)
}
