package main

import (
	"bytes"
	"fmt"
	"net/http"
)

func main() {
    url := "#{DISCORD_WEBHOOK_URL}"
    chatMessage := "Hello, World!"
    var jsonStr = []byte(`{"content": ` + chatMessage + `}`)

    req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    res, err := client.Do(req)
    if err != nil {
        panic(err)
    }
    defer res.Body.Close()
}
