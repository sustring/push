package common

type GetDeviceInput struct {
	Id string
}

type GetDeviceOutput struct {
	Id      string
	Alias   string
	TagList []string
}

type SetDeviceInput struct {
	Id        string
	Alias     string
	CleanTags bool
	AddTags   []string
	DelTags   []string
}

type SetDeviceOutput struct {
}

type UpdateTagInput struct {
	Tag     string
	AddList []string
	DelList []string
}

type UpdateTagOutput struct {
}

type DeleteTagInput struct {
	Tag string
}

type DeleteTagOutput struct {
}

type CheckTagInput struct {
	Tag string
	Id  string
}

type CheckTagOutput struct {
	Result bool
}

type PlatformType int

const (
	ALL PlatformType = iota
	Android
	IOS
)

type AudienceInfo struct {
	AliasList []string
	TagList   []string
}

type PushMessageInput struct {
	Platform PlatformType
	Id       int64
	Type     string
	Alert    string
	//Title        string
	Audience     AudienceInfo
	Presentation bool
	//Extra        map[string]interface{}
}

type PushMessageOutput struct {
	MsgId string
}

type InspectMessageInput struct {
	MsgId []string
}

type InspectMessageOutput struct {
	MsgId           string
	AndroidReceived int
	IOSAPNSReceived int
	IOSAPNSSent     int
	IOSMsgReceived  int
}
