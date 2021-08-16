package MacOS

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"sync"
	"time"
)

//Lista para leitura linha a linha
var Linhas []string

//Lista para Informações armazenadas
var Infos []string

//Cria arquivos com as informações retiradas do computador via Terminal
func Create(wg *sync.WaitGroup, command string, args string) {

	outFile, err := os.OpenFile(command+".out", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		fmt.Println("error creating file", err)
	}
	defer outFile.Close()
	cmd := exec.Command("bash", "-c", command+" "+args)

	out, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println("error attaching command stdout", err)
	}
	go io.Copy(outFile, out)

	err = cmd.Run()
	if err != nil {
		panic(err)
	}

	fmt.Println("Arquivo criado...")
	time.Sleep(time.Second)
	wg.Done()
}

//Lê os arquivos criados pela função Create
func Running() {

	file, err := os.Open("uname.out")
	if err != nil {
		log.Print(err)
	}
	fileScanner := bufio.NewScanner(file)
	Linhas = []string{}

	for fileScanner.Scan() {
		Linhas = append(Linhas, fileScanner.Text())
		if fileScanner.Err() != nil {
			log.Fatalf("Erro SCAN: %v", fileScanner.Err().Error())
		}
	}
	Infos = append(Infos, Linhas[0])
	file.Close()

	file, err = os.Open("sysctl.out")
	if err != nil {
		log.Print(err)
	}
	fileScanner = bufio.NewScanner(file)
	Linhas = []string{}
	for fileScanner.Scan() {
		Linhas = append(Linhas, fileScanner.Text())
	}
	Infos = append(Infos, Linhas[0])
	fmt.Println(Infos[1])
	file.Close()

	file, err = os.Open("hostinfo.out")
	if err != nil {
		log.Print(err)
	}
	fileScanner = bufio.NewScanner(file)
	Linhas = []string{}
	for fileScanner.Scan() {
		Linhas = append(Linhas, fileScanner.Text())
	}
	Infos = append(Infos, Linhas[0])
	fmt.Println(Infos[2])
	file.Close()

	file, err = os.Open("diskutil.out")
	if err != nil {
		log.Print(err)
	}
	fileScanner = bufio.NewScanner(file)
	Linhas = []string{}
	for fileScanner.Scan() {
		Linhas = append(Linhas, fileScanner.Text())
	}
	Infos = append(Infos, Linhas[0])
	fmt.Println(Infos[3])
	file.Close()

	file, err = os.Open("sw_vers.out")
	if err != nil {
		log.Print(err)
	}
	fileScanner = bufio.NewScanner(file)
	Linhas = []string{}
	for fileScanner.Scan() {
		Linhas = append(Linhas, fileScanner.Text())
	}
	Infos = append(Infos, Linhas[0])
	fmt.Println(Infos[4])
	file.Close()

	//Apagando Junk Files
	cmd := exec.Command("rm", "uname.out", "sysctl.out", "hostinfo.out", "diskutil.out", "sw_vers.out")
	stdout, _ := cmd.Output()
	fmt.Println(string(stdout))
}
