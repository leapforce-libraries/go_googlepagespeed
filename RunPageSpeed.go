package googlepagespeed

import (
	"fmt"
	"io/ioutil"
	"net/url"

	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
)

type RunPageSpeedConfig struct {
	URL string
}

func (service *Service) RunPageSpeed(config *RunPageSpeedConfig) (*[]byte, *errortools.Error) {
	if config == nil {
		return nil, errortools.ErrorMessage("RunPageSpeedConfig must not be a nil pointer")
	}

	values := url.Values{}
	values.Set("url", config.URL)

	requestConfig := go_http.RequestConfig{
		URL: service.url(fmt.Sprintf("runPagespeed/?%s", values.Encode())),
	}

	_, response, e := service.get(&requestConfig)
	if e != nil {
		return nil, e
	}

	if response == nil {
		return nil, errortools.ErrorMessage("Nil response returned")
	}
	if response.Body == nil {
		return nil, errortools.ErrorMessage("Nil response body returned")
	}

	defer response.Body.Close()

	b, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, errortools.ErrorMessage(err)
	}

	return &b, nil
}
