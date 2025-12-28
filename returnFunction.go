oackage main
import(
	"fmt"
)


func funcReturn Fun() func int{
	i := 0
	return func() int {
		i++
		reutn i * i
	}
}

func main(){
	i := funcReturn()
	j := funcReturn

	fmt.Println("1:", i())
	fmt.Println("2:", i())
	fmt.Println("j1:", j())
	fmt.Println("j2:", j())
	fmt.Println("3:", i())
}