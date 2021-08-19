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

	functions "github.com/JVLAlves/DinamizeSnipeitDataQuest/Utilities/Functions"
)

//Modelo para execução Powershell
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
	stdout, _, _ := posh.Execute("Get-Location")
	caminho := functions.RegexThis(`[(C:\\Users\)].+[[:alnum:]]?[^(\r\n)]`, stdout)
	fmt.Printf("%#v", caminho)
	log.Fatal()
	caminho = caminho[:len(caminho)-1]
	//caminho = ("\"" + caminho + "\"")
	fmt.Printf("%#v", caminho)
	log.Fatal()
	_, _, _ = posh.Execute("New-Item -path" + "\"" + caminho + "\"" + "-Name \"logfileees\" -ItemType \"directory\"")
	_, _, _ = posh.Execute("New-Item -path \"logfileees\" -Name \"logfileees.txt\" -ItemType \"file\"")
	_, _, _ = posh.Execute("Systeminfo > " + "\"" + caminho + "\\logfileees\\Informacoes_Do_Sistema.txt\"")
	_, _, _ = posh.Execute("Get-WmiObject -Class Win32_Processor -ComputerName . | Select-Object -Property \"name\" > " + "\"" + caminho + "\\logfileees\\cpu.txt\"")
	_, _, _ = posh.Execute("get-WMIobject Win32_LogicalDisk -Filter \"DeviceID = 'C:'\" | Select-Object -Property \"Size\" > " + "\"" + caminho + "\\logfileees\\disk.txt\"")
	_, _, _ = posh.Execute("(Get-Content -path " + "\"" + caminho + "\\logfileees\\cpu.txt\" -TotalCount 6)[3] | Add-Content -path " + "\"" + caminho + "\\logfileees\\logfileees.txt\"")
	_, _, _ = posh.Execute("(Get-Content -path " + "\"" + caminho + "\\logfileees\\disk.txt\" -TotalCount 6)[3] | Add-Content -path " + "\"" + caminho + "\\logfileees\\logfileees.txt\"")
	_, stderr, err := posh.Execute("Add-Content -Path " + "\"" + caminho + "\\logfileees\\logfileees.txt\" -value (Select-String -Path " + "\"" + caminho + "\\logfileees\\Informacoes_Do_Sistema.txt\" -Pattern \"Nome do host:\",\"Nome do sistema operacional:\",\"Memória física total:\")")
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

	//Limpa o Array das Linhas
	Linhas = []string{}

	//Lendo linha a linha
	for fileScanner.Scan() {
		Linhas = append(Linhas, fileScanner.Text())

	}
	// adicionando informação encontrada no arquivo teste a variável
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
	for fileScanner.Scan() {
		Linhas = append(Linhas, fileScanner.Text())

	}
	// adicionando informação encontrada no arquivo teste a variável

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
	for fileScanner.Scan() {
		Linhas = append(Linhas, fileScanner.Text())

	}
	Infos = append(Infos, Linhas[2])
	Infos = append(Infos, Linhas[3])
	Infos = append(Infos, Linhas[4])

	//Tratando o ocasoional erro da leitura do arquivo
	if err := fileScanner.Err(); err != nil {
		log.Fatalf("Error while reading file: %s", err)
	}

	//fechando o arquivo lido
	file.Close()

	re := regexp.MustCompile(`[ ]{2}[[:alnum:]]+[.\-|[:alnum:])]+[[:print:]]+`)

	Infos = []string{}
	for i := 0; i < len(Linhas); i++ {

		Abc = re.FindAllString(Linhas[i], -1)
		justString := strings.Join(Abc, "")
		if justString != "" {
			Infos = append(Infos, justString)
		}
		justString = ""
	}

	Infos = append(Infos, Linhas[7])

	Infos = append(Infos, Linhas[8])

	_, _, _ = posh.Execute("Remove-Item -path $env:userprofile\\logfileees\\logfileees.txt")
	_, _, _ = posh.Execute("Remove-Item -path $env:userprofile\\logfileees\\disk.txt")
	_, _, _ = posh.Execute("Remove-Item -path $env:userprofile\\logfileees\\cpu.txt")
	_, _, _ = posh.Execute("Remove-Item -path $env:userprofile\\logfileees\\Informacoes_Do_Sistema.txt")
	_, _, _ = posh.Execute("Remove-Item -path $env:userprofile\\logfileees\\logfileees.txt")
	_, _, _ = posh.Execute("Remove-Item -path $env:userprofile\\logfileees")

}
