package util

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// ReadJSONRequestBody reads a http Request body and attempts to unmarshal it the object at v
func ReadJSONRequestBody(r *http.Request, v interface{}) error {
	requestContent, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(requestContent, v)
}

// WriteError writes an error back to responsewriter
func WriteError(trace string, w http.ResponseWriter) {
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte(trace))
}

// WriteJSONResponse writes a response to the response writer and sets the header as json
func WriteJSONResponse(trace string, resp interface{}, w http.ResponseWriter) {
	bytes, err := json.Marshal(resp)
	if err != nil {
		w.WriteHeader(http.StatusExpectationFailed)
		w.Write([]byte("Failure to marshal response interface on action : " + trace))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(bytes)
}

// AddIntSlices adds two int slices and returns the result
func AddIntSlices(a []int, b []int) []int {
	if len(a) < len(b) {
		c := a
		a = b
		b = c
	}
	ret := make([]int, len(a))
	var i int
	for i = range b {
		ret[i] = a[i] + b[i]
	}
	// because we need to increment by 1 prior to checking for same-length lists
	for i++; i < len(a); i++ {
		ret[i] = a[i]
	}
	return ret
}
