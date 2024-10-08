package mattermost

import (
	"lifelogger/server/util/converter"
	"lifelogger/server/util/models/blocks"

	"fmt"
	"net/http"
	"os"
)

var Noti = func(title, text string) {
	attachments := []map[string]string{
		{
			"title":  title,
			"text":   text,
			"color":  "#ff5d52",
			"footer": blocks.Now().Format(blocks.DateTimeFormat),
		},
	}
	props := map[string]interface{}{
		"attachments": attachments,
	}
	payload := map[string]interface{}{
		"channel_id": os.Getenv("MATTERMOST_CHANNEL_ID"),
		"props":      props,
	}

	req, _ := http.NewRequest(
		"POST",
		os.Getenv("MATTERMOST_URL"),
		converter.ConvertPayload(payload),
	)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+os.Getenv("MATTERMOST_TOKEN"))

	_, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
	}
}
