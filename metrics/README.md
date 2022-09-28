# Metrics (prometheus)

The metrics integration makes use of the `"github.com/prometheus/client_golang/prometheus` package. It uses the `prometheus.DefaultRegisterer`.

So you can add metrics as simple as in the following code example:

```go
var currentRequests = promauto.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: "fiber_requests_current",
		Help: "The current number of active requests",
	},
	[]string{"method", "path"},
)
```

## UseFiberEndpoint

The `UseFiberEndpoint` adds the `/metrics` route to your fiber app.
