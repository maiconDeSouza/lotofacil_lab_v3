package model

import "time"

type Numbers []int

type TypeStrategy string

const (
	StrategyFixed     TypeStrategy = "fixed"
	StrategyFrequency TypeStrategy = "frequency"
	StrategyRandom    TypeStrategy = "random"
)

type Game struct {
	Number         int          `bson:"number" json:"number"`
	Numbers        Numbers      `bson:"numbers" json:"numbers"`
	TypeStrategy   TypeStrategy `bson:"typeStrategy" json:"typeStrategy"`
	Correct        int          `bson:"correct" json:"correct"`
	CorrectNumbers Numbers      `bson:"correctNumbers" json:"correctNumbers"`
	Createad       time.Time    `bson:"criadoEm" json:"criadoEm"`
}

type TypeGame struct {
	Name  TypeStrategy `bson:"name" json:"name"`
	Score int          `bson:"score" json:"score"`
}

type Scored struct {
	Scores []TypeGame `bson:"scores" json:"scores"`
}

type GameFixed struct {
	Numbers Numbers `bson:"numbers" json:"numbers"`
}

type RequestGame struct {
	Number int `json:"number"`
}

type RequestUpdateGame struct {
	Correct        int     `json:"correct"`
	CorrectNumbers Numbers `json:"correctNumbers"`
}
