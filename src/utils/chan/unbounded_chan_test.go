/**
* @Description：
* @Author: cdx
* @Date: 2022/12/4 2:52 下午
 */

package _chan

import (
	"fmt"
	"sync"
	"testing"
)

func TestMakeUnboundedChan(t *testing.T) {

	ch := NewUnboundedChan(100)
	for i := 1; i < 1000; i++ {
		ch.In <- int64(i)
	}
	close(ch.In)

	var count int64
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()

		for v := range ch.Out {
			fmt.Println("out = ", v.(int64))
			count += v.(int64)
		}
	}()
	wg.Wait()

	fmt.Println("count = ", count)
}
