package utils

import (
	"bytes"
	"math/rand"
	"strconv"
	"time"
)

func GenerateBillCode() string {
	var billCode bytes.Buffer
	currentTime := time.Now()
	billCode.WriteString(currentTime.Format("02012006"))
	billCode.WriteString(strconv.Itoa(rand.Intn(200)))
	return billCode.String()
}
