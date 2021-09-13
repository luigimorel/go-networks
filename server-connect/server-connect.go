package serverconnect

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"net"
	"net/http"
	"time"
)


type StringServer string

func (s StringServer) ServeHTTP(res http.ResponseWriter, req *http.Request) {
res.Write([]byte(string(s)))
}

func createServer(addr string ) http.Server {
	return http.Server{
		Addr: addr ,
		Handler: StringServer("Hello Gogher\n"),
	}
}
const addr = "localhost:7070"

func main() {
	s := createServer(addr)
	go s.ListenAndServe()

	//Connect with plain TCP 
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	_, err = io.WriteString(conn, "GET /HTTP/1.1\r/nHost:localhost:7070\r\n\r\n")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(conn)
	conn.SetReadDeadline(time.Now().Add(time.Second))
	for scanner.Scan(){
		fmt.Println(scanner.Text())
	}

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	s.Shutdown(ctx)
}