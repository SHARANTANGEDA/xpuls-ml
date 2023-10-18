package utils

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/xpuls-com/xpuls-ml/types"
)

func NewUUIDWithPrefix(prefix types.UniqueIdPrefixes) (string, error) {
	id, err := uuid.NewUUID()

	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s-%s", prefix, id.String()), nil
}
