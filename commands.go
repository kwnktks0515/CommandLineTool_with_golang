package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

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

//Cbuild コマンドを実行する
func Cbuild(p []string) {
	for i := range p {
		for _, s := range searchfile(p[i]) {
			cmd := strings.Split(strings.Replace(s, "[name]", p[i], -1), " ")
			result, err := exec.Command(cmd[0], cmd[1:]...).Output()
			if err != nil {
				error("build", "Hwlloo")
				return
			}
			if 0 < len(result) {
				fmt.Print(string(result))
				continue
			}
		}
		// コマンドが見つからなかった場合
	}
}

//Private

func searchfile(s string) []string {
	file, err := os.Open(s)
	if err != nil {
		error("build", "Can`t Open file")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	result := []string{}
	for scanner.Scan() {
		text := scanner.Text()
		for i := 0; i < len(text); i++ {
			if text[i] == ' ' {
				continue
			}
			if text[i] == '/' && text[i+1] == '/' {
				for t := i + 2; t < len(text); t++ {
					if text[t] != ' ' {
						i = t
						break
					}
				}
				result = append(result, string(text[i:]))
				break
			}
			return result
		}
	}
	return nil
}

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
  build:
    設定ファイルまたはファイルの先頭からコマンドを取得し実行する
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
	case "build":
		s = `
Usage:
  run build [option] [...file]
Option:
		`
	}
	return s
}
