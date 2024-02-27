package main

import (
	"bytes"
	"crypto"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

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

	client := http.Client{}

	req, err := http.NewRequest("POST", "https://act.serpro.gov.br:8444", bytes.NewReader(tsq))
	if err != nil {
		log.Fatal(err)
	}

	req.SetBasicAuth("85570540172", "serpro@123")
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

	fmt.Println("Carimbo realizado com sucesso")
	os.WriteFile("tsr_go", resp, 0777)
}
