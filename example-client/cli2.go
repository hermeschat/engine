package main

import (
	"context"
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"hermes/api/pb"
)

func main() {
	//con, err := grpc.Dial("https://chat.paygear.ir:443")

	con, err := grpc.Dial("localhost:9000", grpc.WithInsecure())
	//con, err := grpc.Dial("192.168.41.221:30050", grpc.WithInsecure())
	if err != nil {
		logrus.Fatalf("error : %v", err)
	}
	ctx, _ := context.WithTimeout(context.Background(), time.Hour*2)
	md := metadata.Pairs("Authorization", "bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJhdWQiOiJiOWRjNzEyYzk1MmI0YWFmYjQ4MWFiZWRlMGZlYzRkOCIsImV4cCI6MTU2ODg3OTgxNCwibmJmIjoxNTY2Mjg3ODE0LCJpZCI6IjVjNGMyNjgzYmZkMDJhMmI5MjNhZjhiZiIsIm1lcmNoYW50X3JvbGVzIjp7IjViMmRmZTA0Y2YyNjU2MDAwYzk3YWFlNyI6WyJhZG1pbiJdfSwicm9sZSI6WyJ6ZXVzIl0sImFwcCI6IjU5YmVjM2ZhMGVjYTgxMDAwMWNlZWI4NiIsInN2YyI6eyJhY2NvdW50Ijp7InBlcm0iOjB9LCJjYXNoaWVyIjp7InBlcm0iOjB9LCJjYXNob3V0Ijp7InBlcm0iOjB9LCJjbHViIjp7InBlcm0iOjB9LCJjbHViX3NlcnZpY2UiOnsicGVybSI6MH0sImNvdXBvbiI6eyJwZXJtIjowfSwiY3JlZGl0Ijp7InBlcm0iOjB9LCJkZWxpdmVyeSI6eyJwZXJtIjowfSwiZXZlbnQiOnsicGVybSI6MH0sImZpbGUiOnsicGVybSI6MH0sImdhbWlmaWNhdGlvbiI6eyJwZXJtIjowfSwiZ2VvIjp7InBlcm0iOjB9LCJtZXNzYWdpbmciOnsicGVybSI6MH0sIm5vdGljZXMiOnsicGVybSI6MH0sInBheW1lbnQiOnsicGVybSI6MH0sInByb2R1Y3QiOnsicGVybSI6MH0sInB1c2giOnsicGVybSI6MH0sInFyIjp7InBlcm0iOjB9LCJzZWFyY2giOnsicGVybSI6MH0sInNldHRsZW1lbnQiOnsicGVybSI6MH0sInNvY2lhbCI6eyJwZXJtIjowfSwic3luYyI6eyJwZXJtIjowfSwidHJhbnNwb3J0Ijp7InBlcm0iOjB9LCJ3YXJnIjp7InBlcm0iOjB9fX0.N2nBsDWaRwsOvRshxtI73HjbmgeED3-RfoH9F0d8qzY")
	ctx = metadata.NewOutgoingContext(ctx, md)
	cli := pb.NewHermesClient(con)

	// resp, err := cli.CreateSession(ctx, &pb.CreateSessionRequest{
	// 	ClientType: "Ubuntu",
	// 	UserAgent:  "Terminal",
	// })
	// if err != nil {
	// 	panic(err)
	// }
	// sid := resp.SessionID
	// logrus.Info(sid)
	msgs, err := cli.ListMessages(ctx, &pb.ChannelID{Id: "85403be4-7ceb-4cbd-965f-9bb9efc64963"})
	if err != nil {
		panic(err)
	}
	fmt.Println(msgs)
	return
	eventCli, err := cli.EventBuff(ctx)
	if err != nil {
		panic(err)
	}
	err = eventCli.Send(&pb.Event{Event: &pb.Event_Join{&pb.JoinSignal{SessionId: "7c222aa1-68cd-4a84-b4d5-039941180323"}}})
	if err != nil {
		panic(err)
	}

	logrus.Info("Wait for any event")
	for {
		ev, err := eventCli.Recv()
		if err != nil {
			continue
		}
		switch ev.GetEvent().(type) {
		case *pb.Event_Read:
			logrus.Info("Message has been read")
		case *pb.Event_NewMessage:
			logrus.Info("New Message recieved")
			m := ev.GetNewMessage()
			logrus.Infof("%+v", m)
		case *pb.Event_Dlv:
			logrus.Info("Message delivered")
		}
	}
	//emp, err := cli.Echo(ctx, &pb.Empty{})
	//if err != nil {
	//	panic(err)
	//}
	//logrus.Infof("status is %v", emp.Status)
	//err = eventCli.SendMsg(&pb.Event{
	//	Event: &pb.Event_Read{
	//		Read: &pb.ReadSignal{},
	//	},
	//})
	//if err != nil {
	//	panic(err)
	//}
	time.Sleep(time.Second * 100)
	//_, err = cli.Echo(ctx, &pb.Some{})
	//if err != nil {
	//	logrus.Fatalf("error : %v", err)
	//}
	//resp, err := cli.CreateSession(ctx, &pb.CreateSessionRequest{
	//	ClientType: "Proudly Windows",
	//	UserID:     os.Args[1],
	//})
	//if err != nil {
	//	logrus.Fatalf("error: %v", err)
	//}
	//
	//logrus.Println(resp.SessionID)
	//sid := resp.SessionID
	//sid := "7c222aa1-68cd-4a84-b4d5-039941180323"
	//_, err = cli.Join(ctx, &pb.JoinSignal{
	//	UserID: "amir",
	//	SessionId: sid,
	//})
	//if err != nil {
	//	panic(err)
	//}
	//msgCli, err := cli.NewMessage(context.Background())
	//if err != nil {
	//	panic(err)
	//}
	//cli.ListChannels()
	//err = msgCli.Send(&pb.Message{
	//	MessageType: "1",
	//	From:"amir",
	//	To:"reza",
	//	Body: "hey",
	//})
	//if err != nil {
	//	panic(err)
	//}

	// cli.NewMessage(ctx)
}
