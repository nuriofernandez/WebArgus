package notify

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

var GlaucusSMSServiceUrl = os.Getenv("GLAUCUS_SMS_SERVICE_URL")
var GlaucusSMSServiceToken = os.Getenv("GLAUCUS_SMS_SERVICE_TOKEN")

func sendSms(number, title, content string) {
	payload := strings.NewReader(
		`{` +
			`"recipients": ["` + number + `"],` +
			`"senderTitle": "` + title + `",` +
			`"message": "` + content +
			`"}`,
	)

	client := &http.Client{}
	req, err := http.NewRequest("POST", GlaucusSMSServiceUrl, payload)

	// Authenticate
	req.Header.Set("Authorization", "Bearer "+GlaucusSMSServiceToken)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
