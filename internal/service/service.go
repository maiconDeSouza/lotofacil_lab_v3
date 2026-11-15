package service

import (
	"cmp"
	"errors"
	"net/http"
	"slices"
	"strconv"
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
	CreateNewGameFrequency(n int) (*model.ResponseGameFrequency, *errapp.ErrAppRes)
	UpdateGameFrequency(number int, up model.RequestUpdateGame) (*model.Game, *errapp.ErrAppRes)
	GetGameFrequency(number int) (*model.Game, *errapp.ErrAppRes)
	GetAllGamesFrequency() ([]model.Game, *errapp.ErrAppRes)
	CreateNewGameRandom(n int) (*model.Game, *errapp.ErrAppRes)
	UpdateGameRandom(number int, up model.RequestUpdateGame) (*model.Game, *errapp.ErrAppRes)
	GetGameRandom(number int) (*model.Game, *errapp.ErrAppRes)
	GetAllGamesRandom() ([]model.Game, *errapp.ErrAppRes)
	UpdateScored(up *model.Scored) (*model.Scored, *errapp.ErrAppRes)
	ConfirmResult(n int) (*model.ResponseConfirmResult, *errapp.ErrAppRes)
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
		if errors.Is(err, errapp.ErrNoDocuments) {
			errAppRes := errapp.NewErrAppRes(err, errapp.ErrNoDocuments.Error(), http.StatusBadRequest)
			return nil, errAppRes
		}
		errAppRes := errapp.NewErrAppRes(err, "Erro no servidor", http.StatusInternalServerError)
		return nil, errAppRes
	}

	gamersFixed := []model.Game{}

	for _, v := range *gamers {
		if v.TypeStrategy == model.StrategyFixed {
			gamersFixed = append(gamersFixed, v)
		}
	}

	return gamersFixed, nil
}

func (s *Service) CreateNewGameFrequency(n int) (*model.ResponseGameFrequency, *errapp.ErrAppRes) {
	numbers, errAppRes := s.lottery.GetFrequentNumbers()
	if errAppRes != nil {
		return nil, errAppRes
	}

	g := &model.Game{
		Number:       n,
		Numbers:      numbers.Palpite,
		TypeStrategy: model.StrategyFrequency,
		Createad:     time.Now(),
	}

	if err := s.repoGame.Insert(g); err != nil {
		errAppRes := errapp.NewErrAppRes(err, "Erro no servidor", http.StatusInternalServerError)
		return nil, errAppRes
	}

	rf := &model.ResponseGameFrequency{
		Game:      *g,
		Frequency: numbers.Ranking,
	}

	return rf, nil
}

func (s *Service) UpdateGameFrequency(number int, up model.RequestUpdateGame) (*model.Game, *errapp.ErrAppRes) {
	g := &model.Game{}

	if err := s.repoGame.Update(number, model.StrategyFrequency, up.Correct, up.CorrectNumbers); err != nil {
		if errors.Is(err, errapp.ErrNoDocuments) {
			errAppRes := errapp.NewErrAppRes(err, errapp.ErrNoDocuments.Error(), http.StatusBadRequest)
			return nil, errAppRes
		}
		errAppRes := errapp.NewErrAppRes(err, "Erro no servidor", http.StatusInternalServerError)
		return nil, errAppRes
	}

	if err := s.repoGame.Find(number, model.StrategyFrequency, g); err != nil {
		if errors.Is(err, errapp.ErrNoDocuments) {
			errAppRes := errapp.NewErrAppRes(err, errapp.ErrNoDocuments.Error(), http.StatusBadRequest)
			return nil, errAppRes
		}
		errAppRes := errapp.NewErrAppRes(err, "Erro no servidor", http.StatusInternalServerError)
		return nil, errAppRes
	}

	return g, nil
}

func (s *Service) GetGameFrequency(number int) (*model.Game, *errapp.ErrAppRes) {
	g := &model.Game{}

	if err := s.repoGame.Find(number, model.StrategyFrequency, g); err != nil {
		if errors.Is(err, errapp.ErrNoDocuments) {
			errAppRes := errapp.NewErrAppRes(err, errapp.ErrNoDocuments.Error(), http.StatusBadRequest)
			return nil, errAppRes
		}
		errAppRes := errapp.NewErrAppRes(err, "Erro no servidor", http.StatusInternalServerError)
		return nil, errAppRes
	}

	return g, nil
}

func (s *Service) GetAllGamesFrequency() ([]model.Game, *errapp.ErrAppRes) {
	gamers := &[]model.Game{}

	if err := s.repoGame.FindAll(gamers); err != nil {
		if errors.Is(err, errapp.ErrNoDocuments) {
			errAppRes := errapp.NewErrAppRes(err, errapp.ErrNoDocuments.Error(), http.StatusBadRequest)
			return nil, errAppRes
		}
		errAppRes := errapp.NewErrAppRes(err, "Erro no servidor", http.StatusInternalServerError)
		return nil, errAppRes
	}

	gamersFrequency := []model.Game{}

	for _, v := range *gamers {
		if v.TypeStrategy == model.StrategyFrequency {
			gamersFrequency = append(gamersFrequency, v)
		}
	}

	return gamersFrequency, nil
}

func (s *Service) CreateNewGameRandom(n int) (*model.Game, *errapp.ErrAppRes) {
	g := &model.Game{
		Number:       n,
		Numbers:      s.lottery.GenerateLotofacilNumbers15(),
		TypeStrategy: model.StrategyRandom,
		Createad:     time.Now(),
	}

	if err := s.repoGame.Insert(g); err != nil {
		errAppRes := errapp.NewErrAppRes(err, "Erro no servidor", http.StatusInternalServerError)
		return nil, errAppRes
	}

	return g, nil
}

func (s *Service) UpdateGameRandom(number int, up model.RequestUpdateGame) (*model.Game, *errapp.ErrAppRes) {
	g := &model.Game{}

	if err := s.repoGame.Update(number, model.StrategyRandom, up.Correct, up.CorrectNumbers); err != nil {
		if errors.Is(err, errapp.ErrNoDocuments) {
			errAppRes := errapp.NewErrAppRes(err, errapp.ErrNoDocuments.Error(), http.StatusBadRequest)
			return nil, errAppRes
		}
		errAppRes := errapp.NewErrAppRes(err, "Erro no servidor", http.StatusInternalServerError)
		return nil, errAppRes
	}

	if err := s.repoGame.Find(number, model.StrategyRandom, g); err != nil {
		if errors.Is(err, errapp.ErrNoDocuments) {
			errAppRes := errapp.NewErrAppRes(err, errapp.ErrNoDocuments.Error(), http.StatusBadRequest)
			return nil, errAppRes
		}
		errAppRes := errapp.NewErrAppRes(err, "Erro no servidor", http.StatusInternalServerError)
		return nil, errAppRes
	}

	return g, nil
}

func (s *Service) GetGameRandom(number int) (*model.Game, *errapp.ErrAppRes) {
	g := &model.Game{}

	if err := s.repoGame.Find(number, model.StrategyRandom, g); err != nil {
		if errors.Is(err, errapp.ErrNoDocuments) {
			errAppRes := errapp.NewErrAppRes(err, errapp.ErrNoDocuments.Error(), http.StatusBadRequest)
			return nil, errAppRes
		}
		errAppRes := errapp.NewErrAppRes(err, "Erro no servidor", http.StatusInternalServerError)
		return nil, errAppRes
	}

	return g, nil
}

func (s *Service) GetAllGamesRandom() ([]model.Game, *errapp.ErrAppRes) {
	gamers := &[]model.Game{}

	if err := s.repoGame.FindAll(gamers); err != nil {
		if errors.Is(err, errapp.ErrNoDocuments) {
			errAppRes := errapp.NewErrAppRes(err, errapp.ErrNoDocuments.Error(), http.StatusBadRequest)
			return nil, errAppRes
		}
		errAppRes := errapp.NewErrAppRes(err, "Erro no servidor", http.StatusInternalServerError)
		return nil, errAppRes
	}

	gamersRandom := []model.Game{}

	for _, v := range *gamers {
		if v.TypeStrategy == model.StrategyRandom {
			gamersRandom = append(gamersRandom, v)
		}
	}

	return gamersRandom, nil
}

func (s *Service) UpdateScored(up *model.Scored) (*model.Scored, *errapp.ErrAppRes) {
	sc := &model.Scored{}

	if err := s.repoScored.Find(sc); err != nil {
		if errors.Is(err, errapp.ErrNoDocuments) {
			errAppRes := errapp.NewErrAppRes(err, errapp.ErrNoDocuments.Error(), http.StatusBadRequest)
			return nil, errAppRes
		}
		errAppRes := errapp.NewErrAppRes(err, "Erro no servidor", http.StatusInternalServerError)
		return nil, errAppRes
	}

	for i := range sc.Scores {
		if sc.Scores[i].Name == up.Scores[i].Name {
			sc.Scores[i].Score = sc.Scores[i].Score + up.Scores[i].Score
		}
	}

	if err := s.repoScored.Update(sc); err != nil {
		if errors.Is(err, errapp.ErrNoDocuments) {
			errAppRes := errapp.NewErrAppRes(err, errapp.ErrNoDocuments.Error(), http.StatusBadRequest)
			return nil, errAppRes
		}
		errAppRes := errapp.NewErrAppRes(err, "Erro no servidor", http.StatusInternalServerError)
		return nil, errAppRes
	}
	return sc, nil
}

func (s *Service) ConfirmResult(n int) (*model.ResponseConfirmResult, *errapp.ErrAppRes) {
	contest, errAppRes := s.lottery.GetNumberContest(n)
	if errAppRes != nil {
		return nil, errAppRes
	}

	gamers := &[]model.Game{}

	if err := s.repoGame.FindByNumber(n, gamers); err != nil {
		if errors.Is(err, errapp.ErrNoDocuments) {
			errAppRes := errapp.NewErrAppRes(err, errapp.ErrNoDocuments.Error(), http.StatusBadRequest)
			return nil, errAppRes
		}
		errAppRes := errapp.NewErrAppRes(err, "Erro no servidor", http.StatusInternalServerError)
		return nil, errAppRes
	}

	contestSystem := model.RequestCaixaSystem{
		Numero: contest.Numero,
	}

	for _, v := range contest.ListaDezenas {
		nInt, err := strconv.Atoi(v)
		if err != nil {
			return nil, errapp.NewErrAppRes(err, "Erro a converter string em int", 0)
		}
		contestSystem.ListaDezenas = append(contestSystem.ListaDezenas, nInt)
	}

	for _, v := range contestSystem.ListaDezenas {
		for _, g := range *gamers {
			for _, gn := range g.Numbers {
				if gn == v {
					g.Correct++
					g.CorrectNumbers = append(g.CorrectNumbers, v)
				}
			}
		}
	}

	for _, g := range *gamers {
		if err := s.repoGame.Update(g.Number, g.TypeStrategy, g.Correct, g.CorrectNumbers); err != nil {
			if errors.Is(err, errapp.ErrNoDocuments) {
				errAppRes := errapp.NewErrAppRes(err, errapp.ErrNoDocuments.Error(), http.StatusBadRequest)
				return nil, errAppRes
			}
			errAppRes := errapp.NewErrAppRes(err, "Erro no servidor", http.StatusInternalServerError)
			return nil, errAppRes
		}
	}

	slices.SortFunc(*gamers, func(a, b model.Game) int {
		if a.Correct != b.Correct {
			return cmp.Compare(b.Correct, a.Correct)
		}
		return cmp.Compare(b.Number, a.Number)
	})

	if len(*gamers) != 3 {
		errAppRes := errapp.NewErrAppRes(nil, "Erro no servidor", http.StatusInternalServerError)
		return nil, errAppRes
	}

	sc := &model.Scored{}
	if err := s.repoScored.Find(sc); err != nil {
		if errors.Is(err, errapp.ErrNoDocuments) {
			errAppRes := errapp.NewErrAppRes(err, errapp.ErrNoDocuments.Error(), http.StatusBadRequest)
			return nil, errAppRes
		}
		errAppRes := errapp.NewErrAppRes(err, "Erro no servidor", http.StatusInternalServerError)
		return nil, errAppRes
	}

	for i, v := range *gamers {
		if i == 0 {
			for in, s := range sc.Scores {
				if v.TypeStrategy == s.Name {
					sc.Scores[in].Score++
				}
			}
		}
	}

	if err := s.repoScored.Update(sc); err != nil {
		if errors.Is(err, errapp.ErrNoDocuments) {
			errAppRes := errapp.NewErrAppRes(err, errapp.ErrNoDocuments.Error(), http.StatusBadRequest)
			return nil, errAppRes
		}
		errAppRes := errapp.NewErrAppRes(err, "Erro no servidor", http.StatusInternalServerError)
		return nil, errAppRes
	}

	rc := &model.ResponseConfirmResult{
		Gamers: *gamers,
		Scores: sc.Scores,
	}

	return rc, nil

}
