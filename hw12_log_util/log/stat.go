package log

func StatCalc(input <-chan Record) map[string]map[string]int {
	stat := make(map[string]map[string]int)
	for record := range input {
		processField(&stat, "message", record.Message)
	}

	return stat
}

func processField(stat *map[string]map[string]int, fieldName string, fieldValue string) {
	stats := *stat
	if _, ok := stats[fieldName]; !ok {
		stats[fieldName] = make(map[string]int)
	}
	stats[fieldName][fieldValue]++
}
