package main

import (
	"fmt"
	"github.com/shawnwyckoff/go-utils/sys/chans"
	"sync"
	"time"
)

// 这个demo演示了如何编写合格的线程

// wg 不可以传对象只能传指针，否则，routine中的wg将是另一个wg，和main函数中是两个不同的对象，要避免这种传参方式可以用闭包
func routine(exitMsg chan struct{}, wg *sync.WaitGroup, outputch chan interface{}) {
	defer wg.Add(-1)
	defer fmt.Println("\nRoutine exit")

	tag := 0
	for {
		// Check exit message.
		select {
		case <-exitMsg:
			return
			// 必须要default分支，起到检查exitMsg是否有内容或者是否被close的左右
		default:
		}

		for i := 0; i < 10; i++ {
			tag++
			time.Sleep(time.Millisecond * 100)

			// 下面这一段向outputch输出数据的，不要简单写作outputch <- tag
			// 而要用select，这样可以避免因为outputch满了，写入阻塞，从而导致主进称的exitMsg信号无法及时送达
			select {
			case <-exitMsg:
				return
			case outputch <- tag:
				// 注意，这里不可以开启default分支，假如开启了，会因为outputch写满而直接进入default分支然后忽略了很多数字
				//default:
			}
		}
	}
}

func main() {
	// 向exitMsg写入数据或者close它，都会导致在select语句中可以从exitMsg中取到东西
	exitMsg := make(chan struct{})
	wg := sync.WaitGroup{}
	output := make(chan interface{}, 2)

	wg.Add(1)
	go routine(exitMsg, &wg, output)

	go func() {
		for {
			time.Sleep(time.Second)
			n := <-output
			fmt.Print(n.(int), ",")
		}
	}()
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
	}
	chans.SafeCloseChanStruct(exitMsg)

	wg.Wait()
	fmt.Println("Application exit")
}
