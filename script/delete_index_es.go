package main

import (
	"crypto/tls"
	"net/http"

	"github.com/elastic/go-elasticsearch/v8"
)

func main() {
	cfg := elasticsearch.Config{
		Addresses: []string{
			"https://localhost:9200",
		},
		Username: "elastic",
		Password: "s4QAc6N6JzHm6TFKu9**",
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}

	es_client, _ := elasticsearch.NewClient(cfg)

	es_client.Indices.Delete([]string{"campaign_index"})

}
