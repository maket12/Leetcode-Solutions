type NumMatrix struct {
    PrefixSums [][]int
}

func Constructor(matrix [][]int) NumMatrix {
    if len(matrix) == 0 || len(matrix[0]) == 0 {
        return NumMatrix{PrefixSums: [][]int{}}
    }

    m, n := len(matrix), len(matrix[0])
    
    prefixSums := make([][]int, m + 1)
    for i := range prefixSums {
        prefixSums[i] = make([]int, n + 1)
    }

    for r := 0; r < m; r++ {
        for c := 0; c < n; c++ {
            prefixSums[r+1][c+1] = matrix[r][c] + 
                prefixSums[r][c+1] + 
                prefixSums[r+1][c] - 
                prefixSums[r][c]
        }
    }

    return NumMatrix{PrefixSums: prefixSums}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	row1, col1, row2, col2 = row1 + 1, col1 + 1, row2 + 1, col2 + 1
    return this.PrefixSums[row2][col2] - 
           this.PrefixSums[row1 - 1][col2] - 
           this.PrefixSums[row2][col1 - 1] + 
           this.PrefixSums[row1 - 1][col1 - 1]
}