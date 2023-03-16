package main

import (
	"bufio"
	"flag"
	"fmt"
	"net/http"
	"os/exec"
)

// 服务端口
var port string

// 默认执行命令
var defaultCommand string

// 是否开启自由命令模式（慎重使用）
var enableDangerMode bool

// 是否向前台返回命令执行结果
var enableOutPut bool

func outputMethod(w http.ResponseWriter, command string) {
	w.Header().Set("Content-Type", "text/event-stream; charset=utf-8")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.WriteHeader(http.StatusOK)
	cmd := exec.Command(command)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Fprintln(w, err.Error())
	}

	if err := cmd.Start(); err != nil {
		fmt.Fprintln(w, err.Error())
	}

	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		fmt.Fprintln(w, scanner.Text())
		w.(http.Flusher).Flush()
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(w, err.Error())
	}

	if err := cmd.Wait(); err != nil {
		fmt.Fprintln(w, err.Error())
	}
}

func defaultMethod(w http.ResponseWriter, r *http.Request) {
	var shellCommand = defaultCommand
	// 获取URL的参数
	query := r.URL.Query()
	// 获得URL的id
	command := query.Get("command")
	if !isEmpty(command) {
		if enableDangerMode {
			shellCommand = command
		} else {
			fmt.Fprintf(w, "Free command mode is not enabled!")
		}
	}
	if enableOutPut {
		outputMethod(w, shellCommand)
	} else {
		cmd := exec.Command(shellCommand)
		output, err := cmd.Output()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(output))
	}
}

func main() {
	flag.Parse()
	fmt.Println("Default command: " + defaultCommand)
	fmt.Println("Enable free command mode: " + fmt.Sprintf("%v", enableDangerMode))
	fmt.Println("Default output mode: " + fmt.Sprintf("%v", enableOutPut))
	fmt.Println("--------------------------------RSC Startup Successful!--------------------------------")
	http.HandleFunc("/", defaultMethod)
	//过滤掉尝试获取网站图标的请求，以避免每次访问会执行两次命令
	http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	})
	err := http.ListenAndServe(":"+port, nil) // 设置监听的端口
	if err != nil {
		fmt.Printf("ListenAndServe: %s", err)
	}
}

func isEmpty(str string) bool {
	return len(str) == 0 || str == ""
}

func init() {
	flag.StringVar(&port, "p", "6021", "Change the default server port.")
	flag.StringVar(&defaultCommand, "d", "echo hello world", "The default executed command.")
	flag.BoolVar(&enableDangerMode, "f", false, "Enable the free command mode? (Use with caution)")
	flag.BoolVar(&enableOutPut, "o", false, "Should the command execution result be returned to the front-end?(Use with caution)")
}
