package user

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/paoloposso/jam/src/users"
)

const tableName string = "jam-users"

type UserRepository struct {
	session *session.Session
}

func NewUserRepository() (*UserRepository, error) {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	return &UserRepository{
		session: sess,
	}, nil
}

func (repo *UserRepository) Insert(item users.User) error {
	svc := dynamodb.New(repo.session)

	valuesMap, err := dynamodbattribute.MarshalMap(UserModel{
		Email:     item.Email,
		PK:        "USER#" + item.ID,
		SK:        "USER#" + item.ID,
		BirthDate: item.BirthDate,
		Name:      item.Name,
	})

	if err != nil {
		return err
	}

	input := &dynamodb.PutItemInput{
		Item:      valuesMap,
		TableName: aws.String(tableName),
	}

	fds, err := svc.PutItem(input)

	fmt.Printf("%v", fds)

	return nil
}
