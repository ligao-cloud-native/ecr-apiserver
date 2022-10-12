package metrics

import (
	"flag"
	"github.com/prometheus/client_golang/prometheus"
	"math"
	"math/rand"
	"time"
)

type Metrics struct {
	opsTotal     *prometheus.CounterVec
	rpcDurations *prometheus.SummaryVec
}

func NewMetrics() *Metrics {
	//创建 Prometheus 数据Metric
	ops := prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "myapp_processed_ops_total",
		Help: "The total number of processed events",
	}, []string{"appname", "action"})

	rpcDurations := prometheus.NewSummaryVec(
		prometheus.SummaryOpts{
			Name:       "rpc_durations_seconds",
			Help:       "RPC latency distributions.这个metric的帮助信息,metric的项目作用说明",
			Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
		},
		[]string{"service"},
	)

	//注册定义好的Metric
	prometheus.MustRegister(
		ops,
		rpcDurations)

	return &Metrics{
		opsTotal:     ops,
		rpcDurations: rpcDurations,
	}

}

func (m *Metrics) HandleMetrics() {
	go func() {
		for {
			m.opsTotal.WithLabelValues("myapp", "install").Inc()
			time.Sleep(2 * time.Second)
		}
	}()

	oscillationPeriod := flag.Duration("oscillation-period", 10*time.Minute, "The duration of the rate oscillation period.")
	start := time.Now()
	oscillationFactor := func() float64 {
		return 2 + math.Sin(math.Sin(2*math.Pi*float64(time.Since(start))/float64(*oscillationPeriod)))
	}
	go func() {
		for {
			v := rand.ExpFloat64() / 1e6
			m.rpcDurations.WithLabelValues("exponential").Observe(v)
			time.Sleep(time.Duration(50*oscillationFactor()) * time.Millisecond)
		}
	}()
}
