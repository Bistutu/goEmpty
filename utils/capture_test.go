package utils

import (
	"fmt"
	"testing"
)

func TestCapture(t *testing.T) {
	fmt.Println(getCaptureCode())
}

func TestIsNeedCapture(t *testing.T) {
	fmt.Println(isNeedCaptcha("2018011185"))
}
