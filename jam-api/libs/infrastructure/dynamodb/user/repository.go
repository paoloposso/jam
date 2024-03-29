package userrepo

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/paoloposso/jam/libs/core/customerrors"
	"github.com/paoloposso/jam/libs/users"
)

const tableName string = "jam-users"

type UserRepository struct {
	client *dynamodb.Client
}

func NewUserRepository(dbClient *dynamodb.Client) (*UserRepository, error) {
	return &UserRepository{
		client: dbClient,
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

func (repo UserRepository) insertUserCredentials(user users.User) error {
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

	valuesMap, err := attributevalue.MarshalMap(UserLogin{
		PK:       "ULOGIN#" + user.Email,
		SK:       "ULOGIN#" + user.Email,
		UserID:   user.ID,
		Password: user.Password,
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

func (repo UserRepository) Get(id string) (*users.User, error) {
	output, err := repo.client.GetItem(
		context.TODO(),
		&dynamodb.GetItemInput{
			Key: map[string]types.AttributeValue{
				"PK": &types.AttributeValueMemberS{Value: "USER#" + id},
				"SK": &types.AttributeValueMemberS{Value: "USER#" + id},
			},
			TableName: aws.String(tableName),
		})

	if err != nil {
		return nil, err
	}

	if output.Item == nil {
		return nil, customerrors.CreateArgumentError("User ID not found in the database")
	}

	var model UserInfoModel

	err = attributevalue.UnmarshalMap(output.Item, &model)

	return &users.User{ID: id, Email: model.Email, BirthDate: model.BirthDate}, nil
}
