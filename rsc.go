package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os/exec"
)

// 服务端口
var port string

// 默认执行命令
var defaultCommand string

// 是否开启自由命令模式（慎重使用）
var enableDangerMode bool

func defaultMethod(w http.ResponseWriter, r *http.Request) {
	if enableDangerMode {
		// 获取URL的参数
		query := r.URL.Query()
		// 获得URL的id
		command := query.Get("command")
		Command(command)
	} else {
		Command(defaultCommand)
	}
	fmt.Fprintf(w, "The command was executed successfully！")
}

func Command(command string) {
	cmd := exec.Command("/bin/sh", "-c", command)
	output, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(output))
}

func main() {
	flag.Parse()
	http.HandleFunc("/", defaultMethod)
	//过滤掉尝试获取网站图标的请求，以避免每次访问会执行两次命令
	http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	})
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func init() {
	flag.StringVar(&port, "p", "6021", "Change the default(6021) server port.")
	flag.StringVar(&defaultCommand, "c", "echo hello world", "The default executed command.")
	flag.BoolVar(&enableDangerMode, "f", false, "Enable the free command mode? (Use with caution)")
}
