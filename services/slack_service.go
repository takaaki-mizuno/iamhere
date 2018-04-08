package services

import (
	"bytes"
	"net/http"
	"strings"
)

type slackService struct {
}

// SlackService : Service For Location
func SlackService() *slackService {
	return &slackService{}
}

// GetCountryCode : Get Country Code of current location by using ipinfo api
func (slackService *slackService) SetStatusIcon(countryCode string, token string) error {
	jsonStr := `{"profile":{"status_emoji":":flag-` + strings.ToLower(countryCode) + `:"}}`
	req, err := http.NewRequest(
		"POST",
		"https://slack.com/api/users.profile.set",
		bytes.NewBuffer([]byte(jsonStr)),
	)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", `Bearer `+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return err
}
