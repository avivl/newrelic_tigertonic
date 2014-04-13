package newrelic_tigertonic

import (
	"github.com/rcrowley/go-metrics"
	"github.com/yvasiyarov/newrelic_platform_go"
)

type baseCounterMetrica struct {
	dataSource metrics.Counter
	name       string
	units      string
}

func (metrica *baseCounterMetrica) GetName() string {
	return metrica.name
}

func (metrica *baseCounterMetrica) GetUnits() string {
	return metrica.units
}

func (metrica *baseCounterMetrica) GetValue() (float64, error) {
	return float64(metrica.dataSource.Count()), nil
}

func addCounterMericsToComponent(component newrelic_platform_go.IComponent, counter metrics.Counter, counterName string) {
	ctr := &baseCounterMetrica{
		dataSource: counter,
		name:       counterName,
		units:      "Count",
	}
	component.AddMetrica(ctr)
}
