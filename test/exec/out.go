package main
import (
	"os/exec"
	"fmt"
)
func main() {
	cmd := exec.Command("ls")//查看当前目录下文件
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(out))
}