package handler

import (
	"encoding/json"
	"log"
	"net/http"
)

type responseJson struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

func (m *Repository) SearchAvailability(w http.ResponseWriter, r *http.Request) {
	remoteIp := m.App.Session.GetString(r.Context(), "remote_ip")
	log.Println("Search Availability Get Page from ", remoteIp)

	resp := responseJson{
		Status:  true,
		Message: "Available!",
	}

	out, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		log.Println("Error: ", err)
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(out)
	if err != nil {
		log.Println("Error: ", err)
	}
}
