# Doge-Loader
- Cobalt Strike Shellcode Loader by Golang

- 本人github的项目可能目前主要分为两大类，🐸Frog系列为自动化扫描方向，🐶Doge系列为免杀及内网渗透方向

- Doge-Loader为Doge系列比较重要的项目🐶

# 🐶Doge-Loader
本项目主要作用为Cobalt Strike的shellcode加载器

项目结构如下：
```
main.go
主程序，用于远程下载shellcode并加载

xor.go
异或编码解码器

/enhancement
附加功能，主要用于免杀等

/enhancement/copy.go
程序自删除与自复制

/enhancement/sandbox.go
反沙箱，沙箱内部数据需要自行收集，见我后续项目
```
注：在开源前做了些取舍，去掉了下载的域前置，ip轮询，代理穿透等。仅为功能较为基础的shellcode加载器，可自行增添功能或一起丰富此开源项目。
