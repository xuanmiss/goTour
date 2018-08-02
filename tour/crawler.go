package main

/**
 *Created By miss
 *
 *At 2018/7/31
 */
import (
	"fmt"
	"sync"
)

type Fetcher interface {
	// Fetch 返回 URL 的 body 内容，并且将在这个页面上找到的 URL 放到一个 slice 中。
	Fetch(url string) (body string, urls []string, err error)
}

var (
	//  map  存放爬取的url
	mm = make(map[string]int)
	// 互斥锁
	l sync.Mutex
	//   群组等待     当添加的任务没有完成时（done（））， wait（） 会一直等待    三个方法   Add()  Done()  Wait()
	i sync.WaitGroup
)

func main() {
	i.Add(1)
	Crawl("http://golang.org/", 4, fetcher)

	i.Wait() // 会一直等待直到子线程任务结束

	for k, _ := range mm {
		fmt.Println(k)
	}
	fmt.Println("over")
}

// Crawl 使用 fetcher 从某个 URL 开始递归的爬取页面，直到达到最大深度。
func Crawl(url string, depth int, fetcher Fetcher) {

	defer i.Done() //  和add相对应
	if depth <= 0 {
		return
	}
	_, urls, err := fetcher.Fetch(url)
	if err != nil {
		//  fmt.Println(err)
		return
	}
	// 存入数据  需要同步锁  因为这是在子线程中
	l.Lock()
	if mm[url] == 0 { //  还未爬取过
		mm[url]++ // 存入爬取的url  改变对应的标示
		depth--
		//      fmt.Printf("found: %s %q\n", url, body)
		for _, u := range urls {
			i.Add(1)
			go Crawl(u, depth, fetcher) // 继续爬取
		}
	}
	l.Unlock()

}

// fakeFetcher 是返回若干结果的 Fetcher。
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher 是填充后的 fakeFetcher。
var fetcher = fakeFetcher{
	"http://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"http://golang.org/pkg/",
			"http://golang.org/cmd/",
		},
	},
	"http://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"http://golang.org/",
			"http://golang.org/cmd/",
			"http://golang.org/pkg/fmt/",
			"http://golang.org/pkg/os/",
		},
	},
	"http://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
	"http://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
}
