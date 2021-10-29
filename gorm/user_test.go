package gorm

import (
	"fmt"
	"testing"
	"time"

	"github.com/micro/go-micro/config"
	"github.com/micro/go-micro/config/source/file"
	"github.com/micro/go-micro/v2/logger"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Product struct {
	Code  uint
	Price float64
}

// 使用结构体tag指定列名
type User struct {
	// gorm.Model GORM内置的一个`gorm.Model`结构体, 包含了`ID`, `CreatedAt`, `UpdateAt`, `DeletedAt`

	// 指定主键
	ID        uint      `gorm:"column:user_id;primary_key;auto_increment"`
	Name      string    `gorm:"column:user_name"`
	Birthday  time.Time `gorm:"column:user_birthday"`
	CreatedAt time.Time `gorm:"column:user_create_at"`
	UpdatedAt time.Time `gorm:"column:user_updated_at"`
	// 使用*时,可以输入null
	DeletedAt *time.Time `gorm:"column:user_deleted_at"`
}

func TestInit(t *testing.T) {
	// 加载config.yaml配置文件
	err := config.Load(
		file.NewSource(file.WithPath("./config.yaml")),
	)
	if err != nil {
		panic(err)
	}

	dbType := config.Get("db", "type").String("mysql")
	fmt.Printf("当前的dbType是 %s\n", dbType)

	dbCoreAddr := config.Get("db", "coreDB").String("none")
	fmt.Printf("当前的dbCoreAddr是 %s \n", dbCoreAddr)

	db, err := gorm.Open(dbType, dbCoreAddr)
	// db.SingularTable(true)

	// 开启自动填充
	db.AutoMigrate(&User{})

	defer db.Close()

	if err != nil {
		logger.Errorf("初始化core数据库失败 {%s} ", err)
	}

	// tableName := "users"

	// 创建数据库
	// CreateTable(db, tableName)

	// 插入user数据
	// InsertUser(db, tableName)

	// 更新user数据
	// UpdateUser(db, tableName)

	// 删除user数据
	// DeletedUser(db, tableName)

	// 查找所有的记录
	u := FindAll(db)
	// fmt.Println(u)
	for i, who := range u {
		// 获取每个人的名字
		fmt.Printf("第 %d 个人的名字是: %s \n", i, who.Name)
	}

}

// 查找所有的记录
func FindAll(db *gorm.DB) []User {
	var result []User

	// 降序排列
	// db.Where("user_id >= 0").Order("user_id desc").Find(&result)
	// 升序排列
	db.Where("user_id >= 0").Order("user_id asc").Find(&result)

	// fmt.Println(result)

	return result
}

// delete 名字为feng的记录
func DeletedUser(db *gorm.DB, tableName string) {
	// 删除名字为sun的用户
	db.Where("user_name = ?", "sun").Delete(&User{})
}

// 找到第一个并将其改名为 jinzhu
func UpdateUser(db *gorm.DB, tableName string) {
	t1, err := time.Parse("2006-01-02 15:04:05", "2020-06-13 09:14:00")
	if err != nil {
		panic(err)
	}

	u1 := User{ID: 1, Name: "feng", Birthday: t1}
	db.Model(&u1).Update("user_name", "sun")
}

// 创建一张叫users的表
func CreateTable(db *gorm.DB, tableName string) {
	// 使用User结构体创建名为 "users" 的表
	db.Table(tableName).CreateTable(&User{})
}

// 插入新的User数据
func InsertUser(db *gorm.DB, tableName string) {
	t1, err := time.Parse("2006-01-02 15:04:05", "2016-06-13 09:14:00")
	if err != nil {
		fmt.Printf("时间解析失败")
	}
	fmt.Println(t1)

	t2, err := time.Parse("2006-01-02 15:04:05", "2016-06-14 09:14:00")
	if err != nil {
		fmt.Printf("时间解析失败")
	}
	fmt.Println(t2)

	u1 := User{Name: "fei", Birthday: t1}
	fmt.Printf("%#v\n", u1)
	u2 := User{Name: "liu", Birthday: t2}
	fmt.Printf("%#v\n", u2)

	db.Create(&u1)
	db.Create(&u2)
}
