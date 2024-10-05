package main

import (
    "fmt"
)
//无限循环的生成从2开始的整数，并将它们送到通道ch
func generate(ch chan int) {
    for i := 2; ; i++ {
        ch <- i //将当前数字发送到通道ch
    }
}

//从in通道中读取整数，检查每个整数能否被prime整除（除以prime是否为0）
//能的话就不可能是质数，如果不能被整除，就将该整数发送到out通道
func filter(in chan int, out chan int, prime int) {
    for {
        num := <-in //从in通道接收一个值
        if num%prime != 0 {
            out <- num //不能被整除，就可能是质数，发送到out通道
        }
    }
}

func main() {
    ch := make(chan int)
    go generate(ch)
    //循环 6 次，每次从通道中读取一个质数，然后启动一个新的过滤器协程 filter，该过滤器会接收来自之前通道 (ch) 的数字，并将剩下的数字发送到新的通道 out。
    for i := 0; i < 6; i++ {
        prime := <-ch //从通道ch接收下一个素数，循环的最后ch=out,ch变为筛选后剩余的数
        fmt.Printf("prime:%d\n", prime)
        out := make(chan int)
        go filter(ch, out, prime)//筛选剩下的数，每当找到质数，就会创建一个新的过滤器 这样能把每次过滤的结果保存
        ch = out  
    }
}
// 这个代码实现了什么功能？
// 实现了一个简单的质数生成器，使用“埃拉托斯特尼筛法”来找出前 6 个质数

// 生成整数：从 2 开始生成自然数
// 筛选质数：通过连续的过滤，将非质数（能被已知质数整除的数）剔除，只保留质数
// 输出质数：打印找到的每个质数
// 这个代码利用了 Go 的什么特性？
// 协程（Goroutines）：使用 go 关键字启动新的协程，使得生成和过滤可以并行进行。这种轻量级的线程模型使得并发编程更为高效。

// 通道（Channels）：通过通道在不同的协程之间传递数据，实现了协程间的同步与通信。这使得程序能够在不同阶段之间高效地传递质数和待筛选的数字。

// 无限循环与动态通道重定向：通过不断创建新的通道来处理下一步的筛选，使得程序能灵活应对不同阶段的需求。

// 这个代码相较于普通写法，是否有性能上的提升？
// 并发处理：由于该代码使用了协程，多个质数的筛选可以并行进行。这种并发方式通常比单线程的方式更快，尤其是在处理大量数据时，因为它能充分利用多核 CPU 的能力
// 实时筛选：与传统的逐个检查每个数是否为质数的方法相比，这种方法在每次找到新质数时就开始筛选，可以提高效率
// 内存管理：使用通道传递数据可以减少内存占用，避免一次性将所有数据加载到内存中，这在处理大范围数字时尤其有效