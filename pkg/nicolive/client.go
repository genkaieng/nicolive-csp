package nicolive

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"regexp"
	"strconv"

	"github.com/genkaieng/nicolive-csp/gen/pb"
	"github.com/gorilla/websocket"
)

func getNicoliveMessage(uri string) {
	resp, err := http.Get(uri)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	ch, err := decodeChunkedMessage(resp.Body)
	if err != nil {
		panic(err)
	}
	for chunkedMessage := range ch {
		switch payload := chunkedMessage.Payload.(type) {
		case *pb.ChunkedMessage_Message:
			b, err := json.Marshal(payload)
			if err != nil {
				log.Println("json.Marshal error:", err)
				continue
			}
			log.Println("message", string(b))
		case *pb.ChunkedMessage_State:
			b, err := json.Marshal(payload)
			if err != nil {
				log.Println("json.Marshal error:", err)
				continue
			}
			log.Println("state", string(b))
		case *pb.ChunkedMessage_Signal_:
			log.Println("signal", payload)
			if payload.Signal == pb.ChunkedMessage_Flushed {
				// TODO: 終了処理
			}
		}
	}
}

func connectCommentServer(uri string) {
	var at = "now"
	for {
		resp, err := http.Get(uri + "?at=" + at)
		if err != nil {
			panic(err)
		}

		ch, err := decodeChunkedEntry(resp.Body)
		if err != nil {
			panic(err)
		}
		for chunkedEntry := range ch {
			log.Println("chunked_entry:", chunkedEntry)
			switch entry := chunkedEntry.Entry.(type) {
			case *pb.ChunkedEntry_Backward:
				// handle backward entry
			case *pb.ChunkedEntry_Segment:
				// handle segment entry
				uri := entry.Segment.Uri
				getNicoliveMessage(uri)
			case *pb.ChunkedEntry_Previous:
				// handle previous entry
			case *pb.ChunkedEntry_Next:
				// handle next entry
				at = strconv.FormatInt(entry.Next.At, 10)
			}
		}
		resp.Body.Close()
	}
}

func GetWsUri(lvid string) (string, error) {
	resp, err := http.Get("https://live.nicovideo.jp/watch/" + lvid)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	re := regexp.MustCompile(`&quot;(wss://.+?)&quot;`)
	matches := re.FindSubmatch(body)
	if len(matches) < 2 {
		return "", fmt.Errorf("Failed to extract wss uri from HTML of https://live.nicovideo.jp/watch/%s", lvid)
	}
	uri := matches[1]

	return string(uri), nil
}

func Connect(lvid string) {
	uri, err := GetWsUri(lvid)
	if err != nil {
		panic(err)
	}
	c, _, err := websocket.DefaultDialer.Dial(string(uri), nil)
	if err != nil {
		panic(err)
	}
	defer c.Close()

	// send message: startWatching
	err = c.WriteMessage(websocket.TextMessage, []byte(`{"type":"startWatching","data":{"stream":{"quality":"high","protocol":"hls","latency":"low","chasePlay":false},"room":{"protocol":"webSocket","commentable":false},"reconnect":false}}`))
	if err != nil {
		panic(err)
	}

	for {
		_, msg, err := c.ReadMessage()
		if err != nil {
			log.Println(err)
			break
		}
		log.Println(string(msg))
		if bytes.Contains(msg, []byte(`"type":"ping"`)) {
			err = c.WriteMessage(websocket.TextMessage, []byte(`{"type":"pong"}`))
			if err != nil {
				log.Println("Failed send pong:", err)
			}
			continue
		}
		if bytes.Contains(msg, []byte(`"type":"messageServer"`)) {
			re := regexp.MustCompile(`"viewUri":"([^"]+)"`)
			matches := re.FindSubmatch(msg)
			if len(matches) >= 2 {
				uri := matches[1]
				log.Println("messageServer:", string(uri))
				connectCommentServer(string(uri))
			}
			continue
		}
	}
}
