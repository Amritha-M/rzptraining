package main
import "fmt"

type Matrix struct {
	rows     int
	columns  int
	elements [][]int
}

func (a Matrix) noofrows() int
{
	return a.rows
}

func (a Matrix) noofcols() int
{
	return a.columns
}

func (a *Matrix) setelements(b,i,j int)
{
	a.elements[i][j] = a
}

func (a Matrix) add(c Matrix) [][]int
{
	for i := 0; i < a.rows; i++
	{
		for j := 0; j < a.columns; j++
			a.elements[i][j] +=c.elements[i][j]
	}
	return a.elements[][]
}

func (a Matrix) print() {
		for i := 0; i < a.rows; i++ {
			fmt.Println(a.matrix[i])
		}

	}

	func main() {
		a := Matrix{3, 3, [][]int{{1, 2, 3}{4, 5, 6}{7, 8, 9}}}
		b := Matrix{3, 3, [][]int{{1, 2, 3}{4, 5, 6}{7, 8, 9}}}
		a.print()
		b.print()
		fmt.Println(a.get_rows())
		fmt.Println(a.get_columns())
		fmt.Println(a.add(b))
		a.set_entry(3, 7, 9)
		b.print()
	}

	// example
	