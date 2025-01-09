package hackerank

import (
	"testing"
)

const MOD int64 = 1000000007

type LegoSolver struct {
	rowMemo   map[int]int64    // Memoization for row configurations
	totalMemo map[int]int64    // Memoization for total walls
	solidMemo map[[2]int]int64 // Memoization for solid walls
}

// NewLegoSolver initializes a new solver instance with empty memoization maps
func LegoBlocksSolver() *LegoSolver {
	return &LegoSolver{
		rowMemo:   make(map[int]int64),
		totalMemo: make(map[int]int64),
		solidMemo: make(map[[2]int]int64),
	}
}

// CalculateRowConfigs calculates the number of ways to fill a single row of width m
func (solver *LegoSolver) CalculateRowConfigs(m int) int64 {
	if m == 0 {
		return 1
	}
	if m < 0 {
		return 0
	}
	if val, found := solver.rowMemo[m]; found {
		return val
	}

	result := (solver.CalculateRowConfigs(m-1) +
		solver.CalculateRowConfigs(m-2) +
		solver.CalculateRowConfigs(m-3) +
		solver.CalculateRowConfigs(m-4)) % MOD
	solver.rowMemo[m] = result
	return result
}

// TotalWalls calculates the total possible walls of height n and width m
func (solver *LegoSolver) TotalWalls(n, m int) int64 {
	if val, found := solver.totalMemo[m]; found {
		return val
	}

	rowConfig := solver.CalculateRowConfigs(m)
	result := int64(1)
	for i := 0; i < n; i++ {
		result = (result * rowConfig) % MOD
	}
	solver.totalMemo[m] = result
	return result
}

// SolidWalls calculates the number of solid walls of height n and width m
func (solver *LegoSolver) SolidWalls(n, m int) int64 {
	if val, found := solver.solidMemo[[2]int{n, m}]; found {
		return val
	}

	total := solver.TotalWalls(n, m)
	invalid := int64(0)

	for j := 1; j < m; j++ {
		invalid = (invalid + (solver.SolidWalls(n, j)*solver.TotalWalls(n, m-j))%MOD) % MOD
	}

	solid := (total - invalid + MOD) % MOD
	solver.solidMemo[[2]int{n, m}] = solid
	return solid
}

func TestLegoBlocks(t *testing.T) {
	testCases := []struct {
		n, m     int
		expected int64
	}{
		{2, 2, 3},
		{3, 2, 7},
		{2, 3, 9},
		{4, 4, 3375},
	}

	for _, tc := range testCases {
		solver := LegoBlocksSolver()
		result := solver.SolidWalls(tc.n, tc.m)
		if result != tc.expected {
			t.Errorf("SolidWalls(%d, %d) = %d; expected %d", tc.n, tc.m, result, tc.expected)
		}
	}
}

func BenchmarkSolidWalls(b *testing.B) {
	for i := 0; i < b.N; i++ {
		solver := LegoBlocksSolver()
		solver.SolidWalls(4, 4)
	}
}
