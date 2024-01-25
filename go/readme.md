### 编译选项
- GOARCH, 目标平台的 CPU 架构. 常用的值 amd64, arm64, i386, armhf
- GOOS, 目标平台, 常用的值 linux, windows, drawin (macOS)
- GOARM, 只有 GOARCH 是 arm64 才有效, 表示 arm 的版本, 只能是 5, 6, 7 其中之一
- CGO_ENABLED, 是否支持 CGO 交叉汇编, 值只能是 0, 1, 默认情况下是 0, 启用交叉汇编比较麻烦
- CC, 当支持交叉汇编时(即 CGO_ENABLED=1), 编译目标文件使用的 c 编译器.
- CXX, 当支持交叉汇编时(即 CGO_ENABLED=1), 编译目标文件使用的 c++ 编译器.
- AR, 当支持交叉汇编时(即 CGO_ENABLED=1), 编译目标文件使用的创建库文件命令.


### cgo

- e1

```c
#cgo windows,amd64 CFLAGS: -D__USE_MINGW_ANSI_STDIO=1
// #cgo: 这是一个cgo的指令，用于指导cgo工具如何处理下面的C语言代码。
// windows,amd64: 这是一个条件编译的标签，表示后面的C语言代码仅在Windows平台的amd64架构上编译。这就是在交叉编译时，你希望对不同平台提供不同的实现的一种方式。
// CFLAGS: -D__USE_MINGW_ANSI_STDIO=1: 这是传递给C编译器的选项，其中 -D__USE_MINGW_ANSI_STDIO=1 是一个预处理器宏，用于启用 MinGW 版本的 ANSI 标准 I/O 函数。
// -D 是告诉编译器定义一个宏。
// __USE_MINGW_ANSI_STDIO 是一个预定义的宏，用于控制 MinGW 编译器如何处理标准 I/O 函数的缓冲区。
```


