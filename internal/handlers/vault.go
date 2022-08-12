package handlers

import (
	"encoding/json"
	"github.com/kubefirst/metaphor-go/configs"
	"github.com/kubefirst/metaphor-go/pkg"
	"github.com/rs/zerolog/log"
	"net/http"
)

type VaultHandler struct {
	config     *configs.Config
	httpClient pkg.HTTPDoer
}

type vaultResponse struct {
	SecretOne string `json:"secret_one"`
	SecretTwo string `json:"secret_two"`
}

func NewVault(config *configs.Config, httpClient pkg.HTTPDoer) VaultHandler {
	return VaultHandler{
		config:     config,
		httpClient: httpClient,
	}
}

func (handler VaultHandler) Vault(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if r.Method == http.MethodOptions {
		return
	}

	w.WriteHeader(http.StatusOK)

	vaultData := vaultResponse{
		SecretOne: handler.config.SecretOne,
		SecretTwo: handler.config.SecretTwo,
	}

	jsonData, err := json.Marshal(vaultData)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Err(err).Send()
		return
	}

	_, err = w.Write(jsonData)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Err(err).Send()
		return
	}
}
