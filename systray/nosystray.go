//go:build !systray
// +build !systray

package systray

const WithSystray = false

func Run(_connect func(), _disconnect func()) {

}
