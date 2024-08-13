# go-pprof-demo

https://cloud.tencent.com/developer/article/2324701

```sh
curl "http://localhost:8080/debug/pprof/profile?seconds=30" -o profile.pprof
curl "http://localhost:8080/debug/pprof/heap" -o heap.pprof
curl "http://localhost:8080/debug/pprof/goroutine" -o goroutine.pprof
curl "http://localhost:8080/debug/pprof/goroutine?debug=1" -o goroutine_debug1.txt
curl "http://localhost:8080/debug/pprof/goroutine?debug=2" -o goroutine_debug2.txt
curl "http://localhost:8080/debug/pprof/allocs?debug=1" -o allocs_debug1.txt
```

http://localhost:6060/debug/statsviz/

RSS (Resident Set Size)

RSS 表示驻留集大小（Resident Set Size），即进程实际在物理内存中占用的内存量。它包括所有的代码、数据和堆栈段以及共享库的占用部分。RSS 是一个很重要的指标，因为它反映了进程对系统物理内存的实际使用情况。

VSZ (Virtual Memory Size)

VSZ 表示虚拟内存大小（Virtual Memory Size），即进程使用的虚拟地址空间的总量。它包括所有代码段、数据段、堆栈段以及映射到进程地址空间的所有内存（包括共享库和内存映射文件）。VSZ 并不一定反映进程实际占用的物理内存，因为其中可能包含未实际使用的内存区域。

ps -p 25291 -o %mem,rss,vsz

curl http://127.0.0.1:6060/debug/pprof/heap -o heap_info.out