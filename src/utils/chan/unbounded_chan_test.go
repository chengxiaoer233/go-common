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

	ch := NewUnboundedChan(10)

	for i := 1; i < 100; i++ {
		ch.In <- int64(i)
	}
	close(ch.In)

	var count int64
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()

		for v := range ch.Out {
			count += v.(int64)
		}
	}()
	wg.Wait()

	fmt.Println("count = ", count)
}

func TestMakeUnboundedChan1(t *testing.T) {

	ch := NewUnboundedChan(10)

	// make some goroutine to write into in
	var inWg sync.WaitGroup
	for i := 1; i < 100; i++ {
		inWg.Add(1)
		go func(i int) {
			defer inWg.Done()
			ch.In <- int64(i)
		} (i)
	}
	inWg.Wait()
	close(ch.In)

	var count int64
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()

		for v := range ch.Out {
			count += v.(int64)
		}
	}()
	wg.Wait()

	fmt.Println("count = ", count)
}
