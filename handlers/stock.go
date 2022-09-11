package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/izaakdale/service-dynamo/dao"
)

type StoreResponse struct {
	Message string `json:"message,omitempty"`
}

func (s *Service) StoreStockPrice(w http.ResponseWriter, r *http.Request) {

	var req dao.StockIndex
	json.NewDecoder(r.Body).Decode(&req)

	if req.Attribute == "daily" {
		req.Attribute = "DAILY"
		req.Attribute += "#"
		req.Attribute += time.Now().Format("2006-01-02")
	} else if req.Attribute == "latest" {
		req.Attribute = "LATEST"
		req.Attribute += "#"
		req.Attribute += time.Now().Format("2006-01-02 15:04:05")
	}

	s.DBClient.StoreStockPrice(req)

	payload, err := json.Marshal(StoreResponse{
		Message: "stored",
	})
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(payload)
}
