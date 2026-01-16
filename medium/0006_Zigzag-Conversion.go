package main

import "bytes"

func convert(s string, numRows int) string {
	var buffer bytes.Buffer

	if numRows == 1 || numRows >= len(s) {
		return s
	}

	rows := make([]bytes.Buffer, numRows)
	index, step := 0, 1

	for _, char := range s {
		rows[index].WriteRune(char)

		if index == 0 {
			step = 1
		} else if index == numRows-1 {
			step = -1
		}
		index += step
	}

	for _, row := range rows {
		buffer.WriteString(row.String())
	}

	return buffer.String()
}
