package utils

import (
	"os"
	"net/http"
)

//  ──────────────────────────────────────────────────────────────────────────

func InternalErrorHandler(err error) {
	if err != nil {
		panic(err)
	}
}


func InternalHTTPErrorHandler(err error, w http.ResponseWriter) {
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		panic(err)
	}
}


func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)

    if err == nil { return true, nil }
    if os.IsNotExist(err) { return false, nil }

    return false, err
}


