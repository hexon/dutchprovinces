package dutchprovinces

import (
	_ "embed"
	"strconv"
)

// postalCodeMap can be seen as map from postal code number to the province number. The province index for postal code 6500 is simply byte offset 6500.
// TODO: Include the code to update it.
//
//go:embed bypostalcode.dat
var postalCodeMap []byte

// LookupPostalCode returns the ISO 3166-2 province code of the given postal code.
// Currently only the (four) digits of the postal code are considered, a future version might use the characters if needed. (AFAIK there is one postal code that is in two provinces.)
func LookupPostalCode(code string) (string, bool) {
	if len(code) > 4 {
		code = code[:4]
	}
	n, err := strconv.ParseInt(code, 10, 32)
	if err != nil {
		return "", false
	}
	if n < 0 || n > 9999 {
		return "", false
	}
	switch postalCodeMap[n] {
	case 0:
		// Unknown
		return "", false
	case 1:
		return "NL-DR", true
	case 2:
		return "NL-FL", true
	case 3:
		return "NL-FR", true
	case 4:
		return "NL-GE", true
	case 5:
		return "NL-GR", true
	case 6:
		return "NL-LI", true
	case 7:
		return "NL-NB", true
	case 8:
		return "NL-NH", true
	case 9:
		return "NL-OV", true
	case 10:
		return "NL-UT", true
	case 11:
		return "NL-ZE", true
	case 12:
		return "NL-ZH", true
	default:
		// This should never happen.
		return "", false
	}
}
