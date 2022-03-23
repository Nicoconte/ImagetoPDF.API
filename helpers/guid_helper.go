package helpers

import (
	"strings"

	"github.com/pborman/uuid"
)

func GetGuid() string {
	return strings.Replace(uuid.New(), "-", "", -1)
}
