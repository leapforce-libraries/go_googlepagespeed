package googlepagespeed

import (
	"fmt"
	"net/http"
	"net/url"

	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
)

const (
	apiURL string = "https://www.googleapis.com/pagespeedonline/v5"
)

// Service stores Service configuration
//
type Service struct {
	apiKey      string
	httpService *go_http.Service
}

type ServiceConfig struct {
	APIKey string
}

// methods
//
func NewService(serviceConfig *ServiceConfig) (*Service, *errortools.Error) {
	if serviceConfig == nil {
		return nil, errortools.ErrorMessage("ServiceConfig must not be a nil pointer")
	}

	if serviceConfig.APIKey == "" {
		return nil, errortools.ErrorMessage("APIKey not provided")
	}

	httpService, e := go_http.NewService(&go_http.ServiceConfig{})
	if e != nil {
		return nil, e
	}

	return &Service{
		apiKey:      serviceConfig.APIKey,
		httpService: httpService,
	}, nil
}

func (service *Service) httpRequest(httpMethod string, requestConfig *go_http.RequestConfig) (*http.Request, *http.Response, *errortools.Error) {
	// add API key
	_url, err := url.Parse(requestConfig.URL)
	if err != nil {
		return nil, nil, errortools.ErrorMessage(err)
	}
	query := _url.Query()
	query.Set("key", service.apiKey)

	(*requestConfig).URL = fmt.Sprintf("%s://%s%s?%s", _url.Scheme, _url.Host, _url.Path, query.Encode())

	// add error model
	errorResponse := ErrorResponse{}
	(*requestConfig).ErrorModel = &errorResponse

	request, response, e := service.httpService.HTTPRequest(httpMethod, requestConfig)
	if errorResponse.Error.Message != "" {
		e.SetMessage(errorResponse.Error.Message)
	}

	return request, response, e
}

func (service *Service) url(path string) string {
	return fmt.Sprintf("%s/%s", apiURL, path)
}

func (service *Service) get(requestConfig *go_http.RequestConfig) (*http.Request, *http.Response, *errortools.Error) {
	return service.httpRequest(http.MethodGet, requestConfig)
}
