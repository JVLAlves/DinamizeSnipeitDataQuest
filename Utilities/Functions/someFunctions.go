package functions

import (
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"sync"
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

func CreateDir(wg *sync.WaitGroup) {
	HOME, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("ERROR CAPTING USERHOMEDIR: %v", err)
	}
	os.Setenv("HOME", HOME)
	USERname := os.Getenv("USERNAME")
	os.Mkdir(HOME+"/"+USERname+"_logs", 0777)
	wg.Done()
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
	_, boolean := os.LookupEnv("HOME")
	USERNAME := os.Getenv("USERNAME")
	if !(boolean && errboolean) {
		wg := &sync.WaitGroup{}
		wg.Add(1)
		CreateDir(wg)
		wg.Wait()
	}
	HOME := os.Getenv("HOME")
	HOMELOGS := HOME + "/" + USERNAME + "_logs"
	years, month, day := time.Now().Date()

	Day := strconv.Itoa(day)
	Month := strconv.Itoa(int(month))
	Year := strconv.Itoa(years)

	Daytime := "_" + Day + "_" + Month + "_" + Year

	logname := HOMELOGS + "/Logs" + Daytime + ".txt"

	outFile, err := os.OpenFile(logname, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		log.Println("error creating file", err)
	}
	log.SetOutput(outFile)
}
