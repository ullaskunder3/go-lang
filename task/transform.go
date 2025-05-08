package task

func Transform(input string, modifier func(string) string) string {
	for i := 0; i < 5; i++ {
		input = modifier(input)
	}
	return input + " (done)"
}
