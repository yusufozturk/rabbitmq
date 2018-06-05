package metricsRabbitMQ

import (
	"log"
)

func Gets(debug bool, urlRabbitMQ, authRabbitMQ string) (*[]map[string]interface{}, error) {
	merge := []map[string]interface{}{}

	resultOverview, err := GetOverview(debug, urlRabbitMQ, authRabbitMQ)
	if err != nil {
		if debug {
			log.Println("GetOverview failed", err)
		}
	}

	for _, metric := range *resultOverview {
		merge = append(merge, metric)
	}

	resultQueues, err := GetQueues(debug, urlRabbitMQ, authRabbitMQ)
	if err != nil {
		if debug {
			log.Println("GetQueues failed", err)
		}
	}

	for _, metric := range *resultQueues {
		merge = append(merge, metric)

	}

	return &merge, nil
}
