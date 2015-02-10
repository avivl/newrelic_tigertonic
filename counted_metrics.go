package newrelic_tigertonic

import (
	"strings"

	"github.com/rcrowley/go-metrics"
	"github.com/yvasiyarov/newrelic_platform_go"
)

// Base class for counter metrics
type baseCounterMetrica struct {
	dataSource    metrics.Counter
	name          string
	units         string
	previousValue float64
}

// Metric name
func (metrica *baseCounterMetrica) GetName() string {
	return metrica.name
}

// Metric value
func (metrica *baseCounterMetrica) GetUnits() string {
	return metrica.units
}

// 1minute running average metric
type statusCounterMetrica struct {
	*baseCounterMetrica
}

// Metric value. Previous value is substracted from the current value to get the real last minute value
func (metrica *statusCounterMetrica) GetValue() (float64, error) {
	currentValue := float64(metrica.dataSource.Count())
	value := currentValue - metrica.previousValue
	metrica.previousValue = currentValue
	return value, nil
}

// Helper func to add all our counter metrics to the plugin component
func addCounterMericsToComponent(component newrelic_platform_go.IComponent, counter metrics.Counter, counterName string) {

	// with CountedByStatus sampling, counterName looks like <prefix>-<status>
	// where prefix is the actual counter name, status is http status code
	strParts := strings.Split(counterName, "-")
	realCounterName := strParts[0]
	statusCode := strParts[1] // WARN: will panic on Counted sampling

	ctr := &statusCounterMetrica{
		baseCounterMetrica: &baseCounterMetrica{
			dataSource: counter,
			name:       realCounterName + "/http/status/" + statusCode,
			units:      "Count",
		},
	}
	component.AddMetrica(ctr)

}
