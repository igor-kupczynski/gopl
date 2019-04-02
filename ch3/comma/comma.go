package comma

import (
	"bytes"
	"strings"
)

func comma(x string) string {

	n := len(x)

	// Check sign +/-
	prefix := 0
	if strings.HasPrefix(x, "+") {
		prefix = 1
	} else if strings.HasPrefix(x, "-") {
		prefix = 1
	}

	// Check the decimal part
	postfix := 0
	point := strings.Index(x, ".")
	if point > 0 {
		postfix = n - point
	}

	// Digits to comma separate
	ndigits := n - prefix - postfix

	// Short number, let's short circuit
	if ndigits <= 3 {
		return x
	}

	// Sign if any and the first digit group
	start := ndigits%3 + prefix

	// Result buffer
	var result bytes.Buffer
	result.WriteString(x[:start])

	// Iterate over other digit groups
	for i := start; i < prefix+ndigits; i += 3 {
		result.WriteString(",")
		result.WriteString(x[i : i+3])
	}

	// Write the decimal part if any
	if point > 0 {
		result.WriteString(x[point:])
	}

	return result.String()
}
