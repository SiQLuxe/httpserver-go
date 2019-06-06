package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"strconv"
	"strings"
	"syscall"

	"github.com/fatih/color"
)

func getIp() string {
	addrs, err := net.InterfaceAddrs()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, address := range addrs {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil && strings.HasPrefix(ipnet.IP.String(), "192.168") {
				// fmt.Println("ip:", ipnet.IP.String())
				return ipnet.IP.String()
			}
		}
	}
	return ""
}

func colorPrint(s string, i int) { //设置终端字体颜色
	kernel32 := syscall.NewLazyDLL("kernel32.dll")
	proc := kernel32.NewProc("SetConsoleTextAttribute")
	handle, _, _ := proc.Call(uintptr(syscall.Stdout), uintptr(i))
	fmt.Print(s)
	handle, _, _ = proc.Call(uintptr(syscall.Stdout), uintptr(7))
	CloseHandle := kernel32.NewProc("CloseHandle")
	CloseHandle.Call(handle)
}

func main() {
	for i := 0; i < 100; i++ {
		fs := http.FileServer(http.Dir("."))

		port := strconv.Itoa(8000 + i)

		color.Blue("Starting up httpserver, serving ./ \nAvailable on: \n")
		color.Red("  http://127.0.0.1:%s \n", port)

		color.Red("  http://%s:%s\n", getIp(), port)
		// colorPrint("Starting up httpserver, serving ./ \nAvailable on: \n", 2|1)

		err := http.ListenAndServe(":"+port, fs)
		if err != nil {
			fmt.Println(err.Error())
			continue
		} else {
			break
		}
	}
}
