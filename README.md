# gin-api-boilerplate
A Go RESTful API server with gin and docker

## 安装
安装最新的[Go](https://golang.org/)，设置好`$GOPATH`，将`$GOPATH/bin`设置为可执行PATH。这里使用的是Mac或者Linux环境，如果使用Windows环境，请自行修改，比如使用`Makefile`文件`make`命令执行自动化脚本。

````
$ cd $GOPATH/src
$ git clone https://github.com/ihahoo/gin-api-boilerplate.git
````

包依赖管理使用[Dep](https://github.com/golang/dep)，可以把当前项目的依赖模块安装到`vendor`目录下，需要Go 1.8版本以上。
```
$ go get -u github.com/golang/dep/cmd/dep
```
安装好后，通过下面命令初始化包依赖:
```
$ dep ensure -v
```
当然，`需要翻墙`，否则会失败。

如果不使用`Dep`，也可以使用下面命令安装依赖，但是最好是能正常使用`Dep`
```
$ go get -d ./...
```

## 启动开发环境
安装自动重新启动服务工具[codegangsta/gin](https://github.com/codegangsta/gin)，这样在开发过程中，修改保存了文件后，会自动重启web服务。
```
$ go get github.com/codegangsta/gin
```
启动服务：
```
$ make dev
```
访问`http://localhost:3030`，测试服务是否正常请访问：`http://localhost:3030/hello`

## 构建和启动正式环境
```
$ make build
$ make start
```

## 清理
```
$ make clean
```

## Docker
- 安装[Docker](https://www.docker.com/) (请设置阿里云等镜像加速器, 不用加速器拉镜像速度很慢)
- 安装[Docker Compose](https://github.com/docker/compose/releases) (一般在Windows和Mac下安装了Docker，同时会安装docker-compose)

> 开发环境是通过源代码的方式运行调试；测试和正式环境镜像只有可执行文件和配置文件，无源代码。

### 构建和启动docker开发环境
生成开发用的docker镜像，并且通过镜像中的环境初始化安装依赖(通过Dep)，如果安装依赖失败请翻墙，或者在本机（非docker下）通过Dep安装，一般会生成`vendor`文件夹，开发环境使用项目根目录下的所有文件做为数据卷。
```
$ make docker-image-dev
$ make start-docker-dev
```
或者在项目根文件夹下通过以下命令启动：
````
$ docker-compose up -d
````

通过docker-compose启动api web, postgres, redis三个容器，api web端口为8080, postgres端口为5432, redis端口为6379。数据库可通过本地客户端工具连接进行操作和调试。

测试接口是否正常启动：请访问`http://localhost:8080/hello`看是否有反馈信息。

### 停止开发环境
````
$ make stop-docker-dev
````
或者在项目根文件夹下运行命令：
````
$ docker-compose down
````

### 生成测试环境docker镜像
修改`Makefile`，可以修改镜像名称
````
$ make docker-image-staging
````

### 生成生产环境docker镜像
修改`Makefile`，可以修改镜像名称
````
$ make docker-image
````

## 注意
- 本脚手架使用[gin](https://github.com/gin-gonic/gin)作为HTTP框架，使用方式请查看官网。也可以换成其它HTTP框架，使用方式类似。
- 由于Go的习惯，`$GOPATH/src/gin-api-boilerplate`为本项目的目录，根据自己的项目情况，请修改项目名称和路径，然后相关`gin-api-boilerplate`名称的都需修改，包括代码中`import`部分，也使用了类似`gin-api-boilerplate/lib/log`这样的名称(Go对相对路径支持不是很好，所以只好使用这种带有包名称的路径)
- [golint](https://github.com/golang/lint)为Go语言的Lint工具。详情请看官网。可以设置代码编辑器支持`golint`动态检测。
- 主配置文件采用json格式，文件位置`data/config/app.json`，配置库使用了[viper](https://github.com/spf13/viper)，详情请查看官网。本项目里，封装到`lib/config`中。
- 封装了[Logrus](https://github.com/sirupsen/logrus)库为日志工具，位置在`lib/log`




