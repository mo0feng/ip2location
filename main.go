package main

import (
	"os"
	"fmt"
	"ip2location/lib"
	"strings"
)

func main(){
	err := lib.LoadCz("./qqwry.dat")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	//concurrency limit
	limit := 10000
	//read file
	var ip = new(lib.IpLists)
	ip.F = "./ip_list.list"
	err = ip.LoadIpFile()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("concurrency limit：",limit)
	C:
	var con string
	fmt.Print("contiue Y/N：")
	fmt.Scan(&con)
	switch strings.ToLower(con){
	case "n":
		os.Exit(1)
	case "y":
		break
	default:
		goto C
	}
	var c = lib.NewIpMap()
	c.CountIp(ip.IpList,limit)
}