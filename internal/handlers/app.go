package handlers

import (
	"encoding/json"
	"github.com/kubefirst/metaphor-go/configs"
	"github.com/kubefirst/metaphor-go/pkg"
	"github.com/rs/zerolog/log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

type AppHandler struct {
	config     *configs.Config
	httpClient pkg.HTTPDoer
}

type appInfoResponse struct {
	AppName      string `json:"app_name"`
	CompanyName  string `json:"company_name"`
	ChartVersion string `json:"chart_version"`
	DockerTag    string `json:"docker_tag"`
}

func NewApp(config *configs.Config, httpClient pkg.HTTPDoer) *AppHandler {
	return &AppHandler{
		config:     config,
		httpClient: httpClient,
	}
}

func (handler AppHandler) Info(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if r.Method == http.MethodOptions {
		return
	}

	w.WriteHeader(http.StatusOK)

	appInfo := appInfoResponse{
		AppName:      pkg.AppName,
		CompanyName:  pkg.CompanyName,
		ChartVersion: handler.config.ChartVersion,
		DockerTag:    handler.config.DockerTag,
	}
	jsonData, err := json.Marshal(appInfo)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
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

func (handler AppHandler) Performance(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if r.Method == http.MethodOptions {
		return
	}
	sleepTime := rand.Intn(10)
	log.Debug().Msgf("I'm sleeping for %d seconds", sleepTime)
	time.Sleep(time.Duration(sleepTime) * time.Second)
	w.WriteHeader(http.StatusOK)
}

func (handler AppHandler) Healthz(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if r.Method == http.MethodOptions {
		return
	}
	log.Debug().Msg("I'm health!")
	w.WriteHeader(http.StatusOK)
}

func (handler AppHandler) Kill(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if r.Method == http.MethodOptions {
		return
	}
	log.Warn().Str("endpoint", "kill").Msg("Kill endpoint was called, and the application is being force terminated...")
	os.Exit(1)
}
