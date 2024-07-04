package main

import (
	"PaaS/db"
	"PaaS/environment"
	"PaaS/routes"
	"log"
	"net/http"
	_ "net/http/pprof"
	"runtime"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/shirou/gopsutil/cpu"
)

var (
	cpuUsage = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "cpu_usage",
		Help: "Current CPU usage.",
	})
	memoryUsage = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "memory_usage",
		Help: "Current memory usage.",
	})
)

func init() {
	prometheus.MustRegister(cpuUsage)
	prometheus.MustRegister(memoryUsage)
}

func recordMetrics() {
	go func() {
		for {
			var memStats runtime.MemStats
			runtime.ReadMemStats(&memStats)
			memoryUsage.Set(float64(memStats.Alloc))

			percentages, err := cpu.Percent(0, false)
			if err != nil {
				log.Printf("Error retrieving CPU usage: %v", err)
				continue
			}

			if len(percentages) > 0 {
				cpuUsage.Set(percentages[0])
			}

			time.Sleep(2 * time.Second)
		}
	}()
}

func main() {
	recordMetrics()

	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()
	environment.InitEnv()

	db.InitializePostgres()

	//db.ApplyMigrations()

	r := routes.SetupRouter()

	r.Run(":8080")
	log.Println("Server is running on port 8080")
}
