package handlers

import (
	"encoding/json"
	"github.com/kubefirst/metaphor-go/configs"
	"github.com/kubefirst/metaphor-go/pkg"
	"github.com/rs/zerolog/log"
	"net/http"
)

type KubernetesHandler struct {
	config     *configs.Config
	httpClient pkg.HTTPDoer
}

type kubernetesConfigMapResponse struct {
	ConfigOne string `json:"config_one"`
	ConfigTwo string `json:"config_two"`
}

func NewKubernetesHandler(config *configs.Config, httpClient pkg.HTTPDoer) KubernetesHandler {
	return KubernetesHandler{
		config:     config,
		httpClient: httpClient,
	}
}

func (handler KubernetesHandler) KubernetesConfigMapData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if r.Method == http.MethodOptions {
		return
	}

	w.WriteHeader(http.StatusOK)

	k8sConfigMapData := kubernetesConfigMapResponse{
		ConfigOne: handler.config.ConfigOne,
		ConfigTwo: handler.config.ConfigTwo,
	}

	jsonData, err := json.Marshal(k8sConfigMapData)
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
