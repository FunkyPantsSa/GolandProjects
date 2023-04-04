package main

import (
	"bytes"
	"log"
	"os/exec"
	"time"
)

func main() {
	for {
		var a string
		a = getIp()
		//fmt.Println(getIp())
		if a == "" {
			restart()
			//test
			//fmt.Println("reatrt")
		}
		//fmt.Println(time.Now())
		//fmt.Println("已完成检测", a)
		log.Printf("已完成检测，检测结果： %s", a)
		time.Sleep(500000 * time.Millisecond)

	}
	//fmt.Println(getIp())
}

func getIp() string {
	result, _ := Command("ip addr |grep 174")
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

func restart() {
	Command("systemctl restart keepalived")
	//fmt.Println("restart keepalived")
	log.Printf("restrat keepalived")
	time.Sleep(50000 * time.Millisecond)

}
