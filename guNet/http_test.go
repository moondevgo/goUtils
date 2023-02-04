package network_test

import (
	"fmt"
	// "log"
	// "net/url"
	"testing"
	// "with_go/env"
	"github.com/moondevgo/go-mods/network"
)

// func TestHttpGet(t *testing.T) {
// 	method := "get"
// 	uri := "http://152.67.230.230:9901/users"
// 	// uri := "https://docs.upbit.com/reference/"
// 	header := map[string]string{}
// 	data := map[string]interface{}{}
// 	resp, err := network.Http(method, uri, header, data)
// 	if err != nil {
// 		t.Error("Wrong result")
// 	}
// 	fmt.Printf("%T, %v\n", resp, string(resp[:]))
// }

// // [GoLang 나만의 RESTful API 서버 만들기 (1)](https://okky.kr/article/386116)
// func TestHttpPost(t *testing.T) {
// 	method := "post"
// 	// method := "get"
// 	uri := "http://152.67.230.230:9901/users"
// 	// uri := "https://docs.upbit.com/reference/"
// 	header := map[string]string{}
// 	data := map[string]interface{}{"nickname": "mon", "email": "mon@gmail.com"}
// 	// data := map[string]interface{}{}

// 	resp, err := network.Http(method, uri, header, data)
// 	// resp, err := network.HttpPOST(uri, header, data)
// 	if err != nil {
// 		t.Error("Wrong result")
// 	}
// 	fmt.Printf("%T, %v\n", resp, string(resp))
// }

// func TestHttpPost(t *testing.T) {
// 	method := "post"
// 	uri := "http://127.0.0.1:8888/graphql"
// 	header := map[string]string{}
// 	// data := map[string]interface{}{"query": `user(id: "0x03") {\n id \n name \n }`}
// 	// C:\Dev\inGo\on_test\graphql\graphql_go\examples\social\server>go run server.go
// 	// data := `{"query": user(id: "0x03") {\n id \n name \n }}` // C:\Dev\inGo\on_test\graphql\graphql_go\examples\social\server>go run server.go
// 	data := `{
// 		"query": "query: user(id: "0x03") {\n id\n name \n}"
// 	}`
// 	// data := `{"query": "query { info }", "variables": {}}` // Success! Server: C:\Dev\inGo\on_test\graphql\graphql_go>go run main-11.go
// 	// data := `{
// 	// 	"query": "query { info }",
// 	// 	"variables": {}
// 	// }`
// 	// data := `{
// 	// 	"query": "query { info }"
// 	// }`
// 	// data := `{
// 	// 	"query": "query { info }"
// 	// }`
// 	resp, _ := network.HttpStr(method, uri, header, data)
// 	fmt.Printf("%T, %v\n", resp, string(resp))
// }

// func TestHttpGraphql(t *testing.T) {
// 	uri := "http://127.0.0.1:8888/graphql"
// 	query := `{
// 		user(id: "0x03") {
// 			id
// 			name
// 			address
// 		}
// 		admin(id: "0x01") {
// 			name
// 			role
// 		}
// 		search(text: "Potter") {
// 			... on User {
// 			name
// 			}
// 		}
// 	}`
// 	network.HttpGraphql(uri, query)
// }

// # TOKEN = 'secret_WSVKqRFViNiBEUXv64VmOEgbTKPYlpfzBL7bR0zDn8w'  # monWorkBot
// # token_v2 = 'ac4d0250ee37311a7e5bb1cafc728ed4afd91e018633d346516935f37d02cab6407c31c1eecdf0a195e30ba1f4da20258a51114676b503a2b0cb045d70300da79930145c5b0bf46345c4cc05e0b8'

// # block_id = "da1c09759006449ea6f7be087bc20f08"  # /TestPage//#Archive/Acount DB
// # userId = "1a3d6d74-2978-4800-9f82-1150557d6a7e"  # monblue@snu.ac.kr

// # HEADERS = {
// #     "Authorization": "Bearer " + TOKEN,
// #     "Content-Type": "application/json",
// #     "Notion-Version": "2022-02-22",
// # }

// # BASE_URL = "https://api.notion.com/v1/"

// # ACTIONS = dict(
// #     create = ("POST", "{obj_type}s"),  # database, page
// #     retrieve = ("GET", "{obj_type}s/{obj_id}"),  # database, page, block, user
// #     update = ("PATCH", "{obj_type}s/{obj_id}"),  # database, page, block
// #     delete = ("DELETE", "{obj_type}s/{obj_id}"),  # block
// #     query = ("POST", "{obj_type}s/{obj_id}/query"),  # database / Query a database
// #     children = ("GET", "{obj_type}s/{obj_id}/children"),  # block / Retrieve block children
// #     append = ("PATCH", "{obj_type}s/{obj_id}/children"),  # block / Append block children
// # )

// NOTE: notion
func NotionBearer(data map[string]interface{}) string {
	return "secret_WSVKqRFViNiBEUXv64VmOEgbTKPYlpfzBL7bR0zDn8w"
}

func TestHttpNotion(t *testing.T) {
	header := map[string]string{"Content-Type": "application/json", "Notion-Version": "2022-02-22"}
	data := map[string]interface{}{}
	obj_type := "database"
	obj_id := "c477220bb6d34e20ae0ea288427c2c97"
	uri := fmt.Sprintf("https://api.notion.com/v1/%ss/%s", obj_type, obj_id)
	resp, err := network.HttpBearer("GET", uri, header, data, NotionBearer)
	if err != nil {
		t.Error("Wrong result")
	}
	fmt.Printf("%T, %v\n", resp, string(resp))
}

// func TestHttpGetBearer(t *testing.T) {
// 	method := "get"
// 	// method := "post"
// 	// 전체 계좌 조회
// 	uri := "https://api.upbit.com/v1/accounts"
// 	data := map[string]interface{}{}

// 	// // 주문 가능 정보
// 	// uri := "https://api.upbit.com/v1/orders/chance"
// 	// data := map[string]interface{}{"market": "KRW-BTC"}

// 	// // 주문 하기
// 	// uri := "https://api.upbit.com/v1/orders"
// 	// // data := map[string]interface{}{"market": "KRW-XRP", "side": "ask", "volume": "2", "price": "460", "order_type": "limit"}
// 	// // {"uuid":"8046133f-1e9c-46c7-a6dc-092c99cf9283","side":"bid","ord_type":"limit","price":"600.0","state":"wait","market":"KRW-XRP","created_at":"2022-07-17T10:21:53+09:00","volume":"10.0","remaining_volume":"10.0","reserved_fee":"3.0","remaining_fee":"3.0","paid_fee":"0.0","locked":"6003.0","executed_volume":"0.0","trades_count":0}
// 	// data := map[string]interface{}{"market": "KRW-XRP", "side": "ask", "volume": "11", "price": "500", "order_type": "limit"}
// 	// // {"uuid":"d9f4db6b-bcd4-4609-be6b-730a54cb1604","side":"ask","ord_type":"limit","price":"500.0","state":"wait","market":"KRW-XRP","created_at":"2022-07-17T10:31:19+09:00","volume":"11.0","remaining_volume":"11.0","reserved_fee":"0.0","remaining_fee":"0.0","paid_fee":"0.0","locked":"11.0","executed_volume":"0.0","trades_count":0}

// 	// // 주문 리스트 조회
// 	// // market	마켓 아이디	String
// 	// // uuids	주문 UUID의 목록	Array
// 	// // identifiers	주문 identifier의 목록	Array
// 	// // state	주문 상태
// 	// // - wait : 체결 대기 (default)
// 	// // - watch : 예약주문 대기
// 	// // - done : 전체 체결 완료
// 	// // - cancel : 주문 취소	String
// 	// // states	주문 상태의 목록

// 	// // * 미체결 주문(wait, watch)과 완료 주문(done, cancel)은 혼합하여 조회하실 수 없습니다.	Array
// 	// // page	페이지 수, default: 1	Number
// 	// // limit	요청 개수, default: 100	Number
// 	// // order_by
// 	// uri := "https://api.upbit.com/v1/orders"
// 	// data := map[string]interface{}{"market": "KRW-XRP", "state": "done"}

// 	// // 개별 주문 조회
// 	// uri := "https://api.upbit.com/v1/order"
// 	// data := map[string]interface{}{"uuid": "d9f4db6b-bcd4-4609-be6b-730a54cb1604"}

// 	header := map[string]string{}

// 	// keys := env.GetConfigYaml("coin_apis", "upbit")

// 	// // bearer := network.UpbitJwt(keys["access_key"].(string), keys["secret_key"].(string), data)
// 	// bearer := network.UpbitJwt(keys["access_key"].(string), keys["secret_key"].(string), data)
// 	// // bearer := network.UpbitJwt(keys["access_key"].(string), keys["secret_key"].(string), nil)
// 	// fmt.Printf("%T, %v\n", bearer, bearer)

// 	// resp, err := network.HttpBearer(method, uri, header, data, bearer)
// 	resp, err := network.HttpBearer(method, uri, header, data, network.UpbitJwt)
// 	if err != nil {
// 		t.Error("Wrong result")
// 	}
// 	fmt.Printf("%T, %v\n", resp, string(resp))
// }
