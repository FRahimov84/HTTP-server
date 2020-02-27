package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	file, err := os.OpenFile("server-log.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Printf("can't open log file %e", err)
		return
	}
	defer func() {
		err := file.Close()
		if err != nil {
			fmt.Printf("can't close log file %e", err)
		}
	}()

	log.SetOutput(file)
	log.Print("start application\n")

	host := "0.0.0.0"
	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = "9999"
	}
	err = startServer(fmt.Sprintf("%s:%s", host, port))
	if err != nil {
		log.Fatal(err)
	}

}

func startServer(addr string) error {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Printf("can't listen 0.0.0.0:9999 %v", err)
		return err
	}
	defer func() {
		err = listener.Close()
		if err != nil {
			log.Printf("Can't close Listener: %v", err)
		}
	}()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("can't connect client")
			continue
		}
		fmt.Println("some connection")
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	readString, _ := reader.ReadString('\n')
	split := strings.Split(strings.TrimSpace(readString), " ")
	if len(split) != 3 {
		log.Fatal("")
	}
	meth, requet, protocol := split[0], split[1], split[2]
	if meth == "GET" && protocol == "HTTP/1.1" {
		switch requet {
		case "/":
			all, _ := ioutil.ReadFile("index.html")
			writer := bufio.NewWriter(conn)
			writer.WriteString("HTTP/1.1 200 OK\r\n")
			writer.WriteString(fmt.Sprintf("Content-Length: %d\r\n", len(all)))
			writer.WriteString("Content-Type: text/html\r\n")
			writer.WriteString("Connection: Close\r\n")
			writer.WriteString("\r\n")
			writer.Write(all)
			writer.Flush()
		case "/1.png":
			all, _ := ioutil.ReadFile("1.png")
			writer := bufio.NewWriter(conn)
			writer.WriteString("HTTP/1.1 200 OK\r\n")
			writer.WriteString(fmt.Sprintf("Content-Length: %d\r\n", len(all)))
			writer.WriteString("Content-Type: image/png\r\n")
			writer.WriteString("Connection: Close\r\n")
			writer.WriteString("\r\n")
			writer.Write(all)
			writer.Flush()
		case "/2.jpg":
			all, _ := ioutil.ReadFile("2.jpg")
			writer := bufio.NewWriter(conn)
			writer.WriteString("HTTP/1.1 200 OK\r\n")
			writer.WriteString(fmt.Sprintf("Content-Length: %d\r\n", len(all)))
			writer.WriteString("Content-Type: image/jpg\r\n")
			writer.WriteString("Connection: Close\r\n")
			writer.WriteString("\r\n")
			writer.Write(all)
			writer.Flush()
		case "/index-html.html":
			all, _ := ioutil.ReadFile("index-html.html")
			writer := bufio.NewWriter(conn)
			writer.WriteString("HTTP/1.1 200 OK\r\n")
			writer.WriteString(fmt.Sprintf("Content-Length: %d\r\n", len(all)))
			writer.WriteString("Content-Type: text/html\r\n")
			writer.WriteString("Connection: Close\r\n")
			writer.WriteString("\r\n")
			writer.Write(all)
			writer.Flush()
		case "/123.txt":
			all, _ := ioutil.ReadFile("123.txt")
			writer := bufio.NewWriter(conn)
			writer.WriteString("HTTP/1.1 200 OK\r\n")
			writer.WriteString(fmt.Sprintf("Content-Length: %d\r\n", len(all)))
			writer.WriteString("Content-Disposition: attachment; filename=down.txt\r\n")
			writer.WriteString("Connection: Close\r\n")
			writer.WriteString("\r\n")
			writer.Write(all)
			writer.Flush()
		case "/1.pdf":
			all, _ := ioutil.ReadFile("1.pdf")
			writer := bufio.NewWriter(conn)
			writer.WriteString("HTTP/1.1 200 OK\r\n")
			writer.WriteString(fmt.Sprintf("Content-Length: %d\r\n", len(all)))
			writer.WriteString("Content-Type: application/pdf\r\n")
			writer.WriteString("Connection: Close\r\n")
			writer.WriteString("\r\n")
			writer.Write(all)
			writer.Flush()
		}

	}

}
