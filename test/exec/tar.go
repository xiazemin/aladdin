package main
import (
	"os/exec"
	"fmt"
	"strings"
	"bytes"
	"log"
)
func main() {
	//func Command(name string, arg ...string) *Cmd
	cmd := exec.Command("tr", "a-z", "A-Z")
	cmd.Stdin = strings.NewReader("some input")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("in all caps: %q\n", out.String())//in all caps: "SOME INPUT"
}