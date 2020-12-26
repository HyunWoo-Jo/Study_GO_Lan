package main

import (
	"fmt"
	_ "fmt"
	"log"
	_ "log"
	"net/http"
	_ "net/http"
	"runtime"
	_ "runtime"
	"sync"
	_ "sync"

	"golang.org/x/net/html"
	_ "golang.org/x/net/html" // HTML 파싱 패키지
)

var fetched = struct {
	m map[string]error
	sync.Mutex
}{m: make(map[string]error)}

type result struct {
	url  string
	name string
}

func fetch(url string) (*html.Node, error) {
	res, err := http.Get(url) // url 에서 html 데이터를 가지고옴
	if err != nil {
		log.Println(err)
		return nil, err
	}

	doc, err := html.Parse(res.Body) // res.Body를 넣으면 파싱된 데이터가 리턴됨
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return doc, nil
}

func parseFollowing(doc *html.Node, urls chan string) <-chan string {
	name := make(chan string)
	go func() {
		var f func(*html.Node)
		f = func(n *html.Node) {
			if n.Type == html.ElementNode && n.Data == "img" { // img 태그
				for _, a := range n.Attr {
					if a.Key == "class" && a.Val == "avatar avatar-user" { // class가 avatar avatar-use인 요소
						for _, a := range n.Attr {
							if a.Key == "alt" {
								name <- a.Val
								break
							}
						}
					}
					if a.Key == "class" && a.Val == "gravatar" { // class가 gravatar인 요소
						user := n.Parent.Attr[0].Val                           // 부모 요소의 첫 번째 속성(href)
						urls <- "https://github.com" + user + "?tab=following" // 사용자 이름으로 팔로잉 URL 조합
						break
					}
				}
			}
			for c := n.FirstChild; c != nil; c = c.NextSibling {
				f(c) // 재귀 호출로 자식과 형제를 모두 탐색
			}
		}
		f(doc)
	}()
	return name
}
func worker(done <-chan struct{}, urls chan string, c chan<- result) {
	for url := range urls {
		select {
		case <-done:
			return
		default:
			crawl(url, urls, c)
		}
	}
}

func crawl(url string, urls chan string, c chan<- result) {
	fetched.Lock()
	if _, ok := fetched.m[url]; ok {
		fetched.Unlock()
		return
	}
	fetched.Unlock()

	doc, err := fetch(url)
	if err != nil {
		go func(u string) {
			urls <- u
		}(url)
		return
	}

	fetched.Lock()
	fetched.m[url] = err
	fetched.Unlock()

	name := <-parseFollowing(doc, urls)
	c <- result{url, name}
}

func main() {
	numCPUs := runtime.NumCPU()
	runtime.GOMAXPROCS(numCPUs)

	urls := make(chan string)
	done := make(chan struct{})
	c := make(chan result)

	var wg sync.WaitGroup
	const numWorkers = 10
	wg.Add(numWorkers)
	for i := 0; i < numWorkers; i++ {
		go func() {
			worker(done, urls, c)
			wg.Done()
		}()
	}
	go func() {
		wg.Wait()
		close(c)
	}()
	urls <- "https://github.com/JanDeDobbeleer?tab=following"

	count := 0
	for r := range c {
		fmt.Println(r.name)

		count++
		if count > 100 {
			close(done)
			break
		}
	}
	fmt.Println("End")
}
