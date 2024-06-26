package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type Cotacao struct {
	USDBRL struct {
		Code       string `json:"code"`
		Codein     string `json:"codein"`
		Name       string `json:"name"`
		High       string `json:"high"`
		Low        string `json:"low"`
		VarBid     string `json:"varBid"`
		PctChange  string `json:"pctChange"`
		Bid        string `json:"bid"`
		Ask        string `json:"ask"`
		Timestamp  string `json:"timestamp"`
		CreateDate string `json:"create_date"`
	} `json:"USDBRL"`
}

func main() {
	http.HandleFunc("/cotacao", BuscaCotacaoHandler)
	http.ListenAndServe(":8080", nil)
}

func BuscaCotacaoHandler(w http.ResponseWriter, r *http.Request) {
	cotacao, err := BuscaCotacao()
	if err != nil {
		log.Println("Erro ao buscar a cotação")
	}
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(cotacao)
}

func BuscaCotacao() (*Cotacao, error) {
	req, err := http.Get("https://economia.awesomeapi.com.br/json/last/USD-BRL")
	if err != nil {
		return nil, err
	}
	defer req.Body.Close()
	res, err := io.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}
	var data Cotacao
	err = json.Unmarshal(res, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}
