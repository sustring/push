package jpush

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
)

type PushClient struct {
	*BaseClient
	url string
}

type Platform string

const (
	PlatformAll      Platform = "all"
	PlatformAndroid  Platform = "android"
	PlatformIOS      Platform = "ios"
	PlatformWinPhone Platform = "winphone"
)

type Audience struct {
	Tag            []string `json:"tag,omitempty"`             // max 20
	TagAnd         []string `json:"tag_and,omitempty"`         // max 20
	TagNot         []string `json:"tag_not,omitempty"`         // max 20
	Alias          []string `json:"alias,omitempty"`           // max 1000
	RegistrationId []string `json:"registration_id,omitempty"` // max 1000
	Segment        []string `json:"segment,omitempty"`
	ABTest         []string `json:"abtest,omitempty"`
}

type Notification struct {
	Alert   string               `json:"alert,omitempty"`
	Android *NotificationAndroid `json:"android,omitempty"`
	IOS     *NotificationIOS     `json:"ios,omitempty"`
}

type NotificationAndroid struct {
	Alert             string                 `json:"alert"`
	Title             string                 `json:"title,omitempty"`
	BuilderId         int                    `json:"builder_id,int,omitempty"`
	ChannelId         string                 `json:"channel_id,omitempty"`
	Priority          int                    `json:"priority,omitempty"` //default 0, -2~2
	Category          string                 `json:"category,omitempty"`
	Style             int                    `json:"style,int,omitempty"`
	AlertType         int                    `json:"alert_type,int,omitempty"`
	BigText           string                 `json:"big_text,omitempty"`
	Inbox             map[string]interface{} `json:"inbox,omitempty"`
	BigPicPath        string                 `json:"big_pic_path,omitempty"`
	Extras            map[string]interface{} `json:"extras,omitempty"`
	LargeIcon         string                 `json:"large_icon,omitempty"`
	SmallIconUri      string                 `json:"small_icon_uri,omitempty"`
	Intent            map[string]interface{} `json:"intent,omitempty"`
	UriActivity       string                 `json:"uri_activity,omitempty"`
	UriAction         string                 `json:"uri_action,omitempty"`
	BadgeAddNum       int                    `json:"badge_add_num,omitempty"`
	BadgeClass        string                 `json:"badge_class,omitempty"`
	Sound             string                 `json:"sound,omitempty"`
	ShowBeginTime     string                 `json:"show_begin_time,omitempty"`
	ShowEndTime       string                 `json:"show_end_time,omitempty"`
	DisplayForeground string                 `json:"display_foreground,omitempty"`
}

type NotificationIOS struct {
	Alert            interface{}            `json:"alert"`
	Sound            string                 `json:"sound,omitempty"`
	Badge            int                    `json:"badge,int,omitempty"`
	ContentAvailable bool                   `json:"content-available,omitempty"`
	MutableContent   bool                   `json:"mutable-content,omitempty"`
	Category         string                 `json:"category,omitempty"`
	Extras           map[string]interface{} `json:"extras,omitempty"`
	ThreadId         string                 `json:"thread-id,omitempty"`
}

type Message struct {
	MsgContent  string                 `json:"msg_content"`
	Title       string                 `json:"title,omitempty"`
	ContentType string                 `json:"content_type,omitempty"`
	Extras      map[string]interface{} `json:"extras,omitempty"`
}

type SmsMessage struct {
	Content   string `json:"content"`
	DelayTime int    `json:"delay_time,int,omitempty"`
}

type PushOptions struct {
	SendNo            int                     `json:"sendno,int,omitempty"`
	TimeToLive        int                     `json:"time_to_live,int,omitempty"`
	OverrideMsgId     int64                   `json:"override_msg_id,int64,omitempty"`
	ApnsProduction    bool                    `json:"apns_production"`
	ApnsCollapseId    string                  `json:"apns_collapse_id,omitempty"`
	BigPushDuration   int                     `json:"big_push_duration,int,omitempty"`
	ThirdPartyChannel ThirdPartyChannelOption `json:"third_party_channel,omitempty"`
}

type ThirdPartyChannelOption struct {
	Xiaomi XiaomiChannel `json:"xiaomi,omitempty"`
	Huawei HuaweiChannel `json:"huawei,omitempty"`
	Meizu  MeizuChannel  `json:"meizu,omitempty"`
	Fcm    FcmChannel    `json:"fcm,omitempty"`
	Oppo   OppoChannel   `json:"oppo,omitempty"`
	Vivo   VivoChannel   `json:"vivo,omitempty"`
}

type XiaomiChannel struct {
	Distribution          string `json:"distribution"` // jpush, ospush, secondary_push
	ChannelId             string `json:"channel_id,omitempty"`
	LargeIcon             string `json:"large_icon,omitempty"`
	SmallIconUri          string `json:"small_icon_uri,omitempty"`
	SmallIconColor        string `json:"small_icon_color,omitempty"`
	BigText               string `json:"big_text,omitempty"`
	Style                 int    `json:"style,omitempty"`
	DistributionFcm       string `json:"distribution_fcm,omitempty"`
	DistributionCustomize string `json:"distribution_customize,omitempty"`
}

type HuaweiChannel struct {
	Distribution       string                 `json:"distribution"` // jpush, ospush, secondary_push
	DistributionFcm    string                 `json:"distribution_fcm,omitempty"`
	Importance         string                 `json:"importance,omitempty"`
	LargeIcon          string                 `json:"large_icon,omitempty"`
	SmallIconUri       string                 `json:"small_icon_uri,omitempty"`
	SmallIconColor     string                 `json:"small_icon_color,omitempty"`
	Inbox              map[string]interface{} `json:"inbox,omitempty"`
	Style              int                    `json:"style,omitempty"`
	OnlyUseVendorStyle bool                   `json:"only_use_vendor_style"`
}

type MeizuChannel struct {
	Distribution    string `json:"distribution"` // jpush, ospush, secondary_push
	DistributionFcm string `json:"distribution_fcm,omitempty"`
}

type FcmChannel struct {
	Distribution string `json:"distribution"` // jpush, ospush, secondary_push
}

type OppoChannel struct {
	Distribution    string `json:"distribution"` // jpush, ospush, secondary_push
	ChannelId       string `json:"channel_id,omitempty"`
	DistributionFcm string `json:"distribution_fcm,omitempty"`
	LargeIcon       string `json:"large_icon,omitempty"`
	BigPicPath      string `json:"big_pic_path,omitempty"`
	Style           int    `json:"style,omitempty"`
}

type VivoChannel struct {
	Distribution    string `json:"distribution"` // jpush, ospush, secondary_push
	Classification  string `json:"classification,omitempty"`
	DistributionFcm string `json:"distribution_fcm,omitempty"`
	PushMode        int    `json:"push_mode,omitempty"`
}

type PushCallback struct {
	Url    string                 `json:"url,omitempty"`
	Params map[string]interface{} `json:"params,omitempty"`
	Type   int                    `json:"type,omitempty"`
}

type Notification3rd struct {
	Title       string                 `json:"title,omitempty"`
	Content     string                 `json:"content"`
	ChannelId   string                 `json:"channel_id,omitempty"`
	UriActivity string                 `json:"uri_activity,omitempty"`
	UriAction   string                 `json:"uri_action,omitempty"`
	BadgeAddNum string                 `json:"badge_add_num,omitempty"`
	BadgeClass  string                 `json:"badge_class,omitempty"`
	Sound       string                 `json:"sound,omitempty"`
	Extras      map[string]interface{} `json:"extras,omitempty"`
}

type PushPayload struct {
	Cid             string           `json:"cid,omitempty"`
	Platform        Platform         `json:"platform"`
	Audience        *Audience        `json:"audience,omitempty"`
	Notification    *Notification    `json:"notification,omitempty"`
	Message         *Message         `json:"message,omitempty"`
	SmsMessage      *SmsMessage      `json:"sms_message,omitempty"`
	Options         *PushOptions     `json:"options,omitempty"`
	Callback        *PushCallback    `json:"callback,omitempty"`
	Notification3rd *Notification3rd `json:"notification_3rd,omitempty"`
}

type PushResult struct {
	SendNo int    `json:"sendno,int,omitempty"`
	MsgId  string `json:"msg_id"`
}

func (c PushClient) Push(payload *PushPayload, validate bool) (*PushResult, error) {
	link := c.url + "/v3/push"
	if validate {
		link = c.url + "/v3/push/validate"
	}
	buf, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	fmt.Printf("[Push] %s\n", string(buf))
	resp, err := c.Request("POST", link, bytes.NewReader(buf), false)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() != http.StatusOK {
		return nil, errors.New(resp.Status())
	}
	var out PushResult
	err = json.Unmarshal(resp.Bytes(), &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

type CidPool struct {
	CidList []string `json:"cidlist"`
}

func (c PushClient) GetCidPool(count int, cidType string) ([]string, error) {
	link := c.url + "/v3/push/cid?"
	if count > 0 {
		link += "count=" + strconv.Itoa(count)
	}
	if cidType != "" {
		link += "type=" + cidType
	}
	resp, err := c.Request("GET", link, nil, false)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() != http.StatusOK {
		return nil, errors.New(resp.Status())
	}
	var out CidPool
	err = json.Unmarshal(resp.Bytes(), &out)
	if err != nil {
		return nil, err
	}
	return out.CidList, nil
}

func (c PushClient) DeletePush(msgId string) error {
	link := c.url + "/v3/push/" + msgId
	resp, err := c.Request("DELETE", link, nil, false)
	if err != nil {
		return err
	}
	if resp.StatusCode() != http.StatusOK {
		return errors.New(string(resp.Bytes()))
	}
	return nil
}

func (c PushClient) GroupPush(payload *PushPayload) (map[string]interface{}, error) {
	link := c.url + "/v3/grouppush"
	buf, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	resp, err := c.Request("POST", link, bytes.NewReader(buf), true)
	if err != nil {
		return nil, err
	}
	return resp.Map()
}
