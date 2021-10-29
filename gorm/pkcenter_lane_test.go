package gorm

import (
	"fmt"
	"testing"

	"github.com/micro/go-micro/config"
	"github.com/micro/go-micro/config/source/file"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type pkcenter_camera_lane struct {
	ID       uint `gorm:"column:id"`
	CameraId uint `gorm:"column:camera_id"`
	LaneId   uint `gorm:"column:lane_id"`
}

func LoadConfig() {
	// 加载config.yaml配置文件
	err := config.Load(
		file.NewSource(file.WithPath("./config.yaml")),
	)
	if err != nil {
		panic(err)
	}
	fmt.Println("加载配置文件信息成功")
}

func FindPkcneter() (string, string) {
	// 加载配置
	LoadConfig()

	dbType := config.Get("db", "type").String("mysql")

	dbCoreAddr := config.Get("db", "serviceDB").String("none")
	return dbType, dbCoreAddr

}

func TestOpenGorm(t *testing.T) {
	dbType, dbCoreAddr := FindPkcneter()
	fmt.Printf("当前的dbCoreAddr是 %s\n", dbCoreAddr)
	fmt.Printf("当前的dbType是 %s\n", dbType)

	db, err := gorm.Open(dbType, dbCoreAddr)
	if err != nil {
		panic(err)
	}
	db = db.Debug()

	var lanes []pkcenter_camera_lane
	db.Raw("SELECT * from pkcenter_camera_lane where camera_id = ?", 1114).Scan(&lanes)

	fmt.Println(lanes)

}
