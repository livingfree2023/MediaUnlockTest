package main

import (
	core "MediaUnlockTest/pkg/core"
)

func checkUpdate(force bool) {
	cfg := core.UpdateConfig{
		AppName:         "unlock-test",
		VersionURL:      "https://api.github.com/repos/livingfree2023/MediaUnlockTest/releases/latest",
		BinaryURLPrefix: "https://github.com/livingfree2023/MediaUnlockTest/releases/latest/download/unlock-test",
		Silent:          false,
		ForceUpdate:     force,
	}
	core.CheckUpdate(cfg)
}

func checkUpdateOnly() {
	cfg := core.UpdateConfig{
		AppName:         "unlock-test",
		VersionURL:      "https://api.github.com/repos/livingfree2023/MediaUnlockTest/releases/latest",
		BinaryURLPrefix: "https://github.com/livingfree2023/MediaUnlockTest/releases/latest/download/unlock-test",
		Silent:          false,
		ForceUpdate:     false,
		JustCheck:       true,
	}
	core.CheckUpdate(cfg)
}
