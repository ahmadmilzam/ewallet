# Go Ewallet App

A REST API service for ewallet

## How to
### Start a Live Reload Server
```
$ air
```

### Start a Server
```
$ make run
```

### Create a Migration File
```
$ make migrate-create FILENAME=your_migration_name
```

### Run a Migration UP/DOWN
```
$ make migrate-up

$ make migrate-down
```

## Features
- [x] Package manager: Go mod
- [x] Env config: [Viper](https://github.com/spf13/viper)
- [x] Router: [Gin](https://github.com/gin-gonic/gin)
- [x] Logger: [Zap](https://github.com/uber-go/zap)
- [x] CLI: [Urfave/CLI](https://github.com/urfave/cli)
- [x] Unit test: [Testify](https://github.com/stretchr/testify)
- [x] StatsD instrumentation: [Datadog](https://github.com/DataDog/datadog-go)
- [x] Tracing: [Datadog](https://github.com/DataDog/dd-trace-go)
- [ ] API Documentation: [TBA](https://github.com/)
