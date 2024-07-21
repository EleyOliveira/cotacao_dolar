package main

import (
	"context"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080/cotacao", nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	log.Println(res.Body)
	gravaArquivo(res.Body)

}

func gravaArquivo(cotacao io.ReadCloser) {
	arq, err := os.Create("cotacao.txt")
	if err != nil {
		panic(err)
	}

	log.Println(cotacao)
	cotacaoByte, err := io.ReadAll(cotacao)
	if err != nil {
		panic(err)
	}
	_, err = arq.Write([]byte(cotacaoByte))
	if err != nil {
		panic(err)
	}
	arq.Close()
}
