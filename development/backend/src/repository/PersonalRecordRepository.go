package repository

import (
	"context"
	"oly-backend/domain"
	d "oly-backend/domain"

	"go.mongodb.org/mongo-driver/mongo"
)

type PRRepository interface {
	SavePRs(ctx context.Context, prs d.PersonalRecord) error
	GetPRs(ctx context.Context) (d.PersonalRecord, error)
}

type personalRecordRepository struct {
	collection *mongo.Collection
}

var instance *personalRecordRepository

func InitPRRepository(collection *mongo.Collection) {
	instance = &personalRecordRepository{collection}
}

func GetPRs(ctx context.Context, filter interface{}) (*mongo.SingleResult, error) {
	if instance == nil {
		return nil, mongo.ErrNilDocument
	}
	return instance.collection.FindOne(ctx, filter), nil
}

func SavePRs(ctx context.Context, pr domain.PersonalRecord) error {
	if instance == nil {
		return mongo.ErrNilDocument
	}
	_, err := instance.collection.InsertOne(ctx, pr)
	return err
}
