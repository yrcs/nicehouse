package do

import (
	"github.com/yrcs/nicehouse/pkg/usecase"
)

type Role struct {
	usecase.BaseDO
	Name        string
	Description string
	IsSystem    bool
}
