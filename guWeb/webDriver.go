// https://zetcode.com/golang/chromedp/
// https://github.com/chromedp/examples/blob/master/pdf/main.go
// https://www.burndogfather.com/251
// https://github.com/chromedp/examples/blob/master/submit/main.go

// [크롤러 만들기 (3)](https://sir.kr/so_golang/53)
package guWeb

import (
	"context"
	"log"
	"time"

	"github.com/chromedp/chromedp"
)

func RunWD(url, xpath string, timeOut int, actions ...chromedp.Action) error {
	ctx, cancel := chromedp.NewContext(
		context.Background(),
		chromedp.WithLogf(log.Printf),
	)
	defer cancel()

	ctx, cancel = context.WithTimeout(ctx, time.Duration(timeOut)*time.Second)
	defer cancel()

	err := chromedp.Run(ctx, actions...)
	if err != nil {
		log.Fatal(err)
	}

	return err
}

// * url, xpath -> outerhtml string
func OuterHtmlWD(url, xpath string) (out string) {
	ctx, cancel := chromedp.NewContext(
		context.Background(),
		chromedp.WithLogf(log.Printf),
	)
	defer cancel()

	ctx, cancel = context.WithTimeout(ctx, 15*time.Second) // TODO: timeOut int 입력변수로
	defer cancel()

	err := chromedp.Run(
		ctx,
		chromedp.Navigate(url),
		chromedp.WaitVisible(xpath),
		chromedp.OuterHTML(xpath, &out),
	)
	if err != nil {
		log.Fatal(err)
	}

	return
}

// * url, xpath -> outerhtml string
// func OuterHtmlAfterClick(url, xpathClick, xpathAfter string, timeOut int) (out string) {
func OuterHtmlAfterClick(url, click, xpath string) (out string) {
	ctx, cancel := chromedp.NewContext(
		context.Background(),
		chromedp.WithLogf(log.Printf),
	)
	defer cancel()

	ctx, cancel = context.WithTimeout(ctx, 15*time.Second) // TODO: timeOut int 입력변수로
	// ctx, cancel = context.WithTimeout(ctx, time.Duration(timeOut)*time.Second) // TODO: timeOut int 입력변수로
	defer cancel()

	err := chromedp.Run(
		ctx,
		chromedp.Navigate(url),
		chromedp.WaitVisible(click),
		chromedp.Click(click, chromedp.NodeVisible),
		chromedp.WaitVisible(xpath),
		chromedp.OuterHTML(xpath, &out),
	)
	if err != nil {
		log.Fatal(err)
	}

	return
}

// // * url, xpath -> outerhtml string
// func OuterHtmlAfterLogin(url, xpathClick, xpathAfter string, timeOut int) (out string) {
// 	ctx, cancel := chromedp.NewContext(
// 		context.Background(),
// 		chromedp.WithLogf(log.Printf),
// 	)
// 	defer cancel()

// 	ctx, cancel = context.WithTimeout(ctx, time.Duration(timeOut)*time.Second) // TODO: timeOut int 입력변수로
// 	defer cancel()

// 	err := chromedp.Run(
// 		ctx,
// 		chromedp.Navigate(url),
// 		chromedp.WaitVisible(xpathClick),
//         chromedp.SendKeys("input[name=name]", "Lucia"),
//         chromedp.SendKeys("input[name=message]", "Hello!"),
//         // chromedp.Click("button", chromedp.NodeVisible),
//         chromedp.Submit("input[name=name]"),
// 		chromedp.WaitVisible(xpathClick),
// 		chromedp.OuterHTML(xpathAfter, &out),
// 	)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	return
// }

// * table -> maps by htmlquery
func TableMapsWD(url, xpath string, keys map[int]string) (rsts []map[string]string) {
	return MapsFromTableStr(OuterHtmlWD(url, xpath), keys)
}

// TODO: Click, Login
// type WebDriver struct {
// 	ctx    context.Context
// 	cancel context.CancelFunc
// 	Str    string // 데이터 반환
// }

// func NewCrawler() *WebDriver {
// 	return &WebDriver{}
// }

// // Run 메소드
// func (w *WebDriver) Run(url string) string {
// 	w.ctx, w.cancel = chromedp.NewContext(context.Background())
// 	w.ctx, w.cancel = context.WithTimeout(w.ctx, 30*time.Second)
// 	defer w.cancel()

// 	err := chromedp.Run(
// 		w.ctx,
// 		chromedp.Navigate(url),
// 		chromedp.WaitVisible(`body div#ft`),
// 		chromedp.InnerHTML(`html body div#wrapper div#container_wr div#container div#bo_list form#fboardlist div.tbl_head01.tbl_wrap table tbody tr td.td_subject div.bo_tit a`, &w.Str),
// 	)
// 	if err != nil {
// 		panic(err)
// 	}

// 	return w.Str
// }

// // func main() {
// // 	url := ""
// // 	w := NewCrawler()
// // 	data := w.Run(url)
// // 	log.Println(data)
// // }

// //===================================================
// // [Go언어 chromedp 라이브러리로 웹 크롤링하기](https://streamls.tistory.com/entry/go언어-chromedp-라이브러리로-웹-크롤링하기)
// package main

// import (
// 	"context"
// 	"fmt"
// 	_ "log"
// 	"time"

// 	"github.com/chromedp/cdproto/runtime"
// 	"github.com/chromedp/chromedp"
// )

// func main() {

// 	// chrome 실행 옵션 설정
// 	opts := append(chromedp.DefaultExecAllocatorOptions[:],
// 		chromedp.DisableGPU,
// 		chromedp.Flag("headless", false),	//headless를 false로 하면 브라우저가 뜨고, true로 하면 브라우저가 뜨지않는 headless 모드로 실행됨. 기본값은 true.
// 	)

// 	contextVar, cancelFunc := chromedp.NewExecAllocator(context.Background(), opts...)
// 	defer cancelFunc()

// 	contextVar, cancelFunc = chromedp.NewContext(contextVar)
// 	defer cancelFunc()

// 	// setting 필요 없이 headless 모드로 실행할 경우 아래의 코드로 대체 가능
// 	// contextVar, cancelFunc := chromedp.NewContext(
// 	// 	context.Background(),
// 	// 	chromedp.WithLogf(log.Printf),
// 	// )
// 	//defer cancelFunc()

// 	contextVar, cancelFunc = context.WithTimeout(contextVar, 50*time.Second)	// timeout 값을 설정
// 	defer cancelFunc()

// 	var strVar string
// 	err := chromedp.Run(contextVar,

// 		chromedp.Navigate(`https://www.google.com/`),	// 시작 URL
// 		chromedp.WaitVisible(`body`),		//body 요소를 모두 불러들일 때까지 대기
// 		chromedp.InnerHTML(`pre`, &strVar),	// pre 태그의 내부 텍스트를 strVar변수에 입력
//         chromedp.OuterHTML(`div.srch_employees__list`, &strVar),	// 태크와 그 내부 텍스트를 함께 strVar변수에 입력
// 		chromedp.Click(`.employees-open`, chromedp.NodeVisible),	// .employees-open 클래스의 요소를 클릭
// 		chromedp.Sleep(10*time.Second),		//10초간 대기
// 		chromedp.SetAttributeValue(`input#empSearchKeyWord`, "value", "홍길동", chromedp.NodeVisible),	//input의 value에 '홍길동'을 입력
// 		chromedp.SetAttributes(`input#empSearchKeyWord`, map[string]string{"value": "12345678"}, chromedp.NodeVisible), //input에 value 속성을 추가하고 값을 대입함
// 		chromedp.SendKeys(`input#empSearchKeyWord`, "a"),	//키 입력
// 		chromedp.Evaluate(`window.open('www.naver.com')`, &strVar),	//javascript 실행
// 		chromedp.Value(`div#executiveStatusPop > div.dialog__contents > div.srch_employees > div.srch_employees__container > div.employees__srch > div._fr > div.board-search > input#empSearchKeyWord label`, &strVar),	// 해당 요소의 텍스트값을 취하여 strVar변수에 대입
// 	)
// 	if err != nil {
// 		panic(err)
// 	}
// }

// //==========================================================
// // [[개인프로잭트_레시피추천]golang chromedp](https://velog.io/@qkqk2938/golang-chromedp)

// package main

// import (
// 	"log"
// 	"context"
// 	"time"
// 	"fmt"
// 	"github.com/chromedp/chromedp"
// 	//"github.com/chromedp/cdproto/cdp"
// 	//"github.com/chromedp/cdproto/runtime"
// )

// func main() {
// 	linklist := getLinkList()
// 	for _, val := range linklist{
// 		getDescription(val)
// 	}

// }

// func getDescription(url string){
// 	fmt.Println(url)
// 	contextVar, cancelFunc := chromedp.NewContext(
// 		context.Background(),
// 		chromedp.WithLogf(log.Printf),
// 	)
// 	defer cancelFunc()

// 	contextVar, cancelFunc = context.WithTimeout(contextVar, 30*time.Second)	// timeout 값을 설정
// 	defer cancelFunc()
// 	contextVar, cancelFunc = chromedp.NewContext(contextVar)
// 	defer cancelFunc()

// 	var strVar string
// 	err := chromedp.Run(contextVar,
// 		chromedp.Navigate("https://www.youtube.com"+url),
// 		chromedp.Click("#primary div#primary-inner div#below ytd-watch-metadata div#above-the-fold div#bottom-row div#description tp-yt-paper-button#expand-sizer", chromedp.ByID ),
// 		chromedp.Text("#primary div#primary-inner div#below ytd-watch-metadata div#above-the-fold div#bottom-row div#description", &strVar,chromedp.ByID ),
// 		//chromedp.Text("#primary div#primary-inner div#below ytd-watch-metadata div#above-the-fold div#bottom-row div#description tp-yt-paper-button#expand-sizer", &attr,chromedp.ByQueryAll ),
// 	)
// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Println(strVar)
// }

// func getLinkList() []string {

// 	contextVar, cancelFunc := chromedp.NewContext(
// 		context.Background(),
// 		chromedp.WithLogf(log.Printf),
// 	)
// 	defer cancelFunc()

// 	contextVar, cancelFunc = context.WithTimeout(contextVar, 300*time.Second)	// timeout 값을 설정
// 	defer cancelFunc()

// 	err := chromedp.Run(contextVar,
// 		chromedp.Navigate("https://www.youtube.com/@paik_jongwon/videos"),
// 	)
// 	if err != nil {
// 		panic(err)
// 	}

// 	var oldHeight int
// 	var newHeight int
// 	for {

// 		err = chromedp.Run(contextVar,
// 			chromedp.Evaluate(`window.scrollTo(0,document.querySelector("body ytd-app div#content").clientHeight); document.querySelector("body ytd-app div#content").clientHeight;`, &newHeight),
// 			chromedp.Sleep(700*time.Millisecond),
// 		)
// 		if err != nil {
// 			panic(err)
// 		}
// 		if(oldHeight == newHeight){
// 			break
// 		}
// 		oldHeight = newHeight
// 	}
// 	//var strVar string
// 	//var strTitle string
// 	attr := make([]map[string]string, 0)
// 	//var nodes []cdp.NodeID
// 	err = chromedp.Run(contextVar,

// 		chromedp.AttributesAll("#primary ytd-rich-grid-renderer div#contents ytd-rich-grid-row div#contents ytd-rich-item-renderer #video-title-link", &attr,chromedp.ByQueryAll ),

// 	)
// 	if err != nil {
// 		panic(err)
// 	}

// 	var linklist []string
// 	for _, val := range attr {
// 		linklist = append(linklist, val["href"])
// 	}
// 	fmt.Println(len(linklist))
// 	return linklist
// }
