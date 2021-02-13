package restapi

import (
	"encoding/json"
	"net/http"
	"time"
)

//Version ...
type Version struct {
	ID        string `json:"id"`
	CreatedAt string `json:"createdAt"`
}

//GetVersion ...
func GetVersion(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&Version{
		ID:        "0.0.1",
		CreatedAt: time.Now().UTC().String(),
	})
}
