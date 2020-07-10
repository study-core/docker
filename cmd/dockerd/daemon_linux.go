package main

import systemdDaemon "github.com/coreos/go-systemd/v22/daemon"

// preNotifySystem sends a message to the host when the API is active, but before the daemon is
//
// 当API处于活动状态但在守护程序未启动时，preNotifySystem将消息发送到主机
func preNotifySystem() {
}

// notifySystem sends a message to the host when the server is ready to be used
//
// 准备使用服务器时，notifySystem将消息发送到主机
func notifySystem() {
	// Tell the init daemon we are accepting requests
	go systemdDaemon.SdNotify(false, systemdDaemon.SdNotifyReady)
}
