package mapreduce

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
	"os"
)

type Master struct {
}

// Server func taken from MIT 6.824
func (m *Master) server() {
	rpc.Register(m)
	rpc.HanddleHTTP()
	sockname := masterSock()
	os.Remove(sockname)
	l, e := net.Listen("unix", sockname)
	if e != nil {
		log.Fatal("Listen Error:", e)
	}

	go http.Serve(l, nil)
}

func (m *Master) Done() bool {
	ret := false

	return ret
}

func MakeMaster(file []string, nReduce int) *Master {
	m := Master{}
	m.server()
	return &m
}
