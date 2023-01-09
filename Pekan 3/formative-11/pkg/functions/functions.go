package functions

import (
	"bytes"
	"encoding/json"
)

// Functions untuk cetak JSON menjadi lebih cantik
func PrettyString(str string) (string, error) {
	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, []byte(str), "", "    "); err != nil {
		return "", err
	}
	return prettyJSON.String(), nil
}
