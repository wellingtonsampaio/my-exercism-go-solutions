// Package pascal contains the implementation of the Exercism's go exercise 'Pascals Triangle'
package pascal

// testVersion is the current version of the test
const testVersion = 1

// Triangle compute Pascal's triangle up to the given number of rows.
func Triangle(rows int) (result [][]int) {
	for row := 0; row < rows; row++ {
		// Append a new row
		result = append(result, []int{})

		// All rows start with 1
		result[row] = append(result[row], 1)

		if row > 0 {
			// Calculate each value according to the sum of the (k-1) + k position of the previous row
			for pos := 1; pos < row; pos++ {
				result[row] = append(result[row], result[row-1][pos-1]+result[row-1][pos])
			}
			// All rows > 1 end with 1
			result[row] = append(result[row], 1)
		}
	}
	return
}
