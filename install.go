//go:build !windows

package tornet

import (
	"bufio"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

// Check if the OS is Arch Linux
func isArchLinux() bool {
	filePath := "/etc/os-release"
	file, err := os.Open(filePath)
	if err != nil {
		return false
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}
		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])
		value = strings.Trim(value, "\"")
		if key == "ID" && value == "arch" {
			return true
		}
	}
	return false
}

// Check if tor is installed
func isTornetAvailable() bool {
	cmd := exec.Command("which", "tor")
	out, err := cmd.Output()
	if err != nil {
		return false
	}
	return strings.Contains(string(out), "tor")
}

// Install tornet
func installTornet() error {
	if !isTornetAvailable() {
		switch runtime.GOOS {
		case "linux":
			if isArchLinux() {
				return exec.Command("sudo", "pacman", "-Sy", "tor", "--noconfirm").Run()
			} else {
				if err := exec.Command("sudo", "apt", "update").Run(); err != nil {
					return err
				}
				return exec.Command("sudo", "apt", "install", "tor", "-y").Run()
			}
		case "darwin":
			return exec.Command("brew", "install", "tor").Run()
		default:
			return nil
		}
	}
	return nil
}
