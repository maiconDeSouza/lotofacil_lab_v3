package repo

import (
	"context"
	"errors"
	"time"

	"github.com/maiconDeSouza/lotofacil_lab_v3/internal/errapp"
	"github.com/maiconDeSouza/lotofacil_lab_v3/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type FixedRepoInterface interface {
	Find(f *model.GameFixed) error
}

type FixedRepo struct {
	collection *mongo.Collection
}

func NewFixedRepo(collection *mongo.Collection) *FixedRepo {
	return &FixedRepo{
		collection: collection,
	}
}

func (r *FixedRepo) Find(f *model.GameFixed) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(5)*time.Second)
	defer cancel()

	err := r.collection.FindOne(ctx, bson.M{}).Decode(f)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return errapp.ErrNoDocuments
		}
		return err
	}

	return nil
}
