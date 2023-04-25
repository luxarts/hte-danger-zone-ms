package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	HTTPGetPingCounter = promauto.NewCounter(prometheus.CounterOpts{
		Name: "http_get_ping_counter",
		Help: "Request to GET /ping",
	})
	HTTPResponseCounter = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "http_response_counter",
		Help: "Response to endpoints",
	}, []string{LabelStatusCode, LabelEndpoint, LabelMethod})
)
