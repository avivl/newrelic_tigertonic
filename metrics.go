package newrelic_tigertonic

import (
	"time"

	"github.com/rcrowley/go-metrics"
	"github.com/yvasiyarov/newrelic_platform_go"
)

type baseTimerMetrica struct {
	dataSource metrics.Timer
	name       string
	units      string
}

func (metrica *baseTimerMetrica) GetName() string {
	return metrica.name
}

func (metrica *baseTimerMetrica) GetUnits() string {
	return metrica.units
}

type timerRate1Metrica struct {
	*baseTimerMetrica
}

func (metrica *timerRate1Metrica) GetValue() (float64, error) {
	return metrica.dataSource.Rate1(), nil
}

type timerRateMeanMetrica struct {
	*baseTimerMetrica
}

func (metrica *timerRateMeanMetrica) GetValue() (float64, error) {
	return metrica.dataSource.RateMean(), nil
}

type timerMeanMetrica struct {
	*baseTimerMetrica
}

func (metrica *timerMeanMetrica) GetValue() (float64, error) {
	return metrica.dataSource.Mean() / float64(time.Millisecond), nil
}

type timerMinMetrica struct {
	*baseTimerMetrica
}

func (metrica *timerMinMetrica) GetValue() (float64, error) {
	return float64(metrica.dataSource.Min()) / float64(time.Millisecond), nil
}

type timerMaxMetrica struct {
	*baseTimerMetrica
}

func (metrica *timerMaxMetrica) GetValue() (float64, error) {
	return float64(metrica.dataSource.Max()) / float64(time.Millisecond), nil
}

type timerPercentile75Metrica struct {
	*baseTimerMetrica
}

func (metrica *timerPercentile75Metrica) GetValue() (float64, error) {
	return metrica.dataSource.Percentile(0.75) / float64(time.Millisecond), nil
}

type timerPercentile90Metrica struct {
	*baseTimerMetrica
}

func (metrica *timerPercentile90Metrica) GetValue() (float64, error) {
	return metrica.dataSource.Percentile(0.90) / float64(time.Millisecond), nil
}

type timerPercentile95Metrica struct {
	*baseTimerMetrica
}

func (metrica *timerPercentile95Metrica) GetValue() (float64, error) {
	return metrica.dataSource.Percentile(0.95) / float64(time.Millisecond), nil
}

func addHTTPMericsToComponent(component newrelic_platform_go.IComponent, timer metrics.Timer, timerName string) {
	rate1 := &timerRate1Metrica{
		baseTimerMetrica: &baseTimerMetrica{
			name:       timerName + "/" + "http/throughput/1minute",
			units:      "rps",
			dataSource: timer,
		},
	}
	component.AddMetrica(rate1)
	rateMean := &timerRateMeanMetrica{
		baseTimerMetrica: &baseTimerMetrica{
			name:       timerName + "/" + "http/throughput/rateMean",
			units:      "rps",
			dataSource: timer,
		},
	}
	component.AddMetrica(rateMean)

	responseTimeMean := &timerMeanMetrica{
		baseTimerMetrica: &baseTimerMetrica{
			name:       timerName + "/" + "http/responseTime/mean",
			units:      "ms",
			dataSource: timer,
		},
	}
	component.AddMetrica(responseTimeMean)

	responseTimeMax := &timerMaxMetrica{
		baseTimerMetrica: &baseTimerMetrica{
			name:       timerName + "/" + "http/responseTime/max",
			units:      "ms",
			dataSource: timer,
		},
	}
	component.AddMetrica(responseTimeMax)

	responseTimeMin := &timerMinMetrica{
		baseTimerMetrica: &baseTimerMetrica{
			name:       timerName + "/" + "http/responseTime/min",
			units:      "ms",
			dataSource: timer,
		},
	}
	component.AddMetrica(responseTimeMin)

	responseTimePercentile75 := &timerPercentile75Metrica{
		baseTimerMetrica: &baseTimerMetrica{
			name:       timerName + "/" + "http/responseTime/percentile75",
			units:      "ms",
			dataSource: timer,
		},
	}
	component.AddMetrica(responseTimePercentile75)

	responseTimePercentile90 := &timerPercentile90Metrica{
		baseTimerMetrica: &baseTimerMetrica{
			name:       timerName + "/" + "http/responseTime/percentile90",
			units:      "ms",
			dataSource: timer,
		},
	}
	component.AddMetrica(responseTimePercentile90)

	responseTimePercentile95 := &timerPercentile95Metrica{
		baseTimerMetrica: &baseTimerMetrica{
			name:       timerName + "/" + "http/responseTime/percentile95",
			units:      "ms",
			dataSource: timer,
		},
	}
	component.AddMetrica(responseTimePercentile95)
}
