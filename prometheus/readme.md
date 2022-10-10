##配置项

```
.
├── alertHook #接收告警信息并通过钉钉或企业发送通知的服务(webhook)
│   ├── docker-compose.yml
├── alertmanager
│   └── alertmanager.yml #设置由哪个服务接收告警信息
├── docker-compose.yaml  #grafana数据库储存目录映射到宿主机
├── exporter
│   └── docker-compose.yml #采集服务的指标数据导出接口
├── prometheus
│   ├── alert_rules #告警的触发规则
│   │   ├── mysql.yml
│   │   ├── redis.yml
│   │   └── test.yml
│   └── prometheus.yaml #设置告警的管理器、抓取哪些目标
└── readme.md

````

###修改抓取配置后reload
`curl -X POST http://192.168.80.94:9090/-/reload`

-------

###prometheus服务的IP改变之后，需要更新grafana地址或重装

####向pushgateway主动发送监控指标数据
`echo "vistorCount 20" | curl --data-binary @- http://192.168.80.94:9091/metrics/job/vistorTest
cat <<EOF | curl --data-binary @- http://192.168.80.94:9091/metrics/job/vistorTest
vistorCount 152
carCount 150
EOF`
------

nginx exporter
####查看是否开启stub_status: `nginx -V 2>&1 | grep -o with-http_stub_status_module`
nginx.conf server添加模块:
```
location /nginx_status {
    stub_status;
    access_log off;
    allow all;
}
```

