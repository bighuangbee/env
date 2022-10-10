package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type AlertConfig struct {
	Type string
	Token string
}

const AlertDinDin = "dindin"
const AlertWechat = "wechat"

type Alerts struct {
	Status string `json:"status"`
	Labels struct{
		Alertname string `json:"alertname"`
		Group string `json:"group"`
		Instance string `json:"instance"`
		Job string `json:"job"`
		Status string `json:"status"`
		ExportedJob string `json:"exported_job"`
	}
	Annotations struct{
		Description string `json:"description"`
		Summary string `json:"summary"`
	}
	StartsAt string
	EndsAt string
	GeneratorURL string
	Fingerprint string
}

type AlertReceiver struct {
	Status string `json:"status"`
	Alerts []Alerts
}

var alertConfig AlertConfig

var template = `**********%s**********
告警级别：%s
告警服务：%s
故障主机: %s
告警名称: %s
告警详情: %s
触发阀值：%s
故障时间: %s
`

var alertStatus = map[string]string{
	"firing": "告警通知",
	"resolved": "告警恢复",
}

func SendAlertMsg(alerts []Alerts)error {
	var url string
	if alertConfig.Type == AlertDinDin{
		url = `https://oapi.dingtalk.com/robot/send?access_token=` + alertConfig.Token
	}else if alertConfig.Type == AlertWechat{
		url = `https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=` + alertConfig.Token
	}else {
		return errors.New("alertConfig empty.")
	}


	for _, val := range alerts {
		var msg string
		t := timeFormat(val.StartsAt)
		if timeFormat(val.EndsAt) != ""{
			t += " - " + timeFormat(val.EndsAt)
		}

		fmt.Println(alertStatus[val.Status])
		msg += fmt.Sprintf(template, alertStatus[val.Status], val.Labels.Status, val.Labels.Job, val.Labels.Instance,
			val.Annotations.Summary, val.Annotations.Description, "", t + "\n")

		content := map[string]interface{}{
			"content": fmt.Sprintf(strings.TrimRight(msg, "\n")),
		}
		m := map[string]interface{}{
			"msgtype": "text",
			"text": content,
		}

		data, _ := json.Marshal(m)
		httpPost(url, data)
	}
	return nil
}


func httpPost(url string, data []byte) error{
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil{
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil{
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK{
		body, _ := ioutil.ReadAll(resp.Body)
		return errors.New(string(body))
	}
	return nil
}
