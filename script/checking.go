package main

import (
	"crypto/tls"
	"encoding/json"
	"io"
	"log"
	"net/http"

	model "github.com/crslex/miniProject/model/campaign"
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

	// Check if indices exists
	// resp, _ := es_client.Indices.Exists([]string{"campaign_index"})

	// Check if documents exist in indices
	// resp, _ := es_client.Exists("campaign_index", "1")
	resp, _ := es_client.Get("campaign_index", "1")

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Failed to read response body")
	}

	// var campaign *model.Campaign
	var esResponse *model.CampaignElasticSearchResponse
	if err := json.Unmarshal(bodyBytes, &esResponse); err != nil {
		log.Println("Failed to unmarshal into campaign docs")
	}

	log.Println(esResponse.Source)
}
