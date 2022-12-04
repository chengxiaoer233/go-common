
### 1：golang 无限制长度的channel   
* 使用
    + 初始化一个不限制长度的channel
    ```go
      ch := NewUnboundedChan(100)
    ```
    
    + 往channel中写入数据，支持并发写入
        ```go
          ch.In <- int64(i)
        ```
      
    + 从channel中读出数据
        ```go
      for v := range ch.Out {
      	fmt.Println("out = ", v.(int64))      
      }
      ``` 