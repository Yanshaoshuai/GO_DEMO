package main

type Animal interface {
	sleep()
	eat()
}

//Dog 实现接口时 既有值接收者又有指针接收者
type Dog struct {
}

func (d *Dog) eat() {
	//TODO implement me
	panic("implement me")
}

func (d Dog) sleep() {
	//TODO implement me
}

//Cat 实现接口时只有值接收者
type Cat struct {
}

func (c Cat) eat() {
	//TODO implement me
	panic("implement me")
}

func (c Cat) sleep() {
	//TODO implement me
	panic("implement me")
}

//Tiger 实现接口时只有指针接收者
type Tiger struct {
}

func (t *Tiger) sleep() {
	//TODO implement me
	panic("implement me")
}

func (t *Tiger) eat() {
	//TODO implement me
	panic("implement me")
}

// 不能把nil转成类型
// cannot convert nil to type Dog
//var _ Animal = (Dog)(nil)
//cannot convert nil to type Cat
//var _ Animal = (Cat)(nil)

//可以把nil转成接口
//保证类型实现了接口
var _ Animal = (*Dog)(nil)
var _ Animal = (*Cat)(nil)

func main() {
	//不管有没有用指针接收者实现接口 指针类型都隐式是实现了接口
	//必须显式用值接收者实现所有接口 值类型才算实现了接口

	//既有值接收者又有指针接收者 只有指针类型实现了接口
	var dogPointer Animal = &Dog{}
	//var _ Dog = dogPointer.(Dog)
	var _ *Dog = dogPointer.(*Dog)

	//只有值接收者 指针类型和值类型都实现了接口
	var catPointer Animal = &Cat{}
	//运行时指针赋给值类型报错
	//var _ Cat = catPointer.(Cat)
	var _ Cat = *catPointer.(*Cat)
	var _ *Cat = catPointer.(*Cat)

	var cat Animal = Cat{}
	var catValue Cat = cat.(Cat)
	var _ *Cat = &catValue
	//var _ *Cat = &(cat.(Cat))

	//只有指针接收者 只有指针类型实现了接口
	var tigerPointer Animal = &Tiger{}
	//var _ Tiger = tigerPointer.(Tiger)
	var _ *Tiger = tigerPointer.(*Tiger)

	var animal Animal = &Dog{}

	//判断是否实现某方法
	if dog, ok := animal.(interface{ sleep() }); ok {
		dog.sleep()
	}

	//nil ==>指针 channel 函数 接口 map slice 的零值
	// nil is a predeclared identifier representing the zero value for a
	// pointer, channel, func, interface, map, or slice type.
	// Type must be a pointer, channel, func, interface, map, or slice type

	//interface{} and nil
	var aAnimal Animal
	if aAnimal == nil {
		println("aAnimal is nil")
	}
	var animalInterface interface{} = aAnimal
	if animalInterface == nil { //true 接口的零值是nil
		println("animalInterface is nil")
	}
	var aDog *Dog
	if aDog == nil {
		println("aDog is nil")
	}
	var aDogInterface interface{} = aDog
	if aDogInterface == nil { //false aDogInterface存有类型信息 不是零值
		println("aDogInterface is nil")
	}
}
