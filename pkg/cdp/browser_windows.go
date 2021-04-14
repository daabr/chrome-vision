package cdp

import (
	"os"
	"path/filepath"
)

const (
	localAppData = os.Getenv("LOCALAPPDATA")
	execSuffix   = `Application\chrome.exe`
)

var executables = [...]string{
	filepath.Join(localAppData, `Google\Chrome`, execSuffix),
	filepath.Join(`C:\Program Files`, `Google\Chrome`, execSuffix),
	filepath.Join(`C:\Program Files (x86)`, `Google\Chrome`, execSuffix),

	// Canary.
	filepath.Join(localAppData, `Google\Chrome SxS`, execSuffix),
	filepath.Join(`C:\Program Files`, `Google\Chrome SxS`, execSuffix),
	filepath.Join(`C:\Program Files (x86)`, `Google\Chrome SxS`, execSuffix),

	filepath.Join(localAppData, "Chromium", execSuffix),
	filepath.Join(`C:\Program Files`, "Chromium", execSuffix),
	filepath.Join(`C:\Program Files (x86)`, "Chromium", execSuffix),

	"chrome.exe",
}
