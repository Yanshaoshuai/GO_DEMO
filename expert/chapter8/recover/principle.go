package main

//recover 对应的runtime函数
//func gorecover(argp unintptr) interface{}{
//	gp:=getg()
//	p:=gp._panic
//	if p!=nil&&!p.goexit&&!p.recovered&&argp==uintptr(p.argp){//uintptr(p.argp) recover()必须被defer 直接作用
//		p.recovered=true
//		return p.arg
//	}
//	return nil
//}
