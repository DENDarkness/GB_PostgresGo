package httplib

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/spf13/viper"
)

var client = &http.Client{Timeout: 10 * time.Second}

// func getURL(addr, path string) string {
// 	return fmt.Sprintf("http://%s%s", addr, path)
// }

func GetJSON(url string, target *interface{}) error {

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", viper.GetString("client.token"))

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(target); err != nil {
		return err
	}

	return nil
}

func PostReq(url string) error {

	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", viper.GetString("client.token"))

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}
