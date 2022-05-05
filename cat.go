package cat

import (
	"os"
	"r/cat/internal"
)

func SizeFormat(bytes float64) (_ string) { return internal.SizeFormat(bytes) }

func Stderr(err string) { internal.Stderr(err) }

func String(fp string) string { return internal.String(fp) }

func Bytes(fp string) []byte { return internal.Bytes(fp) }

func MD5sumChunked(fp string) (os.FileInfo, string, error) { return internal.MD5sumChunked(fp) }

func FileExist(fp string) bool { return internal.FileExist(fp) }

func Json(a interface{}) []byte { return internal.Json(a) }

func LanAddress() []string { return internal.LanAddress() }

func SystemInfo() *internal.Environment { return internal.SystemInfo() }
