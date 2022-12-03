/**
* @Description：
* @Author: cdx
* @Date: 2022/12/3 9:49 下午
 */

package _chan

import "testing"

func Test_f1(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			"chan func1 : 控制并发数",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f1()
		})
	}
}

func Test_f2(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			"chan func2.0 : 队列满足大小条件或定时时间到了触发",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f2()
		})
	}
}