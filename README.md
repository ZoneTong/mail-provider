# 支持NTLM协议发送邮件 mail-provider

=============

把smtp封装为一个简单http接口，配置到sender中用来发送报警邮件

## 安装方法

编译方法

```bash
cd $GOPATH/src
mkdir github.com/ZoneTong/ -p
cd github.com/ZoneTong/
# git clone https://github.com/ZoneTong/mail-provider.git
cd mail-provider
go get ./...
./control build
```

编译成功之后，修改cfg.json文件相关信息，使用

```bash
./control start
```

即可启动

## 使用方法

下载之后为源码，需要编译

```bash
curl http://$ip:4000/sender/mail -d "tos=a@a.com,b@b.com&subject=xx&content=yy"
```

curl http://127.0.0.1:10086/sender/mail -d "tos=zhoutong@genomics.cn&subject=Sxx&content=Cyy"
