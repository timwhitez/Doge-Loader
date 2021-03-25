![Doge-Loader](https://socialify.git.ci/timwhitez/Doge-Loader/image?description=1&font=Raleway&forks=1&issues=1&language=1&logo=https%3A%2F%2Favatars1.githubusercontent.com%2Fu%2F36320909&owner=1&pattern=Circuit%20Board&stargazers=1&theme=Light)

- 🐸Frog For Automatic Scan

- 🐶Doge For Defense Evasion&Offensive Security

# Doge-Loader
- Cobalt Strike Shellcode Loader by Golang

- 本项目感谢朋友gd, skate_zhu的帮助

- 此项目不进行后续更新，详见衍生新项目[Doge-Stealth-Framework](https://github.com/timwhitez/Doge-Stealth-Framework)

# 🐶Doge-Loader
本项目主要作用为Cobalt Strike的shellcode加载器

编译：go build -gcflags=-trimpath=$GOPATH -asmflags=-trimpath=$GOPATH -ldflags "-w -s -H windowsgui"

项目结构如下：
```
main.go
主程序，用于远程下载shellcode并加载(需要自行上传shellcode(bin文件)并且修改源码里的url等参数)

xor.go
异或编码解码器

/enhancement
附加功能，主要用于免杀等（需自行在main.go中调用）

/enhancement/copy.go
程序自删除与自复制

/enhancement/sandbox.go
反沙箱，沙箱内部数据需要自行收集，见我后续项目

/local
本地shellcode文件分离加载

/local/bin/main.go
本地shellcode raw格式分离加载

/local/txt/main.go
本地shellcode文本格式异或分离加载（以/x开头）

/local/txt/xor/xor.go
shellcode异或编码
```
注：在开源前做了些取舍，去掉了下载的域前置，cdn ip轮询，代理穿透等功能。

仅为功能较为基础的shellcode加载器，可自行增添功能或协助丰富此开源项目。

# 🚀Star Trend
[![Stargazers over time](https://starchart.cc/timwhitez/Doge-Loader.svg)](https://starchart.cc/timwhitez/Doge-Loader)
