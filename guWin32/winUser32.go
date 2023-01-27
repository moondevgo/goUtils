package guWin

import (
	"syscall"
	"unsafe"
)

// ** CONST For Window(user32)
// ---------------------------------------------------
const (
	// * Window message constants
	WM_DESTROY = 2  // 0x0002
	WM_CLOSE   = 16 // 0x0010
	WM_QUIT    = 18
	WM_USER    = 1024

	// * PeekMessage wRemoveMsg value
	PM_NOREMOVE = 0x000
	PM_REMOVE   = 0x001
	PM_NOYIELD  = 0x002
)

// ** VAR For Window(user32)
// ---------------------------------------------------
var (
	// * Dll
	kernel32 = syscall.NewLazyDLL("kernel32.dll")
	user32   = syscall.NewLazyDLL("user32.dll")

	// * Proc
	pGetModuleHandleW = kernel32.NewProc("GetModuleHandleW")
	pCreateWindowExW  = user32.NewProc("CreateWindowExW")
	pDefWindowProcW   = user32.NewProc("DefWindowProcW")
	pDestroyWindow    = user32.NewProc("DestroyWindow")
	pDispatchMessageW = user32.NewProc("DispatchMessageW")
	pGetMessageW      = user32.NewProc("GetMessageW")
	pPeekMessageW     = user32.NewProc("PeekMessageW")
	pPostQuitMessage  = user32.NewProc("PostQuitMessage")
	pRegisterClassExW = user32.NewProc("RegisterClassExW")
	// pTranslateMessage = user32.NewProc("TranslateMessage")
)

// ** TYPE For Window(user32)
// ---------------------------------------------------
type (
	// HANDLE          uintptr
	// HWND            HANDLE
	HWND uintptr
)

type POINT struct {
	X, Y int32
}

type MSG struct {
	Hwnd    HWND
	Message uint32
	WParam  uintptr
	LParam  uintptr
	Time    uint32
	Pt      POINT
}

type WNDCLASSEX struct {
	Size       uint32
	Style      uint32
	WndProc    uintptr
	ClsExtra   int32
	WndExtra   int32
	Instance   uintptr
	Icon       uintptr
	Cursor     uintptr
	Background uintptr
	MenuName   *uint16
	ClassName  *uint16
	IconSm     uintptr
}

// ** Functions For Window(user32)
// ---------------------------------------------------
// * ModuleHandle
func GetModuleHandle(modulename string) uintptr {
	var mn uintptr
	if modulename == "" {
		mn = 0
	} else {
		pMn, _ := syscall.UTF16PtrFromString(modulename)
		mn = uintptr(unsafe.Pointer(pMn))
	}
	ret, _, _ := pGetModuleHandleW.Call(mn)
	return uintptr(ret)
}

// * WindowProc 정의
func DefWindowProc(hwnd HWND, msg uint32, wParam, lParam uintptr) uintptr {
	ret, _, _ := pDefWindowProcW.Call(
		uintptr(hwnd),
		uintptr(msg),
		wParam,
		lParam,
	)
	return ret
}

// * 클래스 등록
func RegisterClassEx(wndClassEx *WNDCLASSEX) uint16 {
	if wndClassEx != nil {
		wndClassEx.Size = uint32(unsafe.Sizeof(*wndClassEx))
	}
	ret, _, _ := pRegisterClassExW.Call(uintptr(unsafe.Pointer(wndClassEx)))
	return uint16(ret)
}

// * 윈도우 생성(General)
func CreateWindowEx(exStyle uint, className, windowName *uint16,
	style uint, x, y, width, height int, parent HWND, menu uint16,
	instance uintptr, param unsafe.Pointer) HWND {
	ret, _, _ := pCreateWindowExW.Call(
		uintptr(exStyle),
		uintptr(unsafe.Pointer(className)),
		uintptr(unsafe.Pointer(windowName)),
		uintptr(style),
		uintptr(x),
		uintptr(y),
		uintptr(width),
		uintptr(height),
		uintptr(parent),
		uintptr(menu),
		uintptr(instance),
		uintptr(param),
	)
	return HWND(ret)
}

func DestroyWindow(hwnd HWND) bool {
	ret, _, _ := pDestroyWindow.Call(uintptr(hwnd))
	return ret != 0
}

// ** 메시지 처리
// * dispatchMessage
func DispatchMessage(msg *MSG) uintptr {
	ret, _, _ := pDispatchMessageW.Call(uintptr(unsafe.Pointer(msg)))
	return ret
}

// * postQuitMessage
func PostQuitMessage(exitCode int) {
	pPostQuitMessage.Call(uintptr(exitCode))
}

// * getMessage
func GetMessage(msg *MSG, hwnd HWND, msgFilterMin, msgFilterMax uint32) int {
	ret, _, _ := pGetMessageW.Call(
		uintptr(unsafe.Pointer(msg)),
		uintptr(hwnd),
		uintptr(msgFilterMin),
		uintptr(msgFilterMax),
	)
	return int(ret)
}

// * peekMessage
func PeekMessage(msg *MSG, hwnd HWND, wMsgFilterMin, wMsgFilterMax, wRemoveMsg uint32) bool {
	ret, _, _ := pPeekMessageW.Call(
		uintptr(unsafe.Pointer(msg)),
		uintptr(hwnd),
		uintptr(wMsgFilterMin),
		uintptr(wMsgFilterMax),
		uintptr(wRemoveMsg),
	)
	return ret != 0
}
