package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
)

func t1() {
	command := exec.Command("ls", "-al")
	//command.Run()
	out, _ := command.Output()
	fmt.Println(string(out))

}

func t2() {
	attr := new(os.ProcAttr)
	attr.Env = os.Environ()
	attr.Files = []*os.File{os.Stdin, os.Stdout, os.Stderr}
	process, err := os.StartProcess("/bin/pwd", []string{"pwd"}, attr)
	if err != nil {
		fmt.Printf("Error %v starting process!", err) //
		os.Exit(1)
	}
	fmt.Printf("The process id is %v", process)

	fmt.Printf(" pid:%d\n", process.Pid)
	state, err := process.Wait()
	if err != nil {
		fmt.Println("Failed to wait for process:", err)
		return
	}

	// 检查进程退出状态
	if !state.Success() {
		fmt.Println("Process exited with error")
		return
	}

	all, _ := io.ReadAll(os.Stdout)
	//
	fmt.Println("len:", len(all))

	fmt.Printf("输出结果：%v **over", all)

}
func main() {
	t2()
}
