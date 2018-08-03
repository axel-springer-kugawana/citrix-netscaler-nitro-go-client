package nitro

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func responseHandler(resp *http.Response) ([]byte, error) {
	switch resp.Status {
	case "201 Created", "200 OK":
		body, _ := ioutil.ReadAll(resp.Body)
		return body, nil
	case "409 Conflict":
		body, _ := ioutil.ReadAll(resp.Body)
		return body, errors.New("failed: " + resp.Status + " (" + string(body) + ")")

	case "207 Multi Status":
		//This happens in case of Bulk operations, which we do not support yet
		body, _ := ioutil.ReadAll(resp.Body)
		return body, nil
	case "400 Bad Request", "401 Unauthorized", "403 Forbidden",
		"404 Not Found", "405 Method Not Allowed", "406 Not Acceptable",
		"503 Service Unavailable", "599 Netscaler specific error":
		body, _ := ioutil.ReadAll(resp.Body)
		log.Println("[INFO] go-nitro: error = " + string(body))
		return body, errors.New("failed: " + resp.Status + " (" + string(body) + ")")
	default:
		body, err := ioutil.ReadAll(resp.Body)
		return body, err

	}
}

func (c *NitroClient) createHTTPRequest(method string, url string, buff *bytes.Buffer) (*http.Request, error) {
	req, err := http.NewRequest(method, url, buff)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	if c.proxiedNs == "" {
		req.Header.Set("X-NITRO-USER", c.username)
		req.Header.Set("X-NITRO-PASS", c.password)
	} else {
		req.SetBasicAuth(c.username, c.password)
		req.Header.Set("_MPS_API_PROXY_MANAGED_INSTANCE_IP", c.proxiedNs)
	}
	return req, nil
}

func (c *NitroClient) doHTTPRequest(method string, url string, bytes *bytes.Buffer) ([]byte, error) {
	req, err := c.createHTTPRequest(method, c.url+url, bytes)

	resp, err := c.client.Do(req)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		return []byte{}, err
	}

	log.Println("[DEBUG] go-nitro: response Status:", resp.Status)

	ret, err := responseHandler(resp)

	return ret, err
}

func make_url(resourceType string, resourceId string, qs map[string]string) string {
	url := "/nitro/v1/config/" + resourceType

	if len(resourceId) > 0 {
		url = url + "/" + resourceId
	}

	if len(qs) > 0 {
		var args []string

		for key, value := range qs {
			if len(value) > 0 {
				args = append(args, key+"="+value)
			}
		}

		if len(args) > 0 {
			url = url + "?" + strings.Join(args, "&")
		}
	}

	return url
}

func (c *NitroClient) get(resourceType string, resourceId string, qs map[string]string, payload interface{}) error {
	url := make_url(resourceType, resourceId, qs)

	_, err := c.doHTTPRequest("GET", url, bytes.NewBuffer([]byte{}))

	return err
}

func (c *NitroClient) put(resourceType string, resourceId string, qs map[string]string, payload interface{}) error {
	url := make_url(resourceType, resourceId, qs)

	if body, err := json.Marshal(payload); err != nil {
		return err
	} else {
		_, err := c.doHTTPRequest("PUT", url, bytes.NewBuffer(body))

		return err
	}
}

func (c *NitroClient) post(resourceType string, resourceId string, qs map[string]string, payload interface{}) error {
	url := make_url(resourceType, resourceId, qs)

	if body, err := json.Marshal(payload); err != nil {
		return err
	} else {
		_, err := c.doHTTPRequest("POST", url, bytes.NewBuffer(body))

		return err
	}
}

func (c *NitroClient) delete(resourceType string, resourceId string, qs map[string]string) error {
	url := make_url(resourceType, resourceId, qs)

	_, err := c.doHTTPRequest("DELETE", url, bytes.NewBuffer([]byte{}))

	return err
}
