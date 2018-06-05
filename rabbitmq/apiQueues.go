package metricsRabbitMQ

import (
	"errors"
)

// GetQueues returns stats `/api/queues/%2F/` (escape '/' in '%2F') from rabbitMQ HTTP API.
func GetQueues(debug bool, server string, auth string) (*[]map[string]interface{}, error) {
	keyPrefix := "queue."
	result := []map[string]interface{}{}
	respBody := []map[string]interface{}{}

	path := "/api/queues/%2F/"
	err := ApiGet(debug, server, auth, path, &respBody)
	if err == nil && len(respBody) == 0 {
		err = errors.New("got empty response")
	}
	if err != nil {
		part := map[string]interface{}{}
		part["node"] = "unknown"
		part["name"] = "unknown"
		part[keyPrefix+"ping"] = -1
		respBody = append(respBody, part)
		return &respBody, err
	}

	for _, queuesItem := range respBody {
		part := map[string]interface{}{}

		node, ok := queuesItem["node"].(string)
		if ok {
			part["node"] = node
		} else {
			part["node"] = "unknown"
		}

		queueName, ok := queuesItem["name"].(string)
		if ok {
			part["name"] = queueName
		} else {
			part["name"] = "unknown"
		}

		GetValueFromMapSafe(debug, &keysQueue, &queuesItem, &part, keyPrefix)
		result = append(result, part)
	}

	return &result, nil

}
