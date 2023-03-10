Go 运行时是按照“pkg3 -> pkg1 -> pkg2 -> main”的顺序，来对 Go 程序的各个包进行初始化的，而在包内，则是以“常量 -> 变量 -> init 函数”的顺序进行初始化。此外，main 包的两个 init 函数，会按照在源文件 main.go 中的出现次序进行调用。还有一点，pkg1 包和 pkg2 包都依赖 pkg3 包，但根据 Go 语言规范，一个被多个包依赖的包仅会初始化一次，因此这里的 pkg3 包仅会被初始化了一次。

所以简而言之，记住 Go 包的初始化次序并不难，你只需要记住这三点就可以了：
依赖包按“深度优先”的次序进行初始化；
每个包内按以“常量 -> 变量 -> init 函数”的顺序进行初始化；
包内的多个 init 函数按出现次序进行自动调用。