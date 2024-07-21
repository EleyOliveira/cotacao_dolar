package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"

	_ "github.com/mattn/go-sqlite3"
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
	SalvaCotacaoBD(cotacao.USDBRL.Bid)
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(cotacao)

}

func SalvaCotacaoBD(cotacao string) {
	var db *sql.DB
	db, err := sql.Open("sqlite3", "./dbcotacao.sqlite")
	if err != nil {
		log.Println(err)
	}
	defer db.Close()

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Microsecond*1)
	defer cancel()

	_, err = db.ExecContext(ctx, "INSERT INTO cotacao (valor) VALUES ($1)", cotacao)
	if err != nil {
		log.Println(err)
	}

}

func BuscaCotacao() (*Cotacao, error) {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*200)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, "GET", "https://economia.awesomeapi.com.br/json/last/USD-BRL", nil)
	if err != nil {
		log.Println("Erro ao criar requisição")
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
	}
	var data Cotacao
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&data); err != nil {
		return nil, err
	}

	return &data, nil
}
