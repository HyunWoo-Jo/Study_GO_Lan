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

func crawl(url string) {
	fetched.Lock()                   // 뮤택스로 보호
	if _, ok := fetched.m[url]; ok { // Url 중복 처리 여부를 검사
		fetched.Unlock()
		return
	}
	fetched.Unlock()

	doc, err := fetch(url) // URL에서 파싱된 HTML 데이터를 가져옴

	fetched.Lock()
	fetched.m[url] = err // 사용자 정보 출력 ,팔로잉 URL을 구함
	fetched.Unlock()

	urls := parseFlollwing(doc)

	done := make(chan bool)
	for _, u := range urls { //URL 개수만큼
		go func(url string) { // 고루틴 생성
			crawl(url) //재귀 호출
			done <- true
		}(u)
	}
	for i := 0; i < len(urls); i++ {
		<-done // 고루틴으 모두 실행 될 때가지 대기
	}
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

func parseFlollwing(doc *html.Node) []string {
	var urls = make([]string, 0)
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "img" { // img 태그
			for _, a := range n.Attr {
				if a.Key == "class" && a.Val == "avatar left" { // class가 avatar left인 요소
					for _, a := range n.Attr {
						if a.Key == "alt" {
							fmt.Println(a.Val) // 사용자 이름 출력
							break
						}
					}
				}
				if a.Key == "class" && a.Val == "gravatar" { // class가 gravatar인 요소
					user := n.Parent.Attr[0].Val                                // 부모 요소의 첫 번째 속성(href)
					urls = append(urls, "https://github.com"+user+"/following") // 사용자 이름으로 팔로잉 URL 조합
					break
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c) // 재귀 호출로 자식과 형제를 모두 탐색
		}
	}
	f(doc)
	return urls
}

func main() {
	numCPUs := runtime.NumCPU()
	runtime.GOMAXPROCS(numCPUs)

	crawl("https://github.com/pyrasis/following")
}
