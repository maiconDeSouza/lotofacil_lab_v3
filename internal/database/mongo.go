package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/maiconDeSouza/lotofacil_lab_v3/internal/env"
	"github.com/maiconDeSouza/lotofacil_lab_v3/internal/errapp"
	"github.com/maiconDeSouza/lotofacil_lab_v3/internal/lottery"
	"github.com/maiconDeSouza/lotofacil_lab_v3/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Collections struct {
	CollectionGamers  *mongo.Collection
	CollectionFixed   *mongo.Collection
	CollectionsScored *mongo.Collection
}

func NewConnectMongo(env *env.Env, lottery *lottery.Lottery) (*Collections, *errapp.ErrAppRes) {
	uri := fmt.Sprintf(
		"mongodb+srv://%s:%s@cluster0.xhbyyve.mongodb.net/",
		env.MONGO_ROOT_USER,
		env.MONGO_ROOT_PASSWORD,
	)
	clientOptions := options.Client().ApplyURI(uri)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		errAppRes := errapp.NewErrAppRes(err, "❌ Erro ao ligar ao MongoDB.", 0)
		return nil, errAppRes
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		errAppRes := errapp.NewErrAppRes(err, "❌ Não foi possível alcançar o MongoDB.", 0)
		return nil, errAppRes
	}

	db := client.Database(env.MONGO_DB_NAME)
	collectionGamers := db.Collection(env.MONGO_COLLECTION_GAMERS)
	collectionGamerFixed := db.Collection(env.MONGO_COLLECTION_FIXED)
	collectionsScored := db.Collection(env.MONGO_COLLECTION_SCORED)

	fmt.Println("✅ Conectado ao banco de dados com sucesso!")

	collections := &Collections{
		CollectionGamers:  collectionGamers,
		CollectionFixed:   collectionGamerFixed,
		CollectionsScored: collectionsScored,
	}

	count, err := collectionGamerFixed.CountDocuments(ctx, bson.M{})
	if err != nil {
		errAppRes := errapp.NewErrAppRes(err, "❌ Erro ao verificar a coleção de jogo fixo", 0)
		return nil, errAppRes
	}

	if count == 0 {
		log.Println("✅ Criando jogo fixo...")
		l := lottery.GenerateLotofacilNumbers15()
		g := model.GameFixed{
			Numbers: l,
		}

		if _, err := collectionGamerFixed.InsertOne(ctx, g); err != nil {
			errAppRes := errapp.NewErrAppRes(err, "❌ Erro ao inserir jogo fixo", 0)
			return nil, errAppRes
		}
		log.Println("✅ Jogo fixo criado com sucesso:", l)
	}

	count, err = collectionsScored.CountDocuments(ctx, bson.M{})
	if err != nil {
		errAppRes := errapp.NewErrAppRes(err, "❌ Erro ao verificar a coleção de pontuação", 0)
		return nil, errAppRes
	}

	if count == 0 {
		log.Println("✅ Criando pontuação...")
		scoredDoc := model.Scored{
			Scores: []model.TypeGame{
				{
					Name:  model.StrategyFixed,
					Score: 0,
				},
				{
					Name:  model.StrategyFrequency,
					Score: 0,
				},
				{
					Name:  model.StrategyRandom,
					Score: 0,
				},
			},
		}

		if _, err := collectionsScored.InsertOne(ctx, scoredDoc); err != nil {
			errAppRes := errapp.NewErrAppRes(err, "❌ Erro ao inserir pontuação", 0)
			return nil, errAppRes
		}

		log.Println("✅ Pontuação criada com sucesso:", scoredDoc)
	}

	return collections, nil
}
