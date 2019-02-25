## How to use the example

* Edit `/docker/srcipts/prometheus.yml` and change `myip` for your local ip address
* Check the port allocations in `prometheus.sh` and `/src/main.go` to make sure that they're not going to overlap with any ports that you already have exposed.
