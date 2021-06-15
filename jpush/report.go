package jpush

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"strings"
)

type ReportClient struct {
	*BaseClient
	url string
}

type ReceivedDetailResult struct {
	MsgId                 string `json:"msg_id"`
	JPushReceived         int    `json:"jpush_received"`
	AndroidPNSSent        int    `json:"android_pns_sent"`
	IOSAPNSSent           int    `json:"ios_apns_sent"`
	IOSAPNSReceived       int    `json:"ios_apns_received"`
	IOSMsgReceived        int    `json:"ios_msg_received"`
	WPMPNSSent            int    `json:"wp_mpns_sent"`
	QuickappJpushReceived int    `json:"quickapp_jpush_received"`
	QuickappPNSSent       int    `json:"quickapp_pns_sent"`
}

func (c ReportClient) ReceivedDetail(msgIds []string) ([]*ReceivedDetailResult, error) {
	if len(msgIds) == 0 {
		return nil, errors.New("invalid msg id")
	}
	link := c.url + "/v3/received/detail?msg_ids=" + strings.Join(msgIds, ",")
	resp, err := c.Request("GET", link, nil, false)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() != http.StatusOK {
		return nil, errors.New(string(resp.Bytes()))
	}

	var list []*ReceivedDetailResult
	err = json.Unmarshal(resp.Bytes(), &list)
	if err != nil {
		return nil, err
	}
	return list, nil
}

type MessageStatusPayload struct {
	MsgId           int      `json:"msg_id,int"`
	RegistrationIds []string `json:"registration_ids"`
	Date            string   `json:"date,omitempty"` //  format:yyyy-mm-dd
}

type MessageStatusResult struct {
	Status int `json:"status"`
}

func (c ReportClient) MessageStatus(payload *MessageStatusPayload) (map[string]*MessageStatusResult, error) {
	link := c.url + "/v3/status/message"
	buf, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	resp, err := c.Request("POST", link, bytes.NewReader(buf), false)
	if err != nil {
		return nil, err
	}
	params := make(map[string]*MessageStatusResult)
	err = json.Unmarshal(resp.Bytes(), &params)
	if err != nil {
		return nil, err
	}
	return params, nil
}
