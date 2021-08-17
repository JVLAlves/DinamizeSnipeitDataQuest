package functions

import (
	"regexp"
	"strconv"
	"strings"
	"time"
)

func Today() (Daytime string) {

	years, month, day := time.Now().Date()

	Day := strconv.Itoa(day)
	Month := strconv.Itoa(int(month))
	Year := strconv.Itoa(years)

	Daytime = "_" + Day + "_" + Month + "_" + Year

	return Daytime

}

func RegexThis(regex string, target string) (result string) {
	re := regexp.MustCompile(regex)
	filter := re.FindAllString(target, -1)
	dojoin := strings.Join(filter, "")
	return dojoin
}
