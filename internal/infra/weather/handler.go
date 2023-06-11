package weather

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/emur-uy/backend/config"
	"github.com/emur-uy/backend/internal/pkg/entity"
	"github.com/getsentry/sentry-go"
)

func GetWeatherApi(lang string, state string, country string) entity.RequestForecast {
	apiKey := config.Get().ForecastKey
	apiUrl := config.Get().ForecastAPI

	escapeCountry := url.PathEscape(country)
	escapeState := url.PathEscape(state)

	res, err := http.Get(fmt.Sprintf("%slang=%s&key=%s&q=%s&%s&days=3", apiUrl, lang, apiKey, escapeState, escapeCountry))

	if err != nil {
		errMsg := "cannot fetch URL"
		sentry.CaptureMessage(fmt.Sprintf("[GetWeatherAPI]: %s, %s", errMsg, err))
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		errMsg := "unexpected http GET status"
		sentry.CaptureMessage(fmt.Sprintf("[GetWeatherAPI]: %s, %s", errMsg, res.Status))
	}

	var data entity.RequestForecast

	// We could check the resulting content type
	resBodyBytes, err := ioutil.ReadAll(res.Body)
	resBodyString := string(resBodyBytes)

	err = json.Unmarshal([]byte(resBodyString), &data)
	if err != nil {
		errMsg := "cannot decode JSON:"
		sentry.CaptureMessage(fmt.Sprintf("[GetWeatherAPI]: %s, %v", errMsg, err))
	}

	return data
}
