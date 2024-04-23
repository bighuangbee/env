FROM centos:7.6.1810

WORKDIR /root

RUN yum -y install vim make wget gstreamer1


#golang
RUN wget https://golang.google.cn/dl/go1.19.4.linux-amd64.tar.gz -O go.linux-amd64.tar.gz
RUN tar -C /usr/local -xvzf go.linux-amd64.tar.gz && rm -f go.linux-amd64.tar.gz

ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn
ENV PATH=$PATH:/usr/local/go/bin
ENV GODEBUG=cgocheck=0

