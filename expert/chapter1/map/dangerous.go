package main

import "time"

func multiRoutineWrite() {
	m := map[string]string{
		"name":   "Mr.Yan",
		"school": "清华",
	}
	go func() {
		for true {
			m["name"] = "Mr.Wang"
		}
	}()

	go func() {
		for true {
			m["school"] = "北京大学"
		}
	}()
	time.Sleep(time.Second * 10)
}
func multiRoutineRead() {
	m := map[string]string{
		"name":   "Mr.Yan",
		"school": "清华",
	}
	go func() {
		for true {
			println(m["name"])
		}
	}()

	go func() {
		for true {
			println(m["name"])
		}
	}()
	time.Sleep(time.Second * 10)
}
func multiRoutineReadWrite() {
	m := map[string]string{
		"name":   "Mr.Yan",
		"school": "清华",
	}
	go func() {
		for true {
			m["name"] = "Mr.Wang"
		}
	}()

	go func() {
		for true {
			println(m["school"])
		}
	}()
	time.Sleep(time.Second * 10)
}
func main() {
	//multiRoutineWrite()//多协程写会panic
	multiRoutineRead() //多协程读不会panic
	//multiRoutineReadWrite()//多线程读写会panic
}
