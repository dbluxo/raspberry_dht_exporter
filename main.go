package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"net/http"

	"github.com/dbluxo/raspberry_dht_exporter/collector"
	"github.com/go-kit/kit/log/level"
	"github.com/pkg/errors"
	"github.com/prometheus/common/promlog"
	logflag "github.com/prometheus/common/promlog/flag"
	kingpin "gopkg.in/alecthomas/kingpin.v2"

	"log"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {

	var cfg struct {
		dhtPinNumber  int
		dhtSensorType string
	}

	a := kingpin.New(filepath.Base(os.Args[0]), "The Raspberry Pi DHT Prometheus Exporter")

	a.Flag("dht.pin-number", "GPIO pin number from which sensor data will be read.").
		Default("4").IntVar(&cfg.dhtPinNumber)

	a.Flag("dht.sensor-type", "DHT sensor type, one of: [dht11, dht22]").
		Default("dht22").StringVar(&cfg.dhtSensorType)

	var logcfg promlog.Config

	logflag.AddFlags(a, &logcfg)

	_, err := a.Parse(os.Args[1:])
	if err != nil {
		fmt.Fprintln(os.Stderr, errors.Wrapf(err, "Error parsing commandline arguments"))
		a.Usage(os.Args[1:])
		os.Exit(2)
	}

	logger := promlog.New(&logcfg)

	level.Info(logger).Log("msg", "Starting Exporter")
	level.Debug(logger).Log("msg", "GPIO pin number from which sensor data will be read: "+strconv.Itoa(cfg.dhtPinNumber))
	level.Debug(logger).Log("msg", "DHT sensor type: "+cfg.dhtSensorType)

	d := collector.NewDHTCollector()
	prometheus.MustRegister(d)

	http.Handle("/metrics", promhttp.Handler())
	level.Info(logger).Log("msg", "Beginning to serve on port :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
