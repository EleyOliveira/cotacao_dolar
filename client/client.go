package main

import (
	"context"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

type valorDolar struct {
	Valor string `json:"valor"`
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080/cotacao", nil)
	if err != nil {
		log.Println(err)
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	gravaArquivo(res.Body)

}

func gravaArquivo(cotacao io.ReadCloser) {
	arq, err := os.Create("cotacao.txt")
	if err != nil {
		panic(err)
	}

	cotacaoByte, err := io.ReadAll(cotacao)
	if err != nil {
		panic(err)
	}

	posicao, err := arq.Write([]byte("Dólar:"))
	if err != nil {
		panic(err)
	}

	_, err = arq.WriteAt([]byte(cotacaoByte), int64(posicao))
	if err != nil {
		panic(err)
	}
	arq.Close()
}
