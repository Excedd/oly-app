package application

import (
	"context"
	i "oly-backend/aplication/services/interfaces"
	d "oly-backend/domain"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
)

type prsService struct {
	collection *mongo.Collection
	mu         sync.Mutex
}

func NewPRSService(coll *mongo.Collection) i.Service {
	return &prsService{
		collection: coll,
	}
}

func (s *prsService) SetPRs(ctx context.Context, prs d.PersonalRecord) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	_, err := s.collection.DeleteMany(ctx, map[string]interface{}{})
	if err != nil {
		return err
	}

	_, err = s.collection.InsertOne(ctx, prs)
	return err
}

func (s *prsService) GetPRs(ctx context.Context) (d.PersonalRecord, error) {
	var result d.PersonalRecord
	err := s.collection.FindOne(ctx, map[string]interface{}{}).Decode(&result)
	return result, err
}
