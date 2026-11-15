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

type GameRepoInterface interface {
	Insert(game *model.Game) error
	Update(number int, strategy model.TypeStrategy, correct int, correctNumbers model.Numbers) error
	Find(number int, strategy model.TypeStrategy, game *model.Game) error
	FindAll(games *[]model.Game) error
	FindByNumber(n int, gamers *[]model.Game) error
}

type GameRepo struct {
	colletion *mongo.Collection
}

func NewGameRepo(collection *mongo.Collection) *GameRepo {
	return &GameRepo{
		colletion: collection,
	}
}

func (g *GameRepo) Insert(game *model.Game) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(5)*time.Second)
	defer cancel()

	if _, err := g.colletion.InsertOne(ctx, game); err != nil {
		return err
	}
	return nil
}

func (g *GameRepo) Update(number int, strategy model.TypeStrategy, correct int, correctNumbers model.Numbers) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(5)*time.Second)
	defer cancel()

	filter := bson.D{{Key: "number", Value: number}, {Key: "typeStrategy", Value: strategy}}

	up := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "correct", Value: correct},
			{Key: "correctNumbers", Value: correctNumbers},
		}},
	}

	res, err := g.colletion.UpdateOne(ctx, filter, up)
	if err != nil {
		return fmt.Errorf("erro ao atualizar documento: %w", err)
	}

	if res.MatchedCount == 0 {
		return errapp.ErrNoDocuments
	}

	return nil
}

func (g *GameRepo) Find(number int, strategy model.TypeStrategy, game *model.Game) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(5)*time.Second)
	defer cancel()

	filter := bson.M{"number": number, "typeStrategy": strategy}
	if err := g.colletion.FindOne(ctx, filter).Decode(&game); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return errapp.ErrNoDocuments
		}
		return err
	}
	return nil
}

func (g *GameRepo) FindAll(gamers *[]model.Game) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(5)*time.Second)
	defer cancel()

	cursor, err := g.colletion.Find(ctx, bson.M{})
	if err != nil {
		return err
	}
	defer cursor.Close(ctx)

	if err := cursor.All(ctx, gamers); err != nil {
		return err
	}

	if len(*gamers) == 0 {
		return errapp.ErrNoDocuments
	}

	return nil
}

func (g *GameRepo) FindByNumber(n int, gamers *[]model.Game) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(5)*time.Second)
	defer cancel()

	filter := bson.M{"number": n}
	err := g.colletion.FindOne(ctx, filter).Decode(gamers)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return errapp.ErrNoDocuments
		}
		return err
	}
	return nil
}
