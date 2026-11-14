package lottery

import (
	"math/rand/v2"
	"slices"

	"github.com/maiconDeSouza/lotofacil_lab_v3/internal/env"
)

type Lottery struct {
	env *env.Env
}

func NewLottery(env *env.Env) *Lottery {
	return &Lottery{
		env: env,
	}
}

func (s *Lottery) GenerateLotofacilNumbers15() []int {
	l := rand.Perm(25)

	l = l[:15]

	for i := range l {
		l[i]++
	}

	slices.Sort(l)

	return l
}
