package Test

import (
	"encoding/json"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var (
	x  int64
	y  int64
	mu sync.Mutex
	wg sync.WaitGroup
	Ch chan int
)

type A struct {
}

type PostAllUserReceiver struct {
	UserName string `json:"UserName"`
	UserID   string `json:"UserID"`
}

func (s *A) Add(e int) *A {
	fmt.Print(e)
	return &A{}
}

func test() {
	//go RotineLoopMaker(50, Ch)
	//RotineReceiver(Ch)
	user := &PostAllUserReceiver{
		//UserName: "aaa",
		//UserID:   "bbb",
	}
	//fmt.Println(JSONSerialize(user))
	json.Unmarshal([]byte("{\"UserName\":\"min\"}"), user)
	fmt.Printf("res:%#v", user)
}

func DeferTest() {
	/*Doc:::
	The expression must be a function or method call; it cannot be
	parenthesized. Calls of built-in functions are restricted as for expression statements;
	Each time a "defer" statement executes, the function value and parameters to the call
	are evaluated as usual and saved anew but the actual function is not invoked.
	*/
	B := &A{}
	defer B.Add(1).Add(2).Add(3).Add(7)
	//[3]:链式调用要先明确父级，不然defer不清楚如何保存函数关系，所以.Add(3)之前必须初始化：（1）(2)(3)（7）
	defer B.Add(4)
	//[2]:正常的defer，由于先入后出原则，调用顺序在上一句之前：（6）
	B.Add(5).Add(6)
	//[1]:正常调用-（4）(5)
	//打印顺序为1,2,3,5,6，4,7
	//若紧跟for{}阻塞主进程，defer的延时调用也会被阻塞，将只会打印1,2,3,5,6
}

func MutexTest() {
	mu.Lock()
	x++
	mu.Unlock()
	wg.Done()
}

func AtomicTest() {
	defer wg.Done()
	atomic.AddInt64(&x, 1)
}

func MutexCombineAtomicTest() {
	start := time.Now()
	for i := 0; i < 100000; i++ {
		go MutexTest()
		wg.Add(1)
	}
	wg.Wait()
	end := time.Now()
	fmt.Println(x)
	fmt.Println(end.Sub(start))
}

func JSONSerialize(res interface{}) string {
	bytes, _ := json.Marshal(res)
	return string(bytes)
}

func JSONDeSerialize(str string, res *any) {
	json.Unmarshal([]byte(str), res)
}

func RotineReceiver(x chan int) {
	//core := runtime.NumCPU()
	//fmt.Printf("cpu core:%d\n", core)
	//runtime.GOMAXPROCS(core / 2)
	/*
		采用select将会减少一半接收到数据,但是使用default将会覆盖case的判断
	*/
	//for {
	//	elem, _ := <-x
	//	select {
	//	case <-x:
	//		if elem%2 == 0 {
	//			fmt.Printf(" loop:%d\n", elem)
	//		} else {
	//			fmt.Printf(" loop:%d\n", elem)
	//		}
	//	default:
	//		fmt.Printf(" uncatched loop:%d\n", elem)
	//	}
	//}

	for {
		elem, _ := <-x
		if elem%2 == 0 {
			fmt.Printf(" loop:%d\n", elem)
		} else {
			fmt.Printf(" loop:%d\n", elem)
		}
	}
}

func RotineLoopMaker(n int, x chan int) {
	for i := 0; i < n; i++ {
		time.Sleep(time.Millisecond * 100)
		x <- i
		if i == n-1 {
			close(x)
		}
	}
}
