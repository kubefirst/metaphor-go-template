package pkg

import "net/http"

type HTTPDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

const (
	AppName     = "metaphor-go"
	CompanyName = "Kubefirst"
)
