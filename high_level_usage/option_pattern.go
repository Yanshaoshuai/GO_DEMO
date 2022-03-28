package main

type option struct {
	name string
}
type Option func(*option)

func wakeup() Option {
	return func(o *option) {
		println(o.name + " wakeup")
	}
}

func brushTeeth() Option {
	return func(o *option) {
		println(o.name + " brushTeeth")
	}
}

func eatBreakfast() Option {
	return func(o *option) {
		println(o.name + " eatBreakfast")
	}
}

func changeToFullName() Option {
	return func(o *option) {
		o.name = "yan shao shuai"
	}
}

func doMorningThing(opt *option, opts ...Option) {
	for _, f := range opts {
		f(opt)
	}
}
func main() {
	opt := option{"yan"}
	doMorningThing(&opt,
		wakeup(),
		changeToFullName(),
		brushTeeth(),
		eatBreakfast(),
	)
}
