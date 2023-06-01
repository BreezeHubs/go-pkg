package etcdsvcdispkg

// PingRouter MsgId=1的路由
//type PingRouter struct {
//	znet.BaseRouter
//}

// Handle Ping Handle MsgId=1的路由处理方法
//func (r *PingRouter) Handle(request ziface.IRequest) {
//	fmt.Println("req", request.GetMsgID(), string(request.GetData()))
//
//	//fmt.Printf("Received from client: %s\n", request.GetData())
//	if err := request.GetConnection().SendMsg(2, typexpkg.StringToBytes(
//		"192.168.0.1:8080",
//	)); err != nil {
//		log.Printf("Failed to send response: %v\n", err)
//	}
//}
