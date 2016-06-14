package go_strings
import (
"golang.org/x/crypto/bcrypt"
"strings"
	"time"
)

func GenerateUniqueString() (string,error) {
	now := time.Now().Format("20060102_150405_")
	b,err := bcrypt.GenerateFromPassword([]byte(now), 1)
	if err != nil {
		return "",err
	}
	bStr := strings.TrimSpace(string(b))
	bStr = strings.Replace(bStr, "$", "", -1)
	bStr = strings.Replace(bStr, ".", "", -1)
	bStr = strings.Replace(bStr, "/", "", -1)
	return now+bStr, nil
}