package api

import "github.com/nuttchai/go-rest/internal/constant"

func CreatePath(path string) string {
	return constant.BasePath + path
}
