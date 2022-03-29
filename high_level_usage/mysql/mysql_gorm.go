package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"
	"time"
)

type User struct {
	ID      int
	Name    string
	Deleted bool
}
type UserHasPet struct {
	ID   int
	Name string
	Pets []Pet
}

type Pet struct {
	Name string
}

func dbConnect(user, pass, addr, dbName string) (*gorm.DB, error) {
	//root:123456@tcp(127.0.0.1:3306)/go_gin_api?charset=utf8mb4&parseTime=true&loc=Local
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=%t&loc=%s",
		user,
		pass,
		addr,
		dbName,
		true,
		"Local")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		//Logger: logger.Default.LogMode(logger.Info), // 日志配置
	})

	if err != nil {
		return nil, err
	}

	db.Set("gorm:table_options", "CHARSET=utf8mb4")

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	// 设置连接池 用于设置最大打开的连接数，默认值为0表示不限制.设置最大的连接数，可以避免并发太高导致连接mysql出现too many connections的错误。
	sqlDB.SetMaxOpenConns(10)

	// 设置最大连接数 用于设置闲置的连接数.设置闲置的连接数则当开启的一个连接使用完成后可以放在池里等候下一次使用。
	sqlDB.SetMaxIdleConns(60)

	// 设置最大连接超时
	sqlDB.SetConnMaxLifetime(time.Minute * 60)

	// 使用插件

	return db, nil
}

func main() {
	//connect
	db, err := dbConnect("root", "123456", "127.0.0.1:3306", "gorm")
	if err != nil {
		panic(err)
	}
	//insert(db)
	update(db)
}

func insert(db *gorm.DB) {
	var user = User{Name: "Yan shao shuai"}
	//insert one
	result := db.Create(&user)
	err := result.Error
	if err != nil {
		panic(err)
	}
	fmt.Printf("%d user %v insert success \n", result.RowsAffected, user)
	//insert multi
	var users = []User{{Name: "xiyangyang"}, {Name: "lanyangyang"}, {Name: "feiyangyang"}}
	result = db.Create(users)
	if err = result.Error; err != nil {
		panic(err)
	}
	fmt.Printf("%d user %v insert success \n", result.RowsAffected, users)
	//使用map创建
	var userMap = map[string]interface{}{
		"Name": "meiyangyang",
	}
	result = db.Model(&User{}).Create(&userMap)
	if err = result.Error; err != nil {
		panic(err)
	}
	fmt.Printf("%d user %v insert success \n", result.RowsAffected, userMap)

	var userMapSlice = []map[string]interface{}{
		{"Name": "huitailang"},
		{"Name": "hongtailang"},
	}
	result = db.Table("user").Create(userMapSlice)
	if err = result.Error; err != nil {
		panic(err)
	}
	fmt.Printf("%d user %v insert success \n", result.RowsAffected, userMapSlice)

	//insert batch
	//way 1
	//var largeUserSlice = []User{{Name: "yan1"}, ...,{Name: "yan1000"}}
	////每次插入一千条
	//db.CreateInBatches(&largeUserSlice, 1000)
	////way 2
	//thisDb, err := gorm.Open(mysql.Open(
	//	"root:123456@tcp(127.0.0.1:3306)/go_gin_api?charset=utf8mb4&parseTime=true&loc=Local",
	//), &gorm.Config{
	//	CreateBatchSize: 1000,
	//})
	//thisDb.Create(&largeUserSlice)
	////way 3
	//db = db.Session(&gorm.Session{CreateBatchSize: 1000})
	//db.Create(&largeUserSlice)

	// 携带关联
	//var userHasPets=[]UserHasPet{{Name:"xiaoming",Pets: []Pet{pet1,pet2,pet3}},...}//10000条
	//方法一插入 分成10批 每批1000条 每批的3000条pets使用同一条SQL插入
	//方法2 3 分成10批 每批1000条 每批的3000条pets分成三批创建
	//分批插入推荐方法2,3

	//错误发生时忽略 批量插入时很有用
	ignoreErrDB := db.Clauses(clause.OnConflict{DoNothing: true})
	result = ignoreErrDB.Create(&User{
		ID: 1,
	})
	if err = result.Error; err != nil {
		panic(err)
	}
	// Mysql特有忽略错误的方式
	result = db.Clauses(clause.Insert{Modifier: "IGNORE"}).Create(&User{
		ID: 1,
	})
	if err = result.Error; err != nil {
		panic(err)
	}

	//数据冲突时 更新某些字段
	db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoUpdates: clause.Assignments(map[string]interface{}{"deleted": nil}),
	}).Create(&User{ID: 1})

	//冲突时更新多个字段值
	db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoUpdates: clause.AssignmentColumns([]string{"name", "deleted"}),
	}).Create(&User{
		ID:      1,
		Name:    "admin",
		Deleted: false,
	})
	//冲突时除了主键外全部更新为新值
	db.Clauses(clause.OnConflict{UpdateAll: true}).Create(&User{
		ID:      1,
		Name:    "cunzhang",
		Deleted: false,
	})
}

func update(db *gorm.DB) {
	// UPDATE `user` SET `name`='Yanshaoshuai' WHERE `id` = 1 不会带上零值
	db.Model(&User{ID: 1}).Updates(User{Name: "Yanshaoshuai", Deleted: false})
	//带上零值
	db.Model(&User{ID: 1}).Select("*").Updates(User{Name: "Yanshaoshuai", Deleted: false})
}
