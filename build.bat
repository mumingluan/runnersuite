go build -o Runnersuite.Autostart.exe -ldflags "-H=windowsgui -extldflags '-Wl,--manifest,manifest.xml'" Runnersuite.Autostart.go
go build -o Runnersuite.Launcher.exe -ldflags "-H=windowsgui -extldflags '-Wl,--manifest,manifest.xml'" Runnersuite.Launcher.go
go build -o runner.exe -ldflags "-H=windowsgui -extldflags '-Wl,--manifest,manifest.xml'" runner.go
go build -o runonce.exe -ldflags "-H=windowsgui -extldflags '-Wl,--manifest,manifest.xml'" runonce.go