package utils

import (
	"github.com/pkg/errors"
	"github.com/xpuls-com/xpuls-ml/common/constants"
)

func IsNotFound(err error) bool {
	return errors.Is(err, constants.ErrNotFound)
}
