package repo

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/maiconDeSouza/lotofacil_lab_v3/internal/errapp"
	"github.com/maiconDeSouza/lotofacil_lab_v3/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ScoredRepoInterface interface {
	Find(s *model.Scored) error
	Update(s *model.Scored) error
}

type ScoredRepo struct {
	collection *mongo.Collection
}

func NewScoredRepo(collection *mongo.Collection) *ScoredRepo {
	return &ScoredRepo{
		collection: collection,
	}
}

func (r *ScoredRepo) Find(s *model.Scored) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(5)*time.Second)
	defer cancel()

	err := r.collection.FindOne(ctx, bson.M{}).Decode(s)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return errapp.ErrNoDocuments
		}
		return err
	}

	return nil
}

func (r *ScoredRepo) Update(s *model.Scored) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(5)*time.Second)
	defer cancel()

	filter := bson.M{}
	up := bson.M{
		"$set": bson.M{
			"scores": s.Scores,
		},
	}

	res, err := r.collection.UpdateOne(ctx, filter, up)
	if err != nil {
		return fmt.Errorf("erro ao atualizar documento: %w", err)
	}

	if res.MatchedCount == 0 {
		return errapp.ErrNoDocuments
	}

	return nil
}
