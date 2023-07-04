package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	s, _ := reader.ReadString('\n')
	s = strings.TrimSpace(s)
	ss := strings.Split(s, " ")
	s1 := ss[0]
	fmt.Printf("%s ", s1)
	for i := 1; i < len(ss); i++ {
		if !strings.EqualFold(s1, ss[i]) {
			fmt.Printf("%s ", ss[i])
		}
		s1 = ss[i]
	}
}
