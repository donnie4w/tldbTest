package mq

import (
	"testing"

	"github.com/donnie4w/simplelog/logging"
	"github.com/donnie4w/tlmq-go/cli"
	. "github.com/donnie4w/tlmq-go/stub"
)

func TestMq(t *testing.T) {
	sc := &cli.SimpleClient{Url: "ws://db.tlnet.top:5001", Auth: "mymq=123"}
	sc.PubJsonHandler(func(jmb *JMqBean) { logging.Debug("PubJson >> ", jmb) })
	sc.Connect()
	//Subscriber `userinfotopic`
	sc.Sub("userinfotopic")
	//publish  `userinfotopic`
	sc.PubJson("userinfotopic", "this is first pub json msg")
}

func TestMqTls(t *testing.T) {
	sc := &cli.SimpleClient{Url: "ws://db.tlnet.top:5002", Auth: "mymq=123"}
	sc.PubByteHandler(func(mb *MqBean) { logging.Debug("PubByte >> ", mb) })
	sc.Connect()
	//Subscriber `imtopic`
	sc.Sub("imtopic")
	//publish  `imtopic`
	sc.PubByte("imtopic", []byte("this is first pub bytes msg"))
}
