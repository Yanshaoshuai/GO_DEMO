package main

var FruitColor map[string]string

func AddFruit(name, color string) {
	if FruitColor != nil {
		FruitColor[name] = color //add nil map will be panic
	}
}

var StudentScore map[string]int

func GetScore(name string) int {
	score := StudentScore[name] //键不存在返回值的零值
	return score
}
func main() {
	AddFruit("Apple", "Red")
	println(GetScore("Mr.Yan"))
}
