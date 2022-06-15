package helpers

// MinDistance: Given two strings word1 and word2, the function return the minimum number of operations required to convert word1 to word2.
// Example 1:
// Input: word1 = "horse", word2 = "ros"
// Output: 3
// Explanation: 
// horse -> rorse (replace 'h' with 'r')
// rorse -> rose (remove 'r')
// rose -> ros (remove 'e')
func MinDistance(word1 string, word2 string) int {
	// create 2 dimensional slices
	cache := [][]int{}
	// fill each row with a slice of length word2 
	for i := 0; i < len(word1)+1; i++ {
		// Note: 1 was added to the length of word2 to handle base case
		cache = append(cache, make([]int, len(word2)+1))
	}
	// fill last row with integers in decreasing order for base case
	for j := 0; j < len(word2)+1; j++ {
		cache[len(word1)][j] = len(word2) - j
	}
   // fill last column with integers in decreasing order for base case  
	for i := 0; i < len(word1)+1; i++ {
		cache[i][len(word2)] = len(word1) - i
	}
    // loop through each row and column in the table in reverse order
	for i := len(word1) - 1; i >= 0; i-- {
		for j := len(word2) - 1; j >= 0; j-- {
			// if char at row i and col j are thesame then set cache[i][j] to value at diagonal cell i.e i+1 and j+1 
			if word1[i] == word2[j] {
				cache[i][j] = cache[i+1][j+1]
			} else {
				// if char at row i and col j are different then set cache[i][j] to the minimum value at direct neighbouring cells
				// Note: We added 1 here because we had to perform at least one operation i.e (insert, delete, replace) 
				cache[i][j] = 1 + MinOf(cache[i+1][j], cache[i][j+1], cache[i+1][j+1])
			}
		}
	}
	// return value at row 0 and col 0 since it contains the minimum edit performed to match the two strings
	return cache[0][0]
}
func MinOf(vars ...int) int {
	min := vars[0]
	for _, i := range vars {
		if min > i {
			min = i
		}
	}

	return min
}
