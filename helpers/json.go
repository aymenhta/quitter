package helpers

import (
	"encoding/json"
	"net/http"

	"github.com/aymenhta/quitter/config"
)

func EncodeRes(w http.ResponseWriter, res any) {
	encoder := json.NewEncoder(w)
	err := encoder.Encode(res)
	if err != nil {
		config.G.ErrorLog.Println(err)
		http.Error(w, "Could not marshal json", http.StatusInternalServerError)
		return
	}
}

func DecodeReq(w http.ResponseWriter, r *http.Request, dst any) {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(dst)
	if err != nil {
		config.G.ErrorLog.Println(err)
		http.Error(w, "Could not decode request's body", http.StatusInternalServerError)
		return
	}
}
