package gosetlocale

// #include <locale.h>
import (
	"C"
)

const (
	// LcAll is used for all of the locale.
	LcAll = 0
	// LcCollate is used for regular expression matching (it determines the meaning of range expressions and equivalence classes) and string collation.
	LcCollate = 1
	// LcCtype is used for regular expression matching, character classification, conversion, case-sensitive comparison, and wide character functions.
	LcCtype = 2
	// LcMonetary is used for monetary formatting.
	LcMonetary = 3
	// LcNumeric is used for number formatting (such as the decimal point and the thousands separator).
	LcNumeric = 5
	// LcTime is used for time and date formatting.
	LcTime = 6
)

// SetLocale is used to set or query the program's current locale.
func SetLocale(lc C.int, locale string) string {
	param := C.CString(locale)
	ret := C.setlocale(lc, param)
	return C.GoString(ret)
}
