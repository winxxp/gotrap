package gotrap

import (
	"fmt"
	"reflect"
)

// 问题1. f1()中defer调用顺序
// 问题2. f1()调用后的打印的结果
func f1() {
	i := 1
	defer func() {
		fmt.Println(i)
	}() // defer 1

	{
		i := 2
		defer func() {
			fmt.Println(i)
		}() // defer 2
	}

	defer func(i int) {
		fmt.Println(i)
	}(i) // defer 3

	func() {
		defer func() {
			fmt.Println(i)
		}() // defer 4
	}()

	i = 3
	defer func() {
		fmt.Println(i)
	}() // defer 5
}

// 问题3: 函数f2()调用后的返回值
func f2() (n int) {
	defer func() {
		n = 2
	}()

	return 3
}

// 问题4: 函数f3()调用后的返回值
func f3() {
	a := 1
	pa1 := &a

	a, b := 2, 3
	pa2 := &a

	fmt.Println("Q4:", pa1 == pa2, a, b)
}

// 问题5: 函数f4()中arr1与arr2是同一种类型吗？是什么类型
// 问题6: 函数f4()中arr3的len(arr3), cap(arr3)分别是多少
// 问题7: 函数f4()中Q7打印内容
// 问题8: 函数f4()中Q8打印内容
func f4() {
	var arr1 = [...]int{1, 2, 3}
	var arr2 = []int{1, 2, 3}

	var arr3 = arr1[:2:3]

	fmt.Println("Q5:", reflect.TypeOf(arr1), reflect.TypeOf(arr2))
	fmt.Println("Q6:", "len(arr3)=", len(arr3), "cap(arr3)=", cap(arr3))

	arr3[1] = 4
	fmt.Println("Q7:", arr1, arr3)

	arr3 = append(arr3, 5, 6, 7)
	arr3[1] = 9
	fmt.Println("Q8:", arr1, arr3)
}

// 问题9: Q9,Q10打印内容
func f5() {
	a := 1
	f := func() int {
		a++
		return a
	}
	a, b, c := f(), f(), a

	fmt.Println("Q9:", a, b, c)

	n0, n1 := 1, 2
	n0, n1 = n0+n1, n0

	fmt.Println("Q10:", n0, n1)
}

// 问题10: Q11, a,b,c,d的初始化顺序及f6打印内容
var (
	a = c + b
	b = f()
	c = f()
	d = 3
)

func f() int {
	d++
	return d
}

func f6() {
	fmt.Println("Q11:", a, b, c, d)
}
