func RandomGoProgram(n int) string {
	// Create a new slice to store the generated code
	var generatedCode []string

	// Generate n random lines of code
	for i := 0; i < n; i++ {
		// Select a random line type (function, variable, or statement)
		lineType := rand.Intn(3)

		// Generate a random line of code based on the selected type
		switch lineType {
		case 0: // Function
			generatedCode = append(generatedCode, fmt.Sprintf("func %s() {\n", randString(10)))
		case 1: // Variable
			generatedCode = append(generatedCode, fmt.Sprintf("%s := %s\n", randString(10), randInt()))
		case 2: // Statement
			generatedCode = append(generatedCode, fmt.Sprintf("%s()\n", randString(10)))
		}
	}

	// Return the generated code as a string
	return strings.Join(generatedCode, "")
}
