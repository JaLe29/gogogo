package metrics

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Labels struct {
	Method string
	Route  string
	Status string
}

type Activity struct {
	Ip      string
	ProxyId string
	Host    string
}

type Metrics interface {
	HandlerExecutionTime(Labels) prometheus.Observer
	HandleActivity(activity Activity)
	Handler() http.Handler
}

type metrics struct {
	handlerExecutionTime *prometheus.HistogramVec
	handleActivity       *prometheus.CounterVec
}

func New() Metrics {
	m := &metrics{
		handlerExecutionTime: promauto.NewHistogramVec(prometheus.HistogramOpts{
			Name:    "bp_proxy_handler_execution_time_ms",
			Help:    "Total execution time of handler",
			Buckets: []float64{10, 20, 30, 50, 80, 130, 210, 340},
		}, []string{"method", "route", "status"}),
		handleActivity: promauto.NewCounterVec(prometheus.CounterOpts{
			Name: "bp_proxy_activity",
			Help: "Total activity",
		}, []string{"ip", "proxyId", "host"}),
	}

	return m
}

func (m *metrics) HandlerExecutionTime(labels Labels) prometheus.Observer {
	return m.handlerExecutionTime.With(prometheus.Labels{
		"method": labels.Method,
		"route":  labels.Route,
		"status": labels.Status,
	})
}

func (m *metrics) HandleActivity(activity Activity) {
	m.handleActivity.With(prometheus.Labels{
		"ip":      activity.Ip,
		"proxyId": activity.ProxyId,
		"host":    activity.Host,
	}).Inc()
}

func (m *metrics) Handler() http.Handler {
	return promhttp.Handler()
}
