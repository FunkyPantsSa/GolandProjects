
package main

import (
	"github.com/valyala/fasthttp"
	"go.uber.org/zap"
)

var logger *zap.Logger

func init() {
	logger, _ = zap.NewProduction()
}

func fastHTTPHandler(ctx *fasthttp.RequestCtx) {
	logger.Info("hello, go module", zap.ByteString("uri", ctx.RequestURI()))
}

func main() {
	fasthttp.ListenAndServe(":8081", fastHTTPHandler)
}

// cmd: go mod init github.com/bigwhite/hellomodule
/* 一个 module 就是一个包的集合，这些包和 module 一起打版本、发布和分发。go.mod 所在的目录被我们称为它声明的 module 的根目录。不过呢，这个时候的 go.mod 文件内容还比较简单，第一行内容是用于声明 module 路径（module path）的。而且，module 隐含了一个命名空间的概念，module 下每个包的导入路径都是由 module path 和包所在子目录的名字结合在一起构成。比如，如果 hellomodule 下有子目录 pkg/pkg1，那么 pkg1 下面的包的导入路径就是由 module path（github.com/bigwhite/hellomodule）和包所在子目录的名字（pkg/pkg1）结合而成，也就是 github.com/bigwhite/hellomodule/pkg/pkg1

 */
