// Example: collect specify RabbitMQ instance metrics from HTTP API and dump to stdout in JSON format.
package main

import (
	"encoding/json"
	"flag"
	"fmt"

	"github.com/MonitorMetrics/rabbitmq/rabbitmq"
)

var (
	Debug    = flag.Bool("debug", false, "enable debug")
	Addr     = flag.String("addr", "http://127.0.0.1:15672", "instance address")
	Password = flag.String("auth", "guest:guest", "instance passsword")
)

func main() {
	flag.Parse()
	metrics, err := metricsRabbitMQ.Gets(*Debug, *Addr, *Password)
	if err != nil {
		panic(err)
	}

	out, err := json.MarshalIndent(metrics, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(out))
}
