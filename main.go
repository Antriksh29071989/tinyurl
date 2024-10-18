package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
	"tinyurl/handlers"
)

//TIP To run your code, right-click the code and select <b>Run</b>. Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.

var (
	// Define a Prometheus counter
	// Define a Prometheus counter for search requests
	searchCounter = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "http_search_request_counter_total",
			Help: "Total number of HTTP search requests.",
		},
	)

	// Define a Prometheus counter for create tiny url requests
	requestCounter = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "http_create_tiny_url_request_counter_total",
			Help: "Total number of HTTP requests to create tiny url.",
		},
	)
)

func init() {
	// Register the counter with the default registry
	prometheus.MustRegister(requestCounter)
	prometheus.MustRegister(searchCounter)
}
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		searchCounter.Inc()
		handlers.SearchHandler(w, r)
	})

	http.HandleFunc("/resource", func(w http.ResponseWriter, r *http.Request) {
		requestCounter.Inc()
		handlers.ResourceHandler(w, r)
	})
	http.Handle("/metrics", promhttp.Handler())
	// Start the server
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
