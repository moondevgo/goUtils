package guWeb

import (
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	playwright "github.com/playwright-community/playwright-go"
)

const (
	HEADLESS   = false // headless 설정
	SEPARATOR  = "\n__<SEP>__\n"
	WRAP_START = "<div>\n<table>\n"   // ? apps script에서 사용하는 Cheerio 용
	WRAP_END   = "\n</table>\n</div>" // ? apps script에서 사용하는 Cheerio 용
)

// var (
// 	PWO = PWObj{}
// )

// type NodePW interface {
// 	playwright.Page | playwright.Frame
// }

type PWObj struct {
	PW      *playwright.Playwright
	Browser playwright.Browser
	Context playwright.BrowserContext
	Page    playwright.Page
	Frame   playwright.Frame
}

// ** Util Functions
func writeStr(path string, data string) {
	f, _ := os.Create(path)
	f.WriteString(data)
}

func setXpath(xpath string) string {
	if len(xpath) < 2 {
		return ""
	}
	if xpath[0:1] != "x" {
		xpath = "xpath=" + xpath
	}
	return xpath
}

// * 초기화
func InitPlayWright() PWObj {
	// url, iframe, xpath
	pw, err := playwright.Run()
	if err != nil {
		log.Fatalf("could not start playwright: %v", err)
	}

	headless := HEADLESS
	browser, err := pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{Headless: &headless})
	if err != nil {
		log.Fatalf("could not launch browser: %v", err)
	}

	context, err := browser.NewContext()
	if err != nil {
		log.Fatalf("could not launch browser: %v", err)
	}

	page, err := context.NewPage()
	if err != nil {
		log.Fatalf("could not create page: %v", err)
	}

	return PWObj{
		PW:      pw,
		Context: context,
		Browser: browser,
		Page:    page,
		Frame:   nil,
	}
}

// * Url 이동
func (pwo *PWObj) GotoUrl(url string) {
	if _, err := pwo.Page.Goto(url); err != nil {
		log.Fatalf("could not goto: %v", err)
	}
}

// * Frame 이동
func (pwo *PWObj) GotoFrame(xpath string) {
	frameElement, err := pwo.Page.QuerySelector(setXpath(xpath))
	if err != nil {
		log.Fatalf("could not find #notice iframe: %v\n", err)
	}
	frame, err := frameElement.ContentFrame()
	if err != nil {
		log.Fatalf("could not get content frame: %v\n", err)
	}

	pwo.Frame = frame
}

// * Frame 제거
func (pwo *PWObj) LeaveFrame(xpath string, wait string) {
	pwo.Frame = nil
}

// * Wait
// func (pwo *PWObj) Wait(option string) {
func (pwo *PWObj) Wait(xpath string) { // TODO: option: xpath, type 적용
	wait := setXpath(xpath)
	if len(wait) > 2 {
		enabled, _ := pwo.Page.IsVisible(wait)
		for !enabled {
			enabled, _ = pwo.Page.IsVisible(wait)
		}
	}
	// log.Println("-------------------------------------")
}

// * 로그인
func (pwo *PWObj) LoginInPage(form map[string]string) {
	pwo.Page.Fill(setXpath(form["id_xpath"]), form["id"])
	time.Sleep(1 * time.Second)
	pwo.Page.Fill(setXpath(form["pw_xpath"]), form["pw"])
	time.Sleep(1100 * time.Millisecond)
	pwo.Page.Click(setXpath(form["submit_xpath"]))
	time.Sleep(500 * time.Millisecond)
}

// * Response
func (pwo *PWObj) ResponseFromPage(option map[string]string) (rst string) {
	switch option["rType"] {
	case "content":
		rst, _ = pwo.Page.TextContent(setXpath(option["xpath"]))
	case "innerhtml":
		node, _ := pwo.Page.Locator(setXpath(option["xpath"]))
		rst, _ = node.InnerHTML()
	case "outerhtml":
		node, _ := pwo.Page.Locator(setXpath(option["xpath"]))
		rst, _ = node.InnerHTML()
		rst = WRAP_START + rst + WRAP_END
	}
	// log.Printf("\nResponseFromPage rst: %v\n", rst)
	return
}

// * Response
func (pwo *PWObj) ResponseFromFrame(option map[string]string) (rst string) {
	switch strings.ToLower(option["rType"]) {
	case "content":
		rst, _ = pwo.Frame.TextContent(setXpath(option["xpath"]))
	case "innerhtml":
		node, _ := pwo.Frame.Locator(setXpath(option["xpath"]))
		rst, _ = node.InnerHTML()
	case "outerhtml":
		node, _ := pwo.Frame.Locator(setXpath(option["xpath"]))
		rst, _ = node.InnerHTML()
		rst = WRAP_START + rst + WRAP_END
	}
	// log.Printf("\nResponseFromFrame rst: %v\n", rst)
	return
}

func (pwo *PWObj) Response(option map[string]string) (rst string) {
	switch option["loc"] {
	case "page":
		// log.Println("****************ResponseFromPage")
		rst = pwo.ResponseFromPage(option)
	case "frame":
		// log.Println("****************ResponseFromFrame")
		rst = pwo.ResponseFromFrame(option)
	}
	if rst == "" { // ? 비어 있는 경우에도 처리되도록
		rst = " "
	}
	return
}

// 소스 저장
func (pwo *PWObj) SavePage(option map[string]string) {
	writeStr(option["path"], pwo.Response(option))
}

// * Click
func (pwo *PWObj) ClickInPage(option map[string]string) {
	if position, ok := option["position"]; ok {
		xy := strings.Split(position, ",")
		x, _ := strconv.ParseFloat(xy[0], 64)
		y, _ := strconv.ParseFloat(xy[1], 64)
		position := playwright.PageClickOptions{
			Position: &playwright.PageClickOptionsPosition{X: &x, Y: &y},
		}
		pwo.Page.Click(setXpath(option["xpath"]), position)
	} else {
		pwo.Page.Click(setXpath(option["xpath"]))
		log.Printf("\nCicked %v\n", option["xpath"])
	}
}

func (pwo *PWObj) ClickInFrame(option map[string]string) {
	if position, ok := option["position"]; ok {
		xy := strings.Split(position, ",")
		x, _ := strconv.ParseFloat(xy[0], 64)
		y, _ := strconv.ParseFloat(xy[1], 64)
		position := playwright.PageClickOptions{
			Position: &playwright.PageClickOptionsPosition{X: &x, Y: &y},
		}
		pwo.Frame.Click(setXpath(option["xpath"]), position)
	} else {
		node, _ := pwo.Frame.Locator(setXpath(option["xpath"]))
		val, _ := node.TextContent()
		log.Printf("\nClickInFrame node: %v\n", val)

		pwo.Frame.Click(setXpath(option["xpath"]))
	}
}

// * Download
func (pwo *PWObj) Download(option map[string]string) {
	// Open the downloaded file and parse it as a CSV
	// filePath := "file.csv"
	time.Sleep(3 * time.Second)
	filePath := option["path"]
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("could not open downloaded file: %v", err)
	}
	defer file.Close()

	// reader := csv.NewReader(file)
	// var rows [][]string
	// for {
	// 	row, err := reader.Read()
	// 	if err != nil {
	// 		if err == io.EOF {
	// 			break
	// 		}
	// 		log.Fatalf("could not read CSV row: %v", err)
	// 	}
	// 	rows = append(rows, row)
	// }
}

// * 단위 action 실행
// action: act: gotoUrl,gotoFrame|gotoMain,click,login,form  /loc: xpath,url / ret: content/text/innerHtml.../xpath(true,false) / wait: xpath,second,...
func (pwo *PWObj) doAction(option map[string]string) (rst string) {
	rst = ""
	switch option["act"] {
	case "gotoUrl":
		pwo.GotoUrl(option["loc"])
		// pwo.Wait(option["wait"])
	case "gotoFrame":
		pwo.GotoFrame(option["loc"])
	case "wait":
		pwo.Wait(option["xpath"])
	case "clickPage":
		pwo.ClickInPage(option)
	case "clickFrame":
		pwo.ClickInFrame(option)
	case "loginPage":
		pwo.LoginInPage(option)
	case "savePage":
		pwo.SavePage(option)
	case "request":
		rst = pwo.Response(option)
	case "download":
		pwo.Download(option)
	}

	return
}

// * 전체 action 실행
func (pwo *PWObj) DoActions(actions []map[string]string) (rsts string) {
	rsts_ := []string{}
	for _, action := range actions {
		rst := pwo.doAction(action)
		if rst != "" {
			rsts_ = append(rsts_, rst)
		}
	}

	time.Sleep(100 * time.Second)
	// time.Sleep(5 * time.Second)
	if err := pwo.Browser.Close(); err != nil {
		log.Fatalf("could not close browser: %v", err)
	}
	if err := pwo.PW.Stop(); err != nil {
		log.Fatalf("could not stop Playwright: %v", err)
	}

	return strings.Join(rsts_, SEPARATOR)
}
