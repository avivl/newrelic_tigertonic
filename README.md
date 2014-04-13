newrelic_tigertonic
===================

NewRelic agent for [TigerTonic](https://github.com/rcrowley/go-tigertonic) inspired by [Gorelic](https://github.com/yvasiyarov/gorelic).  


Usage
-----

Install dependencies:

```sh
sh bootstrap.sh
```
Then define your service.  The working [example](https://github.com/rounds/newrelic_tigertonic/tree/master/example) may be a more convenient place to start.


```go
agent := newrelic_tigertonic.NewAgent()
agent.Verbose = true
agent.NewrelicLicense = *newrelicLicense
agent.NewrelicName = "TigerTonic Example"
agent.Run()
```
