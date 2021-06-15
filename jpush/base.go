package jpush

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"runtime"
)

type BaseClient struct {
	AppKey            string
	MasterSecret      string
	GroupKey          string
	GroupMasterSecret string
}

func (c BaseClient) GetAuthorization(isGroup bool) string {
	str := c.AppKey + ":" + c.MasterSecret
	if isGroup {
		str = "group-" + c.GroupKey + ":" + c.GroupMasterSecret
	}
	buf := []byte(str)
	return fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString(buf))
}

func (c BaseClient) GetUserAgent() string {
	return fmt.Sprintf("(%s) go/%s", runtime.GOOS, runtime.Version())
}

func (c BaseClient) Request(method, link string, body io.Reader, isGroup bool) (*Response, error) {
	req, err := http.NewRequest(method, link, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", c.GetAuthorization(isGroup))
	req.Header.Set("User-Agent", c.GetUserAgent())
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return &Response{statusCode: resp.StatusCode, status: resp.Status, data: buf}, nil
}

type Response struct {
	statusCode int
	status     string
	data       []byte
}

func (r Response) Array() ([]interface{}, error) {
	list := make([]interface{}, 0)
	err := json.Unmarshal(r.data, &list)
	return list, err
}

func (r Response) Map() (map[string]interface{}, error) {
	result := make(map[string]interface{})
	err := json.Unmarshal(r.data, &result)
	return result, err
}

func (r Response) Bytes() []byte {
	return r.data
}

func (r Response) StatusCode() int {
	return r.statusCode
}

func (r Response) Status() string {
	return r.status
}
