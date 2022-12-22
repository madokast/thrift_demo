package main

import (
	"context"
	"fmt"
	rpc "gtd/go_stu"
	"net"
	"strconv"
	"time"

	"github.com/apache/thrift/lib/go/thrift"
)

// 实现 StudentService
type StudentServiceImpl struct{}

func (s *StudentServiceImpl) GroupByName(ctx context.Context, students []*rpc.Student) (_r map[string][]*rpc.Student, _err error) {
	fmt.Printf("[Go_Server]GroupByName %v\n", students)
	group := make(map[string][]*rpc.Student)
	for _, stu := range students {
		stud, ok := group[stu.Name]
		if ok {
			group[stu.Name] = append(stud, stu)
		} else {
			group[stu.Name] = []*rpc.Student{stu}
		}
	}
	return group, nil
}

func (s *StudentServiceImpl) Length(ctx context.Context, str string) (_r int32, _err error) {
	fmt.Printf("[Go_Server]Length %v\n", str)
	return int32(len(str)), nil
}

func (s *StudentServiceImpl) Distinct(ctx context.Context, values []int32) (_r []int32, _err error) {
	fmt.Printf("[Go_Server]Distinct %v\n", values)
	set := make(map[int32]struct{})
	for _, v := range values {
		set[v] = struct{}{}
	}
	ret := make([]int32, 0, len(set))
	for k := range set {
		ret = append(ret, k)
	}
	return ret, nil
}

func start_server(port int) {
	processor := rpc.NewStudentServiceProcessor(&StudentServiceImpl{})
	serverSocket, err := thrift.NewTServerSocket(net.JoinHostPort("localhost", strconv.Itoa(port)))
	fmt.Println("[Go_Server]start serverSocket. err =", err)
	transportFactory := thrift.NewTTransportFactory()
	protocolFactory := thrift.NewTBinaryProtocolFactoryConf(nil)
	server := thrift.NewTSimpleServer4(processor, serverSocket, transportFactory, protocolFactory)
	err = server.Serve()
	fmt.Println("[Go_Server]start server. err =", err)
}

func start_client(port int) *rpc.StudentServiceClient {
	socket := thrift.NewTSocketConf(net.JoinHostPort("localhost", strconv.Itoa(port)), nil)
	transport := thrift.NewTBufferedTransport(socket, 4096)
	protocol := thrift.NewTBinaryProtocolFactoryConf(nil)
	client := rpc.NewStudentServiceClient(thrift.NewTStandardClient(protocol.GetProtocol(transport), protocol.GetProtocol(transport)))
	err := transport.Open()
	fmt.Println("[Go_Client]started. err =", err)
	return client
}

func main() {
	go start_server(18081)
	time.Sleep(1 * time.Second)
	client := start_client(18081)

	len_r, err := client.Length(context.Background(), "abc")
	fmt.Printf("[Go_Client]Length %v err %v\n", len_r, err)

	dis_r, err := client.Distinct(context.Background(), []int32{1, 1, 2, 2, 3})
	fmt.Printf("[Go_Client]Distinct %v err %v\n", dis_r, err)

	gou_r, err := client.GroupByName(context.Background(), []*rpc.Student{
		{Name: "mdk", Age: 14, Scores: []int32{2, 3}},
		{Name: "mdk", Age: 15, Scores: []int32{3, 4}},
		{Name: "zrx", Age: 14, Scores: []int32{99}},
	})
	fmt.Printf("[Go_Client]GroupByName %v err %v\n", gou_r, err)
}
