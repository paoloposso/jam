package user

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/paoloposso/jam/src/core"
	customerrors "github.com/paoloposso/jam/src/core/custom_errors"
	"github.com/paoloposso/jam/src/users"
)

const tableName string = "jam-users"

type UserRepository struct {
	client *dynamodb.Client
}

func NewUserRepository() (*UserRepository, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO(), func(o *config.LoadOptions) error {
		return nil
	})
	if err != nil {
		panic(err)
	}

	svc := dynamodb.NewFromConfig(cfg)

	return &UserRepository{
		client: svc,
	}, nil
}

func (repo *UserRepository) Insert(user users.User) (err error) {
	err = repo.insertUserCredentials(user)

	if err != nil {
		return err
	}

	err = repo.insertUserInfo(user)

	if err != nil {
		return err
	}
	return nil
}

func (repo *UserRepository) insertUserInfo(user users.User) error {
	valuesMap, err := attributevalue.MarshalMap(UserInfoModel{
		Email:     user.Email,
		PK:        "USER#" + user.ID,
		SK:        "USER#" + user.ID,
		BirthDate: user.BirthDate,
		Name:      user.Name,
	})

	if err != nil {
		return err
	}

	_, err = repo.client.PutItem(
		context.TODO(),
		&dynamodb.PutItemInput{
			Item:      valuesMap,
			TableName: aws.String(tableName),
		})

	return nil
}

func (repo *UserRepository) insertUserCredentials(user users.User) error {
	output, err := repo.client.GetItem(
		context.TODO(),
		&dynamodb.GetItemInput{
			Key: map[string]types.AttributeValue{
				"PK": &types.AttributeValueMemberS{Value: "ULOGIN#" + user.Email},
				"SK": &types.AttributeValueMemberS{Value: "ULOGIN#" + user.Email},
			},
			TableName: aws.String(tableName),
		})

	if output.Item != nil {
		return customerrors.CreateArgumentError("E-mail already exists")
	}

	hashed, err := core.HashPassword(user.Password)

	if err != nil {
		return err
	}

	valuesMap, err := attributevalue.MarshalMap(UserLogin{
		PK:       "ULOGIN#" + user.Email,
		SK:       "ULOGIN#" + user.Email,
		UserID:   user.ID,
		Password: hashed,
	})

	if err != nil {
		return err
	}

	_, err = repo.client.PutItem(
		context.TODO(),
		&dynamodb.PutItemInput{
			Item:      valuesMap,
			TableName: aws.String(tableName),
		})

	return nil
}
