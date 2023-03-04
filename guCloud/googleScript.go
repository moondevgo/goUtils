package guCloud

// [Updating Google Sheets using Golang.](https://dev.to/mediocredevops/playing-with-google-sheets-api-using-golang-14en)
// [Sheets API Guides Go quickstart](https://developers.google.com/sheets/api/quickstart/go)

import (
	"encoding/json"
	"log"

	// basic "github.com/moondevgo/goUtils/guBasic"
	script "google.golang.org/api/script/v1"
)

var (
	G_script GScript
)

type GScript struct {
	Service *script.Service `json:"service"` // Google Service
}

// * Create GScript Obj
func NewScript(options ...string) GScript {
	// options variables
	bot_nick := "moonsats" // 0
	if len(options) > 0 {
		bot_nick = options[0]
	}

	client := ApiClient("script", bot_nick, "")
	log.Printf("NewScript client%v\n", client)
	// ? script.New is deprecated: please use NewService instead. To provide a custom HTTP client, use option.WithHTTPClient. If you are using google.golang.org/api/googleapis/transport.APIKey, use option.WithAPIKey with NewService instead.
	srv, err := script.New(client)
	if err != nil {
		log.Fatalf("Unable to retrieve Sheets client: %v", err)
	}

	gScript := GScript{}
	gScript.Service = srv

	log.Printf("gScript.Service: %v\n", gScript.Service)
	return gScript
}

// * Run GScript
// script_id: 1o_mo-dnbEzCKa0j3VOMj_bUEMtBuYmaCL1KB3oMaq8kvMhB1O0-vAM9n
// function: CheerioElFromUrl
// params := []interface{}{"https://www.seo.incheon.kr/open_content/main/community/news/gosi.jsp"}
// func (gScript *GScript) runGScript(script_id, function string, params map[string]interface{}) [][]interface{} {
func (gScript *GScript) RunGScript(script_id, function string, params []interface{}) {
	// resp, err := gScript.Service.Script.Values.Get(spreadsheetId, readRange).Do()
	// params := map[string]interface{}{
	// 	"myArgument": "Hello, World!",
	// }

	// // Define the script function to run
	// function := "myFunction"
	log.Printf("\n RunGScript gScript: %v\n", gScript)

	// Run the script function
	resp, err := gScript.Service.Scripts.Run(script_id, &script.ExecutionRequest{
		Function:   function,
		Parameters: params,
	}).Do()

	if err != nil {
		log.Fatalf("Failed to run script: %v", err)
	}

	// Unmarshal the response from the script function
	var result map[string]interface{}
	if err := json.Unmarshal([]byte(resp.Response), &result); err != nil {
		log.Fatalf("Failed to unmarshal response: %v", err)
	}

	// Print the result
	log.Println(result)
}

// import (
// 	"context"
// 	"encoding/json"
// 	"fmt"
// 	"log"

// 	"golang.org/x/oauth2/google"
// 	script "google.golang.org/api/script/v1"
// )

// func main() {
// 	// Authenticate and construct the service client
// 	client, err := google.DefaultClient(context.Background(), script.ScriptScope)
// 	if err != nil {
// 		log.Fatalf("Failed to create client: %v", err)
// 	}

// 	service, err := script.New(client)
// 	if err != nil {
// 		log.Fatalf("Failed to create script service: %v", err)
// 	}

// 	// Define the parameters for the script function
// 	params := map[string]interface{}{
// 		"myArgument": "Hello, World!",
// 	}

// 	// Define the script function to run
// 	function := "myFunction"

// 	// Run the script function
// 	resp, err := service.Scripts.Run("<SCRIPT_ID>", &script.ExecutionRequest{
// 		Function: function,
// 		Parameters: params,
// 	}).Do()
// 	if err != nil {
// 		log.Fatalf("Failed to run script: %v", err)
// 	}

// 	// Unmarshal the response from the script function
// 	var result map[string]interface{}
// 	if err := json.Unmarshal([]byte(resp.Response), &result); err != nil {
// 		log.Fatalf("Failed to unmarshal response: %v", err)
// 	}

// 	// Print the result
// 	fmt.Println(result)
// }

// * 초기화 함수
func init() {
	G_script = NewScript()
}
