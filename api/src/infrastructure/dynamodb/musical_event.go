package database

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	musicalevents "github.com/paoloposso/jam/src/musical_events"
)

const tableName string = "jam"

type MusicalEventRepository struct {
	session *session.Session
}

func NewMusicalEventRepository() (*MusicalEventRepository, error) {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	return &MusicalEventRepository{
		session: sess,
	}, nil
}

func (repo *MusicalEventRepository) Create(item musicalevents.MusicalEvent) (id string, err error) {
	svc := dynamodb.New(repo.session)

	attributeValue, err := dynamodbattribute.MarshalMap(item)
	if err != nil {
		return "", err
	}

	input := &dynamodb.PutItemInput{
		Item:      attributeValue,
		TableName: aws.String(tableName),
	}

	_, err = svc.PutItem(input)

	return "", nil
}
