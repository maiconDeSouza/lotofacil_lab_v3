package service

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/maiconDeSouza/lotofacil_lab_v3/internal/errapp"
	"github.com/maiconDeSouza/lotofacil_lab_v3/internal/lottery"
	"github.com/maiconDeSouza/lotofacil_lab_v3/internal/model"
	"github.com/maiconDeSouza/lotofacil_lab_v3/internal/repo"
)

type ServiceInterface interface {
	CreateNewGameFixed(n int) (*model.Game, *errapp.ErrAppRes)
	UpdateGameFixed(number int, up model.RequestUpdateGame) (*model.Game, *errapp.ErrAppRes)
	GetGameFixed(number int) (*model.Game, *errapp.ErrAppRes)
	GetAllGamesFixed() ([]model.Game, *errapp.ErrAppRes)
}

type Service struct {
	repoGame   repo.GameRepoInterface
	repoFixed  repo.FixedRepoInterface
	repoScored repo.ScoredRepoInterface
	lottery    *lottery.Lottery
}

func NewGameService(
	repoGame repo.GameRepoInterface,
	repoFixed repo.FixedRepoInterface,
	repoScored repo.ScoredRepoInterface,
	lottery *lottery.Lottery,
) *Service {
	return &Service{
		repoGame:   repoGame,
		repoFixed:  repoFixed,
		repoScored: repoScored,
		lottery:    lottery,
	}
}

func (s *Service) CreateNewGameFixed(n int) (*model.Game, *errapp.ErrAppRes) {
	fn := &model.GameFixed{}

	if err := s.repoFixed.Find(fn); err != nil {
		if errors.Is(err, errapp.ErrNoDocuments) {
			errAppRes := errapp.NewErrAppRes(err, errapp.ErrNoDocuments.Error(), http.StatusBadRequest)
			return nil, errAppRes
		}
		errAppRes := errapp.NewErrAppRes(err, "Erro no servidor", http.StatusInternalServerError)
		return nil, errAppRes
	}

	g := &model.Game{
		Number:       n,
		Numbers:      fn.Numbers,
		TypeStrategy: model.StrategyFixed,
		Createad:     time.Now(),
	}

	if err := s.repoGame.Insert(g); err != nil {
		errAppRes := errapp.NewErrAppRes(err, "Erro no servidor", http.StatusInternalServerError)
		return nil, errAppRes
	}

	return g, nil
}

func (s *Service) UpdateGameFixed(number int, up model.RequestUpdateGame) (*model.Game, *errapp.ErrAppRes) {
	g := &model.Game{}

	if err := s.repoGame.Update(number, model.StrategyFixed, up.Correct, up.CorrectNumbers); err != nil {
		if errors.Is(err, errapp.ErrNoDocuments) {
			errAppRes := errapp.NewErrAppRes(err, errapp.ErrNoDocuments.Error(), http.StatusBadRequest)
			return nil, errAppRes
		}
		errAppRes := errapp.NewErrAppRes(err, "Erro no servidor", http.StatusInternalServerError)
		return nil, errAppRes
	}

	if err := s.repoGame.Find(number, model.StrategyFixed, g); err != nil {
		if errors.Is(err, errapp.ErrNoDocuments) {
			errAppRes := errapp.NewErrAppRes(err, errapp.ErrNoDocuments.Error(), http.StatusBadRequest)
			return nil, errAppRes
		}
		errAppRes := errapp.NewErrAppRes(err, "Erro no servidor", http.StatusInternalServerError)
		return nil, errAppRes
	}

	return g, nil

}

func (s *Service) GetGameFixed(number int) (*model.Game, *errapp.ErrAppRes) {
	g := &model.Game{}

	if err := s.repoGame.Find(number, model.StrategyFixed, g); err != nil {
		if errors.Is(err, errapp.ErrNoDocuments) {
			errAppRes := errapp.NewErrAppRes(err, errapp.ErrNoDocuments.Error(), http.StatusBadRequest)
			return nil, errAppRes
		}
		errAppRes := errapp.NewErrAppRes(err, "Erro no servidor", http.StatusInternalServerError)
		return nil, errAppRes
	}

	return g, nil
}

func (s *Service) GetAllGamesFixed() ([]model.Game, *errapp.ErrAppRes) {
	gamers := &[]model.Game{}

	if err := s.repoGame.FindAll(gamers); err != nil {
		fmt.Println(gamers)
		if errors.Is(err, errapp.ErrNoDocuments) {
			errAppRes := errapp.NewErrAppRes(err, errapp.ErrNoDocuments.Error(), http.StatusBadRequest)
			return nil, errAppRes
		}
		errAppRes := errapp.NewErrAppRes(err, "Erro no servidor", http.StatusInternalServerError)
		return nil, errAppRes
	}

	fmt.Println(gamers)

	gamersFixed := []model.Game{}

	for _, v := range *gamers {
		if v.TypeStrategy == model.StrategyFixed {
			gamersFixed = append(gamersFixed, v)
		}
	}

	return gamersFixed, nil
}
