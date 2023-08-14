package client

import (
	"testing"

	"github.com/donnie4w/simplelog/logging"
	"github.com/donnie4w/tlcli-go/tlcli"
)

func Test_client(t *testing.T) {
	if client, err := tlcli.NewConnect(false, "db.tlnet.top:7001", "mycli=123"); err == nil {
		//create table
		client.CreateTable("usertest", []string{"name", "age", "type"}, []string{"name", "type"})
		//insert data
		datamap := map[string][]byte{"name": []byte("tom"), "age": []byte("1"), "type": []byte("1")}
		seq, _ := client.Insert("usertest", datamap)
		logging.Debug("id:", seq) //return id
		if data, err := client.SelectById("usertest", seq); err == nil {
			for k, v := range data.TBean {
				logging.Debug(k, "==", string(v))
			}
		}
	}
}

func Test_clientTls(t *testing.T) {
	if client, err := tlcli.NewConnect(false, "db.tlnet.top:7002", "mycli=123"); err == nil {
		//create table
		client.CreateTable("usertest", []string{"name", "age", "type"}, []string{"name", "type"})
		//insert data
		datamap := map[string][]byte{"name": []byte("jerry"), "age": []byte("2"), "type": []byte("2")}
		seq, _ := client.Insert("usertest", datamap)
		logging.Debug("id:", seq) //return id
		if data, err := client.SelectById("usertest", seq); err == nil {
			for k, v := range data.TBean {
				logging.Debug(k, "==", string(v))
			}
		}
	}
}
