package dao

import (
	"github.com/go-xorm/xorm"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type SendEmailInfo struct {
	Id         int `xorm:"not null pk autoincr comment('主键id') Int(11)" form:"id"`
	FromEmail  string `xorm:"not null comment('来源地址') VARCHAR(1000)" form:"from"`
	ToEmail    string `xorm:"not null comment('发送邮件') VARCHAR(1000)" form:"to"`
	CcEmail    string `xorm:"not null comment('cc邮件地址') VARCHAR(1000)" form:"cc"`
	Subject    string `xorm:"not null comment('主题') TEXT" form:"subject"`
	Body       string `xorm:"not null comment('邮件内容') TEXT" form:"body"`
	IsHtml     int `xorm:"not null comment('是否html') Int(11)" form:"isHtml"`
	IsAttach   int `xorm:"not null comment('是否附件') Int(11)" form:"isAttach"`
	Attach     string `xorm:"not null comment('附件地址') VARCHAR(1000)" form:"attach"`
	Ctime      string `xorm:"not null comment('创建时间') Int(11)" form:"ctime"`
	Etime      string `xorm:"not null comment('修改时间') Int(11)" form:"etime"`
	Extension  string `xorm:"comment('扩展信息') Text" form:"extension"`
	Unique_key string `xorm:"not null comment('唯一标识') VARCHAR(500)" form:"unique_key"`
	User_ip    string `xorm:"not null comment('用户ip') varchar(25)" form:"user_ip"`
	Status     int `xorm:"not null comment('发邮件状态') Int(11)" form:"status"`
	Weights    int `xorm:"not null comment('权重') Int(11)" form:"weights"`
}

type SendEmailInfoModule struct {
	engine *xorm.Engine
}

func NewSendEmailInfoModule(engine *xorm.Engine) *SendEmailInfoModule {
	return &SendEmailInfoModule{
		engine: engine,
	}
}
func (d *SendEmailInfoModule) Get(id int) *SendEmailInfo {
	data := &SendEmailInfo{Id:id}
	ok,err := d.engine.Get(data)
	if ok && err == nil {
		return data
	}else {
		return nil
		data.Id = 0
		return data
	}
	return nil
}
func (d *SendEmailInfoModule) GetAll() []SendEmailInfo {
	// 两种构造方式都可以
	datalist := make([]SendEmailInfo, 0)
	//datalist := make(map[int]SendEmailInfo)
	//datalist := []SendEmailInfo{}
	err := d.engine.Desc("id").Find(&datalist)
	if err != nil {
		log.Println(err)
		return datalist
		//return nil
	} else {
		return datalist
	}
	return nil
}
func (d *SendEmailInfoModule) Search(country string) []SendEmailInfo {
	datalist := make([]SendEmailInfo, 0)
	//datalist := []SendEmailInfo{}
	err := d.engine.Where("country=?", country).Desc("id").Find(&datalist)
	if err != nil {
		println(err)
		return datalist
	} else {
		return datalist
	}
}
func (d *SendEmailInfoModule) Delete(id int) error {
	data := &SendEmailInfo{Id: id,Status:5}
	_, err := d.engine.Id(data.Id).Update(data)
	return err
}
func (d *SendEmailInfoModule) Update(data *SendEmailInfo, column []string) error {
	_, err := d.engine.Id(data.Id).MustCols(column...).Update(data)
	return err
}
