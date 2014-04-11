// Metrics output to newrelic.
// based on https://github.com/yvasiyarov/gorelic
package newrelic_tigertonic

import (
	"errors"
	"log"

	"github.com/rcrowley/go-metrics"
	"github.com/yvasiyarov/newrelic_platform_go"
)

const (
	// DefaultNewRelicPollInterval - how often we will report metrics to NewRelic.
	// Recommended values is 60 seconds
	DefaultNewRelicPollInterval = 10

	// DefaultPollIntervalInSeconds - how often we will get  statistic
	// Default value is - every 10 seconds
	DefaultPollIntervalInSeconds = 10

	//DefaultAgentGuid is plugin ID in NewRelic.
	//You should not change it unless you want to create your own plugin.
	DefaultAgentGuid = "com.rounds.TigerTonic"

	//CurrentAgentVersion is plugin version
	CurrentAgentVersion = "0.0.1"

	//DefaultAgentName in NewRelic GUI. You can change it.
	DefaultAgentName = "TigerTonic"
)

//Agent - is NewRelic agent implementation.
//Agent start separate go routine which will report data to NewRelic
type Agent struct {
	NewrelicName         string
	NewrelicLicense      string
	NewrelicPollInterval int
	PollInterval         int
	Verbose              bool
	AgentGUID            string
	AgentVersion         string
	plugin               *newrelic_platform_go.NewrelicPlugin
}

//NewAgent build new Agent objects.
func NewAgent() *Agent {
	agent := &Agent{
		NewrelicName:         DefaultAgentName,
		NewrelicPollInterval: DefaultNewRelicPollInterval,
		Verbose:              false,
		PollInterval:         DefaultPollIntervalInSeconds,
		AgentGUID:            DefaultAgentGuid,
		AgentVersion:         CurrentAgentVersion,
	}
	return agent
}

//Run initialize Agent instance and start harvest go routine
func (agent *Agent) Run() error {
	if agent.NewrelicLicense == "" {
		return errors.New("please, pass a valid newrelic license key")
	}

	agent.plugin = newrelic_platform_go.NewNewrelicPlugin(agent.AgentVersion, agent.NewrelicLicense, agent.NewrelicPollInterval)
	component := newrelic_platform_go.NewPluginComponent(agent.NewrelicName, agent.AgentGUID)
	agent.plugin.AddComponent(component)
	registry := metrics.DefaultRegistry
	registry.Each(func(name string, i interface{}) {
		switch metric := i.(type) {
		case metrics.Timer:
			addTimerMericsToComponent(component, metric, name)
		case metrics.Counter:
			addCounterMericsToComponent(component, metric, name)

		}
	})

	agent.plugin.Verbose = agent.Verbose
	go agent.plugin.Run()
	return nil
}

//Print debug messages
func (agent *Agent) debug(msg string) {
	if agent.Verbose {
		log.Println(msg)
	}
}
