package main

import (
	"log"

	"github.com/go-toast/toast"
)

func notify(title, message string) {
	notification := toast.Notification{
		AppID:   "Runnersuite Runner Test",
		Title:   title,
		Message: message,
	}
	if err := notification.Push(); err != nil {
		log.Fatalf("无法发送通知: %v", err)
	}
}

func main() {
	notify("Runner", "Runner.exe Test OK")
}
