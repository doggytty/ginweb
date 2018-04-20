package corm


import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/go-xorm/core"
	"os"
	"fmt"
)

var Engine *xorm.Engine

func init() {
	var err error
	Engine, err = xorm.NewEngine("mysql", "heika_dev:heika_dev@qwe321@tcp(172.16.2.112:3306)/rrd_verify?charset=utf8")
	if err != nil {
		fmt.Println("初始化数据库失败!")
		os.Exit(2)
	}
	// 控制台打印出生成的SQL语句
	Engine.ShowSQL(true)
	Engine.Logger().SetLevel(core.LOG_DEBUG)
	//err = engine.Sync2(new(LbsPosition))
	//if err != nil {
	//	fmt.Println("同步数据库失败!")
	//	fmt.Println(err)
	//	os.Exit(3)
	//}
}

type LbsPosition struct {
	Id uint64 `xorm:"int pk autoincr"`
	ApplicationId string  `xorm:"varchar(64) notnull unique 'application_id'"`
	UserKey string  `xorm:"varchar(64) notnull 'user_key'"`
	Lng float64  `xorm:"double notnull 'lng'"`
	Lat float64  `xorm:"double notnull 'lat'"`
	Status int  `xorm:"int"`
	FormattedAddress string `xorm:"varchar(500)"`
	Confidence int `xorm:"int"`
	Business string `xorm:"varchar(100)"`
	Country string `xorm:"varchar(50)"`
	Province string `xorm:"varchar(50)"`
	City string `xorm:"varchar(50)"`
	District string `xorm:"varchar(50)"`
	Town string `xorm:"varchar(50)"`
	Street string `xorm:"varchar(50)"`
	StreetNumber string `xorm:"varchar(50)"`
	Adcode string `xorm:"varchar(50)"`
	CountryCode int `xorm:"int"`
	RegionsName string `xorm:"varchar(100)"`
	RegionsTag string `xorm:"varchar(100)"`
	RegionsDirection string `xorm:"varchar(100)"`
	SematicDescription string `xorm:"varchar(100)"`
	CityCode int `xorm:"int"`
	WholeMsg string `xorm:"text"`
}

func Get(id uint64) (*LbsPosition, error) {
	var lp LbsPosition
	exist, err := Engine.ID(id).Get(&lp)
	if err != nil {
		return nil, err
	}
	if exist {
		return &lp, nil
	}
	return nil, nil
}

func WhereByApplicationId(applicationId string) (*LbsPosition, error) {
	lp := new(LbsPosition)
	//user := &User{Id:1}
	//has, err := engine.Get(user)
	exist, err := Engine.Where("application_id=?", applicationId).Get(lp)
	if err != nil {
		return nil, err
	}
	if exist {
		return lp, nil
	}
	return nil, nil
}

func CountByApplicationId(applicationId string) (int64, error) {
	lp := new(LbsPosition)
	num, err := Engine.Where("application_id=?", applicationId).Count(lp)
	if err != nil {
		return 0, err
	}
	return num, nil
}

func InsertLBS(lp *LbsPosition) (int64, error) {
	// 查询是否存在，存在则更新，不存在则删除
	num, err := CountByApplicationId(lp.ApplicationId)
	if err != nil {
		return 9001, err
	}
	if num == 1{
		// update
		num, err = Engine.Where("application_id=", lp.ApplicationId).Omit("application_id").Update(lp)
	} else {
		// insert
		num, err = Engine.Insert(lp)
	}
	if err != nil {
		return 0, err
	}
	return num, nil
}
