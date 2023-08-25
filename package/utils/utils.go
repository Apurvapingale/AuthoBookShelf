package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func ParseBody(r *http.Request, x interface{}) error {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, x)
	if err != nil {
		return err
	}
	return nil
}

func Response(w http.ResponseWriter, data map[string]interface{}) {
	w.Header().Set("Content-Type", "pkglication/json")
	json.NewEncoder(w).Encode(data)
	return
}