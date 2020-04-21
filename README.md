## Go语言基础学习 参考http://c.biancheng.net/资料
Go语言又成为Golang，由Google公司开发的一种编程语言。
Go语言为编译型语言，目前 Docker、Go-Ethereum、Thrraform 和 Kubernetes均为Go语言发开。
Go语言编写遵循以下几个步骤：
1、创建编写Go程序文件
2、使用Go build编译Go程序，生成执行文件
3、执行编译后的执行程序。

一个Go语言项目的目录一般包含以下三个子目录：
src 目录：放置项目和库的源文件；
pkg 目录：放置编译后生成的包/库的归档文件；
bin 目录：放置编译后生成的可执行文件。

#### 第三方依赖管理工具
```
go get github.com/tools/godep
```
使用：
命令	作用
godep save	将依赖包的信息保存到 Godeps.json 文件中
godep go	使用保存的依赖项运行 go 工具
godep get	下载并安装指定的包
godep path	打印依赖的 GOPATH 路径
godep restore	在 GOPATH 中拉取依赖的版本
godep update	更新选定的包或 go 版本
godep diff	显示当前和以前保存的依赖项集之间的差异
godep version	查看版本信息

#### go module使用
Go语言 1.13 及以后的版本则不再需要设置环境变量。通过 GO111MODULE 可以开启或关闭 go module 工具。

GO111MODULE=off 禁用 go module，编译时会从 GOPATH 和 vendor 文件夹中查找包。
GO111MODULE=on 启用 go module，编译时会忽略 GOPATH 和 vendor 文件夹，只根据 go.mod下载依赖。
GO111MODULE=auto（默认值），当项目在 GOPATH/src 目录之外，并且项目根目录有 go.mod 文件时，开启 go module。

window下开启GO111MODULE 的命令
```
set GO111MODULE=on 或者 set GO111MODULE=auto
```
MacOS 或者 Linux 下开启 GO111MODULE 的命令
```
export GO111MODULE=on 或者 export GO111MODULE=auto
```

#### 调试工具delve
安装：
```
go get -u github.com/go-delve/delve/cmd/dlv
```
使用：

#### 标准库常用包
Go语言标准库常用的包及功能
Go语言标准库包名|功能
:--|:--
bufio|带缓冲的 I/O 操作
bytes|实现字节操作
container|封装堆、列表和环形列表等容器
crypto|加密算法
database|数据库驱动和接口
debug|各种调试文件格式访问及调试功能
encoding|常见算法如 JSON、XML、Base64 等
flag|命令行解析
fmt|格式化操作
go|Go语言的词法、语法树、类型等。可通过这个包进行代码信息提取和修改
html|HTML 转义及模板系统
image|常见图形格式的访问及生成
io|实现 I/O 原始访问接口及访问封装
math|数学库
net|网络库，支持 Socket、HTTP、邮件、RPC、SMTP 等
os|操作系统平台不依赖平台操作封装
path|兼容各操作系统的路径操作实用函数
plugin|Go 1.7 加入的插件系统。支持将代码编译为插件，按需加载
reflect|语言反射支持。可以动态获得代码中的类型信息，获取和修改变量的值
regexp|正则表达式封装
runtime|运行时接口
sort|排序接口
strings|字符串转换、解析及实用函数
time|时间接口
text|文本模板及 Token 词法器
#### 第一个go程序，HelloWorld
项目目录：baseDemoProject1

#### 基础语法学习
项目目录：baseDemoProject2

Go语言的变量声明的标准格式为：
var 变量名 变量类型
名字 := 表达式

:=的方式仅能用于
定义变量，同时显式初始化。
不能提供数据类型。
只能用在函数内部。

“_”本身就是一个特殊的标识符，被称为空白标识符,当遇到不用的返回值，而函数被调用时需接受处理，则这时我们使用该标识符
例如： 
```
test,_ :=getName()
func getName() (string,int) {
    return "aaa",0
}
```
匿名变量不占用内存空间，不会分配内存。匿名变量与匿名变量之间也不会因为多次声明而无法使用。

全局变量声明必须以 var 关键字开头，如果想要在外部包中使用全局变量的首字母必须大写。
Go语言程序中全局变量与局部变量名称可以相同，但是函数体内的局部变量会被优先考虑

常用的转义字符包括：
\n：换行符
\r：回车符
\t：tab 键
\u 或 \U：Unicode 字符
\\：反斜杠自身

Go语言的基本类型有：
bool
string
int、int8、int16、int32、int64
uint、uint8、uint16、uint32、uint64、uintptr
byte // uint8 的别名
rune // int32 的别名 代表一个 Unicode 码
float32、float64
complex64、complex128