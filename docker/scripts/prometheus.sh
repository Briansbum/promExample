#!/usr/bin/env bash

docker run -d --rm --name=prometheus \
              --publish=9090:9090 \
              --volume="${PWD}/prometheus.yml:/etc/prometheus/prometheus.yml" \
              prom/prometheus
