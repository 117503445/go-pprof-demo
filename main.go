package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"os"
	"runtime"
	"runtime/debug"
	"time"

	"github.com/arl/statsviz"
	"github.com/shirou/gopsutil/process"
)

var DateTime = "2006-01-02 15:04:05"

// 提升资源层面的可观测性
func main() {
	statsviz.RegisterDefault() // 实时查看 Go 应用程序运行时统计信息(GC，MemStats 等)

	// 打印GC信息
	go printGCStats()
	// pprof
	go func() {
		fmt.Println(http.ListenAndServe(fmt.Sprintf("127.0.0.1:%s", "6060"), nil))
	}()

	// 打印cpu和内存的使用信息
	go func() {

		pid := os.Getpid()
		fmt.Println("当前程序的进程号为：", pid)

		p, _ := process.NewProcess(int32(pid))
		for {
			v, _ := p.CPUPercent()
			if v > 0 {
				memPercent, _ := p.MemoryPercent()
				fmt.Printf("该进程的cpu占用率:%v,内存占用:%v, 时间:%v\n", v, memPercent, time.Now().Format(DateTime))
				println("---------------分割线------------------")
			}
			time.Sleep(5 * time.Second)
		}
	}()

	fmt.Printf("最初！程序中goroutine的数量为:%d\n", runtime.NumGoroutine())

	// for i := 0; i < 1000000; i++ {
	// 	go func() {
	// 		time.Sleep(time.Second * 10)
	// 	}()
	// }

	time.Sleep(5e9)
	fmt.Printf("for循环结束后！程序中goroutine的数量为:%d\n", runtime.NumGoroutine())

	time.Sleep(5e9)

	fmt.Printf("5s后程序中goroutine的数量为:%d\n", runtime.NumGoroutine())

	time.Sleep(5e9)
	fmt.Printf("10s后程序中goroutine的数量为:%d\n", runtime.NumGoroutine())

	time.Sleep(5e9)
	fmt.Printf("15s后程序中goroutine的数量为:%d\n", runtime.NumGoroutine())

	time.Sleep(5e9)
	fmt.Printf("20s后程序中goroutine的数量为:%d\n", runtime.NumGoroutine())

	// 用于阻塞不使程序退出
	select {}

}

func printGCStats() {
	t := time.NewTicker(time.Second)
	s := debug.GCStats{}
	for {
		select {
		case <-t.C:
			debug.ReadGCStats(&s)
			fmt.Printf("gc %d last@%v, PauseTotal %v\n", s.NumGC, s.LastGC, s.PauseTotal)
		}
	}
}
