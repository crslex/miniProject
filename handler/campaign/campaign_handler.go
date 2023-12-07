package campaign

import (
	"net/http"

	"github.com/gorilla/mux"
)

type CampaignHandler struct {
}

func NewCampaignHandler() *CampaignHandler {
	return &CampaignHandler{}
}

func (h *CampaignHandler) RegisterRoutes(r *mux.Router) {
	// r.HandleFunc("/")
}

func (h *CampaignHandler) getCampaignByIDHandler(w http.ResponseWriter, r *http.Request) {

}
