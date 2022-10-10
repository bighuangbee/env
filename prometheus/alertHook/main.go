/**
 * @desc //TODO $
 * @param $
 * @return $
 **/
package main

import (
	"encoding/json"
	"gopackage/common"
	"gopackage/loger"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

var APP_NAME = common.GetPathName()
func init(){
	loger.Setup("/opt/logs/" + APP_NAME)
}

func main()  {
	if os.Getenv("alertType") == "" || os.Getenv("alertToken") == ""{
		loger.Info("alertConfig empty.")
		return
	}
	alertConfig.Type = os.Getenv("alertType")
	alertConfig.Token = os.Getenv("alertToken")

	//a := `{"receiver":"webhook","status":"firing","alerts":[{"status":"resolved","labels":{"alertname":"nodeDown","group":"client-node","instance":"192.168.80.94:9100","job":"node-exporter","status":"High"},"annotations":{"description":"192.168.80.94:9100 of job node-exporter has been down for more than 30s.","summary":"【机器节点下线】Instance 192.168.80.94:9100 down"},"startsAt":"2022-01-26T03:31:17.802305827Z","endsAt":"2022-01-26T03:36:17.802305827Z","generatorURL":"http://7202bd3f529b:9090/graph?g0.expr=up+%3D%3D+0\u0026g0.tab=1","fingerprint":"20ff339993fb7178"},{"status":"firing","labels":{"alertname":"nodeDown","group":"client-node","instance":"192.168.80.94:9101","job":"node-exporter","status":"High"},"annotations":{"description":"192.168.80.94:9101 of job node-exporter has been down for more than 30s.","summary":"【机器节点下线】Instance 192.168.80.94:9101 down"},"startsAt":"2022-01-26T03:31:32.802305827Z","endsAt":"0001-01-01T00:00:00Z","generatorURL":"http://7202bd3f529b:9090/graph?g0.expr=up+%3D%3D+0\u0026g0.tab=1","fingerprint":"120cb5c2008192c9"}],"groupLabels":{},"commonLabels":{"alertname":"nodeDown","group":"client-node","job":"node-exporter","status":"High"},"commonAnnotations":{},"externalURL":"http://80e1c130cfb2:9093","version":"4","groupKey":"{}:{}","truncatedAlerts":0}`
	//var receiver AlertReceiver
	//json.Unmarshal([]byte(a), &receiver)
	//SendAlertMsg(receiver.Alerts)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		msg, _ := ioutil.ReadAll(r.Body)
		loger.Info("receiver:", string(msg))

		var receiver AlertReceiver
		json.Unmarshal(msg, &receiver)
		err := SendAlertMsg(receiver.Alerts)
		if err != nil{
			loger.Error(err)
		}
	})

	loger.Info(APP_NAME+" run succ，listen port:", 9092)
	http.ListenAndServe("0.0.0.0:9092", nil)
}


func timeFormat(str string)string{
	t, err := time.Parse(time.RFC3339, str)
	if err != nil || t.Unix() < 0{
		return ""
	}
	return t.In(time.FixedZone("CST", 8*3600)).Format("2006-01-02 15:04:05")
}
