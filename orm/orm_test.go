package orm

import (
	"testing"

	"github.com/donnie4w/simplelog/logging"
	"github.com/donnie4w/tlorm-go/orm"
)

type UserOrm struct {
	Id      int64
	Name    string `idx`
	Age     int
	Level   byte
	Balance float64
}

func Test_orm(t *testing.T) {
	orm.RegisterDefaultResource(false, "db.tlnet.top:7001", "mycli=123")
	//create table
	orm.Create[UserOrm]()
	//insert data
	seq, _ := orm.Insert(&UserOrm{Name: "tom", Age: 20, Level: 1, Balance: 99.2})
	//select
	u, _ := orm.SelectById[UserOrm](seq)
	logging.Debug(u)
}
