package main

import (
	"bytes"
	"crypto"
	"crypto/tls"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/digitorus/timestamp"
)

func main() {
	file, _ := os.Open("go-data.txt")
	tsq, err := timestamp.CreateRequest(file, &timestamp.RequestOptions{
		Hash:         crypto.SHA256,
		Certificates: true,
	})

	if err != nil {
		log.Fatal(err)
	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	sumCarimbo := 0
	totalCarimbo := 10

	for sumCarimbo < totalCarimbo {
		sumCarimbo++
		req, err := http.NewRequest("POST", "https://act.serpro.gov.br:8444", bytes.NewReader(tsq))
		if err != nil {
			log.Fatal(err)
		}

		req.SetBasicAuth("[[CPF-AQUI]]", "[[SENHA-AQUI]]")
		req.Header.Add("content-type", "application/timestamp-query")

		tsr, err := client.Do(req)
		if err != nil {
			log.Fatal(err)
		}

		if tsr.StatusCode > 200 {
			log.Fatal(tsr.Status)
		}

		resp, err := io.ReadAll(tsr.Body)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Carimbo realizado com sucesso", sumCarimbo)
		fileResp := "tsr_go_" + strconv.Itoa(sumCarimbo)
		os.WriteFile(fileResp, resp, 0644)
	}
}
