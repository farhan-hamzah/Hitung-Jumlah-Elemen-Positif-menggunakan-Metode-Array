package main
import "fmt"

func main(){
	const NMAX int = 100
	var A[NMAX]int
	var n, i, total int

	fmt.Scan(&n)
	for i = 0; i < n; i++{
		fmt.Scan(&A[i])
	}

	for i = 0; i < n; i++{
		if A[i] > 0{
			total++
		}
	}
	fmt.Print(total)
}