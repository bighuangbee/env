FROM ubuntu:20.04

RUN sed -i "s/archive.ubuntu.com/mirrors.aliyun.com/g" /etc/apt/sources.list
RUN apt-get update

ARG DEBIAN_FRONTEND=noninteractive
ENV TZ=Asia/Shanghai

RUN apt-get install -y wget git vim gcc g++ make ffmpeg
#RUN apt install -y software-properties-common gcc-11 #GCC11

#goav
RUN apt-get -y install autoconf automake build-essential libass-dev libfreetype6-dev libsdl1.2-dev libtheora-dev libtool libva-dev libvdpau-dev libvorbis-dev libxcb1-dev libxcb-shm0-dev libxcb-xfixes0-dev pkg-config texi2html zlib1g-dev
RUN apt-get -y install libavdevice-dev libavfilter-dev libswscale-dev libavcodec-dev libavformat-dev libswresample-dev libavutil-dev
#RUN apt -y install yasm

ENV FFMPEG_ROOT=$HOME/ffmpeg
ENV CGO_LDFLAGS="-L$FFMPEG_ROOT/lib/ -lavcodec -lavformat -lavutil -lswscale -lswresample -lavdevice -lavfilter"
ENV CGO_CFLAGS="-I$FFMPEG_ROOT/include"
ENV LD_LIBRARY_PATH=$HOME/ffmpeg/lib

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
