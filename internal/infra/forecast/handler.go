package forecast

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/emur-uy/backend/config"
	"github.com/emur-uy/backend/internal/pkg/entity"
	"github.com/getsentry/sentry-go"
)

// GetForecast fetches the forecast for the provided location parameters
// and returns the forecast data as an entity.RequestForecast object.
// It can return an error if the HTTP request fails, the HTTP status is not OK,
// or the response data cannot be unmarshaled to an entity.RequestForecast.
func GetForecast(lang string, state string, country string) (entity.RequestForecast, error) {
	apiKey := config.Get().ForecastKey
	apiUrl := config.Get().ForecastAPI

	escapeCountry := url.PathEscape(country)
	escapeState := url.PathEscape(state)

	res, err := http.Get(fmt.Sprintf("%slang=%s&key=%s&q=%s,%s&days=3", apiUrl, lang, apiKey, escapeState, escapeCountry))
	if err != nil {
		sentry.CaptureException(err)
		return entity.RequestForecast{}, fmt.Errorf("[GetForecast]: cannot fetch URL, %w", err)
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		errMsg := fmt.Sprintf("[GetForecast]: unexpected http GET status, %s", res.Status)
		sentry.CaptureMessage(errMsg)
		return entity.RequestForecast{}, errors.New(errMsg)
	}

	resBodyBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		sentry.CaptureException(err)
		return entity.RequestForecast{}, fmt.Errorf("[GetForecast]: error reading response body, %w", err)
	}

	var data entity.RequestForecast
	err = json.Unmarshal(resBodyBytes, &data)
	if err != nil {
		sentry.CaptureException(err)
		return entity.RequestForecast{}, fmt.Errorf("[GetForecast]: cannot decode JSON, %w", err)
	}

	return data, nil
}
