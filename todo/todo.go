package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)



func main(){

	fmt.Println("enter command")
	scanner := bufio.NewReader(os.Stdin)

	tasks:=[] string {}
	for {
		
		text,_:=scanner.ReadString('\n')
		text=strings.TrimSpace(text)

		if text == "add"{
			var task string
			fmt.Println("task")
			task,_=scanner.ReadString('\n')
			task=strings.TrimSpace(task)
			
			tasks = append(tasks, task)

		}else if text =="list"{
			for i:=0;i<len(tasks);i++{
				fmt.Printf("%d %s\n",i+1,tasks[i])
			}
		}else if text == "delete"{

			if len(tasks)==0{
				fmt.Println("No tasks")
			}else{
				for i:=0;i<len(tasks);i++{
				fmt.Printf("%d %s\n",i+1,tasks[i])
			}
			var tasknum int 
			fmt.Scanf("%d",&tasknum)

			deletedTask := tasks[tasknum-1]
			tasks = append(tasks[:tasknum-1], tasks[tasknum:]...)
			fmt.Printf("deleted task %s",deletedTask)

			}


		}else if text == "quit"{
			fmt.Println("good bye")
			os.Exit(0)
		}

	}
}