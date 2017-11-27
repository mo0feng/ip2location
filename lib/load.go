package lib

import (
	"time"
	"log"
	"os"
	"errors"
	"io/ioutil"
	"strings"
	"fmt"
)

func LoadCz(file_path string) error {
	IpData.FilePath = file_path
	startTime := time.Now().UnixNano()
	res := IpData.InitIpData()
	if v, ok := res.(error); ok {
		return fmt.Errorf("read ipdata error. %v", v)
	}
	endTime := time.Now().UnixNano()
	log.Printf("ip address base load success. total:%d item, consuming:%.1f ms\n", IpData.IpNum, float64(endTime - startTime) / 1000000)
	return nil
}

type IpLists struct {
	F      string
	IpList []string
}

func (i *IpLists)LoadIpFile() error {
	_, err := os.Stat(i.F)
	if err != nil && os.IsNotExist(err) {
		return errors.New("ip_list.list not exist")
	}
	startTime := time.Now().UnixNano()
	path, err := os.OpenFile(i.F, os.O_RDONLY, 0400)
	if err != nil {
		return err
	}
	defer path.Close()
	tmpData, err := ioutil.ReadAll(path)
	if err != nil {
		return err
	}
	i.IpList = strings.Split(string(tmpData), "\n")
	endTime := time.Now().UnixNano()
	log.Printf("ip_list.list load success. total:%d item, consuming:%.1f ms\n", len(i.IpList), float64(endTime - startTime) / 1000000)
	return nil
}