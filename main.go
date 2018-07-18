package main

import (
	"pb"
	"antnet"
	"github.com/golang/protobuf/proto"
)

var gamerMap = map[string]antnet.IMsgQue{}

func C2SHandlerFunc(msgque antnet.IMsgQue, msg *antnet.Message) bool {
	ppb := msg.C2S().(*pb.Pb)
	antnet.LogInfo("%v", ppb.GamerLoginC2S)
	antnet.LogInfo("%v", ppb.GamerGlobalChatC2S)

	if ppb.GamerLoginC2S != nil{
		gamerMap[ppb.GamerLoginC2S.GetId()] = msgque
		s2c := &pb.Pb {
			GamerLoginS2C: &pb.Pb_GamerLoginS2C{Json: proto.String("登陆成功"),},
		}
		msgque.Send(antnet.NewDataMsg(antnet.PbData(s2c)))
		msgque.SetUser(ppb.GamerLoginC2S.GetId())
	} else if ppb.GamerGlobalChatC2S != nil {
		s2c := &pb.Pb {
			GamerGlobalChatS2C: &pb.Pb_GamerGlobalChatS2C{Data: ppb.GamerGlobalChatC2S.Data,},
		}
		antnet.Send(antnet.NewDataMsg(antnet.PbData(s2c)), func(msgque antnet.IMsgQue) bool {
			return true
		})
	} else if ppb.GamerChatC2S != nil {
		s2c := &pb.Pb{
			GamerChatS2C: &pb.Pb_GamerChatS2C{
				FromId: proto.String(msgque.GetUser().(string)),
				Data: proto.String("玩家不在线"),
			},
		}
		if mq, ok := gamerMap[ppb.GamerChatC2S.GetId()]; ok && mq.Available(){
			s2c.GamerChatS2C.Data = proto.String("发送成功")
			msgque.Send(antnet.NewDataMsg(antnet.PbData(s2c)))

			s2c.GamerChatS2C.Data = ppb.GamerChatC2S.Data
			mq.Send(antnet.NewDataMsg(antnet.PbData(s2c)))
		} else {
			msgque.Send(antnet.NewDataMsg(antnet.PbData(s2c)))
		}
	}

	return true
}

func main() {
	ExtNetHandler := &antnet.DefMsgHandler{}
	PbParser := &antnet.Parser{Type: antnet.ParserTypePB}
	PbParser.RegisterMsg(&pb.Pb{}, nil)
	ExtNetHandler.RegisterMsg(&pb.Pb{}, C2SHandlerFunc)
	antnet.StartServer("ws://:5000", antnet.MsgTypeCmd, ExtNetHandler, PbParser)
	antnet.WaitForSystemExit(nil)
}
