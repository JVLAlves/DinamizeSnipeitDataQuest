package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"runtime"
	"strconv"
	"strings"
	"sync"

	"github.com/JVLAlves/DinamizeSnipeitDataQuest/DataMission/Linux"
	"github.com/JVLAlves/DinamizeSnipeitDataQuest/DataMission/MacOS"
	"github.com/JVLAlves/DinamizeSnipeitDataQuest/DataMission/Windows"
	functions "github.com/JVLAlves/DinamizeSnipeitDataQuest/Utilities/Functions"
	snipe "github.com/JVLAlves/DinamizeSnipeitDataQuest/Utilities/SnipeMethods"
)

//IP do inventário Snipeit
var IP string = "10.20.1.79:8001" //IP do Invetário de TESTE

//Função de execução do programa em MacOS
func forMacOs(f *os.File) {

	//Criando Arquivos via Goroutines
	wg := &sync.WaitGroup{}
	wg.Add(5)
	go MacOS.Create(wg, "uname", "-n")
	go MacOS.Create(wg, "sysctl", "-a |grep machdep.cpu.brand_string |awk '{print $2,$3,$4}'")
	go MacOS.Create(wg, "hostinfo", "|grep memory |awk '{print $4,$5}'")
	go MacOS.Create(wg, "diskutil", "list |grep disk0s2 | awk '{print $5,$6}'")
	go MacOS.Create(wg, "sw_vers", "-productVersion")
	wg.Wait()

	//Realiza o processo de coleta de dados do Sistema MacOS e retorna as informações em um array Infos
	MacOS.Running()

	//Variavel de Contrato
	var mac snipe.CollectionT = snipe.CollectionT{}

	//Populando Struct
	mac.SnipeitCPU11 = MacOS.Infos[1]
	mac.SnipeitHostname10 = MacOS.Infos[0]
	mac.Name = MacOS.Infos[0]

	//Passando Regex antes de popular informação de Memória
	Memrexed := functions.RegexThis(`(^ ?\d{1,3}[,.]?\d*)`, MacOS.Infos[2])
	//Convertendo response de string para float
	Memnum, _ := strconv.ParseFloat(Memrexed, 64)
	//Arredondando valor númerico da variável
	Memround := math.Round(Memnum)
	//Populando campo de memória com o valor tratado
	mac.SnipeitMema3Ria7 = strconv.Itoa(int(Memround)) + "GB"

	//Passando Regex antes de popular informação de HD
	HDrexed := functions.RegexThis(`(^ ?\d{1,3}[,.]?\d*)`, MacOS.Infos[3])
	//Convertendo response de string para float
	HDnum, _ := strconv.ParseFloat(HDrexed, 64)
	//Arredondando valor númerico da variável
	HDround := math.Round(HDnum)
	//Populando campo de HD com o valor tratado
	mac.SnipeitHd9 = strconv.Itoa(int(HDround)) + "GB"

	//Passando Regex antes de popular informação de Asset Tag
	mac.AssetTag = functions.RegexThis(`\d`, MacOS.Infos[0])
	//Caso não haja digitos no campo HOSTNAME (Fonte do Asset Tag), o retorno do sistema é um Asset Tag Default (NO ASSET TAG)
	if mac.AssetTag == "" {
		mac.AssetTag = "No Asset Tag"
		fmt.Fprintf(f, "Nenhum Asset Tag foi definido, pois nenhuma sequência numérica foi encontrada no HOSTNAME: %v", MacOS.Infos[0])

	}

	//Passando Regex antes de popular informação de Sistema Operacional
	SOregexed := functions.RegexThis(`(^\d{2}\.\d+)`, MacOS.Infos[4])
	//Convertendo response de string para float
	numSO, err := strconv.ParseFloat(SOregexed, 64)
	//Tratando erro
	if err != nil {
		log.Fatalf("Erro na conversão do S.O. para float")
	}

	//Arredondamento da Versão MACOSX
	if numSO >= 11.4 && numSO < 12.0 {
		mac.SnipeitSo8 = "11.4"
	}

	//Alternando Versão Númerica (RETIRADA DO SITEMA) para Versão Nominal (DEFINIDA PELA APPLE INC.)
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
	case "11.4":
		mac.SnipeitSo8 = "MacOs Big Sur"
	default:
		mac.SnipeitSo8 = "MacOs"
	}

	//Entrada Default
	var IDmodelo *string = &mac.ModelID
	var IDstatus *string = &mac.StatusID
	var modeloAtivo *string = &mac.SnipeitModel12

	//identificando o Modelo
	*IDmodelo = "8"
	*modeloAtivo = "DNZ-COMPUTER"

	//Status ID
	*IDstatus = "7"

	//Resumo gráfico das informações coletadas.
	fmt.Printf("NOME DO DISPOSITIVO: %v\n", mac.Name)
	fmt.Printf("ASSET TAG: %v\n", mac.AssetTag)
	fmt.Printf("TIPO DE ATIVO: %v\n", mac.ModelID)
	fmt.Printf("MODELO DO ATIVO: %v\n", mac.SnipeitModel12)
	fmt.Printf("STATUS: %v\n\n", mac.StatusID)
	fmt.Printf("DESCRIÇÃO DO ATIVO\n")
	fmt.Printf("HOSTNAME: %v\n", mac.SnipeitHostname10)
	fmt.Printf("S.O.: %v\n", mac.SnipeitSo8)
	fmt.Printf("CPU: %v\n", mac.SnipeitCPU11)
	fmt.Printf("MEMORIA RAM: %v\n", mac.SnipeitMema3Ria7)
	fmt.Printf("DISCO: %v\n\n", mac.SnipeitHd9)

	//Verificando a existência de um ativo semelhante no inventário Snipe it
	if snipe.Verifybytag(mac.AssetTag, IP) {
		fmt.Fprintln(f, "Os dados do Ativo Criado não constam no sistema.")
		fmt.Println("Enviando Ativo para o Snipeit ")

		//Caso o Ativo não exista no sistema, as informações são enviadas para tal.
		snipe.PostSnipe(mac, IP, f)
	} else {
		//caso já exista, o programa procura por disparidades.
		//log.Println("Um Ativo semelhante foi encontrado no sistema.")
		fmt.Print("Asset Tag idêntico encontrado. Iniciando análise de disparidades")
		patch, boolean := snipe.Getbytag(IP, mac.AssetTag, mac, f)
		if boolean {
			//Caso haja disparidades, o processo de PATCH é iniciado.
			fmt.Println("\nPATCH necessário.")
			fmt.Println("\nExecutando PATCH RESQUEST.")

			id := snipe.Getidbytag(mac.AssetTag, IP)
			snipe.Patchbyid(id, IP, patch)

		} else {
			//Caso não haja disparidades... Nada acontece.
			_, _ = fmt.Fprintf(f, "")
			fmt.Fprintln(f, "\nSem alterações")
		}
	}
}

//Função de execução do programa em Windows
func forWindows(f *os.File) {

	//Realiza o processo de coleta de dados do Sistema Windows e retorna as informações em um array Infos
	Windows.MainProgram()

	//Variavel de Contrato
	var win snipe.CollectionT = snipe.CollectionT{}

	//Populando Struct
	win.SnipeitCPU11 = Windows.Infos[3]
	win.SnipeitMema3Ria7 = Windows.Infos[2]
	win.SnipeitSo8 = Windows.Infos[1]
	win.SnipeitHostname10 = Windows.Infos[0]
	win.Name = Windows.Infos[0]

	//Passando Regex antes de popular informação de Memória (COLETA: Primeiros três digitos com espaço em branco)
	Memrexed := functions.RegexThis(`^[ ]*\d{1,3}`, Windows.Infos[2])
	//Passando Regex antes de popular informação de Memória (COLETA: Somente os digitos)
	Memrexed = functions.RegexThis(`\d`, Memrexed)
	//Convertendo response de string para float
	Memnum, _ := strconv.ParseFloat(Memrexed, 64)
	//Arredondando valor númerico da variável
	Memround := math.Round(Memnum)
	//Populando campo de memória com o valor tratado
	win.SnipeitMema3Ria7 = strconv.Itoa(int(Memround)) + "GB"

	//Passando Regex antes de popular informação de HD
	HDrexed := functions.RegexThis(`^[ ]*\d{1,3}`, Windows.Infos[4])
	//Convertendo response de string para float
	HDnum, _ := strconv.ParseFloat(HDrexed, 64)
	//Arredondando valor númerico da variável
	HDround := math.Round(HDnum)
	//Populando campo de HD com o valor tratado
	win.SnipeitHd9 = strconv.Itoa(int(HDround)) + "GB"

	//Passando Regex antes de popular informação de Asset Tag
	win.AssetTag = functions.RegexThis(`\d`, Windows.Infos[0])
	//Caso não haja digitos no campo HOSTNAME (Fonte do Asset Tag), o retorno do sistema é um Asset Tag Default (NO ASSET TAG)
	if win.AssetTag == "" {
		win.AssetTag = "No Asset Tag"
		log.Printf("Nenhum Asset Tag foi defino, pois nenhuma sequência numérica foi encontrada no HOSTNAME: %v", Windows.Infos[0])

	}

	//Entrada Default
	var IDmodelo *string = &win.ModelID
	var IDstatus *string = &win.StatusID
	var modeloAtivo *string = &win.SnipeitModel12

	//identificando o Modelo
	*IDmodelo = "8"
	*modeloAtivo = "DNZ-COMPUTER"
	//Status ID
	*IDstatus = "7"

	//Resumo gráfico das informações coletadas.
	fmt.Printf("NOME DO DISPOSITIVO: %v\n", win.Name)
	fmt.Printf("ASSET TAG: %v\n", win.AssetTag)
	fmt.Printf("TIPO DE ATIVO: %v\n", win.ModelID)
	fmt.Printf("MODELO DO ATIVO: %v\n", win.SnipeitModel12)
	fmt.Printf("STATUS: %v\n\n", win.StatusID)
	fmt.Printf("DESCRIÇÃO DO ATIVO\n")
	fmt.Printf("HOSTNAME: %v\n", win.SnipeitHostname10)
	fmt.Printf("S.O.: %v\n", win.SnipeitSo8)
	fmt.Printf("CPU: %v\n", win.SnipeitCPU11)
	fmt.Printf("MEMORIA RAM: %v\n", win.SnipeitMema3Ria7)
	fmt.Printf("DISCO: %v\n\n", win.SnipeitHd9)

	//Verificando a existência de um ativo semelhante no inventário Snipe it
	if snipe.Verifybytag(win.AssetTag, IP) {
		fmt.Fprintln(f, "Os dados do Ativo Criado não constam no sistema.")
		fmt.Println("Enviando Ativo para o Snipeit ")

		//Caso o Ativo não exista no sistema, as informações são enviadas para tal.
		snipe.PostSnipe(win, IP, f)

		log.Println("Ativo Criado enviado para o sistema.")

	} else {
		//caso já exista, o programa procura por disparidades.
		//log.Println("Um Ativo semelhante foi encontrado no sistema.")
		fmt.Print("Asset Tag idêntico encontrado. Iniciando análise de disparidades")
		patch, boolean := snipe.Getbytag(IP, win.AssetTag, win, f)
		if boolean {
			//Caso haja disparidades, o processo de PATCH é iniciado.
			fmt.Println("\nPATCH necessário.")
			fmt.Println("\nExecutando PATCH RESQUEST.")

			id := snipe.Getidbytag(win.AssetTag, IP)
			snipe.Patchbyid(id, IP, patch)

		} else {
			//Caso não haja disparidades... Nada acontece.
			fmt.Fprintln(f, "\nSem alterações")
		}
	}

}

//Função de execução do programa em Linux
func forLinux(f *os.File) {

	//Realiza o processo de coleta de dados do Sistema Linux e retorna as informações em um array Infos
	Linux.MainProgram()

	//Variavel de Contrato
	var lin snipe.CollectionT = snipe.CollectionT{}

	//Populando Struct
	lin.SnipeitCPU11 = Linux.Infos[0]
	lin.SnipeitSo8 = Linux.Infos[2]
	lin.SnipeitHostname10 = Linux.Infos[3]
	lin.Name = Linux.Infos[3]

	//Passando Regex antes de popular informação de HD (COLETA: Número com vírgula)
	interHD := functions.RegexThis(`(^ ?\d{1,3}[,.]?\d*)`, Linux.Infos[4])
	//Separação do result
	indexHD := strings.Split(interHD, ",")
	//Integração do result utilizando ponto (Padrão para conversão)
	interHD = strings.Join(indexHD, ".")
	//Convertendo response de string para float
	HDnum, _ := strconv.ParseFloat(interHD, 64)
	//Arredondando valor númerico da variável
	HDround := math.Round(HDnum)
	//Populando campo de HD com o valor tratado
	lin.SnipeitHd9 = strconv.Itoa(int(HDround)) + "GB"

	//Passando Regex antes de popular informação de Memória
	intermem := functions.RegexThis(`(^ ?\d{1,3}[,.]?\d*)`, Linux.Infos[1])
	//Convertendo response de string para float
	Memnum, _ := strconv.ParseFloat(intermem, 64)
	//Arredondando valor númerico da variável
	Memround := math.Round(Memnum)
	//Populando campo de memória com o valor tratado
	lin.SnipeitMema3Ria7 = strconv.Itoa(int(Memround)) + "GB"

	//Passando Regex antes de popular informação de Asset Tag
	lin.AssetTag = functions.RegexThis(`\d`, Linux.Infos[3])
	//Caso não haja digitos no campo HOSTNAME (Fonte do Asset Tag), o retorno do sistema é um Asset Tag Default (NO ASSET TAG)
	if lin.AssetTag == "" {
		lin.AssetTag = "No Asset Tag"
		log.Printf("Nenhum Asset Tag foi defino, pois nenhuma sequência numérica foi encontrada no HOSTNAME: %v", Linux.Infos[0])

	}

	//Entrada Default
	var IDmodelo *string = &lin.ModelID
	var IDstatus *string = &lin.StatusID
	var modeloAtivo *string = &lin.SnipeitModel12

	//identificando o Modelo
	*IDmodelo = "8"
	*modeloAtivo = "DNZ-COMPUTER"
	//Status ID
	*IDstatus = "7"

	//Resumo gráfico das informações coletadas.
	fmt.Printf("NOME DO DISPOSITIVO: %v\n", lin.Name)
	fmt.Printf("ASSET TAG: %v\n", lin.AssetTag)
	fmt.Printf("TIPO DE ATIVO: %v\n", lin.ModelID)
	fmt.Printf("MODELO DO ATIVO: %v\n", lin.SnipeitModel12)
	fmt.Printf("STATUS: %v\n\n", lin.StatusID)
	fmt.Printf("DESCRIÇÃO DO ATIVO\n")
	fmt.Printf("HOSTNAME: %v\n", lin.SnipeitHostname10)
	fmt.Printf("S.O.: %v\n", lin.SnipeitSo8)
	fmt.Printf("CPU: %v\n", lin.SnipeitCPU11)
	fmt.Printf("MEMORIA RAM: %v\n", lin.SnipeitMema3Ria7)
	fmt.Printf("DISCO: %v\n\n", lin.SnipeitHd9)

	//Verificando a existência de um ativo semelhante no inventário Snipe it
	if snipe.Verifybytag(lin.AssetTag, IP) {
		fmt.Fprintln(f, "Os dados do Ativo Criado não constam no sistema.")
		fmt.Println("Enviando Ativo para o Snipeit ")

		//Caso o Ativo não exista no sistema, as informações são enviadas para tal.
		snipe.PostSnipe(lin, IP, f)

	} else {
		//caso já exista, o programa procura por disparidades.
		//log.Println("Um Ativo semelhante foi encontrado no sistema.")
		fmt.Print("Asset Tag idêntico encontrado. Iniciando análise de disparidades")
		patch, boolean := snipe.Getbytag(IP, lin.AssetTag, lin, f)
		if boolean {
			//Caso haja disparidades, o processo de PATCH é iniciado.
			fmt.Println("\nPATCH necessário.")
			fmt.Println("\nExecutando PATCH RESQUEST.")

			id := snipe.Getidbytag(lin.AssetTag, IP)
			snipe.Patchbyid(id, IP, patch)

		} else {
			//Caso não haja disparidades... Nada acontece.
			fmt.Fprintln(f, "\nSem alterações")
		}
	}

}

//função Principal do programa
func main() {

	//Cria tanto a pasta para logs quanto o arquivo inicial de logs
	f := functions.ActiveLogs()

	//Log de inicialização
	log.Printf("Inicio de execução.")

	//Identificando sistema operacional
	switch runtime.GOOS {
	case "darwin":
		forMacOs(f)
	case "linux":
		forLinux(f)

	case "windows":
		forWindows(f)
	default:
		log.Fatalf("Erro em econtrar o Sistema Operacional")
	}

	//mensagem de encerramento
	fmt.Println("\n\nObrigado pela paciência! (FIM)")
	log.Printf("Fim de execução.")
}
