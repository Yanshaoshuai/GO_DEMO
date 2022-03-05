package varible

var (
	//首字母大写包外可访问
	Astr string = "abc"
	b    int    = 16
)

func varible() {
	a := 1
	println(a)
	const name = "yanshaoshuai"
}
func GetSumAndPlus(a int, b int) (sum int, plus int) {
	//return a + b, a * b
	sum = a + b
	plus = a * b
	return sum, plus
}
