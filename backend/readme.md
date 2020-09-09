**开发环境热启动**

```
# 安装
go get github.com/silenceper/gowatch

# 执行
gowatch
```

**去掉未注释警告**

项目根目录下创建文件夹`.vscode`，新建文件`settings.json`，内容如下：

```
{
    "go.lintFlags": ["--disable=all", "--enable=errcheck"]
}
```
