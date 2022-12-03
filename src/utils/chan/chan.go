/**
* @Description：
* @Author: cdx
* @Date: 2022/12/3 9:38 下午
 */

package _chan

import (
	"fmt"
	"sync"
	"time"
)

// （1）控制每秒最多只有n和协程去处理某个任务
// 有100个并发，但是最多同时只能10个执行某个任务
func f1() {

	var wg sync.WaitGroup
	var ch = make(chan struct{}, 10)
	to := time.Now()
	for i := 0; i < 100; i++ {
		wg.Add(1)

		go func(i int, wg *sync.WaitGroup) {
			defer wg.Done()

			ch <- struct{}{} // 往chan中插入数据，同时最多只能insert 10个，模拟控制并发数
			fmt.Println("insert into success, i=", i)
			time.Sleep(1 * time.Second) // sleep 1s
			<-ch                        // channel缓冲减一，可被其他协程获取到

		}(i, &wg)
	}
	wg.Wait()

	// 每秒最大并发10个，需要10s可以执行完这100个任务
	fmt.Println("time use=", time.Since(to))
}

// （2）模拟tcp的nagle算法，上游任务来了后积攒一批，满足一定数目或者时间到了就执行
// 通过channel和定时器
var AddChan = make(chan string, 10000)
func f2() {

	// 临时tmpUrls及预置的队列阈值大小
	size := 100
	tmpUrls := make([]string, 0, size)

	// 定时器，每60s触发一次
	timeStamp := 6 * time.Second
	ticker := time.NewTicker(timeStamp)
	defer ticker.Stop()

	for {
		select {
		// 定时时间到了
		case <-ticker.C:
			if len(tmpUrls) > 0 {

				// 模拟业务，do something
				// tmoUrls 重新清空
				tmpUrls = make([]string, 0, size)
			}

		// AddChan就是所谓的全局队列，有数据到达了
		case url, ok := <-AddChan:
			if !ok { // 通道关闭

				// 模拟业务，do something

				// tmoUrls 重新清空
				tmpUrls = make([]string, 0, size)
				return

			} else { // 有数据写入

				tmpUrls = append(tmpUrls, url)
				if len(tmpUrls) >= size { // 大于设定的值，条件触发

					// 模拟业务，do something

					// tmoUrls 重新清空
					tmpUrls = make([]string, 0, size)

					// 重置定时器（这一步不能少）
					ticker.Reset(timeStamp)
				}
			}
		}
	}
	return
}

// 通过时间戳和加锁判断队列
func f3() {

}
