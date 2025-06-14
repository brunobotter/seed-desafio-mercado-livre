package data

import (
	"github.com/brunobotter/mercado-livre/configs/mapping"
	"github.com/brunobotter/mercado-livre/internal/data/datasql"
	"github.com/brunobotter/mercado-livre/internal/domain/contract"
)

func Connect(cfg *mapping.Config) (contract.DataManager, error) {
	return datasql.Instance(cfg)
}
