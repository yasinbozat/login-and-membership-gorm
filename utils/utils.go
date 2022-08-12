package utils

import (
	"crypto/md5"
	"encoding/hex"

	"github.com/shirou/gopsutil/host"
)

var (
	Stat2, _ = host.Info()
	Stat, _  = host.Info()
)

func GetMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func GetHWID() (text string) {
	text = GetMD5Hash(Stat2.HostID)
	return text
}
