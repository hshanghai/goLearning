package main

import "fmt"
import (
	"./node"
)



//结构
type comment struct {
	sex string
}
type person struct {
	//嵌入结构
	comment
	Name string
	Age  int
}
type teacher struct {
	//嵌入结构
	comment
	Name string
	Age  int
	sex  string
}

//扩展已有的包方法 , 第一：组合的方式， 第二：定义别名
type myTreeNode struct {
	node *node.TreeNode
}

func (myNode *myTreeNode) postOrder()  {
	if myNode == nil || myNode.node == nil{
		return
	}
	left :=  myTreeNode{myNode.node.Left}
	ring :=  myTreeNode{myNode.node.Right}
	left.postOrder()
	ring.postOrder()
	myNode.node.Print()
}

// 类型别名  绑定方法 method
type Mnt int

//接口

//反射reflection  引用包 reflect

func main() {

	//fmt.Println("hello,world")

	//控制语句
	// a1 := 2
	// if a1 > 3 {
	// 	fmt.Println("a1>1")
	// } else {
	// 	fmt.Println("a1<1")
	// }

	// if a2 := 1; a2 > 4 {
	// 	fmt.Println("a2>4")
	// } else {
	// 	fmt.Println("a2<4")
	// }

	//数组
	// fmt.Println("数组")
	// var a [5]int
	// var b = [2]int{2}
	// c := [3]int{3}
	// fmt.Println(a)
	// fmt.Println(b)
	// fmt.Println(c)

	//切片 slice
	// fmt.Println("//切片")
	// var s1 []int
	// fmt.Println(s1)
	// var s2 = make([]int, 2, 10)
	// fmt.Println(s2)
	// s3 := []int{1, 2, 3}
	// fmt.Println(s3)
	// arr1 := [3]int{1, 2, 3}
	// s4 := arr1[0:2]
	// s4_2 := arr1[1:]
	// fmt.Println(s4, s4_2)
	// s4[1] = 88
	// fmt.Println(s4, s4_2)
	// s5 := append(s4, 4, 5, 6, 7)
	// fmt.Println(s5)

	//map
	// var m map[int]string
	// m = map[int]string{}
	// fmt.Println(m)
	// //var n = make(map[int]string)
	// n := make(map[int]string)
	// n[1] = "ok"
	// fmt.Println(n)
	// an := n[1]
	// fmt.Println(an)
	// delete(n, 1)
	// an1 := n[1]
	// fmt.Println("an1:" + an1)
	// //多维map赋值钱需要make初始化
	// mm := map[int]map[int]string{}
	// fmt.Println("mm的值是：",mm)
	// aa, ok := mm[2][1]
	// fmt.Println(aa, ok)
	// if !ok {
	// 	mm[2] = make(map[int]string)
	// }
	// mm[2][1] = "good"
	// aa, ok = mm[2][1]
	// fmt.Println(aa, ok)
	// //两个map键值调换
	// m_a := map[int]string{1: "a", 2: "b", 3: "c"}
	// m_1 := make(map[string]int)
	// for k, v := range m_a {
	// 	m_1[v] = k
	// }
	// fmt.Println(m_a, m_1)

	//函数
	// a := 1
	// A(&a)
	// fmt.Println(a)
	// //变量函数类型
	// name1 := Name_fun
	// name1()
	// //匿名函数
	// name_no := func() {
	// 	fmt.Println("func name_no")
	// }
	// name_no()
	// //闭包
	// pak_fun := Add(10)
	// fmt.Println(pak_fun(1))
	// fmt.Println(pak_fun(2))
	 //defer  反顺序调用，即使函数出现严重错误也会执行，支持匿名函数调用
	 //defer fmt.Println("a")
	 //defer fmt.Println("b")
	 //for i := 0; i < 3; i++ {
	 //	defer fmt.Println(i) //i是参数，遵循值拷贝，所以一次输出0，1，2
	 //	defer func() {
	 //		fmt.Println(i) // i的地址，在循环执行完后i等于3，所以地址引用都是3
	 //	}() //后面加（）是直接调用  类似 a()
	 //}

	////结构
	//a := person{
	//	Name: "上海",
	//}
	////结构地址
	//b := &person{}
	//a.Age = 20
	//a.Name = "上海"
	//fmt.Println(a,b)
	//B(a)   //参数复制，函数里修改参数的值，外部调用没有修改
	//BB(&a) //地址参数，函数里修改地址参数的值，外部调用也是修改的值
	//BB(b)
	//fmt.Println(a)
	//匿名结构
	//aa := struct {
	//	Name string
	//	Age  int
	//}{
	//	Name: "海上",
	//	Age:  34,
	//}
	//fmt.Println(aa)
	//c := person{Name: "大海", Age: 90, comment: comment{sex: "男"}}
	//d := person{}
	//d.Name = "张三"
	//d.Age = 12
	//d.sex = "女"
	//fmt.Println(c)
	//fmt.Println(d)
	//
	//e := teacher{}
	//e.Age = 22
	//e.Name = "小花"
	//e.sex = "女"
	//
	//f := teacher{Name: "小妞", comment: comment{sex: "女"}}
	//fmt.Println(e)
	//fmt.Println(f)
	//

	////方法  method
	//var mmm Mnt
	//mmm.Increase(100)
	//mmm.Increase(100)
	//fmt.Println(mmm)
	//var nnn Mnt
	//fmt.Println(nnn)
	//var selfMoth Mnt
	//selfMoth.printlns("ssss")

	fmt.Println("树形结构")
	root := node.TreeNode{Value:1}
	root.Left = &node.TreeNode{Value:2}
	root.Right = &node.TreeNode{Value:3}
	root.Left.Left = &node.TreeNode{Value:4}
	//root.Left.right = &treeNode{Value:5}
	root.Left.Right = node.CreateNode(5)
	//root.right.left = &treeNode{Value:6}
	//没有提前创建结构
	root.Right.Left = new(node.TreeNode)
	root.Right.Left.SetValue(6)
	root.Right.Right = &node.TreeNode{Value:7}

	//root.Print()
	//root.Right.Left.Print()

	//go语音的nil 可以被调用
	//var pRoot *treeNode
	//pRoot.setValue(2)

	root.Traverse()
	fmt.Println()
	myRoot := myTreeNode{&root}
	myRoot.postOrder()
}

//函数 func  ，参数的int，string等类型传入函数等于值拷贝，而slice参数传入等于指针，函数内修改影响函数外部，如果int等类型也想达到类似的效果，
//就传入相应的内存地址 &a ，参数类型 a *int

func A(a *int) {
	*a = 2
	fmt.Println(*a)
}
func Name_fun() {
	fmt.Println(" func name_fun")
}
func Add(x int) func(y int) int {
	fmt.Println("%p\n", &x)
	return func(y int) int {
		return x + y
	}
}
func B(per person) {
	per.Age = 30
	fmt.Println("B", per)
}
func BB(per *person) {
	per.Age = 40
	fmt.Println(per)
}

//方法 method
func (mm *Mnt) Increase(num int) {
	*mm += Mnt(num)
	//fmt.Println(mmm)
}

func (mnt *Mnt) printlns(str string){
	fmt.Println(str)
}
