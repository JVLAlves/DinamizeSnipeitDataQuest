package windows

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
var linhas = []string{}
var infos = []string{}

//var linhasEditadas = []string{}
var abc = []string{}

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

func main() {
	posh := New()

	//Aplicanddo os comandos literais que serão executados no powershell
	stdout, stderr, err := posh.Execute("New-Item -path \"$env:userprofile\" -Name \"logfileees\" -ItemType \"directory\"")
	stdout, stderr, err = posh.Execute("New-Item -path \"logfileees\" -Name \"logfileees.txt\" -ItemType \"file\"")
	stdout, stderr, err = posh.Execute("Systeminfo > \"$env:userprofile\\logfileees\\Informacoes_Do_Sistema.txt\"")
	stdout, stderr, err = posh.Execute("Get-WmiObject -Class Win32_Processor -ComputerName . | Select-Object -Property \"name\" > \"$env:userprofile\\logfileees\\cpu.txt\"")
	stdout, stderr, err = posh.Execute("get-WMIobject Win32_LogicalDisk -Filter \"DeviceID = 'C:'\" | Select-Object -Property \"Size\" > \"$env:userprofile\\logfileees\\disk.txt\"")
	stdout, stderr, err = posh.Execute("(Get-Content -path \"$env:userprofile\\logfileees\\cpu.txt\" -TotalCount 6)[3] | Add-Content -path \"$env:userprofile\\logfileees\\logfileees.txt\"")
	stdout, stderr, err = posh.Execute("(Get-Content -path \"$env:userprofile\\logfileees\\disk.txt\" -TotalCount 6)[3] | Add-Content -path \"$env:userprofile\\logfileees\\logfileees.txt\"")
	stdout, stderr, err = posh.Execute("Add-Content -Path \"$env:userprofile\\logfileees\\logfileees.txt\" -value (Select-String -Path \"$env:userprofile\\logfileees\\Informacoes_Do_Sistema.txt\" -Pattern \"Nome do host:\",\"Nome do sistema operacional:\",\"Memória física total:\")")
	fmt.Println(stdout)
	fmt.Println(stderr)

	if err != nil {
		fmt.Println(err)

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

	//Limpa o Array das linhas
	linhas = []string{}

	//Lendo linha a linha
	for fileScanner.Scan() {
		linhas = append(linhas, fileScanner.Text())

	}
	// adicionando informação encontrada no arquivo teste a variável
	//fmt.Println(linhas)
	infos = append(infos, linhas[3])

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

	//Limpa o Array das linhas
	linhas = []string{}

	//Lendo linha a linha
	for fileScanner.Scan() {
		linhas = append(linhas, fileScanner.Text())

	}
	// adicionando informação encontrada no arquivo teste a variável

	infos = append(infos, linhas[3])

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

	//Limpa o Array das linhas

	//Lendo linha a linha
	for fileScanner.Scan() {
		//fmt.Println(fileScanner.Text())
		linhas = append(linhas, fileScanner.Text())
		//fmt.Println("Print lendo logfiles linha a linha\n\n\n", linhas)

	}
	infos = append(infos, linhas[2])
	infos = append(infos, linhas[3])
	infos = append(infos, linhas[4])
	//Tratando o ocasoional erro da leitura do arquivo
	if err := fileScanner.Err(); err != nil {
		log.Fatalf("Error while reading file: %s", err)
	}

	//fechando o arquivo lido
	file.Close()

	re := regexp.MustCompile(`[ ]{2}[[:alnum:]]+[.\-|[:alnum:])]+[[:print:]]+`)
	//fmt.Println("\n\n\n", infos)

	infos = []string{}
	for i := 0; i < len(linhas); i++ {

		abc = re.FindAllString(linhas[i], -1)
		justString := strings.Join(abc, "")
		if justString != "" {
			infos = append(infos, justString)
		}
		justString = ""
		//fmt.Println("Print linha a linha da variável linha \n", linhas[i])
	}
	infos = append(infos, linhas[7])

	infos = append(infos, linhas[8])

	fmt.Println("print inteiro da infos\n\n\n", infos)
	//fmt.Println("print inteiro da infos\n\n\n", len(infos))

	stdout, stderr, err = posh.Execute("Remove-Item -path $env:userprofile\\logfileees\\logfileees.txt")
	stdout, stderr, err = posh.Execute("Remove-Item -path $env:userprofile\\logfileees\\disk.txt")
	stdout, stderr, err = posh.Execute("Remove-Item -path $env:userprofile\\logfileees\\cpu.txt")
	stdout, stderr, err = posh.Execute("Remove-Item -path $env:userprofile\\logfileees\\Informacoes_Do_Sistema.txt")
	stdout, stderr, err = posh.Execute("Remove-Item -path $env:userprofile\\logfileees\\logfileees.txt")
	stdout, stderr, err = posh.Execute("Remove-Item -path $env:userprofile\\logfileees")

}
