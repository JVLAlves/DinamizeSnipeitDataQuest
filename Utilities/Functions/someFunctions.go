package functions

import (
	"log"
	"os"
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

//Cria arquivos com as informações retiradas do computador via Terminal
func ActiveLogs() {

	var errboolean bool = true
	_, err := os.Stat(os.Getenv("USERNAME") + "_logs")
	if os.IsNotExist(err) {
		errboolean = false
	}
	if err != nil {
		errboolean = false
	}
	HOME, boolean := os.LookupEnv("HOME")
	USERNAME := os.Getenv("USERNAME")
	if !(boolean && errboolean) {

		HOME, err := os.UserHomeDir()
		if err != nil {
			log.Fatalf("ERROR CAPTING USERHOMEDIR: %v", err)
		}
		os.Setenv("HOME", HOME)
		USERname := os.Getenv("USERNAME")
		os.Mkdir(HOME+"/"+USERname+"_logs", 0777)
	}

	HOMELOGS := HOME + "/" + USERNAME + "_logs"
	path := HOMELOGS
	years, month, day := time.Now().Date()

	Day := strconv.Itoa(day)
	Month := strconv.Itoa(int(month))
	Year := strconv.Itoa(years)

	Daytime := "_" + Day + "_" + Month + "_" + Year

	logname := path + "/Logs" + Daytime + ".txt"

	outFile, err := os.OpenFile(logname, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		log.Println("error creating file", err)
	}
	log.SetOutput(outFile)
}
