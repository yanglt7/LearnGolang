# golang 11 并发编程



### 11.1 概述

#### 11.1.1 并发和并行
#### 11.1.2 Go 语言并发优势

### 11.2 goroutine

#### 11.2.1 goroutine 是什么
#### 11.2.2 创建 goroutine
#### 11.2.3 主 goroutine 先退出
#### 11.2.4 runtime 包

##### 11.2.4.1 Gosched
##### 11.2.4.2 Goexit
##### 11.2.4.3 GOMAXPROCS

### 11.3 channel

#### 11.3.1 channel 类型
#### 11.3.2 无缓冲的 channel
#### 11.3.3 有缓冲的 channel
#### 11.3.4 range 和 close
#### 11.3.4 单方向的 channel
#### 11.3.6 定时器
##### 11.3.6.1 Timer
##### 11.3.6.2 Ticker

### 11.4 select

#### 11.4.1 select 作用
#### 11.4.2 超时