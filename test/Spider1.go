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

func DoWork(start, end int) {
	fmt.Printf("正在爬取 %d 到 %d 的页面", start, end)

	for i := start; i <= end; i++ {
		url := "http://202.116.83.50/hope/Diaries/Index_" + strconv.Itoa(i) + ".aspx"
		fmt.Println("url = ", url)
		res, err := HttpGet(url)
		if err != nil {
			fmt.Println("Http Get err = ", err)
			continue
		}

		filename := strconv.Itoa(i) + ".html"
		f, err1 := os.Create(filename)
		if err1 != nil {
			fmt.Println("os.Create err1 = ", err1)
			return
		}

		f.WriteString(res)
		defer f.Close()
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
