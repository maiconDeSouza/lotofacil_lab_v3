package lottery

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"math/rand/v2"
	"net/http"
	"slices"
	"strconv"
	"time"

	"github.com/maiconDeSouza/lotofacil_lab_v3/internal/env"
	"github.com/maiconDeSouza/lotofacil_lab_v3/internal/errapp"
	"github.com/maiconDeSouza/lotofacil_lab_v3/internal/model"
)

type Lottery struct {
	env    *env.Env
	client *http.Client
}

func NewLottery(env *env.Env) *Lottery {
	return &Lottery{
		env: env,
		client: &http.Client{
			Timeout: time.Duration(5) * time.Second,
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			},
		},
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

func (l *Lottery) GetLastContest() (*model.RequestCaixa, *errapp.ErrAppRes) {
	req, err := http.NewRequest(http.MethodGet, l.env.URL, nil)
	if err != nil {
		errapp := errapp.NewErrAppRes(err, "Erro ao criar requisição", 0)
		return nil, errapp
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36")
	req.Header.Set("Accept", "application/json")

	res, err := l.client.Do(req)
	if err != nil {
		errapp := errapp.NewErrAppRes(err, "Erro ao acessar a API", 0)
		return nil, errapp
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		errapp := errapp.NewErrAppRes(nil, fmt.Sprintf(
			"Erro ao acessar a API retornou status %d, mas era esperado %d", res.StatusCode, http.StatusOK,
		), res.StatusCode,
		)
		return nil, errapp
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		errapp := errapp.NewErrAppRes(err, "Erro ao ler resposta da API", 0)
		return nil, errapp
	}

	contest := model.RequestCaixa{}
	if err := json.Unmarshal(body, &contest); err != nil {
		errapp := errapp.NewErrAppRes(err, "API retornou json inesperado", 0)
		return nil, errapp
	}

	return &contest, nil
}

func (l *Lottery) GetNumberContest(c int) (*model.RequestCaixa, *errapp.ErrAppRes) {
	url := fmt.Sprintf("%s/%d", l.env.URL, c)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		errapp := errapp.NewErrAppRes(err, "Erro ao criar requisição", 0)
		return nil, errapp
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36")
	req.Header.Set("Accept", "application/json")

	res, err := l.client.Do(req)
	if err != nil {
		errapp := errapp.NewErrAppRes(err, "Erro ao acessar a API", 0)
		return nil, errapp
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		errapp := errapp.NewErrAppRes(nil, fmt.Sprintf(
			"Erro ao acessar a API retornou status %d, mas era esperado %d", res.StatusCode, http.StatusOK,
		), res.StatusCode,
		)
		return nil, errapp
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		errapp := errapp.NewErrAppRes(err, "Erro ao ler resposta da API", 0)
		return nil, errapp
	}

	contest := model.RequestCaixa{}
	if err := json.Unmarshal(body, &contest); err != nil {
		errapp := errapp.NewErrAppRes(err, "API retornou json inesperado", 0)
		return nil, errapp
	}

	return &contest, nil
}

func (l *Lottery) GetFrequentNumbers() (*model.ResultFrequency, *errapp.ErrAppRes) {
	concurso, errRes := l.GetLastContest()
	if errRes != nil {
		return nil, errRes
	}

	m := make(map[string]int)
	for x := 1; x <= 25; x++ {
		m[fmt.Sprintf("%02d", x)] = 0
	}

	numero := concurso.Numero
	limite := numero - 29 // Busca os últimos 30 concursos (inclusivo)

	for x := numero; x >= limite; x-- {
		c, errRes := l.GetNumberContest(x)
		if errRes != nil {
			return nil, errRes
		}
		for _, n := range c.ListaDezenas {
			m[n]++
		}
	}

	ranking := l.ordenarFrequencia(m)

	palpite, err := l.generateFrequent(ranking)
	if err != nil {
		// Ajuste a criação do ErrApp conforme o construtor do seu pacote errapp
		return nil, errapp.NewErrAppRes(err, "Deu ruim", 0)
	}

	pInt := []int{}
	for _, v := range palpite {
		i, err := strconv.Atoi(v)
		if err != nil {
			return nil, errapp.NewErrAppRes(err, "Erro a converter string em int", 0)
		}
		pInt = append(pInt, i)
	}

	return &model.ResultFrequency{
		Palpite: pInt,
		Ranking: ranking,
	}, nil
}

func (l *Lottery) ordenarFrequencia(m map[string]int) []model.DezFrequency {
	ranking := make([]model.DezFrequency, 0, len(m))

	for num, qtd := range m {
		ranking = append(ranking, model.DezFrequency{
			Numero:     num,
			Quantidade: qtd,
		})
	}

	slices.SortFunc(ranking, func(a, b model.DezFrequency) int {
		if a.Quantidade != b.Quantidade {
			return b.Quantidade - a.Quantidade
		}
		if a.Numero < b.Numero {
			return -1
		}
		if a.Numero > b.Numero {
			return 1
		}
		return 0
	})

	return ranking
}

func (l *Lottery) generateFrequent(numbers []model.DezFrequency) ([]string, error) {
	if len(numbers) != 25 {
		return nil, fmt.Errorf("quantidade inválida de dezenas: esperado 25, recebido %d", len(numbers))
	}

	// Define as 5 faixas e quantos números sortear de cada uma (5+4+3+2+1 = 15)
	tiers := []struct {
		items    []model.DezFrequency
		quantity int
	}{
		{numbers[0:5], 5},
		{numbers[5:10], 4},
		{numbers[10:15], 3},
		{numbers[15:20], 2},
		{numbers[20:25], 1},
	}

	var result []string

	for _, t := range tiers {
		var pool []string
		for _, item := range t.items {
			pool = append(pool, item.Numero)
		}

		drawn, err := l.drawNumbers(pool, t.quantity)
		if err != nil {
			return nil, fmt.Errorf("erro ao sortear faixa: %w", err)
		}
		result = append(result, drawn...)
	}

	if len(result) != 15 {
		return nil, fmt.Errorf("erro ao gerar números: total diferente de 15")
	}

	// Ordena o palpite final em ordem crescente ("01", "02", "05"...)
	slices.Sort(result)

	return result, nil
}

func (l *Lottery) drawNumbers(source []string, quantity int) ([]string, error) {
	if quantity > len(source) {
		return nil, fmt.Errorf("quantidade solicitada (%d) é maior que o número de elementos disponíveis (%d)", quantity, len(source))
	}
	if quantity <= 0 {
		return nil, fmt.Errorf("a quantidade a sortear deve ser maior que zero")
	}

	copiedNumbers := make([]string, len(source))
	copy(copiedNumbers, source)

	rand.Shuffle(len(copiedNumbers), func(i, j int) {
		copiedNumbers[i], copiedNumbers[j] = copiedNumbers[j], copiedNumbers[i]
	})

	return copiedNumbers[:quantity], nil
}
