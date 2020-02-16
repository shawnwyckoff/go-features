------------- library提供方 -------------

1、编写library代码，然后编译

library源码路径为{original-src-path}，比如github.com/djz/libdemo
编译有两种方法
go install // 把生成的.a文件输出到了$GOPATH/pkg/darwin_amd64
go build -buildmode=archive -o libdemo.a // 把生成的.a文件输出到当前目录（貌似不管用）


2、生成头文件

将函数体全部删掉，需要return的函数随便return什么都不影响



------------- library使用方 -------------

1、拷贝静态库和头文件

将静态库.a拷贝到$GOPATH/pkg/$PLATFORM/{user-choose-path}

将头文件.go拷贝到$GOPATH/src/{user-choose-path}/libdemo

主程序要使用{user-choose-path}/libdemo作为该库的import的路径，比如libs123/github.com/djz/libdemo或者github.com/djz/libdemo，它和{original-src-path}可以相同也可以不同，只要最后一级包名（package）相同即可。


2、编译最终程序

如果是前者，按如下方法编译
go tool compile -I $GOPATH/pkg/darwin_amd64 myapp.go
go tool link -o myapp -L $GOPATH/pkg/darwin_amd64 myapp.o

如果是后者，按如下方法编译（貌似不管用）
go tool compile -I . myapp.go
go tool link -o myapp -L . myapp.o




参考文章：
https://segmentfault.com/q/1010000014609549
https://www.jianshu.com/p/b3fb41d7c33f
