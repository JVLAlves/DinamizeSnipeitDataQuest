package functions

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"regexp"
	"runtime"
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

type PowerShell struct {
	powerShell string
}

func New() *PowerShell {
	ps, _ := exec.LookPath("powershell.exe")
	return &PowerShell{
		powerShell: ps}
}

//Definindo os Argumentos necessários para executar um comendo no powershell
func (p *PowerShell) Execute(args ...string) (stdOut string, stdErr string, err error) {
	args = append([]string{"-NoProfile", "-NonInteractive"}, args...)
	cmd := exec.Command(p.powerShell, args...)

	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err = cmd.Run()
	stdOut, stdErr = stdout.String(), stderr.String()
	return
}

//Cria o diretório
func CreateDir(wg *sync.WaitGroup) {
	switch runtime.GOOS {
	case "darwin":
		//forMacOs()
	case "linux":
		LinuxDir(wg)

	case "windows":
		//WindowsDir(wg)
	default:
		fmt.Println("ERROR! Could not found the Operating System!")
		time.Sleep(time.Second * 1)
		fmt.Println("Aborting")
		for i := 0; i < 4; i++ {
			time.Sleep(time.Second * 1)
			fmt.Print(".")
		}
		time.Sleep(time.Second * 3)
		log.Fatal()
	}
}

func WindowsDir(wg *sync.WaitGroup) {
	//invocando o PowerShell
	posh := New()

	//Aplicanddo os comandos literais que serão executados no powershell
	stdout, _, _ := posh.Execute("$env:userprofile")
	home := stdout
	os.Setenv("HOME", home[:len(home)-2])
	username := os.Getenv("USERNAME")
	_, _, _ = posh.Execute("New-Item -path \"$env:userprofile\" -Name \"" + username + "_logs" + "\" -ItemType \"directory\"")
	wg.Done()
}

func LinuxDir(wg *sync.WaitGroup) {

	cmd := exec.Command("bash", "-c", "echo"+" "+"$HOME > ./.Pathfinder.txt")
	err := cmd.Run()
	if err != nil {
		log.Fatalf("PANIC DE CRIAÇÃO DO PATHFINDER: %v", err)
	}
	file, err := os.Open(".Pathfinder.txt")
	if err != nil {
		log.Print(err)
	}
	fileScanner := bufio.NewScanner(file)
	Linhas := []string{}
	for fileScanner.Scan() {
		Linhas = append(Linhas, fileScanner.Text())
		if fileScanner.Err() != nil {
			log.Fatalf("Erro SCAN: %v", fileScanner.Err().Error())
		}
	}
	os.Setenv("HOME", Linhas[0])
	username := os.Getenv("USERNAME")
	file.Close()
	cmd = exec.Command("rm", "./.Pathfinder.txt")
	_, _ = cmd.Output()

	cmd = exec.Command("bash", "-c", "mkdir"+" "+"$HOME/"+username+"_logs")
	err = cmd.Run()
	if err != nil {
		panic(err)
	}
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
	HOME, boolean := os.LookupEnv("HOME")
	USERNAME := os.Getenv("USERNAME")
	HOMELOGS := HOME + "/" + USERNAME + "_logs"

	if !(boolean && errboolean) {
		wg := &sync.WaitGroup{}
		wg.Add(1)
		CreateDir(wg)
		wg.Wait()
	}

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
