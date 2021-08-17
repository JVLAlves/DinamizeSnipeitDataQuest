package Linux

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

//Variáveis de armazenamento dos dados da máquina
var Linhas = []string{}
var Infos = []string{}

func MainProgram() {
	// Abrindo o Arquivo CPU
	file, err := os.Open("/proc/cpuinfo")
	if err != nil {
		log.Fatalf("Error when opening file: %s", err)
	}

	//Lendo o Arquivo CPU
	fileScanner := bufio.NewScanner(file)

	//Lendo linha a linha
	for fileScanner.Scan() {
		fmt.Println(fileScanner.Text())
		Linhas = append(Linhas, fileScanner.Text())

	}
	// adicionando informação encontrada no arquivo CPU a variável
	var infostemp []string
	infostemp = append(infostemp, Linhas[4])

	re := regexp.MustCompile(`(Intel).+`)
	for i := 0; i < len(infostemp); i++ {
		Abc := re.FindAllString(infostemp[i], -1)
		justString := strings.Join(Abc, "")
		if justString != "" {
			Infos = append(Infos, justString)
		}
		justString = ""
	}

	//Tratando o ocasoional erro da leitura do arquivo
	if err := fileScanner.Err(); err != nil {
		log.Fatalf("Error while reading file: %s", err)
	}

	//fechando o arquivo lido
	file.Close()

	//Executa o comando Script para escrever a sessão do terminal em arquivo txt (Tamanho do disco)
	cmd := exec.Command("script", "-c", "free -h |grep Mem |awk '{print $2}'", "tamanhoDoHd.txt")
	_, _ = cmd.Output()

	// abrindo o arquiuvo criado "tamanhoDoHd.txt"
	file, err = os.Open("tamanhoDoHd.txt")

	//Tratando o ocasoional erro da leitura do arquivo
	if err != nil {
		log.Fatalf("Error when opening file: %s", err)
	}

	//Lendo o arquivo "tamanhoDoHd.txt"
	fileScanner = bufio.NewScanner(file)

	//Limpa o Array das Linhas
	Linhas = []string{}

	//Lendo linha a linha
	for fileScanner.Scan() {
		fmt.Println(fileScanner.Text())
		Linhas = append(Linhas, fileScanner.Text())

	}

	// adicionando informação encontrada no arquivo "tamanhoDoHd.txt" a variável
	Infos = append(Infos, Linhas[1])

	//Tratando o ocasoional erro da leitura do arquivo
	if err := fileScanner.Err(); err != nil {
		log.Fatalf("Error while reading file: %s", err)
	}

	//fechando o arquivo lido
	file.Close()

	//Executa o comando Script para escrever a sessão do terminal em arquivo txt (S.O.)
	cmd = exec.Command("script", "-c", "lsb_release -d |grep Description |awk '{print $2,$3,$4}'", "SO.txt")
	_, _ = cmd.Output()

	// abrindo o arquiuvo criado "S0.txt"
	file, err = os.Open("SO.txt")

	//Tratando o ocasoional erro da leitura do arquivo
	if err != nil {
		log.Fatalf("Error when opening file: %s", err)
	}

	//Lendo o arquivo "SO.txt"
	fileScanner = bufio.NewScanner(file)

	//Limpa o Array das Linhas
	Linhas = []string{}

	//Lendo linha a linha
	for fileScanner.Scan() {
		fmt.Println(fileScanner.Text())
		Linhas = append(Linhas, fileScanner.Text())

	}

	// adicionando informação encontrada no arquivo "SO.txt" a variável
	Infos = append(Infos, Linhas[1])

	//Tratando o ocasoional erro da leitura do arquivo
	if err := fileScanner.Err(); err != nil {
		log.Fatalf("Error while reading file: %s", err)
	}

	//fechando o arquivo lido
	file.Close()

	//Executa o comando Script para escrever a sessão do terminal em arquivo txt (Hostname)
	cmd = exec.Command("script", "-c", "hostname", "hostname.txt")
	_, _ = cmd.Output()

	// abrindo o arquiuvo criado "Hostname.txt"
	file, err = os.Open("hostname.txt")

	//Tratando o ocasoional erro da leitura do arquivo
	if err != nil {
		log.Fatalf("Error when opening file: %s", err)
	}

	//Lendo o arquivo "Hostname.txt"
	fileScanner = bufio.NewScanner(file)

	//Limpa o Array das Linhas
	Linhas = []string{}

	//Lendo linha a linha
	for fileScanner.Scan() {
		fmt.Println(fileScanner.Text())
		Linhas = append(Linhas, fileScanner.Text())

	}

	Infos = append(Infos, Linhas[1])

	// adicionando informação encontrada no arquivo "Hostname.txt" a variável
	if err := fileScanner.Err(); err != nil {
		log.Fatalf("Error while reading file: %s", err)
	}

	//fechando o arquivo lido
	file.Close()

	//Executa o comando Script para escrever a sessão do terminal em arquivo txt (Tamanho do Disco)
	cmd = exec.Command("script", "-c", "lsblk |grep disk |awk '{print $4}'", "tamanhoDoDisco.txt")
	_, _ = cmd.Output()

	// abrindo o arquiuvo criado "tamanhoDoDisco.txt"
	file, err = os.Open("tamanhoDoDisco.txt")

	//Tratando o ocasoional erro da leitura do arquivo
	if err != nil {
		log.Fatalf("Error when opening file: %s", err)
	}

	fileScanner = bufio.NewScanner(file)
	Linhas = []string{}

	//Lendo linha a linha
	for fileScanner.Scan() {
		fmt.Println(fileScanner.Text())
		Linhas = append(Linhas, fileScanner.Text())

	}

	// adicionando informação encontrada no arquivo "tamanhoDoDisco.txt" a variável
	Infos = append(Infos, Linhas[1])

	//Tratando o ocasoional erro da leitura do arquivo
	if err := fileScanner.Err(); err != nil {
		log.Fatalf("Error while reading file: %s", err)
	}

	//fechando o arquivo lido
	file.Close()

	cmd = exec.Command("rm", "tamanhoDoHd.txt", "SO.txt", "hostname.txt", "tamanhoDoDisco.txt")
	_, _ = cmd.Output()

	cmd = exec.Command("rm", "tamanhoDoHd.txt", "SO.txt", "hostname.txt", "tamanhoDoDisco.txt")
	_, _ = cmd.Output()
}
