package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

// go run main.go http://www.baidu.com http://weibo.com http://qq.com
func main() {
	start := time.Now()
	// 加第二个参数后，通道就是非阻塞的
	ch := make(chan string, len(os.Args[1:]))

	for _, url := range os.Args[1:] {
		// 启动一个 goroutine
		go fetch(url, ch)
	}

	for range os.Args[1:] {
		// 从通道 ch 接收数据
		fmt.Println(<-ch)
	}

	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)

	if err != nil {
		// 发送数据到通道 ch
		ch <- fmt.Sprint(err)
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	// 关闭资源，避免泄漏
	resp.Body.Close()

	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}

	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}
