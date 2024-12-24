package handlers

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type CLIh struct {
	h *Handlers
}

func NewCLInterface(h *Handlers) *CLIh {
	return &CLIh{h: h}
}

func help() {
	fmt.Println("Опции:")
	fmt.Println("prodcl help")
	fmt.Println("prod help")
	fmt.Println("enumcl help")
	fmt.Println("unit help")
	fmt.Println("enumpos help")
	fmt.Println("param help")
	fmt.Println("paramcl help")
	fmt.Println("parampr help")
	fmt.Println("exit")
}

func (cli *CLIh) StartCli() {
	for true {
		fmt.Println("запущен интерфейс командной строки")
		fmt.Println("введите help для вывода списка команд")

		in := bufio.NewReader(os.Stdin)

		line, _ := in.ReadString('\n')
		command, args := proccessInput(line)

		if command == "exit" {
			fmt.Println("выход из программы")
			break
		} else if command == "help" {
			help()
		} else if command == "prodcl" {
			cli.h.ProductClass.ProductClassHandle(args)
		} else if command == "product" {
			cli.h.Product.ProductHandle(args)
		} else if command == "enumcl" {
			cli.h.EnumClassifier.Handle(args)
		} else if command == "unit" {
			cli.h.Unit.UnitHandle(args)
		} else if command == "enumpos" {
			cli.h.EnumPosition.EnumPositionHandle(args)
		} else if command == "param" {
			cli.h.Param.ParamHandle(args)
		} else if command == "paramcl" {
			cli.h.ParamClass.ParamClassHandle(args)
		} else if command == "parampr" {
			cli.h.ParamProduct.ParamProductHandle(args)
		}
	}
}

func proccessInput(input string) (command string, args []string) {
	s := strings.Split(input, " ")
	command = s[0]
	args = s[1:]
	return command, args
}
