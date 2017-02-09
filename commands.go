package main

import "fmt"

//Public

//Cversion バージョンを表示する
func Cversion(v string) {
	fmt.Printf("This version is %s\n", v)
}

//Chelp ヘルプを表示
func Chelp(p []string) {
	switch len(p) {
	case 1:
		fmt.Println(helplist(""))
	case 2:
		if s := helplist(p[1]); s != "" {
			fmt.Println(helplist(p[1]))
		} else {
			error("help", "unknown command "+p[1], "Run 'run help'")
		}
	default:
		error("help", "many parameter", "try again 'run help'")
	}
}

//Chello Hello Worldと表示する
func Chello(p []string) {
	if len(p) == 0 {
		fmt.Println("Hello World")
		return
	}
	for i := range p {
		fmt.Printf("Hello %s\n", p[i])
	}
}

//Private

func helplist(t string) string {
	s := ""
	switch t {
	case "":
		s = `
Usage:
  run command [option] [...parameter]
Command:
  version:
    バージョンを表示する
  help:
    ヘルプを表示
  hello:
    Hello WorldまたはHello [...parameter]を出力
Options:
	`
	case "version":
		s = `
Usage:
  run version [option]
Option:
	`
	case "hello":
		s = `
Usage:
	run hello [option] [...parameter]
Option:
		`
	case "run":
		s = `HAIFIA`
	}
	return s
}
