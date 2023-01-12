package main

import "fmt"

/*
	解释器模式: 给定一个语言，定义它的文法的一种表示，并定义一个解释器，这个解释器使用该表示来解释语言中的句子。
	如果一种特定类型的问题发生的频率足够高，那么可能旧值得将该问题的各个实例表述为一个简单语言的句子。可以构建一个解释器，
	该解释器通过解释这些句子来解决该问题。

	比如正则表达式，解释器为正则表达式定义了一个文法，如何表示一个特定的正则表达式，以及如何解释这个正则表达式。

	通常当有一个语言需要解释执行，并且你可将语言中的句子表示为一个抽象语法树时，可使用解释器模式
*/

type Interpreter interface {
	Interpret(input string) (output string)
}

type TerminalInterpreter struct{}

func (t *TerminalInterpreter) Interpret(input string) (output string) {
	fmt.Printf("终端解释器解释表达式: %s\n", input)
	return input
}

func main() {
	var interpreter Interpreter = &TerminalInterpreter{}
	interpreter.Interpret("hello world!")
}
