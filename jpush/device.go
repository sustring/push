package jpush

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

type DeviceClient struct {
	*BaseClient
	url string
}

type Device struct {
	Tags   []string `json:"tags"`
	Alias  string   `json:"alias"`
	Mobile string   `json:"mobile"`
}

func (c DeviceClient) DeviceView(registrationId string) (*Device, error) {
	link := c.url + "/v3/devices/" + registrationId
	resp, err := c.Request("GET", link, nil, false)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() != http.StatusOK {
		return nil, errors.New(resp.Status())
	}
	var out Device
	err = json.Unmarshal(resp.Bytes(), &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

type DeviceSettingPayload struct {
	Tags   interface{} `json:"tags"` // empty string or DeviceSettingRequestTags
	Alias  string      `json:"alias"`
	Mobile string      `json:"mobile"`
}

type DeviceSettingRequestTags struct {
	Add    []string `json:"add,omitempty"`
	Remove []string `json:"remove,omitempty"`
}

func (c DeviceClient) DeviceSet(registrationId string, payload *DeviceSettingPayload) error {
	link := c.url + "/v3/devices/" + registrationId
	buf, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	resp, err := c.Request("POST", link, bytes.NewReader(buf), false)
	if err != nil {
		return err
	}
	if resp.StatusCode() != http.StatusOK {
		return errors.New(resp.Status())
	}
	return nil
}

type Alias struct {
	RegistrationIds []string `json:"registration_ids"`
}

func (c DeviceClient) AliasGet(alias string, platforms []string) (*Alias, error) {
	link := c.url + "/v3/aliases/" + alias
	if len(platforms) > 0 {
		link += "?platform=" + strings.Join(platforms, ",")
	}
	resp, err := c.Request("GET", link, nil, false)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() != http.StatusOK {
		return nil, errors.New(resp.Status())
	}
	var out Alias
	err = json.Unmarshal(resp.Bytes(), &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c DeviceClient) AliasDelete(alias string) error {
	link := c.url + "/v3/aliases/" + alias
	resp, err := c.Request("DELETE", link, nil, false)
	if err != nil {
		return err
	}
	if resp.StatusCode() != http.StatusOK {
		return errors.New(resp.Status())
	}
	return nil
}

func (c DeviceClient) AliasUnbind(alias string, registrationIds []string) error {
	link := c.url + "/v3/aliases/" + alias
	params := make(map[string]interface{})
	params["registration_ids"] = map[string][]string{"remove": registrationIds}
	buf, err := json.Marshal(params)
	if err != nil {
		return err
	}
	resp, err := c.Request("POST", link, bytes.NewReader(buf), false)
	if err != nil {
		return err
	}
	if resp.StatusCode() != http.StatusOK {
		return fmt.Errorf("%s", string(resp.Bytes()))
	}
	return nil
}

type Tags struct {
	Tags []string `json:"tags"`
}

func (c DeviceClient) TagsGet() (*Tags, error) {
	link := c.url + "/v3/tags/"
	resp, err := c.Request("GET", link, nil, false)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() != http.StatusOK {
		return nil, errors.New(resp.Status())
	}
	var out Tags
	err = json.Unmarshal(resp.Bytes(), &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

type TagCheckResult struct {
	Result bool `json:"result"`
}

func (c DeviceClient) TagCheck(tag, registrationId string) (bool, error) {
	link := c.url + "/v3/tags/" + tag + "/registration_ids/" + registrationId
	resp, err := c.Request("GET", link, nil, false)
	if err != nil {
		return false, err
	}
	if resp.StatusCode() != http.StatusOK {
		return false, errors.New(resp.Status())
	}
	var out TagCheckResult
	err = json.Unmarshal(resp.Bytes(), &out)
	if err != nil {
		return false, err
	}
	return out.Result, nil
}

type TagUpdatePayload struct {
	Add    []string `json:"add,omitempty"`
	Remove []string `json:"remove,omitempty"`
}

func (c DeviceClient) TagUpdate(tag string, payload *TagUpdatePayload) error {
	link := c.url + "/v3/tags/" + tag
	params := make(map[string]interface{})
	params["registration_ids"] = payload
	buf, err := json.Marshal(params)
	if err != nil {
		return err
	}
	resp, err := c.Request("POST", link, bytes.NewReader(buf), false)
	if err != nil {
		return err
	}
	if resp.StatusCode() != http.StatusOK {
		return errors.New(resp.Status())
	}
	return nil
}

func (c DeviceClient) TagDelete(tag string, platforms []string) error {
	link := c.url + "/v3/tags/" + tag
	if len(platforms) > 0 {
		link += "?platform=" + strings.Join(platforms, ",")
	}
	resp, err := c.Request("DELETE", link, nil, false)
	if err != nil {
		return err
	}
	if resp.StatusCode() != http.StatusOK {
		return errors.New(resp.Status())
	}
	return nil
}
