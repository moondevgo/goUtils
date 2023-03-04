package guCloud

// [Updating Google Sheets using Golang.](https://dev.to/mediocredevops/playing-with-google-sheets-api-using-golang-14en)
// [Sheets API Guides Go quickstart](https://developers.google.com/sheets/api/quickstart/go)

import (
	"log"

	basic "github.com/moondevgo/goUtils/guBasic"
	"google.golang.org/api/sheets/v4"
)

var (
	G_gSheet GSheet
)

type GSheet struct {
	Service *sheets.Service `json:"service"` // Google Service
}

// * Create GSheet Obj
func NewGSheet(options ...string) GSheet {
	// options variables
	bot_nick := "moonsats" // 0
	user_nick := ""        // 0
	if len(options) > 0 {
		bot_nick = options[0]
	}
	if len(options) > 1 {
		bot_nick = ""
		user_nick = options[1]
		log.Printf("\nelse if bot_nick: |%v|, user_nick: |%v|\n", bot_nick, user_nick)
	}

	// log.Printf("\nlen: %v, bot_nick: |%v|, user_nick: |%v|\n", len(options), bot_nick, user_nick)
	client := ApiClient("sheets", bot_nick, user_nick)

	srv, err := sheets.New(client)
	if err != nil {
		log.Fatalf("Unable to retrieve Sheets client: %v", err)
	}

	gSheet := GSheet{}
	gSheet.Service = srv

	return gSheet
}

// * Read GSheet All Fields
func (gSheet *GSheet) readGSheet(spreadsheetId, readRange string) [][]interface{} {
	resp, err := gSheet.Service.Spreadsheets.Values.Get(spreadsheetId, readRange).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve data from sheet: %v", err)
	}

	if len(resp.Values) == 0 {
		log.Println("No data found.")
		return nil
	}

	return resp.Values
}

// * Read GSheet
func (gSheet *GSheet) ReadGSheet(spreadsheetId, readRange string, options ...[]string) []map[string]string {
	if len(options) > 0 {
		return basic.MapsFromInterfaces(gSheet.readGSheet(spreadsheetId, readRange), options[0])
	}
	return basic.MapsFromInterfaces(gSheet.readGSheet(spreadsheetId, readRange), []string{})
}

func (gSheet *GSheet) writeGSheet(data [][]interface{}, spreadsheetId, writeRange string) {
	var vr sheets.ValueRange

	for _, d := range data {
		vr.Values = append(vr.Values, d)
	}

	_, err := gSheet.Service.Spreadsheets.Values.Update(spreadsheetId, writeRange, &vr).ValueInputOption("RAW").Do()
	if err != nil {
		log.Fatalf("Unable to retrieve data from sheet. %v", err)
	}
}

// TODO: fields 지정 순서대로 쓰도록 구현
// TODO: sheet가 없을 때는 생성하도록
// * data []map[string]string -> GSheet
func (gSheet *GSheet) WriteGSheet(data []map[string]string, spreadsheetId, writeRange string, options ...[]string) {
	if len(options) > 0 {
		gSheet.writeGSheet(basic.InterfacesFromMaps(data, options[0]), spreadsheetId, writeRange)
	}
	gSheet.writeGSheet(basic.InterfacesFromMaps(data), spreadsheetId, writeRange)
}

// * 초기화 함수
func init() {
	// G_gSheet = NewGSheet()
	G_gSheet = NewGSheet("", "moondevgoog")
}
