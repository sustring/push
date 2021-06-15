package jpush

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

const (
	IOSAppKey          = "ed27f827f3695f9530d13f8d"
	IOSMasterSecretKey = "31e4eba08c6a128d8fe5ad64"
	IOSRegistrationID  = "171976fa8a523c1dfdd"

	AndroidAppKey         = "60823c3e0d364f99832722ad"
	AndroidMasterSecret   = "54f21a5ea295367e8524a257"
	AndroidRegistrationId = "140fe1da9e038c6b343"
)

var client = NewClient(AndroidAppKey, AndroidMasterSecret, "", "")

func TestClientGetCidPool(t *testing.T) {
	data, err := client.GetCidPool(0, "push")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(data)
}

func TestPushRegistrationId(t *testing.T) {
	audience := Audience{
		RegistrationId: []string{AndroidRegistrationId},
	}

	notification := Notification{}
	notification.Android = &NotificationAndroid{
		Alert: "jpush registration_id test",
		Title: "a test from kitty for registration_id test",
		Extras: map[string]interface{}{
			"msg_id":   123,
			"msg_type": 6,
		},
	}

	client := NewClient(AndroidAppKey, AndroidMasterSecret, "", "")
	res, err := client.Push(&PushPayload{
		Cid:          "60823c3e0d364f99832722ad-eb56f386-79ef-4036-85cc-4bdaf6c1dbcc",
		Platform:     PlatformAndroid,
		Audience:     &audience,
		Notification: &notification,
	}, false)
	if err != nil {
		t.Fatalf("err: %+v", err)
	}

	t.Logf("res: %+v\n", res)
}

func TestPushAlias(t *testing.T) {
	audience := Audience{
		Alias: []string{"qiuqiankun"},
	}

	notification := Notification{}
	notification.Android = &NotificationAndroid{
		Alert: "jpush alias test",
		Title: "a test from kitty for alias test",
		Extras: map[string]interface{}{
			"msg_id":   123,
			"msg_type": 6,
		},
	}

	client := NewClient(AndroidAppKey, AndroidMasterSecret, "", "")
	res, err := client.Push(&PushPayload{
		Platform:     PlatformAndroid,
		Audience:     &audience,
		Notification: &notification,
	}, false)
	if err != nil {
		t.Fatalf("err: %+v", err)
	}

	t.Logf("res: %+v\n", res)
}

func TestPushTag(t *testing.T) {
	audience := Audience{
		Tag: []string{"mobile"},
	}

	notification := Notification{}
	notification.Android = &NotificationAndroid{
		Alert: "jpush tag test",
		Title: "a test from kitty for tag test",
		Extras: map[string]interface{}{
			"msg_id":   123,
			"msg_type": 6,
		},
	}

	client := NewClient(AndroidAppKey, AndroidMasterSecret, "", "")
	res, err := client.Push(&PushPayload{
		Platform:     PlatformAndroid,
		Audience:     &audience,
		Notification: &notification,
	}, false)
	if err != nil {
		t.Fatalf("err: %+v", err)
	}

	t.Logf("res: %+v\n", res)
}

func TestReportReceived(t *testing.T) {
	res, err := client.ReceivedDetail([]string{"67554217262909280"})
	if err != nil {
		t.Fatalf("err: %+v", err)
	}

	t.Logf("res: %+v\n", res)
}

func TestClientDeviceView(t *testing.T) {
	res, err := client.DeviceView(AndroidRegistrationId)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(res)
}

func TestClientDeviceRequest(t *testing.T) {
	tags := &DeviceSettingRequestTags{
		Add: []string{"mobile"},
	}
	req := &DeviceSettingPayload{
		Alias: "qiuqiankun",
		Tags:  tags,
	}
	err := client.DeviceSet(AndroidRegistrationId, req)
	if err != nil {
		t.Error(err)
		return
	}
}

func TestClientDeviceGetWithAlias(t *testing.T) {
	result, err := client.AliasGet("qiuqiankun", nil)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(result)
}

func TestClientDeviceDeleteAlias(t *testing.T) {
	err := client.AliasDelete("qiuqiankun")
	if err != nil {
		t.Error(err)
		return
	}
}

func TestClientDeviceGetTags(t *testing.T) {
	result, err := client.TagsGet()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(result)
}

func TestClientDeviceCheckDeviceWithTag(t *testing.T) {
	result, err := client.TagCheck("mobile", AndroidRegistrationId)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(result)
}

func TestClientDeviceBindTags(t *testing.T) {
	req := &TagUpdatePayload{
		Add: []string{AndroidRegistrationId},
	}
	err := client.TagUpdate("mobile", req)
	if err != nil {
		t.Error(err)
		return
	}
}

func TestClientDeviceDeleteTag(t *testing.T) {
	err := client.TagDelete("mobile", nil)
	if err != nil {
		t.Error(err)
		return
	}
}
