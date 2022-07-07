package utils

import (
	"io"
	"io/fs"
	"math/rand"
	"mime/multipart"
	"os"
	"regexp"
	"strconv"
	"time"
	"unsafe"
)

var regFormatHp *regexp.Regexp

const letterBytes = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"

const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
	letterLen     = 36                   // len(letterBytes)
)

var src = rand.NewSource(time.Now().UnixNano())

func init() {
	regFormatHp = regexp.MustCompile(`(^\+?628)|(^0?8){1}`)
}

func RandString(n int) string {
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < letterLen {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return *(*string)(unsafe.Pointer(&b))
}

func Rand4DigitInt() string {
	return strconv.FormatInt((rand.Int63n(8999) + 1000), 10)
}

func Rand6DigitInt() string {
	return strconv.FormatInt((rand.Int63n(899999) + 100000), 10)
}

func FormatPhoneTo62(phone string) string {
	formatPhone := regFormatHp.ReplaceAllString(phone, "628")
	return formatPhone
}

func CreateFolder(folderPath string, perm fs.FileMode) error {
	var err error
	if _, err = os.Stat(folderPath); os.IsNotExist(err) {
		err = os.MkdirAll(folderPath, perm)
		if err != nil {
			return err
		}
	}
	return nil
}

func UploadImage(path string, file *multipart.FileHeader) error {
	var err error
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	// Destination
	dst, err := os.Create(path)
	if err != nil {
		return err
	}
	defer dst.Close()

	// Copy
	_, err = io.Copy(dst, src)
	if err != nil {
		return err
	}

	return err
}

func RemoveImage(path string) error {
	var err error

	_, err = os.Stat(path)
	if !os.IsNotExist(err) {
		err = os.Remove(path)
		if err != nil {
			return err
		}
	}

	return err
}
