package lib

import (
	"sync"
	"time"
	"fmt"
	"os/exec"
	"os"
	"runtime"
	"log"
	"bufio"
)

type IpMap struct {
	Lock       *sync.RWMutex
	Data      map[string]int
}

func NewIpMap() *IpMap {
	return &IpMap{
		Lock:new(sync.RWMutex),
		Data:make(map[string]int),
	}
}

var clear map[string]func()

func init(){
	clear = make(map[string]func())
	clear["linux"] = func() {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}
func CallClear() {
	value, ok := clear[runtime.GOOS]
	if ok {
		value()
	} else {
		panic("Your platform is unsupported! I can't clear terminal screen :(")
	}
}

func (m *IpMap)CountIp(ips []string, limit int) error {
	totalLen := len(ips)
	startTime := time.Now().UnixNano()
	errIp := ""
	for i := 0; i <= totalLen; i+= limit {
		fLen := limit
		if (totalLen-i) < limit {
			fLen = totalLen - i
			log.Println(fLen)
		}
		var chs = make([]chan string, fLen)
		for j := 0; j < fLen; j++ {
			chs[j] = make(chan string)
			go Ip2(ips[i+j], chs[j])
		}
		m.Lock.Lock()
		tmpI := m.Data
		m.Lock.Unlock()
		for _,ch := range chs {
			select {
			case rs := <- ch:
				tmpI[rs]++
			case <-time.After(10 * time.Second):
				log.Println("time out")
				break
			}
		}
		m.Lock.Lock()
		m.Data = tmpI
		m.Lock.Unlock()
		CallClear()
		fmt.Println("processing ...",fmt.Sprintf("%.2f",float64(i)/float64(totalLen)*100),"%", i, totalLen)
	}
	endTime := time.Now().UnixNano()
	CallClear()
	fmt.Println("process:",totalLen,"item,","consuming:", float64(endTime - startTime) / 1000000,"ms")
	m.Lock.Lock()
	tmpR := m.Data
	m.Lock.Unlock()
	
	txtList := ""
	for k,v := range tmpR {
		txtList +=  k+";"+fmt.Sprintf("%v",v)+"\n"
	}
	fName := fmt.Sprintf("%v.txt",time.Now().Unix())
	f, _ := os.OpenFile(fName ,os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	w := bufio.NewWriter(f)
	w.WriteString(txtList)
	w.Flush()
	f.Close()
	
	fmt.Println("out file:", fName)
	fmt.Println("error ip:", errIp)
	return nil
}

func Ip2(ip string, ch chan string) {
	qqWry := NewQQwry()
	rl := qqWry.Find(ip)
	ch <- rl.Country+";"+rl.Area
}