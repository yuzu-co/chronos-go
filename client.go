package chronos

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// HTTP Transport for Chronos API
// Exposed purposefully to allow setting certificates etc
var Client *http.Client

// Chronos API Endpoint
var Host string = os.Getenv("CHRONOS_URL")

func init() {
	Client = &http.Client{}
}

// Issue GET request against API and unmarshal response to v
func Get(uri string, v interface{}) error {

	res, err := http.Get(Host + uri)
	if err != nil {
		panic(err.Error())
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err.Error())
	}

	err = json.Unmarshal(body, &v)
	if err != nil {
		panic(err.Error())
	}

	return nil
}

// Issue POST request against API and unmarshal response to v
func Post(uri string, body interface{}) error {
	data, err := json.Marshal(body)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", Host+uri, bytes.NewBuffer(data))
	if err != nil {
		return err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	res, err := Client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != 204 {
		return fmt.Errorf("got %d", res.StatusCode)
	}

	return nil
}

// Issue PUT request against API
func Put(uri string) error {
	req, err := http.NewRequest("PUT", Host+uri, nil)
	if err != nil {
		return err
	}

	res, err := Client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != 204 {
		return fmt.Errorf("got %d", res.StatusCode)
	}

	return nil
}

// Issue DELETE request against API
func Delete(uri string) error {
	req, err := http.NewRequest("DELETE", Host+uri, nil)
	if err != nil {
		return err
	}

	res, err := Client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != 204 {
		return fmt.Errorf("got %d", res.StatusCode)
	}

	return nil
}
