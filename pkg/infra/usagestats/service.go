package usagestats

import (
	"context"
	"time"

	"github.com/maksimmernikov/grafana/pkg/bus"
	"github.com/maksimmernikov/grafana/pkg/login/social"
	"github.com/maksimmernikov/grafana/pkg/services/sqlstore"

	"github.com/maksimmernikov/grafana/pkg/log"
	"github.com/maksimmernikov/grafana/pkg/registry"
	"github.com/maksimmernikov/grafana/pkg/setting"
)

var metricsLogger log.Logger = log.New("metrics")

func init() {
	registry.RegisterService(&UsageStatsService{})
}

type UsageStatsService struct {
	Cfg      *setting.Cfg       `inject:""`
	Bus      bus.Bus            `inject:""`
	SQLStore *sqlstore.SqlStore `inject:""`

	oauthProviders map[string]bool
}

func (uss *UsageStatsService) Init() error {

	uss.oauthProviders = social.GetOAuthProviders(uss.Cfg)
	return nil
}

func (uss *UsageStatsService) Run(ctx context.Context) error {
	uss.updateTotalStats()

	onceEveryDayTick := time.NewTicker(time.Hour * 24)
	everyMinuteTicker := time.NewTicker(time.Minute)
	defer onceEveryDayTick.Stop()
	defer everyMinuteTicker.Stop()

	for {
		select {
		case <-onceEveryDayTick.C:
			uss.sendUsageStats(uss.oauthProviders)
		case <-everyMinuteTicker.C:
			uss.updateTotalStats()
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}
