package config

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"strconv"
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
	Environment            string
	Debug                  bool
	ListenTo               targetServer
	DBInfo                 dbServerInfo
	LogPath                string
	CertificateFilePath    string
	CertificateKeyFilePath string
	CookieAuthName         string
	CryptoKey              []byte
	MaxSizeUploadPhotoByte int64
	DataPerPage            int
	PhotoDirectory         string
	PhotoincRunningLimit   int64
	PhotoAccessUrl         string
)

func init() {
	var err error

	err = godotenv.Load()
	if err != nil {
		fmt.Println("Failed load env Err: " + err.Error())
		panic(err)
	}

	Environment = os.Getenv("DEBUG")
	Debug, err = strconv.ParseBool(os.Getenv("DEBUG"))
	if err != nil {
		fmt.Println("Failed parse DEBUG Err: " + err.Error())
		panic(err)
	}
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
	DataPerPage, err = strconv.Atoi(os.Getenv("MIN_DATA_PER_PAGE"))
	if err != nil {
		panic(err)
	}

	hasher := md5.New()
	hasher.Write([]byte(os.Getenv("CRYPTO_KEY")))
	CryptoKey = []byte(hex.EncodeToString(hasher.Sum(nil)))

	MaxSizeUploadPhotoByte, err = strconv.ParseInt(os.Getenv("MAX_SIZE_UPLOAD_PHOTO_BYTE"), 10, 64)
	if err != nil {
		panic(err)
	}
	PhotoDirectory = os.Getenv("PHOTO_DIRECTORY")
	PhotoincRunningLimit, err = strconv.ParseInt(os.Getenv("PHOHOINC_RUNNING_LIMIT"), 10, 64)
	if err != nil {
		panic(err)
	}
	PhotoAccessUrl = os.Getenv("PHOTO_ACCESS_URL")

}
