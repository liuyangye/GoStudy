package main

import "fmt"

func main() {
	var x1 float64
	var y1 float64
	var x2 float64
	var y2 float64
	var k1 float64
	fmt.Print("请输入x1:")
	fmt.Scanln(&x1)
	fmt.Print("请输入y1:")
	fmt.Scanln(&y1)
	fmt.Print("请输入x2:")
	fmt.Scanln(&x2)
	fmt.Print("请输入y2:")
	fmt.Scanln(&y2)
	if x1 == x2 {
		fmt.Print("第一条直线斜率不存在。")
	} else if x1 != x2 {
		k1 = (y2 - y1) / (x2 - x1)
		fmt.Printf("第一条直线斜率为：%f\n", k1)
	}

	var x3 float64
	var y3 float64
	var x4 float64
	var y4 float64
	var k2 float64
	fmt.Print("请输入x3:")
	fmt.Scanln(&x3)
	fmt.Print("请输入y3:")
	fmt.Scanln(&y3)
	fmt.Print("请输入x4:")
	fmt.Scanln(&x4)
	fmt.Print("请输入y4:")
	fmt.Scanln(&y4)
	if x3 == x4 {
		fmt.Print("第二条直线斜率不存在。")
	} else if x3 != x4 {
		k2 = (y4 - y3) / (x4 - x3)
		fmt.Printf("第二条直线斜率为：%f\n", k2)
	}

	if x1 == x2 && x3 == x4 {
		fmt.Print("两条直线的斜率都不存在，此种情况判定两条直线平行。")
	} else if k1 == k2 {
		fmt.Print("两条直线的斜率均存在且相等，此种情况判定两条直线平行。")
	} else if k1 != k2 {
		fmt.Print("两条直线的斜率均存在，但不相等，此种情况判定两条直线不平行。")
	}
}
