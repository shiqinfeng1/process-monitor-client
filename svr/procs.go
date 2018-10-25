package svr

import (
	"fmt"
	"os/exec"

	"github.com/shiqinfeng1/process-monitor-client/conf"
	"github.com/shiqinfeng1/process-monitor-client/xlog"
)

//StartCheck 开始自动检测进程
func StartCheck(command string) {
	cmdStr := fmt.Sprintf(
		"nohup %s 1>/dev/null 2>"+xlog.ErrorFile+"&",
		command,
	)
	cmd := exec.Command("sh", "-c", cmdStr)
	_, err := cmd.CombinedOutput()
	if err != nil {
		xlog.Fatal(xlog.ErrorFile, "err", err)
	}
}

//StopCheck 停止自动检测进程
func StopCheck(command string) {
	cmdStr := fmt.Sprintf(
		"ps -ef|grep -v grep|grep \"%s\" |cut -c 9-15|xargs kill -9",
		command,
	)
	err := exec.Command("sh", "-c", cmdStr).Run()
	if err != nil {
		xlog.Fatal(xlog.ErrorFile, "err", err)
	}
	fmt.Print("StopCheck Process: monitor", " stop ")
	fmt.Printf("%c[1;40;32m%s%c[0m\n", 0x1B, "[success]", 0x1B)
}

//StartProc 启动进程
func StartProc(conf *conf.Config) {
	config := *conf
	command := config.Command
	logfile := config.Logfile
	cmdStr := fmt.Sprintf(
		"nohup %s 1>/dev/null 2>"+logfile+"&",
		command,
	)
	if ok := CheckProc(command); !ok {
		err := exec.Command("sh", "-c", cmdStr).Run()
		if err != nil {
			fmt.Print("Process:", config.ProcessName, " start ")
			fmt.Printf("%c[1;40;31m%s%c[0m\n", 0x1B, "[fail]", 0x1B)
			xlog.Fatal(logfile, err)
		}
		fmt.Print("Process:", config.ProcessName, " start ")
		fmt.Printf("%c[1;40;32m%s%c[0m\n", 0x1B, "[success]", 0x1B)
	} else {
		fmt.Print("Process:", config.ProcessName)
		fmt.Printf("%c[1;40;32m%s%c[0m\n", 0x1B, " is already start...", 0x1B)
	}
}

//StopProc 停止进程
func StopProc(conf *conf.Config) {
	config := *conf
	fmt.Printf("Stop Config=%+v", config)
	command := config.Command
	logfile := config.Logfile
	if ok := CheckProc(command); !ok {
		fmt.Print("Process:", config.ProcessName)
		fmt.Printf("%c[1;40;32m%s%c[0m\n", 0x1B, " is already stop...", 0x1B)
	} else {
		cmdStr := fmt.Sprintf(
			"ps -ef| grep -v grep|grep \"%s\"|cut -c 9-15|xargs kill -9",
			command,
		)
		err := exec.Command("sh", "-c", cmdStr).Run()
		if err != nil {
			fmt.Print("Process:", config.ProcessName, " stop ")
			fmt.Printf("%c[1;40;31m%s%c[0m\n", 0x1B, "[fail]", 0x1B)
			xlog.Fatal(logfile, err)
		}
		fmt.Print("Process:", config.ProcessName, " stop ")
		fmt.Printf("%c[1;40;32m%s%c[0m\n", 0x1B, "[success]", 0x1B)
	}
}

//GetProc 获取当前进程信息
func GetProc(conf *conf.Config) {
	config := *conf
	command := config.Command
	logfile := config.Logfile
	if ok := CheckProc(command); !ok {
		fmt.Println("Process:", config.ProcessName, " is already stop...")
	} else {
		cmdStr := fmt.Sprintf(
			"ps -ef| grep -v grep|grep \"%s\"",
			command,
		)
		cmd := exec.Command("sh", "-c", cmdStr)
		out, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Print("Process:", config.ProcessName, " Status ")
			fmt.Printf("%c[1;40;31m%s%c[0m\n", 0x1B, "[fail]", 0x1B)
			xlog.Fatal(logfile, err)
		}
		fmt.Println("Process:", config.ProcessName, " Status:")
		fmt.Println(string(out))
	}
}

//RestartProc 重启进程
func RestartProc(conf *conf.Config) {
	//stop proc
	StopProc(conf)
	//start proc
	StartProc(conf)
}

//CheckProc 检查进程是否已经启动
func CheckProc(command string) (ok bool) {
	cmdStr := fmt.Sprintf(
		"ps aux| grep -v grep|grep \"%s\" |awk '{print $2}'",
		command,
	)
	cmd := exec.Command("sh", "-c", cmdStr)
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("err:", err)
	}
	if string(out) == "" {
		return false
	}
	return true
}
