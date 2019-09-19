# WIP:  Raspberry Pi DHT Prometheus Exporter 

### Build
```
go build -v -i -o ./bin/raspberry_dht_exporter
```

### Install as a service
```
sudo cp dht.service /lib/systemd/system/
sudo systemctl start dht
sudo systemctl status dht
sudo systemctl enable dht

```

### Manual usage
```
Usage: dht_exporter [ ... ]

Parameters:

  -device int
    	Sensor type, either 11 or 22 for DHT11/DHT22 (default 22)
  -listen-address string
    	Address on which to expose metrics. (default ":9330")
  -names string
    	File mapping GPIOs to names (default "names.yaml")
  -path string
    	Path under which to expose metrics. (default "/metrics")
  -version
    	Print version information.
```
# Todo:

- Parse all parameters correctly