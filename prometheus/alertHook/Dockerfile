FROM alpine:3.14
WORKDIR /app
RUN mkdir /usr/share/zoneinfo/Asia -p
ADD Shanghai /usr/share/zoneinfo/Asia/Shanghai
ENV LANG zh_CN.UTF-8
ADD app_bin /app/app_bin
CMD ["./app_bin"]
