// +build !windows

package main

import (
	"io"

	"github.com/sirupsen/logrus"
)

// daemon 的主入口 (linux)
func runDaemon(opts *daemonOptions) error {

	// 初始化一个 cli 实例
	daemonCli := NewDaemonCli()

	// 启动 cli
	return daemonCli.start(opts)
}

func initLogging(_, stderr io.Writer) {
	logrus.SetOutput(stderr)
}
