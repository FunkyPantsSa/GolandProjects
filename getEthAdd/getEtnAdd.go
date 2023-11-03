package main

import (
	"bytes"
	"os/exec"
	"time"
)

func getIp() string {
	result, _ := Command("ip addr |grep inet6")
	//test
	//result, err := Command("cat 1.txt")
	//if err != nil {
	//	fmt.Println(err)
	//	return string("error:")

	//}
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
