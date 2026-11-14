package repo

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type ScoredRepoInterface interface {
}

type ScoredRepo struct {
	collection *mongo.Collection
}

func NewScoredRepo(collection *mongo.Collection) *ScoredRepo {
	return &ScoredRepo{
		collection: collection,
	}
}
