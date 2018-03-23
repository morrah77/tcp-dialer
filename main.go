package main

import (
	"bufio"
	"flag"
	"net"
	"os"
	"time"
)

type Conf struct {
	KeyTtl     time.Duration
	ListenAddr string
}

var conf Conf

func init() {
	flag.StringVar(&conf.ListenAddr, `dial-addr`, `localhost:12345`, `Address to dial`)
	flag.Parse()
}

func main() {
	conn, err := net.Dial(`tcp`, conf.ListenAddr)
	if err != nil {
		println(err.Error())
		return

	}
	println(`Dialed to`, conf.ListenAddr)
	defer conn.Close()
	for {
		consoleReader := bufio.NewReader(os.Stdin)
		println(`Waiting for command...`)
		s, err := consoleReader.ReadString('\n');
		if err != nil {
			println(`An error occured during reading`)
		}
		tcpReader := bufio.NewReader(conn)
		_, err = conn.Write([]byte(s))
		if err != nil {
			println(`An error occured during sending`)
		}
		println(`Will read from connection...`)
		resp, err := tcpReader.ReadString('\n')
		if err != nil {
			println(`An error occured during receiving`)
		}
		println(`received: `, resp)
	}
}
