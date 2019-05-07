package sqlstore

import (
	"github.com/maksimmernikov/grafana/pkg/bus"
	m "github.com/maksimmernikov/grafana/pkg/models"
)

func init() {
	bus.AddHandler("sql", GetDBHealthQuery)
}

func GetDBHealthQuery(query *m.GetDBHealthQuery) error {
	return x.Ping()
}
