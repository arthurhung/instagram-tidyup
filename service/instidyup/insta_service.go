package instidyup

import (
	"bytes"
	"encoding/base64"
	"os"

	"github.com/ahmdrz/goinsta/v2"
	"github.com/arthurhung/instagram-tidyup/pkg/logging"
	"github.com/arthurhung/instagram-tidyup/pkg/util"
)

// type ConfigFile struct {
// 	ID        int64          `json:"id"`
// 	User      string         `json:"username"`
// 	DeviceID  string         `json:"device_id"`
// 	UUID      string         `json:"uuid"`
// 	RankToken string         `json:"rank_token"`
// 	Token     string         `json:"token"`
// 	PhoneID   string         `json:"phone_id"`
// 	Cookies   []*http.Cookie `json:"cookies"`
// }

// InsTidyUp ...
// type InsTidyUp struct {
// 	Insta *goinsta.Instagram
// }

// GetInsta ...
func GetInsta(username, password string) *goinsta.Instagram {
	insta := goinsta.New(
		username,
		password)
	err := insta.Login()
	if err != nil {
		logging.Error("ERROR:", err)
		return nil
	}
	logging.Info("Successfully logged in")
	util.GenerateToken(username, password)
	return insta
}

// type MongoWriter struct {
// 	sess *mgo.Session
// }

// func (mw *MongoWriter) Write(p []byte) (n int, err error) {
// 	c := mw.sess.DB("").C("log")
// 	err = c.Insert(bson.M{
// 		"created": time.Now(),
// 		"msg":     string(p),
// 	})
// 	if err != nil {
// 		return
// 	}
// 	return len(p), nil
// }

// StoreSession  store user login session to database
func StoreSession(insta *goinsta.Instagram) {

	result, _ := ExportAsBase64String(insta)
	logging.Info(result)
	os.Setenv("IG", result)

}

// ExportAsBytes ...
func ExportAsBytes(insta *goinsta.Instagram) ([]byte, error) {
	buffer := &bytes.Buffer{}
	err := goinsta.Export(insta, buffer)
	if err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}

// ExportAsBase64String exports selected *Instagram object as base64 encoded string
func ExportAsBase64String(insta *goinsta.Instagram) (string, error) {
	bytes, err := ExportAsBytes(insta)
	if err != nil {
		return "", err
	}

	sEnc := base64.StdEncoding.EncodeToString(bytes)
	return sEnc, nil
}

func ImportFromBytes(inputBytes []byte) (*goinsta.Instagram, error) {
	return goinsta.ImportReader(bytes.NewReader(inputBytes))
}

// ImportFromBase64String imports instagram configuration from a base64 encoded string.
func ImportFromBase64String(base64String string) (*goinsta.Instagram, error) {
	sDec, err := base64.StdEncoding.DecodeString(base64String)
	if err != nil {
		return nil, err
	}

	return ImportFromBytes(sDec)
}
