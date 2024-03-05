package dql

import (
	"github.com/AmicieDeMonfourt/sqlparser/domain"
)

type Select struct {
	Expr    string
	From    From
	Join    Join
	Where   Where
	GrouBy  GroupBy
	Having  Having
	OrderBy OrderBy
	Limit   Limit
}

func (s *Select) Execute(set *domain.Set) (*domain.Set, error) {

	return set, nil
}
