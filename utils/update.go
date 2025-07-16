package utils

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"runtime"
	"strings"
)

const RepoURL = "https://github.com/tu-usuario/incognito-chain/releases/latest "

func CheckForUpdate() bool {
	resp, err := http.Get(RepoURL)
	if err != nil {
		return false
	}
	defer resp.Body.Close()

	latestTag := resp.Request.URL.String()
	tag := strings.Split(latestTag, "/")[len(strings.Split(latestTag, "/"))-1]
	return tag != "v1.0.0"
}

func DownloadUpdate(version string) error {
	osName := runtime.GOOS
	var url string

	switch osName {
	case "windows":
		url = fmt.Sprintf("https://github.com/tu-usuario/incognito-chain/releases/download/%s/incognito-node-windows-amd64.exe ", version)
	case "linux":
		url = fmt.Sprintf("https://github.com/tu-usuario/incognito-chain/releases/download/%s/incognito-node-linux-amd64 ", version)
	default:
		return fmt.Errorf("plataforma no soportada")
	}

	resp, _ := http.Get(url)
	defer resp.Body.Close()

	outFile := "incognito-node-new"
	if osName == "windows" {
		outFile += ".exe"
	}

	out, _ := os.Create(outFile)
	io.Copy(out, resp.Body)
	return nil
}
