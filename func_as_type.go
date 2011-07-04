package main
import(
	"fmt"
)

type IO func(input string) (output string)

var printer = func (my1o IO) {
	fmt.Println(my1o("input"));
}

func main() {
	printer(func(i string) string{
		return i + " outputed"
	})
}
