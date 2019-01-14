package splitter

func Split(records [][]string) (questions []string, answers []string) {
	for _, s := range records {
		questions = append(questions, s[0])
		answers = append(answers, s[1])
	}
	return
}