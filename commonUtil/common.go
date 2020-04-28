package commonUtil

import (
	"math/rand"
	"regexp"
	"time"
)

func CheckMailFormat(mail string) bool {
	return regexp.MustCompile(`^[A-Za-z0-9]+([_\.][A-Za-z0-9]+)*@([A-Za-z0-9\-]+\.)+[A-Za-z]{2,6}$`).MatchString(mail)
}

func CheckPhoneFormat(phone string) bool {
	return regexp.MustCompile(`^[1](([3][0-9])|([4][5-9])|([5][0-3,5-9])|([6][5,6])|([7][0-8])|([8][0-9])|([9][1,8,9]))[0-9]{8}$`).MatchString(phone)
}

//return rand num between [min, max]
func RandBetween(min, max int) int {
	if min >= max {
		panic("min bigger than max")
	}
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min+1) + min
}
