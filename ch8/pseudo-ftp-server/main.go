package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"os"
	"path"
	"strings"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	wd := "."

	defer conn.Close()
	input := bufio.NewScanner(conn)
	for input.Scan() {
		line := input.Text()
		args := strings.SplitN(line, " ", 2)
		switch args[0] {
		case "cd":
			if len(args) < 2 {
				fmt.Fprintf(conn, "Missing argument for %s\n", args[0])
				continue
			}
			arg := strings.TrimSpace(args[1])
			if strings.HasPrefix(arg, "/") {
				wd = arg
			} else {
				wd = path.Join(wd, arg)
			}
			wd = path.Clean(wd)
			fmt.Fprintf(conn, "Changed directory: %s\n", wd)
		case "ls":
			if len(args) < 2 {
				args = append(args, ".")
			}
			var fpath string
			arg := strings.TrimSpace(args[1])
			if strings.HasPrefix(arg, "/") {
				fpath = arg
			} else {
				fpath = path.Join(wd, arg)
			}
			fpath = path.Clean(fpath)
			files, err := ioutil.ReadDir(fpath)
			if err != nil {
				log.Print(err)
				fmt.Fprintf(conn, "Can't list %s: %v\n", fpath, err)
				continue
			}
			for _, f := range files {
				fmt.Fprintf(conn, "%s\n", f.Name())
			}
		case "get":
			if len(args) < 2 {
				fmt.Fprintf(conn, "Missing argument for %s\n", args[0])
				continue
			}
			var fpath string
			arg := strings.TrimSpace(args[1])
			if strings.HasPrefix(arg, "/") {
				fpath = arg
			} else {
				fpath = path.Join(wd, arg)
			}
			fpath = path.Clean(fpath)
			file, err := os.Open(fpath)
			if err != nil {
				log.Print(err)
				fmt.Fprintf(conn, "Can't open %s: %v\n", fpath, err)
				continue
			}
			_, _ = io.Copy(conn, file)
		case "close":
			fmt.Fprintf(conn, "Bye!\n")
			return
		}
	}
}
