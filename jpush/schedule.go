package jpush

import (
	"bytes"
	"encoding/json"
	"strconv"
)

type ScheduleClient struct {
	*BaseClient
	url string
}

const (
	ScheduleTimeUnitDay   = "day"
	ScheduleTimeUnitWeek  = "week"
	ScheduleTimeUnitMonth = "month"
)

type SchedulePayload struct {
	Cid     string       `json:"cid,omitempty"`
	Name    string       `json:"name"`
	Enabled bool         `json:"enabled"`
	Trigger *Trigger     `json:"trigger"`
	Push    *PushPayload `json:"push"`
}

type Trigger struct {
	Single     *TriggerSingle     `json:"single,omitempty"`
	Periodical *TriggerPeriodical `json:"periodical,omitempty"`
}

type TriggerSingle struct {
	Timer string `json:"time,omitempty"`
}

type TriggerPeriodical struct {
	Start     string      `json:"start,omitempty"`
	End       string      `json:"end,omitempty"`
	Time      string      `json:"time,omitempty"`
	TimeUnit  string      `json:"time_unit,omitempty"`
	Frequency int         `json:"frequency,int,omitempty"`
	Point     interface{} `json:"point,omitempty"`
}

func (c PushClient) ScheduleCreateTask(req *SchedulePayload) (map[string]interface{}, error) {
	link := c.url + "/v3/schedules"
	buf, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	resp, err := c.Request("POST", link, bytes.NewReader(buf), false)
	if err != nil {
		return nil, err
	}
	return resp.Map()
}

func (c PushClient) ScheduleGetList(page int) (map[string]interface{}, error) {
	link := c.url + "/v3/schedules"
	if page > 0 {
		link += "?page=" + strconv.Itoa(page)
	}
	resp, err := c.Request("GET", link, nil, false)
	if err != nil {
		return nil, err
	}
	return resp.Map()
}

func (c PushClient) ScheduleView(id string) (map[string]interface{}, error) {
	link := c.url + "/v3/schedules/" + id
	resp, err := c.Request("GET", link, nil, false)
	if err != nil {
		return nil, err
	}
	return resp.Map()
}

func (c PushClient) ScheduleUpdate(id string, req *SchedulePayload) (map[string]interface{}, error) {
	link := c.url + "/v3/schedules/" + id
	buf, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	resp, err := c.Request("PUT", link, bytes.NewReader(buf), false)
	if err != nil {
		return nil, err
	}
	return resp.Map()
}

func (c PushClient) ScheduleDelete(id string) ([]byte, error) {
	link := c.url + "/v3/schedules/" + id
	resp, err := c.Request("DELETE", link, nil, false)
	if err != nil {
		return nil, err
	}
	return resp.Bytes(), nil
}
