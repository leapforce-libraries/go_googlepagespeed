package googlepagespeed

import (
	"fmt"
	"io/ioutil"
	"net/url"

	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
)

type Category string

const (
	CategoryUnspecified   Category = "CATEGORY_UNSPECIFIED"
	CategoryAccessibility Category = "ACCESSIBILITY"
	CategoryBestPractices Category = "BEST_PRACTICES"
	CategoryPerfomance    Category = "PERFORMANCE"
	CategoryPWA           Category = "PWA"
	CategorySEO           Category = "SEO"
)

type Strategy string

const (
	StrategyUnspecified Strategy = "STRATEGY_UNSPECIFIED"
	StrategyDesktop     Strategy = "DESKTOP"
	StrategyMobile      Strategy = "MOBILE"
)

type RunPageSpeedConfig struct {
	Category     *Category
	Locale       *string
	Strategy     *Strategy
	URL          string
	UTMCampaign  *string
	UTMSource    *string
	CaptchaToken *string
}

func (service *Service) RunPageSpeed(config *RunPageSpeedConfig) (*[]byte, *errortools.Error) {
	if config == nil {
		return nil, errortools.ErrorMessage("RunPageSpeedConfig must not be a nil pointer")
	}

	values := url.Values{}
	values.Set("url", config.URL)

	if config.Category != nil {
		values.Set("category", string(*config.Category))
	}

	if config.Locale != nil {
		values.Set("locale", string(*config.Locale))
	}

	if config.Strategy != nil {
		values.Set("strategy", string(*config.Strategy))
	}

	if config.UTMCampaign != nil {
		values.Set("utm_campaign", string(*config.UTMCampaign))
	}

	if config.UTMSource != nil {
		values.Set("utm_source", string(*config.UTMSource))
	}

	if config.CaptchaToken != nil {
		values.Set("captchaToken", string(*config.CaptchaToken))
	}

	requestConfig := go_http.RequestConfig{
		URL:           service.url(fmt.Sprintf("runPagespeed/?%s", values.Encode())),
		ResponseModel: nil,
	}
	fmt.Println(service.url(fmt.Sprintf("runPagespeed/?%s", values.Encode())))

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
