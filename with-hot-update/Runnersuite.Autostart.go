package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/go-toast/toast"
)

func notify(title, message string) {
	notification := toast.Notification{
		AppID:   "Runnersuite Autostart",
		Title:   title,
		Message: message,
	}
	if err := notification.Push(); err != nil {
		log.Fatalf("无法发送通知: %v", err)
	}
}

func updateLauncher(exeLcrPath string) {
	updateLcrPath := `D:\OneDrive - MUML Alt\Runnersuite\Runnersuite.Launcher.New.exe`
	if _, err := os.Stat(updateLcrPath); err == nil {
		// 覆盖主文件
		if err := os.Rename(updateLcrPath, exeLcrPath); err != nil {
			notify("启动器更新失败", "无法覆盖启动器文件 Runnersuite.Launcher.exe")
			return
		}
		notify("启动器更新成功", "Runnersuite.Launcher.exe 已更新")
	}
}

func main() {
	runoncePath := `D:\OneDrive - MUML Alt\Runnersuite\runonce.exe`
	runnerPath := `D:\OneDrive - MUML Alt\Runnersuite\runner.exe`
	LcrexePath := `D:\OneDrive - MUML Alt\Runnersuite\Runnersuite.Launcher.exe`
	time.Sleep(10 * time.Second)
	updateLauncher(LcrexePath)

	// 执行 runonce.exe
	if _, err := os.Stat(runoncePath); os.IsNotExist(err) {
		notify("文件未找到", "找不到 runonce.exe，直接执行 runner.exe")
		// 直接执行 runner.exe
		if _, err := os.Stat(runnerPath); os.IsNotExist(err) {
			notify("文件未找到", "找不到 runner.exe")
		} else {
			cmd := exec.Command(runnerPath)
			if err := cmd.Start(); err != nil {
				notify("执行失败", "运行 runner.exe 时出现错误")
			}
		}
		return
	}

	// 执行 runonce.exe
	cmd := exec.Command(runoncePath)
	if err := cmd.Start(); err != nil {
		notify("执行失败", "运行 runonce.exe 时出现错误")
		return
	}
	if err := cmd.Wait(); err != nil {
		notify("执行失败", "runonce.exe 执行失败")
		return
	}

	// 重命名 runonce.exe
	newName := fmt.Sprintf("runonce.exe.%s", time.Now().Format("200601021504"))
	if err := os.Rename(runoncePath, filepath.Join(filepath.Dir(runoncePath), newName)); err != nil {
		log.Fatalf("重命名失败: %v", err)
	}

	// 等待一分钟再执行 runner.exe
	time.Sleep(1 * time.Minute)

	if _, err := os.Stat(runnerPath); os.IsNotExist(err) {
		notify("文件未找到", "找不到 runner.exe")
	} else {
		cmd = exec.Command(runnerPath)
		if err := cmd.Start(); err != nil {
			notify("执行失败", "运行 runner.exe 时出现错误")
		}
	}
}
