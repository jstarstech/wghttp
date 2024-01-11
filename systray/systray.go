//go:build systray
// +build systray

package systray

import (
	"os"

	systraylib "github.com/getlantern/systray"
	"github.com/zhsj/wghttp/systray/icon"
)

const WithSystray = true

var disconnect func()
var connect func()

func Run(_connect func(), _disconnect func()) {
	connect = _connect
	disconnect = _disconnect

	onExit := func() {
		os.Exit(1)
	}

	systraylib.Run(onReady, onExit)
}

func onReady() {
	systraylib.SetIcon(icon.Data)
	systraylib.SetTitle("WGHTTP")
	systraylib.SetTooltip("WGHTTP")

	mDisconnect := systraylib.AddMenuItem("Disconnect", "Stop proxy and vpn")
	mConnect := systraylib.AddMenuItem("Connect", "Connect to vpn and sstart proxy")
	mQuit := systraylib.AddMenuItem("Quit", "Quit the whole app")
	mConnect.Hide()

	go func() {
		for {
			select {
			case <-mConnect.ClickedCh:
				go connect()
				mDisconnect.Show()
				mConnect.Hide()
			case <-mDisconnect.ClickedCh:
				disconnect()
				mDisconnect.Hide()
				mConnect.Show()
			case <-mQuit.ClickedCh:
				systraylib.Quit()
			}
		}
	}()
}
