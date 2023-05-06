package main

import (
	"time"

	"github.com/jlaffaye/ftp"
	"github.com/sirupsen/logrus"
)

func main() {
	c, err := ftp.Dial("ip:port", ftp.DialWithTimeout(10*time.Second), ftp.DialWithDisabledEPSV(false))
	//ftp.DialWithDisabledEPSV   ftp的主被动模式
	if err != nil {
		logrus.Fatalf("init err:%s", err.Error())
	}
	err = c.Login("user", "pwd")
	if err != nil {
		logrus.Fatalf("login err:%s", err.Error())
	}
	var walkPath = "path"

	entries, err := c.List(walkPath)
	if err != nil {
		logrus.Fatalf("list err:%s", err.Error())
	}
	for _, entry := range entries {
		logrus.Println(entry.Name)
	}
	var fileCount, notFile int
	w := c.Walk(walkPath)
	for {
		if w.Next() {
			entry := w.Stat()
			if entry != nil {
				if entry.Type != ftp.EntryTypeFile {
					notFile++
					continue
				}
				fileCount++
			}

		} else {
			break
		}
	}
	if err := c.Quit(); err != nil {
		logrus.Errorf("quit err:%s", err.Error())
	}
	logrus.Infof("walk %s done. notFileCount:%d fileCount:%d", walkPath, notFile, fileCount)
}
