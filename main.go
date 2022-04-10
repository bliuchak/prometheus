package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

func main() {
	// general metric to expose on /metrics endpoint inside app
	c := promauto.NewCounter(
		prometheus.CounterOpts{
			Name:        "up",
			Help:        "help text 123",
			ConstLabels: prometheus.Labels{"job": "api"},
		},
	)
	c.Add(1)

	http.ListenAndServe(":80", promhttp.Handler())
}

//func main() {
//	// pushgateway only metric
//	g := prometheus.NewGauge(prometheus.GaugeOpts{
//		Name: "hello_world",
//		Help: "some text",
//	})
//
//	p := push.New("pushgateway:9091", "push_api_job").Collector(g)
//
//	for {
//		select {
//		case <-time.Tick(5 * time.Second):
//			g.Inc()
//			err := p.Push()
//			if err != nil {
//				panic(err)
//			}
//
//			log.Println("push metrics to pg")
//		}
//	}
//}
