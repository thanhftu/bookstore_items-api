package httputils

import (
	"encoding/json"
	"net/http"

	"github.com/thanhftu/bookstore_utils-go/resterrors"
)

// ResponseJSON response to user
func ResponseJSON(w http.ResponseWriter, statuscode int, body interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statuscode)
	json.NewEncoder(w).Encode(body)
}

// ResponseError response Err to user
func ResponseError(w http.ResponseWriter, err resterrors.RestErr) {
	ResponseJSON(w, err.Status(), err)
}
