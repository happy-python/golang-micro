![架构图](http://qiniu.rocbj.com/Jietu20190920-165458.png)

# 初始化

    go mod init

# 安装

    go get -u github.com/micro/go-micro
    go get -u github.com/satori/go.uuid
    go get -u github.com/dgrijalva/jwt-go
    go get -u github.com/jinzhu/gorm
    go get -u gopkg.in/mgo.v2

# 参考

> https://www.yinzige.com/2018/05/10/microservices-part-1-introduction-and-consignment-service/

> https://blog.dingkewz.com/post/tech/go_ewan_microservices_in_golang_part_1/

> https://github.com/micro/examples/tree/master/greeter

# 运行 API 网关示例

### 下载示例
    git clone https://github.com/micro/examples

### API Handler


 ##### 运行服务
    
    go run examples/greeter/srv/main.go
    

 ##### 运行 api
    
    go run examples/greeter/api/api.go
    
 
 ##### 启动 micro api
    
    micro api --handler=api
    
 
 ##### 向 api 发起 http 请求
    
    curl "http://localhost:8080/greeter/say/hello?name=jack"
    
 
 ##### 得到输出结果

    {
        "message": "Hello jack"
    }
    

### RPC Handler
 
 ##### 运行服务
    
    go run examples/greeter/srv/main.go
    
 
 ##### 运行 api
    
    go run examples/greeter/api/rpc/rpc.go
    
 
 ##### 启动 micro api
    
    micro api
    
 
 ##### 向 api 发起 http 请求
    
    curl -H 'Content-Type: application/json' -d '{"name": "jack"}' http://localhost:8080/greeter/hello
    
 
 ##### 得到输出结果
    
    {
        "msg": "Hello jack"
    }