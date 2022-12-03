/**
* @Description：
* @Author: cdx
* @Date: 2022/12/2 9:20 下午
 */

package md5

import "testing"

func TestFileMd5(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"md5file test",
			args{filename: "./md5file"},
			"098f6bcd4621d373cade4e832627b4f6",
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FileMd5(tt.args.filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("FileMd5() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("FileMd5() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringMd5(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"string md5 1",
			args{s: "1111"},
			"b59c67bf196a4758191e42f76670ceba",
		},

		{
			"string md5 2",
			args{s: "2222"},
			"934b535800b1cba8f96a5d72f72f1611",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringMd5(tt.args.s); got != tt.want {
				t.Errorf("StringMd5() = %v, want %v", got, tt.want)
			}
		})
	}
}
