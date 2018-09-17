# GO 简易教程

## 基础

- go 声明变量不使用报错
- go 导入包不使用报错

### 变量和声明

```
//case1
var power int //声明变量
power = 9000  //赋值
//case2
var power int = 9000 //声明变量 + 赋值
//case3
power := 9000 //`:=` 自动推断变量类型,并赋值
//case4
power := 9000
power := 9001 // 编译器错误, 相同变量不能被声明两次（在相同作用域下）
//case5
power := 9000
power = 9001 // 这是对的
//case6
power := 1000
name, power := "Goku", 9000 //多个变量赋值的时候，只要其中有一个变量是新的，就可以使用`:=`
```

### 运行 go 代码

```
go run main.go //编译+执行
go run --work main.go //会打印编译路径
go build main.go // 编译
```

### 入口函数

每个 `Go` 程序都是由包组成的。

`go`程序入口必须是 `main` 包 , `main` 函数.

```
package main

func main() {
	println("hello world")
}
```

### 导入包

`import`可以导入第三方库

```
package main

// 不建议使用
import "fmt"
import "math"
// 推荐使用
// 打包的导入语句是更好的形式
import (
  "fmt"
  "os"
)

func main() {
  if len(os.Args) != 2 {
    os.Exit(1)
  }
  fmt.Println("It's over", os.Args[1])
}
```

> ⚠️ 在 go 中导入包却没有使用，是会编译报错的

### 函数声明

声明

```
// 无返回值函数
func log(message string) {
}
// 返回一个参数
func add(a int, b int) int {
}
// 返回多个参数
func power(name string) (int, bool) {
}
//如果参数具有相同的类型, 则除了最后一个类型之外，其他都可以省略。
func add(a, b int) int {

}
```

使用

```
value, exists := power("goku")
if exists == false {
  // 处理错误情况
}
```

如果只关注其中一个返回值,可以使用空白标识符`_`

```
_, exists := power("goku")
if exists == false {
  // handle this error case
}
```

## 简易教程

### 声明和初始化

Go 不是像 C ++，Java，Ruby 和 C＃一样的面向对象的（OO）语言。没有对象和继承的概念，没有多态和重载。

Go 所具有的是结构体的概念，可以将一些方法和结构体关联。

```
type Saiyan struct {
  Name string
  Power int
}

goku := Saiyan{
  Name: "Goku",
  Power: 9000,
}
```

> ⚠️ 上述结构末尾的逗号 `,` 是必需的。没有它的话，编译器就会报错。

不必设置所有或哪怕一个字段,就像未赋值的变量其值默认为 0 一样，字段也是如此。

```
goku := Saiyan{}

// or

goku := Saiyan{Name: "Goku"}
goku.Power = 9000
```

也可以省略字段名，但是为了可读性,建议写上

```
goku := Saiyan{"Goku", 9000}
```

## go 数据类型

基本类型，如：int、float、bool、string；
结构化的（复合的），如：struct、array、slice、map、channel；
描述类型的行为的，如：interface。·