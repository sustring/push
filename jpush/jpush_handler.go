package jpush

import (
	"errors"
	"github.com/sustring/push/common"
)

const (
	PushUrl   = "https://api.jpush.cn"
	ReportUrl = "https://report.jpush.cn"
	DeviceUrl = "https://device.jpush.cn"
)

type Client struct {
	*PushClient
	*DeviceClient
	*ReportClient
	*ScheduleClient
}

func NewClient(appKey, masterSecret, groupKey, groupMasterSecret string) *Client {
	base := &BaseClient{
		AppKey:            appKey,
		MasterSecret:      masterSecret,
		GroupKey:          groupKey,
		GroupMasterSecret: groupMasterSecret,
	}
	return &Client{
		&PushClient{
			BaseClient: base,
			url:        PushUrl,
		},
		&DeviceClient{
			BaseClient: base,
			url:        DeviceUrl,
		},
		&ReportClient{
			BaseClient: base,
			url:        ReportUrl,
		},
		&ScheduleClient{
			BaseClient: nil,
			url:        PushUrl,
		},
	}
}

func (c Client) SetDevice(in *common.SetDeviceInput) (*common.SetDeviceOutput, error) {
	payload := &DeviceSettingPayload{
		Alias: in.Alias,
	}
	if in.CleanTags {
		payload.Tags = ""
	} else {
		payload.Tags = &DeviceSettingRequestTags{
			Add:    in.AddTags,
			Remove: in.DelTags,
		}
	}

	err := c.DeviceSet(in.Id, payload)
	if err != nil {
		return nil, err
	}
	return &common.SetDeviceOutput{}, nil
}

func (c Client) GetDevice(in *common.GetDeviceInput) (*common.GetDeviceOutput, error) {
	_, err := c.DeviceView(in.Id)
	if err != nil {
		return nil, err
	}

	return &common.GetDeviceOutput{
		Id:      "",
		Alias:   "",
		TagList: nil,
	}, nil
}

func (c Client) UpdateTag(in *common.UpdateTagInput) (*common.UpdateTagOutput, error) {
	payload := &TagUpdatePayload{
		Add:    in.AddList,
		Remove: in.DelList,
	}

	err := c.TagUpdate(in.Tag, payload)
	if err != nil {
		return nil, err
	}
	return &common.UpdateTagOutput{}, nil
}

func (c Client) DeleteTag(in *common.DeleteTagInput) (*common.DeleteTagOutput, error) {
	err := c.TagDelete(in.Tag, nil)
	if err != nil {
		return nil, err
	}
	return &common.DeleteTagOutput{}, nil
}

func (c Client) CheckTag(in *common.CheckTagInput) (*common.CheckTagOutput, error) {
	_, err := c.TagCheck(in.Tag, in.Id)
	if err != nil {
		return nil, err
	}
	return &common.CheckTagOutput{}, nil
}

func (c Client) PushMessage(in *common.PushMessageInput) (*common.PushMessageOutput, error) {
	payload := &PushPayload{}

	payload.Audience = &Audience{}
	for _, tag := range in.Audience.TagList {
		if tag != "" {
			payload.Audience.Tag = append(payload.Audience.Tag, tag)
		}
	}
	for _, alias := range in.Audience.AliasList {
		if alias != "" {
			payload.Audience.Alias = append(payload.Audience.Alias, alias)
		}
	}

	extra := map[string]interface{}{"msg_id": in.Id, "msg_type": in.Type}
	if !in.Presentation {
		payload.Message = &Message{
			MsgContent: in.Alert,
			//Title:       in.Alert,
			//ContentType: "text",
			Extras: extra,
		}
	}

	if in.Platform == common.ALL {
		payload.Platform = PlatformAll
		if in.Presentation {
			payload.Notification = &Notification{
				Alert: in.Alert,
				Android: &NotificationAndroid{
					//Alert:  in.Alert,
					Extras: extra,
				},
				IOS: &NotificationIOS{
					//Alert:    in.Alert,
					ThreadId: in.Type,
					Extras:   extra,
				},
			}
		}
	} else if in.Platform == common.Android {
		payload.Platform = PlatformAndroid
		if in.Presentation {
			payload.Notification = &Notification{
				Android: &NotificationAndroid{
					Alert:  in.Alert,
					Extras: extra,
				},
			}
		}
	} else if in.Platform == common.IOS {
		payload.Platform = PlatformIOS
		if in.Presentation {
			payload.Notification = &Notification{
				IOS: &NotificationIOS{
					Alert:    in.Alert,
					ThreadId: in.Type,
					Extras:   extra,
				},
			}
		}
	} else {
		return nil, errors.New("invalid input params")
	}

	_, err := c.Push(payload, false)
	if err != nil {
		return nil, err
	}

	return &common.PushMessageOutput{}, nil
}

func (c Client) InspectMessage(in *common.InspectMessageInput) (*common.InspectMessageOutput, error) {
	_, err := c.ReceivedDetail(in.MsgId)
	if err != nil {
		return nil, err
	}
	return &common.InspectMessageOutput{}, nil
}
