package guBasic

import (
	"bytes"
	"log"
	"strconv"
	"strings"
	"syscall"
	"unsafe"

	"github.com/mitchellh/mapstructure"
	"golang.org/x/text/encoding/korean"
	"golang.org/x/text/transform"
)

// ** Data(slice, string, ...) 처리
// * slice 에서 element의 index를 구함
// func IndexOf[T comparable](element T, data []T) int {
func IndexOf[T comparable](element T, data []T) int {
	for k, v := range data {
		if element == v {
			return k
		}
	}
	return -1 //not found.
}

// func IndexOf(element string, data []string) int {
// 	for k, v := range data {
// 		if element == v {
// 			return k
// 		}
// 	}
// 	return -1 //not found.
// }

// 문자열이 숫자로만 이루어졌는지 여부
func IsDigit(s string) bool {
	for _, c := range s {
		if c < '0' || c > '9' {
			return false
		}
	}
	return true
}

// * slice에서 s번째 요소 삭제
func RemoveByIndex[T comparable](slice []T, s int) []T {
	return append(slice[:s], slice[s+1:]...)
}

// * slice에서 값이 v인 요소 삭제
func RemoveByValue[T comparable](slice []T, v T) []T {
	for i, v_ := range slice {
		if v_ == v {
			return append(slice[:i], slice[i+1:]...)
		}
	}
	return slice
}

// * string slice 내에 str이 있는지 여부
func IsInSlice(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

// * slice 내에 e가 있는지 여부(Generic Type)
func Contains[T comparable](s []T, e T) bool {
	for _, v := range s {
		if v == e {
			return true
		}
	}
	return false
}

// * 공백문자 n개를 가지는 buff string
func BuffStr(n int) string {
	return strings.Repeat(" ", n)
}

// ** Convert Data Type
// * string -> int
func IntFromStr(str string) int {
	_sizes := strings.Split(strings.Trim(str, " "), ".") // TODO: float, double 소수점이 있는 경우에도 적용되도록 하는 다른 방법은?
	i, _ := strconv.Atoi(_sizes[0])
	return i
	// i, _ := strconv.Atoi(str)
	// return i
}

// * string -> float32
func Float32FromStr(str string) float32 {
	if s, err := strconv.ParseFloat(str, 32); err == nil {
		return float32(s)
	}
	return float32(0)
}

// * string -> float64
func FloatFromStr(str string) float64 {
	if s, err := strconv.ParseFloat(str, 64); err == nil {
		return s
	}
	return float64(0)
}

func StrFromInt(i int) string {
	return strconv.Itoa(i)
}

// * Ansi -> Utf
func AnsiToUtf(src string) string {
	got, _, _ := transform.String(korean.EUCKR.NewDecoder(), strings.Trim(src, " "))
	return string(bytes.Trim([]byte(got), "\x00")) // ? bytes, fmt 사용하지 않고 "\0" 제거하는 방법은
}

// * bytes -> dtring(trimed)
func TrimStrFromBytes(data []byte) string {
	// ! byte trim("\x00") -> string trim(" ") -> byte trim("\x00")
	// * t1857OutBlock [116 49 56 53 55 79 117 116 66 108 111 99 107 0 0 0 188]
	return string(bytes.Trim(data, "\x00"))
}

// ** Functions For Map
// * Map keys
func Keys[T comparable, U any](m map[T]U) (r []T) {
	for k, _ := range m {
		r = append(r, k)
	}
	return r
}

// * Map Values
func Values[T comparable, U any](m map[T]U) (r []U) {
	for _, v := range m {
		r = append(r, v)
	}
	return r
}

// * items(maps []map[string]interface{} ex) mysql 리턴값) -> field 단일 필드 []interface{}
func FlattenMaps(items []map[string]interface{}, field string) (flattens []interface{}) {
	flattens = []interface{}{}
	for _, item := range items {
		flattens = append(flattens, item[field])
	}
	return
}

// * items(slices [][]interface{} ex) QueryOutput) -> field 단일 필드 []interface{}
func FlattenSlices(items [][]interface{}, index int) (flattens []interface{}) {
	flattens = []interface{}{}
	for _, item := range items {
		flattens = append(flattens, item[index])
	}
	return
}

// * []interface{} -> []string
func StrsFromInterfaces(data []interface{}) (r []string) {
	r = []string{}
	for _, d := range data {
		r = append(r, d.(string))
	}
	return r
}

// * []T -> []interface{}
func ConvertToInterfaces[T comparable](data []T) []interface{} {
	interfaces := []interface{}{}
	for _, d := range data {
		interfaces = append(interfaces, d)
	}
	return interfaces
}

// * [][]T -> [][]interface{} (For google sheets sheets.ValueRange.Values)
func ConvertToInterfaces2[T comparable](datas [][]T) [][]interface{} {
	interfaces := [][]interface{}{}
	for _, data := range datas {
		interfaces = append(interfaces, ConvertToInterfaces(data))
	}
	return interfaces
}

// func ConvertToInterfaces(data []string) []interface{} {
//   interfaces := []interface{}{}
//   for _, d := range data {
//     interfaces = append(interfaces, d)
//   }
//   return interfaces
// }

// func ConvertToInterfaces2(datas [][]string) [][]interface{} {
//   interfaces := [][]interface{}{}
//   for _, data := range datas {
//     interfaces = append(interfaces, ConvertToInterfaces(data))
//   }
//   return interfaces
// }

// * interfaces []interface{} -> header
// fields: interfaces 중에서 사용할 field 이름 slice
func HeaderFromInterfaces(interfaces []interface{}, fields []string) map[int]string {
	header := map[int]string{}
	if len(fields) > 0 {
		for i, v := range interfaces {
			if Contains(fields, v.(string)) {
				header[i] = v.(string)
			}
		}
	} else {
		for i, v := range interfaces {
			header[i] = v.(string)
		}
	}
	return header
}

// * interfaces [][]interface{} -> maps []map[string]string
// [[symbol time open high low close volume] [KRW-BTC 2022-07-11 10:00 12000 14000 11000 13000 56.5] [KRW-BTC 2022-07-11 10:10 13000 13500 12000 14000 120.7] [KRW-BTC 2022-07-12 10:20 14000 15000 12500 13500 310.9]]
// [map[close:13000 high:14000 low:11000 open:12000 symbol:KRW-BTC time:2022-07-11 10:00 volume:56.5] map[close:13000 high:14000 low:11000 open:12000 symbol:KRW-BTC time:2022-07-11 10:00 volume:56.5]
// func MapsFromInterfaces(interfaces [][]interface{}, header []string) (rows []map[string]string) {
func MapsFromInterfaces(interfaces [][]interface{}, options ...[]string) (rows []map[string]string) {
	fields := []string{}
	if len(options) > 0 {
		fields = options[0]
	}
	header := HeaderFromInterfaces(interfaces[0], fields)
	// log.Printf("***header: %v\n", header)
	rows = []map[string]string{}

	for _, intfs := range interfaces[1:] {
		row := map[string]string{}
		for i, intf := range intfs {
			if Contains(Keys(header), i) {
				row[header[i]] = intf.(string)
			}
		}
		rows = append(rows, row)
	}
	return rows
}

// * maps []map[string]string -> interfaces [][]interface{}
func InterfacesFromMaps(rows []map[string]string, options ...[]string) (interfaces [][]interface{}) {
	var fields []string
	if len(options) > 0 {
		fields = options[0]
	} else {
		fields = Keys(rows[0])
	}

	interfaces = append(interfaces, ConvertToInterfaces(fields)) // ? header
	for _, row := range rows {
		_interface := []interface{}{}
		for _, field := range fields {
			_interface = append(_interface, row[field])
		}
		interfaces = append(interfaces, _interface)
	}

	return interfaces
}

// ** Map
// * map 합치기(뒤에 map이 앞에 map을 덮어씀)
func MergeMaps(ms ...map[string]string) map[string]string {
	res := map[string]string{}
	for _, m := range ms {
		for k, v := range m {
			res[k] = v
		}
	}
	return res
}

// * map -> struct
func StructFromMap[T any](map_ map[string]interface{}, data *T) T {
	if err := mapstructure.Decode(map_, &data); err != nil {
		log.Println(err)
	}
	return *data
	// log.Println(*data)
}

// ** Functions(pointer <- data)
// * bool -> uintptr
func UintptrFromBool(b bool) uintptr {
	if b {
		return uintptr(1)
	} else {
		return uintptr(0)
	}
}

// * bytes -> uintptr
func UintptrFromBytes(_bytes []byte) uintptr {
	return uintptr(unsafe.Pointer(&append(_bytes, 0)[0]))
}

// * str -> uintptr
func UintptrFromStr(str string) uintptr {
	return uintptr(unsafe.Pointer(&append([]byte(str), 0)[0]))
}

// * utf -> uintptr
func UintptrFromUtf(str string) uintptr {
	s16, _ := syscall.UTF16PtrFromString(str)
	// return voidptr(unsafe.Pointer(s16))
	return uintptr(unsafe.Pointer(s16))
}

// * str -> ptr(unsafe.Pointer)
func PtrFromStr(str string) unsafe.Pointer {
	return unsafe.Pointer(&append([]byte(str), 0)[0])
}

// ** Functions(data <- pointer)
// * uintptr -> []byte
//   - 문자열 끝(/0)이 나오면 반환, 최대 크기 4096
func BytesFromUintptr(uptr uintptr) []byte {
	bytes := make([]byte, 0)

	for i := 0; i < 4096; i++ { // 4096: max byte
		b := *(*byte)(unsafe.Pointer(uptr + uintptr(i)*unsafe.Sizeof(byte(0))))
		if b == 0 { // NOTE: 문자열 끝(/0)
			return bytes
		}

		bytes = append(bytes, b)
	}
	return bytes
}

// * uintptr -> str
func StrFromUintptr(ptr uintptr) string {
	return string(BytesFromUintptr(ptr))
}

// * ptr(unsafe.Pointer) -> str
func StrFromPtr(ptr unsafe.Pointer) string {
	return string(BytesFromUintptr(uintptr(ptr)))
}

// * uintptr -> []byte
//   - size: bytes 크기
func BytesFromPtrWithSize(ptr unsafe.Pointer, size int) []byte {
	bytes := make([]byte, 0)

	for i := 0; i < size; i++ {
		ptr := (*byte)(unsafe.Pointer(uintptr(ptr) + uintptr(i)*unsafe.Sizeof(byte(0))))
		_byte := *ptr
		bytes = append(bytes, _byte)
	}

	return bytes
}

// ** Functions(한글 처리)
// * []byte -> Kor
func KorFromBytes(bytes []byte) string {
	idxNull := strings.Index(string(bytes), "\x00")

	if idxNull >= 0 {
		bytes = bytes[:idxNull]
	}

	bytes_utf8, err := korean.EUCKR.NewDecoder().Bytes(bytes)
	if err != nil {
		if len(bytes) > 0 {
			return KorFromBytes(bytes[:len(bytes)-1])
		}

		return string(bytes)
	}

	return string(bytes_utf8)
}

// * uintptr -> Kor
func KorFromUintptr(ptr uintptr) string {
	return KorFromBytes(BytesFromUintptr(ptr))
}

// * uintptr -> Kor
func KorFromStr(str string) string {
	return KorFromBytes([]byte(str))
}
