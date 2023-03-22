package authrepo

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/paoloposso/jam/libs/core/customerrors"
)

const tableName string = "jam-users"

type AuthRepository struct {
	client *dynamodb.Client
}

func NewRepository(dbClient *dynamodb.Client) (*AuthRepository, error) {
	return &AuthRepository{
		client: dbClient,
	}, nil
}

func (repo AuthRepository) GetUserByEmail(email string) (userId, password string, err error) {
	dbResult, err := repo.client.GetItem(context.Background(), &dynamodb.GetItemInput{
		Key: map[string]types.AttributeValue{
			"PK": &types.AttributeValueMemberS{Value: "ULOGIN#" + email},
			"SK": &types.AttributeValueMemberS{Value: "ULOGIN#" + email},
		},
		TableName: aws.String(tableName),
	})

	if err != nil {
		return "", "", err
	}

	if dbResult.Item == nil {
		return "", "", customerrors.CreateUnauthorizedError()
	}

	var model UserModel

	err = attributevalue.UnmarshalMap(dbResult.Item, &model)

	if err != nil {
		return "", "", err
	}

	return model.UserID, model.Password, nil
}
