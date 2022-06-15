# Server configuration

Tracetest can be configured using a config.yaml file placed on the same directory as its executable. It is useful to configure some aspects of how tracetest should behave. This section is dedicated to explain the options we currently have available.

## Configuration file example
```yaml
# Connection string to the postgres instance
postgresConnString: "host=localhost user=postgres password=postgres port=5432 sslmode=disable"

# Instance of jaeger that will be used to retrieve the trace of the service under test
jaegerConnectionConfig:
  endpoint: localhost:16685
  tls:
    insecure: true

# Configure how traces should be pooled from the tracing storage.
poolingConfig:
    # How long tracetest can wait for a trace to be complete? After this period, the pooling process will timeout
    # and the test will be marked as failed.
    maxWaitTimeForTrace: 90s

    # How much time tracetest should wait before trying to fetch the trace since the last execution? 
    retryDelay: 5s

# Server configuration
server:
  # Enables you to add a prefix to the server path. So, instead of running tracetest on http://localhost:8080, it would run on http://localhost:8080/tracetest instead.
  pathPrefix: /tracetest
  httpPort: 8080

# Google analytics configuration
googleAnalytics:
  enabled: false
  measurementId: ""
  secretKey: ""

# How tracetest should generate telemetry data.
telemetry:
  serviceName: tracetest
  sampling: 100
  jaeger:
    host: localhost
    port: 6831
  exporters:
    - console
    - jaeger
```

## Providing a configuration when running a container
```cmd
docker run --volume "`pwd`/my-config-file.yaml:/app/config.yaml" kubeshop/tracetest
```