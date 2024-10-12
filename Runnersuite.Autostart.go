package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

func main() {
	runoncePath := `D:\OneDrive - MUML Alt\Runnersuite\runonce.exe`
	runnerPath := `D:\OneDrive - MUML Alt\Runnersuite\runner.exe`
	// 执行 runonce.exe
	if _, err := os.Stat(runoncePath); os.IsNotExist(err) {
		log.Printf("找不到 runonce.exe, 直接运行 runner.exe")
		// 直接执行 runner.exe
		if _, err := os.Stat(runnerPath); os.IsNotExist(err) {
			log.Fatalf("找不到 runner.exe")
		} else {
			cmd := exec.Command(runnerPath)
			if err := cmd.Start(); err != nil {
				log.Fatalf("运行 runner.exe 时出现错误")
			}
		}
		return
	}

	// 执行 runonce.exe
	cmd := exec.Command(runoncePath)
	if err := cmd.Start(); err != nil {
		log.Printf("运行 runonce.exe 时出现错误")
		return
	}
	// 重命名 runonce.exe
	newName := fmt.Sprintf("runonce.exe.%s", time.Now().Format("200601021504"))
	if err := os.Rename(runoncePath, filepath.Join(filepath.Dir(runoncePath), newName)); err != nil {
		log.Panicf("重命名失败: %v", err)
	}

	// 等待一分钟再执行 runner.exe
	time.Sleep(1 * time.Minute)

	if _, err := os.Stat(runnerPath); os.IsNotExist(err) {
		log.Fatalf("找不到 runner.exe")
	} else {
		cmd = exec.Command(runnerPath)
		if err := cmd.Start(); err != nil {
			log.Fatalf("运行 runner.exe 时出现错误")
		}
	}
}
