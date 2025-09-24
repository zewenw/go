package examples

import (
	"encoding/base64"
	"fmt"
	"testing"
)

func TestBase64(t *testing.T) {

	t.Run("Base64 Demo", func(t *testing.T) {
		data := "abc123!?$*&()'-=@~"
		sEnc := base64.StdEncoding.EncodeToString([]byte(data))
		fmt.Println(sEnc)

		sDec, _ := base64.StdEncoding.DecodeString(sEnc)
		fmt.Println(string(sDec))
		fmt.Println()

		uEnc := base64.URLEncoding.EncodeToString([]byte(data))
		fmt.Println(uEnc)
		uDec, _ := base64.URLEncoding.DecodeString(uEnc)
		fmt.Println(string(uDec))
	})
}
