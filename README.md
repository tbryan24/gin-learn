#  windows配置make命令

csdn：https://blog.csdn.net/qq_34432348/article/details/126576336

在学***kratos***时发现有个很好用的命令，即make，比如:

```Go
make api
```

![点击并拖拽以移动](data:image/gif;base64,R0lGODlhAQABAPABAP///wAAACH5BAEKAAAALAAAAAABAAEAAAICRAEAOw==)

但是你会发现这东西默认只能在linux和Mac上能用，作为windows用户就很难受了，你要是不难受就敲完整命令吧，比如（$(API_PROTO_FILES)即你的proto文件路径）：

```Go
protoc --proto_path=./api \
    --proto_path=./third_party \
    --go_out=paths=source_relative:./api \
    --go-http_out=paths=source_relative:./api \
    --go-grpc_out=paths=source_relative:./api \
    --openapi_out=fq_schema_naming=true,default_response=false:. \
    $(API_PROTO_FILES)
```

![点击并拖拽以移动](data:image/gif;base64,R0lGODlhAQABAPABAP///wAAACH5BAEKAAAALAAAAAABAAEAAAICRAEAOw==)

反正我是受不了，为啥别人Mac能用，我不能用，就因为我买不起Mac吗，我不服。

解决方案：

1、首先下载MinGw（干什么用的自行百度）：[Download MinGW - Minimalist GNU for Windows from SourceForge.net](https://sourceforge.net/projects/mingw/files/latest/download?source=files)

2、点击安装后打开将Base Setup全部mark，然后点击installation->Apply Changes，等待安装完就行（时间可能有点长）。

3、然后将bin目录添加进环境变量，比如我的路径是***C:\MinGW\bin***，该目录下有一个mingw-get.exe，这就是我上面说的命令可执行文件了。

4、作为有强迫症的我肯定不愿意执行mingw-get这样的命令，我是想用make这样的命令，所以复制一份mingw-get.exe，将其改为make.exe，到此为止就大功告成了。

5、但是有时候你在***kratos***里执行***make api***时会发现要报错，他会提示**C:\不是内部或外部命令，也不是可运行的程序**，这里是因为你的git安装到的是默认目录，即**C:\Program Files\Git**，因为实际上**kratos**里的makefile文件里用到了Git_Bash，而这个路径正好是你的git-bash.exe的路径，可能是**C:\Program Files\Git**这个路径Program Files中空格无法解析的原因（我猜的），我的方法就是将Git重新安装，然后自定义目录，如**C:\install\Git**，这样就完美解决了。

```Go
GOHOSTOS:=$(shell go env GOHOSTOS)
GOPATH:=$(shell go env GOPATH)
VERSION=$(shell git describe --tags --always)

ifeq ($(GOHOSTOS), windows)
    #the `find.exe` is different from `find` in bash/shell.
    #to see https://docs.microsoft.com/en-us/windows-server/administration/windows-commands/find.
    #changed to use git-bash.exe to run find cli or other cli friendly, caused of every developer has a Git.
    Git_Bash= $(subst cmd\,bin\bash.exe,$(dir $(shell where git)))
    INTERNAL_PROTO_FILES=$(shell $(Git_Bash) -c "find internal -name *.proto")
    API_PROTO_FILES=$(shell $(Git_Bash) -c "find api -name *.proto")
else
    INTERNAL_PROTO_FILES=$(shell find internal -name *.proto)
    API_PROTO_FILES=$(shell find api -name *.proto)
endif
```

![点击并拖拽以移动](data:image/gif;base64,R0lGODlhAQABAPABAP///wAAACH5BAEKAAAALAAAAAABAAEAAAICRAEAOw==)