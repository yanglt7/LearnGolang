package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
)

func HttpGet(url string) (res string, err error) {
	resp, err1 := http.Get(url)
	if err1 != nil {
		err = err1
		return
	}

	defer resp.Body.Close()

	buf := make([]byte, 4*1024)

	for {
		n, _ := resp.Body.Read(buf)
		if n == 0 {
			break
		}

		res += string(buf[:n])
	}

	return
}

func SpidePage(i int, page chan<- int) {
	url := "http://202.116.83.50/hope/Diaries/Index_" + strconv.Itoa(i) + ".aspx"
	fmt.Printf("正在爬取第%d页网页%s\n", i, url)
	res, err := HttpGet(url)
	if err != nil {
		fmt.Println("Http Get err = ", err)
		return
	}

	filename := strconv.Itoa(i) + ".html"
	f, err1 := os.Create(filename)
	if err1 != nil {
		fmt.Println("os.Create err1 = ", err1)
		return
	}

	f.WriteString(res)
	defer f.Close()

	page <- i
}

func DoWork(start, end int) {
	fmt.Printf("正在爬取 %d 到 %d 的页面\n", start, end)

	page := make(chan int)

	for i := start; i <= end; i++ {
		go SpidePage(i, page)
	}

	for i := start; i <= end; i++ {
		fmt.Printf("第%d个页面爬取完成\n", <-page)
	}

}

func main() {
	var start, end int
	fmt.Printf("请输入要爬取的起始页：")
	fmt.Scan(&start)
	fmt.Printf("请输入要爬取的终止页：")
	fmt.Scan(&end)

	DoWork(start, end)

}
