import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Manca il caldo iniziale")
	}
	saldo, _ := strconv.Atoi(os.Args[1])
	var spese int

	for saldo > 0 {
		fmt.Scan(&spese)

		saldo -= spese

		fmt.Println("Saldo: ", saldo)

	}
	fmt.Println("Saldo finale: ", saldo)
}
