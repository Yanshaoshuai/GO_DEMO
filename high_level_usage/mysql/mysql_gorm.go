package main

import (
	"errors"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"time"
)

type User struct {
	ID      int
	Name    string
	Age     int
	Deleted bool
}
type UserHasPet struct {
	ID   int
	Name string
	Pets []Pet
}

type Pet struct {
	ID   string
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
			//表前缀
			TablePrefix:   "",
			SingularTable: true,
		},
		//Logger: logger.Default.LogMode(logger.Info), // 日志配置
		Logger: logger.Default.LogMode(logger.Info),
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
	//update(db)
	//delete(db)
	query(db)
	//subQuery(db)
	//joinSelect(db)
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
	//NOTE 默认 禁止无条件删除/更新
	db.Model(&User{}).Update("name", "YanShaoShuai")
	//临时允许无条件删除/更新
	//db.Session(&gorm.Session{AllowGlobalUpdate: true}).Model(&User{}).Update("name","Yanshaoshuai")

	//select for update
	var users []User
	sql := db.Clauses(clause.Locking{Strength: "UPDATE"}).Find(&users).Statement.SQL.String()
	log.Println(sql)
}

func delete(db *gorm.DB) {
	//永久删除
	db.Unscoped().Where("name='Yanshaoshuai'").Delete(&User{})
}

func query(db *gorm.DB) {
	//查询按照ID升序的第一条记录
	var firstUser User
	db.First(&firstUser)
	log.Printf("%v\n", firstUser)

	//取一条数据 没有排序
	var takeUser User
	db.Take(&takeUser)
	log.Printf("%v\n", takeUser)

	//按照ID升序获取最后一条数据
	var lastUser User
	db.Last(&lastUser)
	log.Printf("%v\n", lastUser)

	var users []User
	result := db.Find(&users, []int{1, 2, 3})
	//判读是否没有找到
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		log.Fatal(gorm.ErrRecordNotFound.Error())
	}
	log.Printf("%v\n", users)
	var user User
	db.Where("name=?", "huitailang").Find(&user)
	log.Printf("%v\n", user)
	user = User{}
	db.Where("name=?", "huitailang").Where("deleted=?", false).Find(&user)
	log.Printf("%v\n", user)
	db.Where("name in ?", []string{"huitailang", "hongtailang"}).Find(&users)
	log.Printf("%v\n", users)
	db.Where("(name,id) in ?", [][]interface{}{{"huitailang", 25}, {"hongtailang", 26}}).Find(&users)
	log.Printf("%v\n", users)
	db.Where([]int{1, 2, 3}).Find(&user)
	log.Printf("%v\n", &user)
	users = []User{}
	//struct查询 忽略零值
	db.Where(&User{Name: "huitailang", Deleted: false}).Find(&users)
	log.Printf("%v\n", users)
	users = []User{}
	db.Where(map[string]interface{}{"name": "huitailang", "deleted": false}).Find(&users)
	log.Printf("%v\n", users)
	//带上零值
	//1用where map方式 如上
	//方法二指定字段
	users = []User{}
	db.Where(&User{Name: "huitailang", Deleted: false}, "name", "deleted").Find(&users)
	log.Printf("%v\n", users)
	//方法三 用string
	users = []User{}
	db.Where("name=? and deleted=?", "huitailang", false).Find(&users)
	log.Printf("%v\n", users)

	//单个字段 结果转成数组
	var names []string
	db.Model(&User{}).Pluck("name", &names)
	log.Printf("%v\n", names)

	//分批处理
	//db.Model(&User{}).FindInBatches(&users, 10, func(tx *gorm.DB, batch int) error {
	//	log.Printf("第%d批次 :%v\n", batch, users)
	//	return tx.Error
	//})
	//命名参数
	user = User{}
	result = db.Session(&gorm.Session{DryRun: true}).Model(&User{}).
		Where("name=@name", map[string]interface{}{"name": "huitailang"}).Find(&user)
	log.Printf("sql:%s", result.Statement.SQL.String())
	log.Printf("%v\n", user)
}

func subQuery(db *gorm.DB) {
	var users []User
	sub := db.Table("user").Select("AVG(id)")
	db.Where("id>(?)", sub).Find(&users)
	log.Printf("%v\n", users)

	var avgAge []float64
	sub = db.Select("Avg(id)").Where("name like ?", "hui%").Table("user")
	db.Model(&User{}).Select("Avg(age) as avgage").Group("name").Having("Avg(age)<(?)", sub).Pluck("avgage", &avgAge)
	log.Printf("%v\n", avgAge)

	//使用gorm.DB作为Table参数
	sub = db.Model(&User{}).Select("name", "age")
	var user User
	db.Table("(?) as u", sub).Where("age<?", 18).Find(&user)
	log.Printf("%v\n", user)

	users = []User{}
	sub1 := db.Model(&User{}).Select("name")
	sub2 := db.Model(&Pet{}).Select("name")
	db.Raw("? UNION ?", sub1, sub2).Scan(&users)

	statement := db.Where(
		db.Where("pizza=?", "pepperoni").Where(
			db.Where("size=?", "small").Or("size=?", "medium"),
		),
	).Or(
		db.Where("pizza=?", "hawaiian").Where("size=?", "xlarge"),
	).Find(&Pizza{}).Statement
	log.Printf("%v", statement.SQL.String())

}

func joinSelect(db *gorm.DB) {
	joins := db.Session(&gorm.Session{DryRun: true}).
		Joins("left join company on company.user_id=user.id").
		Joins("left join manager on manager.user_id=user.id").Find(&User{})
	log.Printf("%v\n", joins.Statement.SQL.String())
}

type Pizza struct {
}
