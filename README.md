# sendemail-golang
golang send email
https://www.cnblogs.com/apocelipes/archive/2018/08/25/9534885.html
golang包管理解决之道——go modules初探
golang的包管理是一直是为人诟病之处，从golang1.5引入的vendor机制，到准官方工具dep，目前为止还没一个简便的解决方案。

不过现在go modules随着golang1.11的发布而和我们见面了，这是官方提倡的新的包管理，乃至项目管理机制，可以不再需要GOPATH的存在。

 

go modules的初始化

现在modules机制仍在早期阶段，所以golang提供了一个环境变量“GO111MODULE”，默认值为auto，如果当前目录里有go.mod文件，就使用go modules，否则使用旧的GOPATH和vendor机制，因为在modules机制下go get只会下载go modules，这一行为会在以后版本中成为默认值，这里我们保持auto即可，如果你想直接使用modules而不需要从GOPATH过度，那么把“GO111MODULE”设置为on。

modules和传统的GOPATH不同，不需要包含例如src，bin这样的子目录，一个源代码目录甚至是空目录都可以作为module，只要其中包含有go.mod文件。

我们就用一个空目录来创建我们的第一个module：

要初始化modules，需要使用如下命令（假设已经安装配置好golang1.11）：

go mod init [module name]
我们的module叫test，所以就是：

go mod init test
初始完成后会在目录下生成一个go.mod文件，里面的内容只有一行“module test”。

 

包管理

那么我们怎么进行包管理呢？别担心，当我们使用go build，go test以及go list时，go会自动得更新go.mod文件，将依赖关系写入其中。

如果你想手动处理依赖关系，那么使用如下的命令：

go mod tidy
这条命令会自动更新依赖关系，并且将包下载放入cache。

下面我们使用chromedp的一个简单example作为实验代码main.go，看下go modules是如何处理包的依赖关系的。

我们手动运行go mod tidy：
## 返回码必须遵循HTTPd的规范

## server端的API设计
创建（注册）用户：URL:/user Method:POST,SC:201,400,500
200 ok，GET的请求
201 创建成功
400 创建失败
500 内部错误

用户登录:URL/user/:username Method:POST,SC:200,400,500
获取用户基本信息：URL:/user/:username Method:GET,SC:200,400,401,403,500
401:验证不通过，表示并没有验证。
403:验证不通过，表示验证了，但是不具备条件，比如权限

用户注销：URL:/user/:username Method:DELETE,sc:204,400,401,403,500
204 标准的成功，并不用返回给用户

List all videos:URL:/user/:username/videos Method:GET sc:200,400,500
Get one video:URL:/user/:username/videos:/vid-id Method:Get sc:200,400,500
Delete one video:URL:/user/:username/videos/vid-id Method:DELETE,SC:204,400,401,403,500

Show comments:URL:/videos/vid-id/comments Method:Get sc:200,400,500
Post a comment:URL:/videos/vid-id/comments Method:Post SC:201,400,500
Delete a comment:URL:/videos/vid-id/comments Method:DELETE,SC:204,400,401,403,500

https://godoc.org/gopkg.in/gomail.v2
