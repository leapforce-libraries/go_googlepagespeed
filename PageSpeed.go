package googlepagespeed

import (
	"fmt"
	"net/http"
	"net/url"

	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
)

type Category string

const (
	CategoryUnspecified   Category = "CATEGORY_UNSPECIFIED"
	CategoryAccessibility Category = "ACCESSIBILITY"
	CategoryBestPractices Category = "BEST_PRACTICES"
	CategoryPerformance   Category = "PERFORMANCE"
	CategoryPwa           Category = "PWA"
	CategorySeo           Category = "SEO"
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
	Url          string
	UtmCampaign  *string
	UtmSource    *string
	CaptchaToken *string
}

type PageSpeed struct {
	LighthouseResult struct {
		Categories struct {
			Performance struct {
				Score float64 `json:"score"`
			} `json:"performance"`
		} `json:"categories"`
	} `json:"lighthouseResult"`
}

func (service *Service) RunPageSpeed(config *RunPageSpeedConfig) (*PageSpeed, *errortools.Error) {
	if config == nil {
		return nil, errortools.ErrorMessage("RunPageSpeedConfig must not be a nil pointer")
	}

	values := url.Values{}
	values.Set("url", config.Url)

	if config.Category != nil {
		values.Set("category", string(*config.Category))
	}

	if config.Locale != nil {
		values.Set("locale", string(*config.Locale))
	}

	if config.Strategy != nil {
		values.Set("strategy", string(*config.Strategy))
	}

	if config.UtmCampaign != nil {
		values.Set("utm_campaign", string(*config.UtmCampaign))
	}

	if config.UtmSource != nil {
		values.Set("utm_source", string(*config.UtmSource))
	}

	if config.CaptchaToken != nil {
		values.Set("captchaToken", string(*config.CaptchaToken))
	}

	pageSpeed := PageSpeed{}

	requestConfig := go_http.RequestConfig{
		Method:        http.MethodGet,
		Url:           service.url(fmt.Sprintf("runPagespeed/?%s", values.Encode())),
		ResponseModel: &pageSpeed,
	}

	_, _, e := service.httpRequest(&requestConfig)
	if e != nil {
		return nil, e
	}

	return &pageSpeed, nil
}
