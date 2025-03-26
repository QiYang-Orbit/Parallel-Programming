/*
This program demonstrates several key concepts in Go programming:
1. Goroutines (concurrent execution)
2. Time handling
3. Recursive functions
4. Format printing
5. Animation effects
6. Type system

The program calculates the nth Fibonacci number while showing a spinning animation.
This is a common pattern in long-running computations to indicate that the program is still working.
*/

package main

import (
	"fmt"    // fmt package provides formatted I/O functions
	"time"   // time package provides time-related functionality
)

/*
spinner function creates an infinite loading animation
Parameters:
    - delay: time.Duration - specifies how long to wait between each character change
             time.Duration is Go's built-in type for time intervals
             Example: 100 * time.Millisecond means 100 milliseconds

The function runs in an infinite loop and displays a sequence of characters
that appear to rotate, creating a loading animation effect.
*/
func spinner(delay time.Duration) {
	for {  // infinite loop
		// range over the string `-\|/` to cycle through each character
		// the backticks ` create a raw string literal, preserving the backslash
		for _, r := range `-\|/` {
			// \r moves cursor to the start of the line
			// %c formats the character
			fmt.Printf("\r%c", r)
			// Sleep for the specified duration
			time.Sleep(delay)
		}
	}
}

/*
fib function calculates the nth Fibonacci number recursively
Parameters:
    - x: int - the position in the Fibonacci sequence to calculate
Returns:
    - int: the Fibonacci number at position x

The Fibonacci sequence is: 0, 1, 1, 2, 3, 5, 8, 13, ...
Each number is the sum of the previous two numbers.
*/
func fib(x int) int {
	// Base cases: fib(0) = 0, fib(1) = 1
	if x < 2 {
		return x
	}
	// Recursive case: fib(n) = fib(n-1) + fib(n-2)
	return fib(x-1) + fib(x-2)
}

/*
main function is the entry point of the program
It demonstrates concurrent execution by running the spinner animation
in parallel with the Fibonacci calculation.
*/
func main() {
	// Start the spinner animation in a goroutine
	// The 'go' keyword launches the function in a separate goroutine
	// This allows the animation to run concurrently with the main calculation
	go spinner(100 * time.Millisecond)

	// Calculate the 45th Fibonacci number
	// const declares a constant value that cannot be changed
	const n = 45
	// Calculate the result
	fibN := fib(n)

	// Print the result
	// \r moves cursor to start of line to overwrite the spinner
	// %d is the format verb for integers
	// \n adds a newline at the end
	fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)
}

/*
Key Concepts Explained:

1. Goroutines:
   - Lightweight threads managed by Go runtime
   - Started with the 'go' keyword
   - Much lighter than OS threads
   - Communicate through channels (not shown in this example)

2. time.Duration:
   - Go's built-in type for time intervals
   - Provides type safety for time-related operations
   - Common units: time.Millisecond, time.Second, time.Minute
   - Example: 100 * time.Millisecond = 100ms

3. Format Printing:
   - fmt.Printf for formatted output
   - \r for carriage return (move cursor to start of line)
   - %c for character formatting
   - %d for integer formatting
   - \n for newline

4. Constants:
   - Declared with const keyword
   - Value must be known at compile time
   - Cannot be changed after declaration
   - Type is inferred from the value

5. Recursion:
   - Function calls itself
   - Must have base cases to prevent infinite recursion
   - Used here to calculate Fibonacci numbers
   - Note: This is not the most efficient way to calculate Fibonacci numbers

6. String Literals:
   - Backticks ` create raw string literals
   - Preserve backslashes and newlines
   - Used here for the spinner characters

7. Range:
   - Used to iterate over strings, slices, maps, etc.
   - Returns index and value for each element
   - _ is used to ignore the index in this case

8. Infinite Loops:
   - Created with for without condition
   - Must have a way to break out (not shown in this example)
   - Used here for continuous animation
*/


/*
该程序演示了 Go 编程中的几个关键概念：
1. 协程（并发执行）
2. 时间处理
3. 递归函数
4. 格式化打印
5. 动画效果
6. 类型系统

该程序在显示旋转动画的同时计算第 n 个斐波那契数。
这是长时间运行的计算中常见的模式，用于指示程序仍在工作。
*/

package main

import (
	"fmt"    // fmt 包提供格式化的 I/O 函数
	"time"   // time 包提供与时间相关的功能
)

/*
spinner 函数创建一个无限加载动画
参数：
    - delay: time.Duration - 指定每个字符更改之间等待的时间
             time.Duration 是 Go 内置的时间间隔类型
             示例：100 * time.Millisecond 表示 100 毫秒

该函数在无限循环中运行，并显示一系列字符
这些字符看起来在旋转，创建加载动画效果。
*/
func spinner(delay time.Duration) {
	for {  // 无限循环
		// 遍历字符串 `-\|/` 以循环显示每个字符
		// 反引号 ` 创建原始字符串字面量，保留反斜杠
		for _, r := range `-\|/` {
			// \r 将光标移到行首
			// %c 格式化字符
			fmt.Printf("\r%c", r)
			// 休眠指定的时间
			time.Sleep(delay)
		}
	}
}

/*
fib 函数递归计算第 n 个斐波那契数
参数：
    - x: int - 要计算的斐波那契序列中的位置
返回：
    - int: 位置 x 处的斐波那契数

斐波那契序列为：0, 1, 1, 2, 3, 5, 8, 13, ...
每个数字是前两个数字的和。
*/
func fib(x int) int {
	// 基本情况：fib(0) = 0, fib(1) = 1
	if x < 2 {
		return x
	}
	// 递归情况：fib(n) = fib(n-1) + fib(n-2)
	return fib(x-1) + fib(x-2)
}

/*
main 函数是程序的入口点
它通过并行运行旋转动画来演示并发执行。
*/
func main() {
	// 在协程中启动旋转动画
	// 'go' 关键字在单独的协程中启动函数
	// 这允许动画与主计算并发运行
	go spinner(100 * time.Millisecond)

	// 计算第 45 个斐波那契数
	// const 声明一个不能更改的常量值
	const n = 45
	// 计算结果
	fibN := fib(n)

	// 打印结果
	// \r 将光标移到行首以覆盖旋转动画
	// %d 是整数的格式化动词
	// \n 在末尾添加换行符
	fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)
}
/*
关键概念解释：

1. 协程（Goroutines）:
   - 由 Go 运行时管理的轻量级线程
   - 使用 'go' 关键字启动
   - 比操作系统线程轻量得多
   - 通过通道进行通信（此示例中未显示）

2. time.Duration:
   - Go 内置的时间间隔类型
   - 为时间相关操作提供类型安全
   - 常用单位：time.Millisecond, time.Second, time.Minute
   - 示例：100 * time.Millisecond = 100 毫秒

3. 格式化打印:
   - fmt.Printf 用于格式化输出
   - \r 表示回车（将光标移到行首）
   - %c 用于字符格式化
   - %d 用于整数格式化
   - \n 用于换行

4. 常量:
   - 使用 const 关键字声明
   - 值必须在编译时已知
   - 声明后不能更改
   - 类型由值推断

5. 递归:
   - 函数调用自身
   - 必须有基例以防止无限递归
   - 此处用于计算斐波那契数
   - 注意：这不是计算斐波那契数的最有效方法

6. 字符串字面量:
   - 反引号 ` 创建原始字符串字面量
   - 保留反斜杠和换行符
   - 此处用于旋转字符

7. Range:
   - 用于迭代字符串、切片、映射等
   - 返回每个元素的索引和值
   - 此处使用 _ 忽略索引

8. 无限循环:
   - 使用没有条件的 for 创建
   - 必须有退出方式（此示例中未显示）
   - 此处用于连续动画
*/

