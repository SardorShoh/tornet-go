package tornet

import (
	"os/exec"
	"runtime"
)

// Start tornet
func startTornet() error {
	switch runtime.GOOS {
	case "linux":
		return exec.Command("sudo", "systemctl", "start", "tor").Run()
	case "darwin":
		return exec.Command("brew", "services", "start", "tor").Run()
	default:
		return nil
	}
}

// Stop tornet
func stopTornet() error {
	switch runtime.GOOS {
	case "linux":
		return exec.Command("sudo", "systemctl", "stop", "tor").Run()
	case "darwin":
		return exec.Command("brew", "services", "stop", "tor").Run()
	default:
		return nil
	}
}

// Restart tornet
func restartTornet() error {
	switch runtime.GOOS {
	case "linux":
		return exec.Command("sudo", "systemctl", "restart", "tor").Run()
	case "darwin":
		return exec.Command("brew", "services", "restart", "tor").Run()
	default:
		return nil
	}
}

// Check if tornet is running
func isTornetRunning() bool {
	cmd := exec.Command("pgrep", "-x", "tor")
	out, err := cmd.Output()
	if err != nil {
		return false
	}
	return len(out) > 0
}
