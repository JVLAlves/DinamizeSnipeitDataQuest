package main

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

//Definindo Tipo para popular com as informações do computador.
type MacOSt struct {
	SnipeitSo8        string `json:"_snipeit_so_8"`
	SnipeitModel12    string `json:"_snipeit_modelo_12"`
	SnipeitHostname10 string `json:"_snipeit_hostname_10"`
	SnipeitHd9        string `json:"_snipeit_hd_9"`
	SnipeitCPU11      string `json:"_snipeit_cpu_11"`
	SnipeitMema3Ria7  string `json:"_snipeit_mema3ria_7"`
}

//Lista para leitura linha a linha
var linhas []string

//Informações armazenadas
var infos []string

//Cria arquivos com as informações retiradas do computador via Terminal
func create(wg *sync.WaitGroup, command string, args string) {

	outFile, err := os.OpenFile(command+".out", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0755)
	defer outFile.Close()
	if err != nil {
		fmt.Println("error creating file", err)
	}

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
func running() {

	file, err := os.Open("uname.out")
	if err != nil {
		log.Print(err)
	}
	fileScanner := bufio.NewScanner(file)
	linhas = []string{}

	for fileScanner.Scan() {
		linhas = append(linhas, fileScanner.Text())
		if fileScanner.Err() != nil {
			log.Fatalf("Erro SCAN: %v", fileScanner.Err().Error())
		}
	}
	infos = append(infos, linhas[0])
	file.Close()

	file, err = os.Open("sysctl.out")
	if err != nil {
		log.Print(err)
	}
	fileScanner = bufio.NewScanner(file)
	linhas = []string{}
	for fileScanner.Scan() {
		linhas = append(linhas, fileScanner.Text())
	}
	infos = append(infos, linhas[0])
	fmt.Println(infos[1])
	file.Close()

	file, err = os.Open("hostinfo.out")
	if err != nil {
		log.Print(err)
	}
	fileScanner = bufio.NewScanner(file)
	linhas = []string{}
	for fileScanner.Scan() {
		linhas = append(linhas, fileScanner.Text())
	}
	infos = append(infos, linhas[0])
	fmt.Println(infos[2])
	file.Close()

	file, err = os.Open("diskutil.out")
	if err != nil {
		log.Print(err)
	}
	fileScanner = bufio.NewScanner(file)
	linhas = []string{}
	for fileScanner.Scan() {
		linhas = append(linhas, fileScanner.Text())
	}
	infos = append(infos, linhas[0])
	fmt.Println(infos[3])
	file.Close()

	file, err = os.Open("sw_vers.out")
	if err != nil {
		log.Print(err)
	}
	fileScanner = bufio.NewScanner(file)
	linhas = []string{}
	for fileScanner.Scan() {
		linhas = append(linhas, fileScanner.Text())
	}
	infos = append(infos, linhas[0])
	fmt.Println(infos[4])
	file.Close()

	return
}

func main() {

	//Criando Arquivos via Goroutines
	wg := &sync.WaitGroup{}

	wg.Add(5)
	go create(wg, "uname", "-n")
	go create(wg, "sysctl", "-a |grep machdep.cpu.brand_string |awk '{print $2,$3,$4}'")
	go create(wg, "hostinfo", "|grep memory |awk '{print $4,$5}'")
	go create(wg, "diskutil", "list |grep disk0s2 | awk '{print $5,$6}'")
	go create(wg, "sw_vers", "-productVersion")
	wg.Wait()

	running()

	//Verificação das informações "Appendadas"
	fmt.Println(infos)

	var mac MacOSt = MacOSt{}

	//Populando Struct MacOSt
	mac.SnipeitCPU11 = infos[1]
	mac.SnipeitMema3Ria7 = infos[2]
	mac.SnipeitSo8 = infos[4]
	mac.SnipeitHostname10 = infos[0]
	mac.SnipeitHd9 = infos[3]

	//Alternando Versão Númerica para Versão Nominal
	switch mac.SnipeitSo8 {

	case "10.7":
		mac.SnipeitSo8 = "MacOs Lion"
	case "10.8":
		mac.SnipeitSo8 = "MacOs Mountain Lion"
	case "10.9":
		mac.SnipeitSo8 = "MacOs Mavericks"
	case "10.10":
		mac.SnipeitSo8 = "MacOs Yosemite"
	case "10.11":
		mac.SnipeitSo8 = "MacOs El Capitan"
	case "10.12":
		mac.SnipeitSo8 = "MacOs Sierra"
	case "10.13":
		mac.SnipeitSo8 = "MacOs High Sierra"
	case "10.14":
		mac.SnipeitSo8 = "MacOs Mojave"
	case "10.15":
		mac.SnipeitSo8 = "MacOs Catalina"
	case "11.0":
		mac.SnipeitSo8 = "MacOs Big Sur"
	default:
		mac.SnipeitSo8 = "MacOs"
	}

	//Somente alguns prints; Sem utilidade pratica para o código.
	fmt.Printf("\nHOSTNAME: %v\n", mac.SnipeitHostname10)
	fmt.Printf("S.O.: %v\n", mac.SnipeitSo8)
	fmt.Printf("CPU: %v\n", mac.SnipeitCPU11)
	fmt.Printf("MEMORIA RAM: %v\n", mac.SnipeitMema3Ria7)
	fmt.Printf("DISCO: %v\n", mac.SnipeitHd9)

}
