newrelic_tigertonic
===================

NewRelic agent for [TigerTonic](https://github.com/rcrowley/go-tigertonic) inspired by [Gorelic](https://github.com/yvasiyarov/gorelic).  


Usage
-----

```go
agent := newrelic_tigertonic.NewAgent()
agent.Verbose = true
agent.NewrelicLicense = *newrelicLicense
agent.NewrelicName = "TigerTonic Example"
agent.Run()
```
