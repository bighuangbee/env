### 预览地址
#### 原图： http://oss.hidrone.cn:8802/hidronedev/test/car_1.jpg
#### 动态缩略图： http://oss.hidrone.cn:8803/200x200/hidronedev/test/car_1.jpg


###dashboad无法修改密码
```bash
先备份/data/minio/data, 删除config，然后在compose.yml设置新密码后重新启动
rm /data/minio/data/.minio.sys/config/ -rf
```

动态缩略图需要对Bucket开放读写权限 set bucket as public permission. 

```sh
$ mc config host add minio http://localhost:9000 MINIO_ACCESS_KEY MINIO_SECRET_KEY
$ mc policy --recursive public minio/test
```


## Reference

* [Nginx: Real time image resizing and caching](https://github.com/sergejmueller/sergejmueller.github.io/wiki/Nginx:-Real-time-image-resizing-and-caching)
* [NGINX reverse proxy image resizing + AWS S3](https://medium.com/merapar/nginx-reverse-proxy-image-resizing-aws-cece1db5da01)
* [Nginx dynamic image resizing with caching](https://stumbles.id.au/nginx-dynamic-image-resizing-with-caching.html)
* [High‑Performance Caching with NGINX and NGINX Plus](https://www.nginx.com/blog/nginx-high-performance-caching/)
* [NGINX and NGINX Plus Deliver Responsive Images Without the Headaches](https://www.nginx.com/blog/responsive-images-without-headaches-nginx-plus/)
