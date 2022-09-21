package usecase

import (
	"errors"
	"math/rand"

	"github.com/fenriz07/100-prisoners-riddle-GOLANG/app/entities"
)

const (
	MaxOfAll      = 100
	MaxToAttempts = 51
)

func GameUseCase() error {

	//rand.Seed(time.Now().Unix())

	prisioners := make([]entities.Prisioner, MaxOfAll)
	boxes := make(entities.Box, MaxOfAll)

	secretValues := rand.Perm(MaxOfAll)

	for k, v := range secretValues {
		prisioners[k].Value = k
		boxes[k] = v
	}

	for _, p := range prisioners {

		boxToSearch := p.Value

		for i := 0; i < MaxToAttempts; i++ {
			//Open the box
			valueIntoBox := boxes[boxToSearch]

			//If the value is the same to the prisioner break the search.
			if valueIntoBox == p.Value {
				break
			}

			boxToSearch = valueIntoBox

			if i == 50 {
				return errors.New("a prision failed")
			}
		}

	}

	return nil
}
