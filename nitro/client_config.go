package nitro

import (
	"bytes"
	"encoding/json"
)

func (c *NitroClient) deleteResource(resourceType string, name string) error {
	return c.deleteResourceWithArgs(resourceType, name, "")
}

func (c *NitroClient) deleteResourceWithArgs(resourceType string, name string, args string) error {
	url := "/nitro/v1/config/" + resourceType + "/" + name

	if len(args) > 0 {
		url = url + "?args=" + args
	}

	_, err := c.doHTTPRequest("DELETE", url, bytes.NewBuffer([]byte{}))

	return err
}

func (c *NitroClient) getResource(resourceType string, name string, results interface{}) error {
	return c.getResourceWithArgs(resourceType, name, "", results)
}

func (c *NitroClient) getResourceWithArgs(resourceType string, name string, args string, results interface{}) error {
	url := "/nitro/v1/config/" + resourceType + "/" + name

	if len(args) > 0 {
		url = url + "?args=" + args
	}

	bytes, err := c.doHTTPRequest("GET", url, bytes.NewBuffer([]byte{}))

	if err != nil {
		return err
	}

	return json.Unmarshal(bytes, results)
}

func (c *NitroClient) listResources(resourceType string, results interface{}) error {
	url := "/nitro/v1/config/" + resourceType

	bytes, err := c.doHTTPRequest("GET", url, bytes.NewBuffer([]byte{}))

	if err != nil {
		return err
	}

	return json.Unmarshal(bytes, results)
}

func (c *NitroClient) addResource(resourceType string, resource interface{}) error {
	url := "/nitro/v1/config/" + resourceType

	payload := map[string]interface{}{
		resourceType: resource,
	}

	body, err := json.Marshal(payload)

	if err != nil {
		return err
	}

	_, err = c.doHTTPRequest("POST", url, bytes.NewBuffer(body))

	return err
}

func (c *NitroClient) renameResource(resourceType string, nameField string, name string, newName string) error {
	url := "/nitro/v1/config/" + resourceType + "?action=rename"

	resource := map[string]string{
		nameField: name,
		"newName": newName,
	}

	payload := map[string]interface{}{
		resourceType: resource,
	}

	body, err := json.Marshal(payload)

	if err != nil {
		return err
	}

	_, err = c.doHTTPRequest("POST", url, bytes.NewBuffer(body))

	return err
}

func (c *NitroClient) unsetResource(resourceType string, nameField string, name string, fields []string) error {
	url := "/nitro/v1/config/" + resourceType + "?action=unset"

	resource := map[string]interface{}{
		nameField: name,
	}

	for _, field := range fields {
		resource[field] = true
	}

	payload := map[string]interface{}{
		resourceType: resource,
	}

	body, err := json.Marshal(payload)

	if err != nil {
		return err
	}

	_, err = c.doHTTPRequest("POST", url, bytes.NewBuffer(body))

	return err
}

func (c *NitroClient) enableResource(resourceType string, nameField string, name string) error {
	url := "/nitro/v1/config/" + resourceType + "?action=enable"

	resource := map[string]interface{}{
		nameField: name,
	}

	payload := map[string]interface{}{
		resourceType: resource,
	}

	body, err := json.Marshal(payload)

	if err != nil {
		return err
	}

	_, err = c.doHTTPRequest("POST", url, bytes.NewBuffer(body))

	return err
}

func (c *NitroClient) disableResource(resourceType string, nameField string, name string) error {
	url := "/nitro/v1/config/" + resourceType + "?action=disable"

	resource := map[string]interface{}{
		nameField: name,
	}

	payload := map[string]interface{}{
		resourceType: resource,
	}

	body, err := json.Marshal(payload)

	if err != nil {
		return err
	}

	_, err = c.doHTTPRequest("POST", url, bytes.NewBuffer(body))

	return err
}
