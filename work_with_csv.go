package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	in := bufio.NewReader(os.Stdin)
	x, _ := in.ReadString('\n')

	//x:="1 149,34643235;1 179,3344455666"

	split := strings.Split(x, ";")
	q := split[0]
	q1 := strings.ReplaceAll(q, " ", "")
	q2 := strings.ReplaceAll(q1, ",", ".")
	qFloat, err := strconv.ParseFloat(q2, 64)
	if err != nil {
		panic(err)
	}

	w := split[1]
	w1 := strings.ReplaceAll(w, " ", "")
	w2 := strings.ReplaceAll(w1, ",", ".")
	wFloat, err := strconv.ParseFloat(w2, 64)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%.4f", qFloat/wFloat)
}
