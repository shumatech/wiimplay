package main

import (
    "errors"
    "log"
    "os"
    "path/filepath"
    "runtime"
    "github.com/shumatech/wiimplay/config"
    "github.com/shumatech/wiimplay/ui"

    "github.com/gotk3/gotk3/glib"
    "github.com/gotk3/gotk3/gtk"
)

var (
    Version string
    Build string
)

func fatalErr(err error) {
    if err != nil {
        log.Fatal(err.Error())
    }
}

func main() {
    runtime.LockOSThread()

    log.Printf("WiiM Play %s (Build: %s)", Version, Build)

    const appID = "com.shumatech.wiimplay"

    configDir := filepath.Join(glib.GetUserConfigDir(), "wiimplay")
    _, err := os.Stat(configDir)
    if errors.Is(err, os.ErrNotExist) {
        err = os.MkdirAll(configDir, 0755)
    }
    fatalErr(err)

    config, err := config.NewConfig(configDir)
    fatalErr(err)

    application, err := gtk.ApplicationNew(appID, glib.APPLICATION_FLAGS_NONE)
    fatalErr(err)

    application.Connect("activate", func() {
        _, err = ui.NewController(application, config, Version, Build)
        fatalErr(err)
    })

    os.Exit(application.Run(os.Args))
}
