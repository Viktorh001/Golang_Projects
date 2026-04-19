package main

import (
	"fmt"
	"strings"
)

func main() {

	for {
		var name string
		fmt.Println("what is your name?")
		fmt.Scanln(&name)

		if name == "stop" {
			fmt.Println("You stopped the program!")
			break
		}
		if strings.ContainsAny(name, ",b,c,d") {
			fmt.Println("You will be a millionaire in four years")
			continue
		} else if strings.ContainsAny(name, "e,f,g") {
			fmt.Println("You are going to be a millionaire is ten years!!")
			continue
		} else if strings.ContainsAny(name, "h,i,j") {
			fmt.Println("You have a bright future but you should workhard to get it!")
			continue
		} else if strings.ContainsAny(name, "k,l,m") {
			fmt.Println("You are smart but you dont think so!")
			continue
		} else if strings.ContainsAny(name, "n,o,p,q") {
			fmt.Println("You are going to be a millionaire in two years!")
			continue
		} else if strings.ContainsAny(name, "r,s,t") {
			fmt.Println("You are to be a millionaire in six weeks")
			continue
		} else if strings.ContainsAny(name, "u,v,w,a") {
			fmt.Println("You should probably be a millionaire!")
			continue
		} else if strings.ContainsAny(name, "x,y,z") {
			fmt.Println("You will be a millionaire if twenty years!")
			continue
		} else {
			fmt.Println("Please enter your name")
			continue
		}
	}
}
