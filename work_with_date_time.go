// На стандартный ввод подаются данные о продолжительности периода (формат приведен в примере). 
// Кроме того, вам дана дата в формате Unix-Time: 1589570165 в виде 
// константы типа int64 (наносекунды для целей преобразования равны 0).

// Требуется считать данные о продолжительности периода, преобразовать их 
// в тип Duration, а затем вывести (в формате UnixDate) дату и время, 
// получившиеся при добавлении периода к стандартной дате.


package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

const now int64 = 1589570165

func main() {
	in := bufio.NewReader(os.Stdin)
	input, err := in.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(err)
	}
	split := strings.Split(input, " ")
	minutes, seconds := split[0], split[2]
	minInt, err := strconv.Atoi(minutes)
	if err != nil {
		panic(err)
	}
	secInt, err := strconv.Atoi(seconds)
	if err != nil {
		panic(err)
	}
	secs := int64(secInt + (minInt * 60))
	resTime := time.Unix(secs+now, 0)
	need := resTime.In(time.UTC)

	res := need.Format(time.UnixDate)
	fmt.Println(res)

}
