package core

import (
	"errors"
	"fmt"
	pb "github.com/proto/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
	"sync"
	"time"
)
var appServe *AppServer
type AppServer struct {
	Conn map[string]*grpc.ClientConn //key:address
	Client map[string]pb.DMineMainClient //key address
	//HashIndex map[string]*hashring.Node //key address
	//HashRing *hashring.Consistent //hashRing
	Index int
	ArrAddress []string
	heartBeat  time.Duration
	s sync.RWMutex
}
//初始化appServer
func init(){
	appServe = &AppServer{
		Conn:make(map[string]*grpc.ClientConn),
		Client:make(map[string]pb.DMineMainClient),
		//HashIndex:make(map[string]*hashring.Node),
		heartBeat:    15*time.Second,
		ArrAddress: []string{},
	}
	//appServe.HashRing = hashring.NewConsistent()
	go appServe.healthCheck()

}

func GetRandServer() (pb.DMineMainClient,error){
	return appServe.getServer()
	//return appServe.getRandServer()
}
func GetAllServer() map[string]pb.DMineMainClient{
	return appServe.getAllServer()
}
func Register(address string) error{
	return appServe.register(address)
}

//随机获取server
//func(as *AppServer) getRandServer() (pb.DMineMainClient,error){
//	as.s.Lock()
//	defer as.s.Unlock()
//	value := utils.GenerateKey(128)
//	node := as.HashRing.Get(value)
//	if node.Ip == "" {
//		return nil,errors.New("node节点为空")
//	}
//	fmt.Println("aaa",node.Ip)
//	return as.Client[node.Ip],nil
//}

//顺序获取server
func (as *AppServer) getServer()(pb.DMineMainClient,error){
	as.s.Lock()
	as.s.Unlock()
	if len(as.ArrAddress) == 0 {
		return nil,errors.New("node子节点为空")
	}
	if as.Index>= len(as.ArrAddress){
		as.Index = 0
	}
	address := as.ArrAddress[as.Index]
	as.Index ++
	return as.Client[address],nil

}
//获取所有server
func(as *AppServer) getAllServer() map[string]pb.DMineMainClient{
	as.s.Lock()
	defer as.s.Unlock()
	return as.Client
}


//服务注册
func(as *AppServer) register(address string) error{
	as.s.Lock()
	defer as.s.Unlock()
	if _,ok := as.Conn[address];ok{
		return errors.New("该地址已经注册过")
	}
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err!=nil {
		fmt.Println("注册失败",err)
		return err
	}

	as.Conn[address] = conn
	as.Client[address] = pb.NewDMineMainClient(conn)
	as.ArrAddress = append(as.ArrAddress, address)
	//生成随机数
	//rand.Seed(time.Now().UnixNano())
	//x := time.Now().Unix()
	//strInt64 := strconv.FormatInt(x, 10)
	//intRand,_ := strconv.Atoi(strInt64)
	//加入hash环
	//node := hashring.NewNode(intRand, address, 0, "", 1)
	//as.HashRing.Add(node)
	//as.HashIndex[address] = node
	return nil
}

//移除服务
func (as *AppServer) unRegister(address string) {
	as.s.Lock()
	defer as.s.Unlock()
	delete(as.Conn,address)
	delete(as.Client,address)
	//删除数组元素
	for i:=0;i<len(as.ArrAddress);i++{
		if as.ArrAddress[i] == address{
			as.ArrAddress = append(as.ArrAddress[:i], as.ArrAddress[i+1:]...)
			break
		}
	}

	//node := as.HashIndex[address]
	//as.HashRing.Remove(node)
	//delete(as.HashIndex,address)
}

// 心跳
func (as *AppServer) healthCheck() {
	ticker := time.NewTicker(as.heartBeat)
	for {
		select {
		case <- ticker.C:{
			for k, v := range as.Conn {
				if v.GetState() == connectivity.TransientFailure {
					fmt.Println("grpc断连",k)
					as.unRegister(k)
				}
			}
		}

		}
	}
}

