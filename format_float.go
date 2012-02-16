package main
import(
	"fmt"
	"strconv"
)

func main() {
	var v float64
	v = 1234567890

	fmt.Println(strconv.FormatFloat(v, 'f', 0, 64))
	fmt.Println(strconv.FormatFloat(v, 'f', 2, 64))
	fmt.Println(strconv.FormatFloat(v, 'G', 2, 64))
	fmt.Println(strconv.FormatFloat(v, 'g', 2, 64))
	fmt.Println(strconv.FormatFloat(v, 'b', 2, 64))
	fmt.Println(strconv.FormatFloat(v, 'e', 2, 64))
	fmt.Println(strconv.FormatFloat(v, 'E', 2, 64))

}