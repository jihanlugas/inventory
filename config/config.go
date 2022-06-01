package config

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/joho/godotenv"
	"os"
	"strconv"
)

const (
	Production = "PRODUCTION"
	FormatTime = "2006-01-02"
)

type targetServer struct {
	Address string
	Port    string
}

type dbServerInfo struct {
	Host     string
	Port     string
	Username string
	Password string
	DbName   string
}

var (
	ListenTo               targetServer
	Environment            string
	DBInfo                 dbServerInfo
	LogPath                string
	CertificateFilePath    string
	CertificateKeyFilePath string
	CookieAuthName         string
	CryptoKey              []byte
	MaxSizeUploadPhotoByte int64
	DataPerPage            int
)

func init() {
	var err error

	err = godotenv.Load()
	if err != nil {
		panic(err)
	}

	Environment = os.Getenv("ENVIRONMENT")
	ListenTo = targetServer{
		Address: os.Getenv("LISTEN_ADDRESS"),
		Port:    os.Getenv("LISTEN_PORT"),
	}
	DBInfo = dbServerInfo{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Username: os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASS"),
		DbName:   os.Getenv("DB_NAME"),
	}
	LogPath = os.Getenv("LOG_PATH")
	CertificateFilePath = os.Getenv("CERTIFICATE_FILE_PATH")
	CertificateKeyFilePath = os.Getenv("CERTIFICATE_KEY_FILE_PATH")
	CookieAuthName = os.Getenv("COOKIE_AUTH_NAME")
	MaxSizeUploadPhotoByte, err = strconv.ParseInt(os.Getenv("MAX_SIZE_UPLOAD_PHOTO_BYTE"), 10, 64)
	if err != nil {
		panic(err)
	}
	DataPerPage, err = strconv.Atoi(os.Getenv("MIN_DATA_PER_PAGE"))
	if err != nil {
		panic(err)
	}

	hasher := md5.New()
	hasher.Write([]byte(os.Getenv("CRYPTO_KEY")))
	CryptoKey = []byte(hex.EncodeToString(hasher.Sum(nil)))
}
