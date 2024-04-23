FROM ubuntu:18.04

RUN sed -i "s/archive.ubuntu.com/mirrors.aliyun.com/g" /etc/apt/sources.list
RUN apt-get update --fix-missing

ARG DEBIAN_FRONTEND=noninteractive
ENV TZ=Asia/Shanghai

RUN apt-get install -y wget git vim make

RUN apt-get install -y software-properties-common
RUN add-apt-repository -y ppa:ubuntu-toolchain-r/test
RUN apt install -y gcc-11 g++-11
RUN update-alternatives --install /usr/bin/gcc gcc /usr/bin/gcc-11 11 && \
    update-alternatives --install /usr/bin/g++ g++ /usr/bin/g++-11 11

#golang
RUN wget https://golang.google.cn/dl/go1.17.13.linux-amd64.tar.gz
RUN tar -C /usr/local -xvzf go1.17.13.linux-amd64.tar.gz

ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn
ENV PATH=$PATH:/usr/local/go/bin
ENV GODEBUG=cgocheck=0

#ptz ocr
RUN apt-get install -y software-properties-common
RUN add-apt-repository "deb http://security.ubuntu.com/ubuntu xenial-security main"
RUN apt update
RUN apt install -y libjasper1 libjasper-dev


#moziAR
RUN apt-get install -y libosmesa6  libglapi-mesa mesa-common-dev

WORKDIR /root/

#git.hiscene.net
# --build-arg ssh_prv_key="$(cat ~/.ssh/id_rsa)" --build-arg ssh_pub_key="$(cat ~/.ssh/id_rsa.pub)"
ARG ssh_prv_key
ARG ssh_pub_key

RUN go env -w GOPROXY=https://goproxy.cn && \
    go env -w GOPRIVATE=git.hiscene.net && \
    git config --global url."git@git.hiscene.net:".insteadOf "https://git.hiscene.net/"
RUN mkdir -p /root/.ssh && \
    chmod 0700 /root/.ssh
RUN echo "$ssh_prv_key" > /root/.ssh/id_rsa && \
    echo "$ssh_pub_key" > /root/.ssh/id_rsa.pub && \
    chmod 600 /root/.ssh/id_rsa && \
    chmod 600 /root/.ssh/id_rsa.pub && \
    ssh-keyscan -t rsa git.hiscene.net >> /root/.ssh/known_hosts

ENV PATH=/app/mozi_render/libs/ptz_ocr/libs:$PATH
ENV LD_LIBRARY_PATH=/app/mozi_render/libs/ptz_ocr/libs:$LD_LIBRARY_PATH

EXPOSE 8109
EXPOSE 9109
