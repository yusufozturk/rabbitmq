package metricsRabbitMQ

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

func ApiGet(debug bool, urlServer string, auth string, path string, respBody interface{}) error {
	var err error
	var content []byte
	var resp *http.Response

	url := fmt.Sprintf("%s%s", urlServer, path)
	timeout := time.Duration(2 * time.Second)
	headerAuth := base64.StdEncoding.EncodeToString([]byte(auth))

	client := http.Client{
		Timeout: timeout,
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		goto FAILED
	}
	req.Header.Set("Authorization", fmt.Sprintf("Basic %s", strings.TrimSpace(headerAuth)))

	resp, err = client.Do(req)
	if err != nil {
		goto FAILED
	}

	defer resp.Body.Close()
	content, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		goto FAILED
	}

	err = json.Unmarshal(content, respBody)
	if err != nil {
		if debug {
			log.Println("json.Unmarshal failed", string(content))
		}
		goto FAILED
	}

	return nil

FAILED:
	return err
}

// GetValueFromMapSafe returns value from `map[]` by key, key compose in format 'a.b.c..'.
func GetValueFromMapSafe(debug bool, keys *[]string, m *map[string]interface{}, dst *map[string]interface{}, keyPrefix string) {
	for _, key := range *keys {
		subkeys := strings.Split(key, ".")
		if len(subkeys) > 1 {
			// access sub map
			// such as `message_stats.deliver_get` => obj["message_stats"]["deliver_get"]
			var value map[string]interface{}
			value, ok := (*m)[subkeys[0]].(map[string]interface{})
			if !ok {
				if debug {
					log.Println("get value failed, key", key)
				}
				continue
			}

			for idx, subkey := range subkeys[1:] {
				if idx == (len(subkeys[1:]) - 1) {
					(*dst)[keyPrefix+key] = value[subkey]
					continue
				}

				value, ok = value[subkey].(map[string]interface{})
				if !ok {
					if idx > (len(subkeys) - 1) {
						if debug {
							log.Println("get value failed, key", key)
						}
						continue
					} else {
						(*dst)[keyPrefix+key] = value[subkey]
					}
				}
			}

		} else {
			value, ok := (*m)[key]
			if !ok {
				if debug {
					log.Println("get value failed, key", key)
				}
			} else {
				(*dst)[keyPrefix+key] = value
			}
		}
	}
}
