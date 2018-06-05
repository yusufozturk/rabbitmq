package metricsRabbitMQ

// GetOverview returns stats `/api/overview/` from rabbitMQ HTTP API.
func GetOverview(debug bool, server string, auth string) (*[]map[string]interface{}, error) {
	keyPrefix := "overview."
	result := []map[string]interface{}{}
	respBody := map[string]interface{}{}

	path := "/api/overview/"
	err := ApiGet(debug, server, auth, path, &respBody)
	if err != nil {
		part := map[string]interface{}{}
		part["node"] = "unknown"
		part["name"] = "unknown"
		part[keyPrefix+"ping"] = -1
		result = append(result, part)
		return &result, err
	}

	part := map[string]interface{}{}
	node, ok := respBody["node"].(string)
	if ok {
		part["node"] = node
	} else {
		part["node"] = "unknown"
	}

	queueName, ok := respBody["name"].(string)
	if ok {
		part["name"] = queueName
	} else {
		part["name"] = "unknown"
	}

	GetValueFromMapSafe(debug, &keysOverview, &respBody, &part, keyPrefix)
	result = append(result, part)
	return &result, nil
}
