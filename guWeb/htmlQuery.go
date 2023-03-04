package guWeb

import (
	// "io"
	// "net/http"
	// "log"

	"strings"

	"github.com/antchfx/htmlquery"

	// "github.com/chromedp/chromedp"
	basic "github.com/moondevgo/goUtils/guBasic"
	"golang.org/x/net/html"
)

// ** Sub Functions
// * str -> nodes
func NodeByStr(str string) *html.Node {
	html, _ := htmlquery.Parse(strings.NewReader(str))
	return html
}

// * str, xpath -> nodes
func FindNodeFromStr(str, xpath string) []*html.Node {
	html, _ := htmlquery.Parse(strings.NewReader(str))
	return htmlquery.Find(html, xpath)
}

// * str, xpath -> nodes
func FindNodeOneFromStr(str, xpath string) *html.Node {
	html, _ := htmlquery.Parse(strings.NewReader(str))
	return htmlquery.FindOne(html, xpath)
}

// * url, xpath -> nodes
func FindNodeFromUrl(url, xpath string) []*html.Node {
	html, _ := htmlquery.LoadURL(url)
	return htmlquery.Find(html, xpath)
}

// * url -> 1 node
func NodeByUrl(url string) *html.Node {
	html, _ := htmlquery.LoadURL(url)
	return html
}

// * url, xpath -> 1 node
func FindNodeOneFromUrl(url, xpath string) *html.Node {
	html, _ := htmlquery.LoadURL(url)
	return htmlquery.FindOne(html, xpath)
}

// * url -> 1 node
func NodeByDoc(path string) *html.Node {
	html, _ := htmlquery.LoadDoc(path)
	return html
}

// * path(file), xpath -> nodes
func FindNodeFromDoc(path, xpath string) []*html.Node {
	html, _ := htmlquery.LoadDoc(path)
	return htmlquery.Find(html, xpath)
}

// * path(file), xpath -> 1 node
func FindNodeOneFromDoc(path, xpath string) *html.Node {
	html, _ := htmlquery.LoadDoc(path)
	return htmlquery.FindOne(html, xpath)
}

// * child node
func Node(root *html.Node, xpath string) []*html.Node {
	return htmlquery.Find(root, xpath)
}

// * children node
func NodeOne(root *html.Node, xpath string) *html.Node {
	return htmlquery.FindOne(root, xpath)
}

// * node -> string
func FindText(node *html.Node) string {
	return strings.TrimSpace(htmlquery.InnerText(node))
}

// * root, xpath -> string
func FindTextChild(root *html.Node, xpath string) string {
	node := htmlquery.FindOne(root, xpath)
	if node == nil { // * xpath 요소가 없는 경우
		return ""
	}
	return FindText(node)
}

// * node, attr -> string / attr: ex) "href"
func FindAttr(node *html.Node, attr string) string {
	return htmlquery.SelectAttr(node, attr)
}

func FindVal(node *html.Node, target string) string {
	switch target {
	case "text":
		return FindText(node)
	case "content": // TODO: "text"와 다르게
		return FindText(node)
	default:
		return FindAttr(node, target) // target ex) "href"
	}
}

// * node -> html
func HtmlFromNode(node *html.Node) string {
	return htmlquery.OutputHTML(node, true)
}

// func FindValCB(node *html.Node, target string, callback func(string) string) string {
// 	return callback(FindVal(node, target))
// }

// func FindValChild(node *html.Node, xpath string, target interface{}) string {
// 	if target_, ok := target.(int); !ok {
// 		return FindText(htmlquery.Find(node, xpath)[target_])
// 	} else {
// 		return FindVal(htmlquery.FindOne(node, xpath), target.(string))
// 	}
// }

// * nodes, keys(indexes) -> map text string
// keys map[string]int{"<key0>": index0, ...}
func MapTextByIndexes(rst map[string]string, nodes []*html.Node, keys map[string]int) map[string]string {
	if len(nodes) < len(basic.Keys(keys)) { // TODO: vals중 가장 큰값으로
		return nil
	}
	for key, i := range keys {
		rst[key] = FindText(nodes[i])
	}
	return rst
}

// * nodes, keys(indexes), attr -> map attr string
func MapAttByIndexes(rst map[string]string, nodes []*html.Node, keys map[string]int, attr string) map[string]string {
	if len(nodes) < len(basic.Keys(keys)) { // TODO: vals중 가장 큰값으로
		return nil
	}
	for key, i := range keys {
		rst[key] = FindAttr(nodes[i], attr)
	}
	return rst
}

// * nodes, keys(indexes) -> rst map[string]string
// keys: ex) map[string][]interface{}{"symbol": []interface{}{1, "text", FindSymbol},...}
func MapValByIndexes(rst map[string]string, nodes []*html.Node, keys map[string][]interface{}) map[string]string {
	// ? nodes의 개수가 적은 경우 return nil
	values := basic.Values(keys)
	// log.Printf("MapValByIndexes values: %v", values)
	switch values[0][0].(type) {
	case int:
		for _, value := range values {  // index중 가장 큰값 > node 개수 - 1
			if value[0].(int) > len(nodes) - 1 {
				return nil
			}
		}
	default:
		if len(nodes) < len(basic.Keys(keys)) { // xpath 개수 > node 개수
			return nil
		}
	}

	for key, setting := range keys {
		i := setting[0].(int)
		val := ""
		if len(setting) > 2 {
			val = FindVal(nodes[i], setting[2].(string))
		} else {
			val = FindText(nodes[i])
		}

		if len(setting) > 1 {
			val = setting[1].(func(string) string)(val)
		}
		rst[key] = val
	}
	return rst
}

// * nodes, keys(xpaths) -> map text string
// keys map[string]int{"<key0>": xpath0, ...}
func MapTextByXpaths(rst map[string]string, node *html.Node, keys map[string]string) map[string]string {
	for key, xpath := range keys {
		rst[key] = FindText(htmlquery.FindOne(node, xpath))
	}
	return rst
}

// * nodes, keys(xpaths), attr -> map attr string
func MapAttrByXpaths(rst map[string]string, node *html.Node, keys map[string]string, attr string) map[string]string {
	for key, xpath := range keys {
		rst[key] = FindAttr(htmlquery.FindOne(node, xpath), attr)
	}
	return rst
}

// * nodes, keys(xpaths) -> rst map[string]string
// keys: ex) map[string][]interface{}{"symbol": []interface{}{"./td[2]/a", "href", FindSymbol},...}
func MapValByXpaths(rst map[string]string, node *html.Node, keys map[string][]interface{}) map[string]string {
	for key, setting := range keys {
		xpath := setting[0].(string)
		val := ""
		if len(setting) > 2 {
			val = FindVal(NodeOne(node, xpath), setting[2].(string))
		} else {
			val = FindText(NodeOne(node, xpath))
		}

		if len(setting) > 1 {
			val = setting[1].(func(string) string)(val)
		}
		rst[key] = val
	}
	return rst
}

// // {<xpath>, <rType>, <callback>}
// // {"0", "text", f1},
// func FindValChildren(root *html.Node, xpath string, setting ...interface{}) (val string) {
// 	nodes := htmlquery.Find(root, xpath)
// 	xpath_ := setting[0]
// 	target_ := setting[1].(string)
// 	if i, ok := xpath_.(int); !ok {
// 		val = FindVal(nodes[i], target_)
// 	} else {
// 		val = FindVal(nodes[0], target_)
// 	}

// 	if setting[2] != nil {
// 		val = setting[2].(func(string) string)(val)
// 	}

// 	return
// }

// // {<xpath>, <rType>, <callback>}
// // {"0", "text", f1},
// func FindValList(root *html.Node, xpath string, keys [][]interface{}) (vals []map[string]string) {
// 	for
// 	nodes := htmlquery.Find(root, xpath)
// 	xpath_ := setting[0]
// 	target_ := setting[1].(string)
// 	if i, ok := xpath_.(int); !ok {
// 		val = FindVal(nodes[i], target_)
// 	} else {
// 		val = FindVal(nodes[0], target_)
// 	}

// 	if setting[2] != nil {
// 		val = setting[2].(func(string) string)(val)
// 	}

// 	return
// }

// ** Page
func FindHtml(url, xpath string) string {
	return HtmlFromNode(FindNodeOneFromUrl(url, xpath))
}

// xpaths := map[int]string{"key0": "xpath0", ...}
func MapFromNodes(nodes []*html.Node, keys map[string]int) (rst map[string]string) {
	rst = map[string]string{}
	for key, i := range keys {
		rst[key] = FindText(nodes[i])
	}
	return
}

// xpaths := map[int]string{"key0": "xpath0", ...}
func MapFromNode(node *html.Node, xpaths map[string]string) (rst map[string]string) {
	rst = map[string]string{}
	for key, xpath := range xpaths {
		rst[key] = FindTextChild(node, xpath)
	}
	return
}

// func MapsFromNode(node *html.Node, xpaths map[string]string) (rst map[string]string) {
// 	for key, xpath := range xpaths {
// 		rst[key] = FindTextChild(node, xpath)
// 	}
// 	return
// }

// func HtmlFromNode(node *html.Node, self bool) string {
// 	return htmlquery.OutputHTML(node, self)
// }

// ** Table
//   - table(html string), keys -> []map[string]string
//     keys := map[int]string{0: "nation", 1: "indexName", 2: "price", 3: "diff", 4: "diff_rate", 6: "time"}
func MapsFromTableNode(node *html.Node, keys map[int]string) (rsts []map[string]string) {
	rsts = []map[string]string{}
	for _, tr := range htmlquery.Find(node, `.//tr`) {
		rst := map[string]string{}
		tds := htmlquery.Find(tr, "./td")
		if len(tds) < len(basic.Keys(keys)) { // TODO: keys의 가장 큰 int와 비교
			continue
		}
		for i, k := range keys {
			rst[k] = FindText(tds[i])
		}
		rsts = append(rsts, rst)
	}
	return
}

//   - table(html string), keys -> []map[string]string
//     keys := map[int]string{0: "nation", 1: "indexName", 2: "price", 3: "diff", 4: "diff_rate", 6: "time"}
func MapsFromTableStr(docStr string, keys map[int]string) (rsts []map[string]string) {
	// rsts = []map[string]string{}
	node, _ := htmlquery.Parse(strings.NewReader(docStr))

	return MapsFromTableNode(node, keys)
}

// * table -> maps by htmlquery
func TableMapsFromUrl(url, xpath string, keys map[int]string) (rsts []map[string]string) {
	html, _ := htmlquery.LoadURL(url)
	node := htmlquery.FindOne(html, xpath)
	return MapsFromTableNode(node, keys)
}
