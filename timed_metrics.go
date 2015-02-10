package newrelic_tigertonic

import (
	"time"

	"github.com/rcrowley/go-metrics"
	"github.com/yvasiyarov/newrelic_platform_go"
)

// Base class for timer metrics
type baseTimerMetrica struct {
	dataSource metrics.Timer
	name       string
	units      string
}

// Metric name
func (metrica *baseTimerMetrica) GetName() string {
	return metrica.name
}

// Metric units
func (metrica *baseTimerMetrica) GetUnits() string {
	return metrica.units
}

// 1minute running average metric
type timerRate1Metrica struct {
	*baseTimerMetrica
}

// Metric value
func (metrica *timerRate1Metrica) GetValue() (float64, error) {
	return metrica.dataSource.Rate1(), nil
}

// Mean rate metrica (from server start!)
type timerRateMeanMetrica struct {
	*baseTimerMetrica
}

func (metrica *timerRateMeanMetrica) GetValue() (float64, error) {
	return metrica.dataSource.RateMean(), nil
}

// Real 1minute mean rate metrica
type timer1MinRateMeanMetrica struct {
	*baseTimerMetrica
	previousCount float64
	previousTime  time.Time
}

// Real metrica that calculates last minute RPS without taking into account the history data
func (metrica *timer1MinRateMeanMetrica) GetValue() (float64, error) {
	currentTime := time.Now()
	currentCount := float64(metrica.dataSource.Count())
	value := (currentCount - metrica.previousCount) / currentTime.Sub(metrica.previousTime).Seconds()
	metrica.previousCount = currentCount
	metrica.previousTime = currentTime
	return value, nil
}

// Mean value metrica (from server start)
type timerMeanMetrica struct {
	*baseTimerMetrica
}

func (metrica *timerMeanMetrica) GetValue() (float64, error) {
	return metrica.dataSource.Mean() / float64(time.Millisecond), nil
}

// Min value metrica (from server start)
type timerMinMetrica struct {
	*baseTimerMetrica
}

func (metrica *timerMinMetrica) GetValue() (float64, error) {
	return float64(metrica.dataSource.Min()) / float64(time.Millisecond), nil
}

// Max value metrica (from server start)
type timerMaxMetrica struct {
	*baseTimerMetrica
}

func (metrica *timerMaxMetrica) GetValue() (float64, error) {
	return float64(metrica.dataSource.Max()) / float64(time.Millisecond), nil
}

// 75 percentile value metrica (from server start)
type timerPercentile75Metrica struct {
	*baseTimerMetrica
}

func (metrica *timerPercentile75Metrica) GetValue() (float64, error) {
	return metrica.dataSource.Percentile(0.75) / float64(time.Millisecond), nil
}

// 90 percentile metrica (from server start)
type timerPercentile90Metrica struct {
	*baseTimerMetrica
}

func (metrica *timerPercentile90Metrica) GetValue() (float64, error) {
	return metrica.dataSource.Percentile(0.90) / float64(time.Millisecond), nil
}

// 95 percentile metrica (from server start)
type timerPercentile95Metrica struct {
	*baseTimerMetrica
}

func (metrica *timerPercentile95Metrica) GetValue() (float64, error) {
	return metrica.dataSource.Percentile(0.95) / float64(time.Millisecond), nil
}

// Helper func to add all our timer metrics to the plugin component
func addTimerMericsToComponent(component newrelic_platform_go.IComponent, timer metrics.Timer, timerName string) {
	rate1 := &timerRate1Metrica{
		baseTimerMetrica: &baseTimerMetrica{
			name:       timerName + "/" + "http/throughput/1minute",
			units:      "rps",
			dataSource: timer,
		},
	}
	component.AddMetrica(rate1)

	realRate1Minute := &timer1MinRateMeanMetrica{
		baseTimerMetrica: &baseTimerMetrica{
			name:       timerName + "/" + "http/throughput/realRate1Min",
			units:      "rps",
			dataSource: timer,
		},
	}
	component.AddMetrica(realRate1Minute)

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
