package list

// CalcMaxColumnWidth - Calculate the maximum width of elements in the list
func CalcMaxColumnWidth(list *[]string) (maxWidth int) {
	if list != nil {
		for _, line := range *list {
			if width := len(line); width > maxWidth {
				maxWidth = width
			}
		}
	}
	return maxWidth
}
