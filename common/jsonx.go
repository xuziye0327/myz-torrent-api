package common

import "encoding/json"

// JsonToString returns json string of obj
func JsonToString(obj interface{}) string {
	bs, err := json.Marshal(obj)
	if err != nil {
		return err.Error()
	}

	return string(bs)
}
