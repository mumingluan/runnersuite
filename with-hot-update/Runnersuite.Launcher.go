package main

import (
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/go-toast/toast"
)

func notify(title, message string) {
	notification := toast.Notification{
		AppID:   "Runnersuite Launcher",
		Title:   title,
		Message: message,
	}
	if err := notification.Push(); err != nil {
		log.Fatalf("无法发送通知: %v", err)
	}
}

func updateAutostart(exePath string) {
	updatePath := `D:\OneDrive - MUML Alt\Runnersuite\Runnersuite.Autostart.New.exe`
	if _, err := os.Stat(updatePath); err == nil {
		// 覆盖主文件
		if err := os.Rename(updatePath, exePath); err != nil {
			notify("更新失败", "无法覆盖主文件 Runnersuite.Autostart.exe")
			return
		}
		notify("更新成功", "Runnersuite.Autostart.exe 已更新")
	}
}

func main() {
	// 等待900秒
	time.Sleep(900 * time.Second)

	exePath := `D:\OneDrive - MUML Alt\Runnersuite\Runnersuite.Autostart.exe`

	// 检查并更新 autostart.exe
	updateAutostart(exePath)

	// 启动主文件
	cmd := exec.Command(exePath)
	if err := cmd.Start(); err != nil {
		notify("启动失败", "无法启动 Runnersuite.Autostart.exe")
	}
}
