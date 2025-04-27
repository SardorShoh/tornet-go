package tornet

import "resty.dev/v3"

func Initialize() error {
	if err := installTornet(); err != nil {
		return err
	}
	if !isTornetRunning() {
		return startTornet()
	}
	return restartTornet()
}

// Get tornet ip
func GetIp() string {
	c := resty.New()
	defer c.Close()
	c.SetProxy("socks5://127.0.0.1:9050")
	resp, err := c.R().Get("https://api.ipify.org")
	if err != nil {
		return ""
	}
	return resp.String()
}

// Change tornet ip
func ChangeIp() error {
	return restartTornet()
}

func Stop() error {
	return stopTornet()
}
