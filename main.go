package main

func newString(arr []string) (answer string) {
	var words [][]rune
	var max int
	for _, value := range arr {
		r := []rune(value)
		words = append(words, r)
		if len(r) > max {
			max = len(r) // used for index range
		}
	}
	fmt.Println(max)
	v := make([][]bool, len(arr))
	for i := range v {
		v[i] = make([]bool, max+1)
	}

	// loop over max length of longest word
	for i, _ := range words {
		for j := 0; j < max; j++ {
			answer = recursion(words, v, answer, i, j)
		}
	}
	return answer
}

func recursion(arr [][]rune, v [][]bool, a string, i int, j int) string {
	if len(arr) <= i {
		return a
	}

	if len(arr[i]) <= j {
		return a
	}
	fmt.Printf("%v %v\n", i, j)
	if v[i][j] {
		return a
	}
	v[i][j] = true
	a += string(arr[i][j])
	a = recursion(arr, v, a, i+1, j)
	return a

}
