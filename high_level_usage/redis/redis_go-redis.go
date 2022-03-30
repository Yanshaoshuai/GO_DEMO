package main

import (
	"github.com/go-redis/redis/v7"
	"log"
	"strconv"
	"time"
)

func init() {
	//主从模式(哨兵模式)
	/*db = redis.NewFailoverClient(&redis.FailoverOptions{
		MasterName:    "master",
		SentinelAddrs: []string{"192.168.1.2:6379", "192.168.1.3:6379"},
	})*/
	//集群模式
	/*db = redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: []string{"192.168.1.2:6379", "192.168.1.3:6379"},
	})*/

	//单机模式
	db = redis.NewClient(&redis.Options{
		Addr:         "127.0.0.1:6379",
		Password:     "",
		DB:           0,
		MaxRetries:   3,
		PoolSize:     10,
		MinIdleConns: 5,
	})
	if _, err := db.Ping().Result(); err != nil {
		log.Fatal(err)
	}
}

var db *redis.Client

func main() {
	defer db.Close()
	log.Println("String Test...")
	db.Set("STRING_CACHE", "test", time.Second*60)
	data, err := db.Get("STRING_CACHE").Result()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("data:", data)

	log.Println("Hash Test...")
	db.HSet("HASH_CACHE", "name", "zhangsan")
	db.HSet("HASH_CACHE", "age", "20")
	total, err := db.HLen("HASH_CACHE").Result()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("total:", total)
	values, err := db.HGetAll("HASH_CACHE").Result()
	if err != nil {
		log.Fatal(err)
	}
	for key, value := range values {
		log.Println(key, value)
	}

	//List
	log.Println("List Test...")
	db.LPush("LIST_CACHE", 1, 2, 3, 4, 5, 6)
	total, err = db.LLen("LIST_CACHE").Result()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("List total:", total)
	//先进先出
	log.Println("queue:")
	for i := 0; i < int(total); i++ {
		data, err := db.RPop("LIST_CACHE").Result()
		if err != nil {
			log.Fatal(err)
		}
		log.Println("data", data)
	}
	//先进后出
	db.LPush("LIST_CACHE", 1, 2, 3, 4, 5, 6)
	log.Println("stack:")
	for i := 0; i < int(total); i++ {
		data, err := db.LPop("LIST_CACHE").Result()
		if err != nil {
			log.Fatal(err)
		}
		log.Println("data", data)
	}
	//Set
	log.Println("Set Test...")
	db.SAdd("SET_CACHE", 1, 2, 3, 4, 5, 6)
	total, err = db.SCard("SET_CACHE").Result()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("total:", total)
	members, err := db.SMembers("SET_CACHE").Result()
	if err != nil {
		log.Println(err)
	}
	for _, member := range members {
		log.Println("member:", member)
	}

	log.Println("ORDERED SET: ")
	scoresOfStudents := []*redis.Z{
		{Score: 60.0, Member: "张三"},
		{Score: 75.0, Member: "小明"},
		{Score: 99.0, Member: "小红"},
		{Score: 82.0, Member: "小蓝"},
	}
	db.ZAdd("ORDERED_CACHE", scoresOfStudents...)
	//inf 无穷
	total, err = db.ZCount("ORDERED_CACHE", "-inf", "+inf").Result()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("total:", total)
	//由低到高排序
	log.Println("order by score asc")
	ascValues, err := db.ZRangeByScore("ORDERED_CACHE", &redis.ZRangeBy{
		Min: "70",
		Max: "100",
	}).Result()
	for _, value := range ascValues {
		log.Println("asc value:", value)
	}
	//由高到低排序
	descValues, err := db.ZRevRangeByScore("ORDERED_CACHE", &redis.ZRangeBy{
		Min: "70",
		Max: "100",
	}).Result()
	for _, value := range descValues {
		log.Println("desc value:", value)
	}

	//Pipeline
	//不是在事务中运行的
	log.Println("Pipeline:")
	var cmd *redis.IntCmd
	_, err = db.Pipelined(func(pipe redis.Pipeliner) error {
		cmd = pipe.Incr("INCR_KEY")
		pipe.Expire("INCR_KEY", time.Second*10)
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("pipe:", cmd.Val())
	//事务
	_, err = db.TxPipelined(func(pipe redis.Pipeliner) error {
		cmd = pipe.Incr("INCR_TX_KEY")
		pipe.Expire("INCR_TX_KEY", time.Second*10)
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("tx pipe:", cmd.Val())

	//发布/订阅
	//发布
	sub := db.Subscribe("chatroom")
	log.Println("public/subscribe")
	go func() {
		for {
			msg, err := sub.ReceiveMessage()
			if err != nil {
				log.Fatal(err)
			}
			if msg != nil {
				log.Printf("receive message:%s from %s", msg.Payload, msg.Channel)
			}
		}
	}()
	for i := 0; i < 10000; i++ {
		db.Publish("chatroom", "msg "+strconv.Itoa(i))
	}
}
