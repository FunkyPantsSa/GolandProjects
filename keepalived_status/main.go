package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"time"
)

func main() {
	for {
		var a string
		a = getIp()
		fmt.Println(getIp())
		if a == "" {
			restart()
			//test
			//fmt.Println("reatrt")
		}
		time.Sleep(5000 * time.Millisecond)
		fmt.Println("ok: ", a)
	}
	//fmt.Println(getIp())
}

func getIp() string {
	result, err := Command("ip addr |grep 172")
	//test
	//result, err := Command("cat 1.txt")
	if err != nil {
		return string("error:")
		fmt.Println(err)
	}
	//cores, err := strconv.Atoi(strings.Trim(result, "\r\n"))
	//if err != nil {
	//	return 0, err
	//}
	return result

}

func Command(cmd string) (string, error) {
	// logger.Infof("执行命令:%s", cmd)
	var buffer bytes.Buffer
	c := exec.Command("bash", "-c", cmd)
	c.Stdout = &buffer
	err := c.Run()
	if err != nil {
		return "", err
	}
	// 暂停50毫秒
	time.Sleep(50 * time.Millisecond)
	return buffer.String(), nil
}

func restart() {
	Command("systemctl restart keepalived")
	fmt.Println("restart keepalived")
}
