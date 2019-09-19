package collector

import (
	"log"

	dht "github.com/d2r2/go-dht"
	"github.com/prometheus/client_golang/prometheus"
)

type dhtCollector struct {
	temperatureMetric *prometheus.Desc
	humidityMetric    *prometheus.Desc
}

func NewDHTCollector() *dhtCollector {
	return &dhtCollector{
		temperatureMetric: prometheus.NewDesc("dht_temperature",
			"Returns the measured temperature",
			nil, nil,
		),
		humidityMetric: prometheus.NewDesc("dht_humidity",
			"Returns the measured humidity",
			nil, nil,
		),
	}
}

func (collector *dhtCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- collector.temperatureMetric
	ch <- collector.humidityMetric
}

func (collector *dhtCollector) Collect(ch chan<- prometheus.Metric) {

	var temperatureValue float64
	var humidityValue float64

	sensorType := dht.DHT22

	temperature, humidity, _, err :=
		dht.ReadDHTxxWithRetry(sensorType, 4, false, 10)
	if err != nil {
		log.Fatal(err)
	}

	temperatureValue = float64(temperature)
	humidityValue = float64(humidity)

	ch <- prometheus.MustNewConstMetric(collector.temperatureMetric, prometheus.GaugeValue, temperatureValue)
	ch <- prometheus.MustNewConstMetric(collector.humidityMetric, prometheus.GaugeValue, humidityValue)

}
