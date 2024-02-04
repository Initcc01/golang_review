package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*
参考链接：https://blog.csdn.net/weixin_44211968/article/details/124632136
*/
func main() {

}

/*
4.2、读取多行字符串
param:多个测试用例，每个测试用例一行。每行通过空格隔开，有n个字符，n＜100
return:对于每组测试用例，输出一行排序过的字符串，每个字符串通过空格隔开
*/
func Bufio3() {
	inputs := bufio.NewScanner(os.Stdin)
	for inputs.Scan() {
		data := strings.Split(inputs.Text(), " ")
		sort.Strings(data)
		fmt.Println(strings.Join(data, " "))
	}
}

/*
4、读字符串
4.1、读一行字符串
param:输入有两行，第一行 n; 第二行是n个字符串，字符串之间用空格隔开
return:输出一行排序后的字符串，空格隔开，无结尾空格
*/
func Bufio2() {
	in := bufio.NewScanner(os.Stdin) //使用bufio创建新的扫描器in
	in.Scan()                        //从标准输入读取数据
	for in.Scan() {                  //使用scan方法从输入中扫描下一行
		str := in.Text() //并在使用text()获取当前文本字符串时返回该行内容
		s := strings.Split(str, " ")
		sort.Strings(s)                   //排序
		fmt.Println(strings.Join(s, " ")) //将切片连接成字符串
	}
}

/*
3.1、必须整行读（用fmt、os、bufio、strconv、strings实现）

	每行数字数量不固定，且不知道数量

param:输入数据有多组, 每行表示一组输入数据。每行不定有n个整数，空格隔开。(1 <= n <= 100)。
return:每组数据输出求和的结果

	这时候需要一整行一整行的读取，这时需要用到bufio包
*/
func Bufio1() {
	inputs := bufio.NewScanner(os.Stdin)
	for inputs.Scan() { //每次读入一行
		data := strings.Split(inputs.Text(), " ") //通过空格分开，并存入字符串切片
		var sum int
		for i := range data {
			val, _ := strconv.Atoi(data[i])
			sum += val
		}
		fmt.Println(sum)
	}
}

/*
2.3、每行的数字数量不固定，但知道数量，知道行数
*/
func FmtScan6() {
	var t int
	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		var num, sum, a int
		fmt.Scan(&num)
		for j := 0; j < t; j++ {
			fmt.Scan(&a)
			sum += a
		}
		fmt.Println(sum)
	}

}

/*
2.2、每行的数字数量不固定，但知道数量，不知道行数

	无结束标志（没得读了就结束）

param：输入数据有多组, 每行表示一组输入数据。

	每行的第一个整数为整数的个数n(1 <= n <= 100)。
	接下来n个正整数, 即需要求和的每个正整数。

return：每组数据输出求和的结果
*/
func FmtScan5() {
	var t, crr, sum int
	for {
		n, _ := fmt.Scan(&t)
		if n == 0 {
			break
		} else {
			sum = 0
			fmt.Scan(&crr)
			sum += crr
		}
		fmt.Println(sum)
	}
}

/*
2.1、每行的数字数量不固定，但知道数量，不知道行数

	有结束标志

param：输入数据包括多组。

	每组数据一行,每行的第一个整数为整数的个数n(1 <= n <= 100), n为0的时候结束输入。
	接下来n个正整数,即需要求和的每个正整数。

return：每组数据输出求和的结果
*/
func FmtScan3() {
	var t int
	for {
		var sum int

		fmt.Scanln(&t)
		if t == 0 {
			break
		} else {
			a := make([]int, t)
			for i := 0; i < t; i++ {
				fmt.Scanln(&a[i])
			}
			//求和
			for j := 0; j < t; j++ {
				sum += a[j]
			}
		}
		fmt.Println(sum)
	}
}
func FmtScan4() {
	var t, crr, sum int
	for {
		fmt.Scanln(&t)
		if t == 0 {
			break
		} else {
			for i := 0; i < t; i++ {
				fmt.Scanln(&crr)
				sum += crr
			}
			fmt.Println(sum)
		}
	}
}

/*
1.2、输入的每行数字数量固定，知道行数
param：输入第一行包括一个数据组数t

	接下来每行包括两个正整数a,b
*/
func FmtScan2() {
	var t, a, b int
	fmt.Scanln(&t)
	for i := 0; i < t; i++ {
		fmt.Scanln(&a, &b)
		fmt.Println(a + b)
	}
}

/*
1.1、输入的每行数字数量固定，不知道行数
param：输入包括两个正整数 a,b
return：n 是输入参数的数量
*/
func FmtScan1() {
	a := 0
	b := 0
	for {
		n, _ := fmt.Scanln(&a, &b)
		if n == 0 {
			break
		} else {
			fmt.Printf("a + b 的值为 %d\n", a+b)
		}
	}

}
