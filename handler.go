package handler

import (
	"github.com/sustring/push/common"
	"github.com/sustring/push/jpush"
)

type API interface {
	SetDevice(in *common.SetDeviceInput) (*common.SetDeviceOutput, error)
	GetDevice(in *common.GetDeviceInput) (*common.GetDeviceOutput, error)
	UpdateTag(in *common.UpdateTagInput) (*common.UpdateTagOutput, error)
	DeleteTag(in *common.DeleteTagInput) (*common.DeleteTagOutput, error)
	CheckTag(in *common.CheckTagInput) (*common.CheckTagOutput, error)
	PushMessage(in *common.PushMessageInput) (*common.PushMessageOutput, error)
	InspectMessage(in *common.InspectMessageInput) (*common.InspectMessageOutput, error)
}

func NewJPushClient(appKey, masterSecret string) API {
	return jpush.NewClient(appKey, masterSecret, "", "")
}
