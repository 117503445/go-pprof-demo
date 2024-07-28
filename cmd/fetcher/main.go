package main

import (
	"os"
	"time"

	"github.com/imroc/req/v3"
)

// curl "http://localhost:8080/debug/pprof/heap" -o heap.pprof
// curl "http://localhost:8080/debug/pprof/goroutine" -o goroutine.pprof
// curl "http://localhost:8080/debug/pprof/goroutine?debug=1" -o goroutine_debug1.txt
// curl "http://localhost:8080/debug/pprof/goroutine?debug=2" -o goroutine_debug2.txt
// curl "http://localhost:8080/debug/pprof/allocs?debug=1" -o allocs_debug1.txt
// curl "http://localhost:8080/debug/pprof/profile?seconds=30" -o profile.pprof

const host = "http://localhost:8080"

func main() {
	for {
		dir := "./data/" + time.Now().Format("2006-01-02-15-04-05")
		if os.MkdirAll(dir, 0755) != nil {
			panic("failed to create dir")
		}

		cases := []struct {
			url string
			out string
		}{
			{"/debug/pprof/heap", "/heap.pprof"},
			{"/debug/pprof/goroutine", "/goroutine.pprof"},
			{"/debug/pprof/goroutine?debug=1", "/goroutine_debug1.txt"},
			{"/debug/pprof/goroutine?debug=2", "/goroutine_debug2.txt"},
			{"/debug/pprof/allocs?debug=1", "/allocs_debug1.txt"},
			{"/debug/pprof/profile?seconds=30", "/profile.pprof"},
		}

		for _, c := range cases {
			req.R().SetOutputFile(dir + c.out).MustGet(host + c.url)
		}
		time.Sleep(1 * time.Minute)
	}
}
