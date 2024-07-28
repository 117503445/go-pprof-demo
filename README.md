# go-pprof-demo

```sh
curl "http://localhost:8080/debug/pprof/profile?seconds=30" -o profile.pprof
curl "http://localhost:8080/debug/pprof/heap" -o heap.pprof
curl "http://localhost:8080/debug/pprof/goroutine" -o goroutine.pprof
curl "http://localhost:8080/debug/pprof/goroutine?debug=1" -o goroutine_debug1.txt
curl "http://localhost:8080/debug/pprof/goroutine?debug=2" -o goroutine_debug2.txt
curl "http://localhost:8080/debug/pprof/allocs?debug=1" -o allocs_debug1.txt
```
