package Windows

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

type PowerShell struct {
	powerShell string
}

//Iniciando as variáveis array
var Linhas = []string{}
var Infos = []string{}

//var linhasEditadas = []string{}
var Abc = []string{}

//invocando o PowerShell
func New() *PowerShell {
	ps, _ := exec.LookPath("powershell.exe")
	return &PowerShell{
		powerShell: ps,
	}
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

func MainProgram() {
	posh := New()

	//Aplicanddo os comandos literais que serão executados no powershell
	fmt.Println("Creating logfileees Directory")
	_, _, _ = posh.Execute("New-Item -path \"$env:userprofile\" -Name \"logfileees\" -ItemType \"directory\"")
	fmt.Println("Creating logfileees.txt")
	_, _, _ = posh.Execute("New-Item -path \"logfileees\" -Name \"logfileees.txt\" -ItemType \"file\"")
	fmt.Println("Colecting Systeminfos information and appending to Informacoes_Do_Sistema.txt")
	_, _, _ = posh.Execute("Systeminfo > \"$env:userprofile\\logfileees\\Informacoes_Do_Sistema.txt\"")
	fmt.Println("greping CPU")
	_, _, _ = posh.Execute("Get-WmiObject -Class Win32_Processor -ComputerName . | Select-Object -Property \"name\" > \"$env:userprofile\\logfileees\\cpu.txt\"")
	fmt.Println("greping Disk")
	_, _, _ = posh.Execute("get-WMIobject Win32_LogicalDisk -Filter \"DeviceID = 'C:'\" | Select-Object -Property \"Size\" > \"$env:userprofile\\logfileees\\disk.txt\"")
	fmt.Println("Appending CPU to logfileees.txt")
	_, _, _ = posh.Execute("(Get-Content -path \"$env:userprofile\\logfileees\\cpu.txt\" -TotalCount 6)[3] | Add-Content -path \"$env:userprofile\\logfileees\\logfileees.txt\"")
	fmt.Println("Appending Disk to logfileees.txt")
	_, _, _ = posh.Execute("(Get-Content -path \"$env:userprofile\\logfileees\\disk.txt\" -TotalCount 6)[3] | Add-Content -path \"$env:userprofile\\logfileees\\logfileees.txt\"")
	fmt.Println("Adding some other content")
	stdout, stderr, err := posh.Execute("Add-Content -Path \"$env:userprofile\\logfileees\\logfileees.txt\" -value (Select-String -Path \"$env:userprofile\\logfileees\\Informacoes_Do_Sistema.txt\" -Pattern \"Nome do host:\",\"Nome do sistema operacional:\",\"Memória física total:\")")
	fmt.Println(stdout)
	fmt.Println(stderr)

	if err != nil {
		fmt.Printf("THERE AREA THIS ERROR:\n%v\n", err)

	}
	var (
		caminhocpu  = "logfileees\\cpu.txt"
		caminhodisk = "logfileees\\cpu.txt"
		caminholog  = "logfileees\\logfileees.txt"
	)

	////////////////////////////////////////////ABRINDO ARQUIVO PARA LER A INFORMAÇÃO DE CPU//////////////////////////////////////////////////////////////////////////////////
	file, err := os.Open(caminhocpu)
	if err != nil {
		log.Fatalf("Error when opening file cpu: %s", err)
	}

	//Lendo o Arquivo CPU
	fileScanner := bufio.NewScanner(file)

	//Limpa o Array das Linhas
	Linhas = []string{}

	//Lendo linha a linha
	fmt.Println("Reading CPU line by line")
	for fileScanner.Scan() {
		Linhas = append(Linhas, fileScanner.Text())

	}
	// adicionando informação encontrada no arquivo teste a variável
	//fmt.Println(Linhas)
	fmt.Println("Appending Linhas from CPU to Infos")
	Infos = append(Infos, Linhas[3])

	//Tratando o ocasoional erro da leitura do arquivo
	if err := fileScanner.Err(); err != nil {
		log.Fatalf("Error while reading file disk: %s", err)
	}

	//fechando o arquivo lido
	file.Close()

	////////////////////////////////////////////ABRINDO ARQUIVO PARA LER A INFORMAÇÃO DE DISK///////////////////////////////////////////////
	file, err = os.Open(caminhodisk)
	if err != nil {
		log.Fatalf("Error when opening file disk: %s", err)
	}

	//Lendo o Arquivo DISK
	fileScanner = bufio.NewScanner(file)

	//Limpa o Array das Linhas
	Linhas = []string{}

	//Lendo linha a linha
	fmt.Println("Reading Disk line by line")
	for fileScanner.Scan() {
		Linhas = append(Linhas, fileScanner.Text())

	}
	// adicionando informação encontrada no arquivo teste a variável
	fmt.Println("Appending Linhas from Disk to Infos")
	Infos = append(Infos, Linhas[3])

	//Tratando o ocasoional erro da leitura do arquivo
	if err := fileScanner.Err(); err != nil {
		log.Fatalf("Error while reading file: %s", err)
	}

	//fechando o arquivo lido
	file.Close()

	////////////////////////////////////////////ABRINDO ARQUIVO PARA LER A INFORMAÇÃO RESTANTES/////////////////////////////////////////////////////
	//////////////////////////////////////////////////////////HOSTNAME///SO///MEMÓRIA//////////////////////////////////////////////////////////////

	// Abrindo o Arquivo teste
	file, err = os.Open(caminholog)
	if err != nil {
		log.Fatalf("Error when opening file others files: %s", err)
	}

	//Lendo o Arquivo LOGFILEEES
	fileScanner = bufio.NewScanner(file)

	//Limpa o Array das Linhas

	//Lendo linha a linha
	fmt.Println("Reading Others line by line")
	for fileScanner.Scan() {
		//fmt.Println(fileScanner.Text())
		Linhas = append(Linhas, fileScanner.Text())
		//fmt.Println("Print lendo logfiles linha a linha\n\n\n", Linhas)

	}
	fmt.Println("Appending Linhas from others to Infos")
	Infos = append(Infos, Linhas[2])
	Infos = append(Infos, Linhas[3])
	Infos = append(Infos, Linhas[4])
	//Tratando o ocasoional erro da leitura do arquivo
	if err := fileScanner.Err(); err != nil {
		log.Fatalf("Error while reading file: %s", err)
	}

	//fechando o arquivo lido
	file.Close()
	fmt.Println("Making Regex")
	re := regexp.MustCompile(`[ ]{2}[[:alnum:]]+[.\-|[:alnum:])]+[[:print:]]+`)
	//fmt.Println("\n\n\n", Infos)

	Infos = []string{}
	fmt.Println("Regexing")
	for i := 0; i < len(Linhas); i++ {

		Abc = re.FindAllString(Linhas[i], -1)
		justString := strings.Join(Abc, "")
		if justString != "" {
			Infos = append(Infos, justString)
		}
		justString = ""
		//fmt.Println("Print linha a linha da variável linha \n", Linhas[i])
	}
	fmt.Println("Final appends 1")
	Infos = append(Infos, Linhas[7])
	fmt.Println("Final appends 2")
	Infos = append(Infos, Linhas[8])

	fmt.Println("print inteiro da Infos\n\n\n", Infos)
	//fmt.Println("print inteiro da Infos\n\n\n", len(Infos))

	fmt.Println("Removing archives")
	_, _, _ = posh.Execute("Remove-Item -path $env:userprofile\\logfileees\\logfileees.txt")
	_, _, _ = posh.Execute("Remove-Item -path $env:userprofile\\logfileees\\disk.txt")
	_, _, _ = posh.Execute("Remove-Item -path $env:userprofile\\logfileees\\cpu.txt")
	_, _, _ = posh.Execute("Remove-Item -path $env:userprofile\\logfileees\\Informacoes_Do_Sistema.txt")
	_, _, _ = posh.Execute("Remove-Item -path $env:userprofile\\logfileees\\logfileees.txt")
	_, _, _ = posh.Execute("Remove-Item -path $env:userprofile\\logfileees")

}
