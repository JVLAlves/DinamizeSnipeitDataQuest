package main

import (
	"fmt"
	"sync"

	"github.com/JVLAlves/DinamizeSnipeitDataQuest/DataQuest/MacOS"
)

func forMacOs() {

	//Criando Arquivos via Goroutines
	wg := &sync.WaitGroup{}
	wg.Add(5)
	go MacOS.Create(wg, "uname", "-n")
	go MacOS.Create(wg, "sysctl", "-a |grep machdep.cpu.brand_string |awk '{print $2,$3,$4}'")
	go MacOS.Create(wg, "hostinfo", "|grep memory |awk '{print $4,$5}'")
	go MacOS.Create(wg, "diskutil", "list |grep disk0s2 | awk '{print $5,$6}'")
	go MacOS.Create(wg, "sw_vers", "-productVersion")
	wg.Wait()

	// Lendo Arquivos
	MacOS.Running()

	//Verificação das informações "Appendadas"
	fmt.Println(MacOS.Infos)

	var mac MacOS.MacOSt = MacOS.MacOSt{}

	//Populando Struct MacOSt
	mac.SnipeitCPU11 = MacOS.Infos[1]
	mac.SnipeitMema3Ria7 = MacOS.Infos[2]
	mac.SnipeitSo8 = MacOS.Infos[4]
	mac.SnipeitHostname10 = MacOS.Infos[0]
	mac.SnipeitHd9 = MacOS.Infos[3]

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
	case "11.4":
		mac.SnipeitSo8 = "MacOs Big Sur"
	default:
		mac.SnipeitSo8 = "MacOs"
	}

	//Entrada Personalizada
	var IDmodelo *string = &mac.ModelID
	var IDstatus *string = &mac.StatusID
	var Assettag *string = &mac.AssetTag
	var name *string = &mac.Name
	var modeloAtivo *string = &mac.SnipeitModel12

	//Input Manual: Tipo de Ativo
	fmt.Println("Digite o Tipo de Ativo (Exemplo:Teclado/Desktop/MacBook/Leito_de_Cartão_Mesa): ")
	fmt.Scanf("%v", IDmodelo)

	//identificando o Modelo
	switch *IDmodelo {
	case "Leitor_de_Cartão_Mesa":
		*IDmodelo = "1"
	case "Leitor_de_Cartão_Porta":
		*IDmodelo = "2"
	case "Mouse_Sem_Fio":
		*IDmodelo = "3"
	case "Roteador":
		*IDmodelo = "4"
	case "Roteador_Wireless":
		*IDmodelo = "5"
	case "Notebook":
		*IDmodelo = "6"
	case "Celular_Samsumg":
		*IDmodelo = "7"
	case "Desktop":
		*IDmodelo = "8"
	case "Vostro":
		*IDmodelo = "9"
	case "Bravia":
		*IDmodelo = "10"
	case "Default":
		*IDmodelo = "11"
	case "DP720":
		*IDmodelo = "12"
	case "Telefone_Grandstream":
		*IDmodelo = "13"
	case "NP350":
		*IDmodelo = "14"
	case "Samsung_Default":
		*IDmodelo = "15"
	case "NP530":
		*IDmodelo = "16"
	case "NP370":
		*IDmodelo = "17"
	case "Ideapad":
		*IDmodelo = "18"
	case "P 250":
		*IDmodelo = "19"
	case "Inspirion":
		*IDmodelo = "20"
	case "Asus_X":
		*IDmodelo = "21"
	case "MacBook":
		*IDmodelo = "22"
	case "SMS":
		*IDmodelo = "23"
	case "OfficeJet":
		*IDmodelo = "24"
	case "LaserJet":
		*IDmodelo = "25"
	case "Asus":
		*IDmodelo = "26"
	case "D11":
		*IDmodelo = "27"
	case "XPS":
		*IDmodelo = "28"
	case "C3":
		*IDmodelo = "29"
	case "Multilaser_Desk":
		*IDmodelo = "30"
	case "Zelman_Desk":
		*IDmodelo = "31"
	case "TV_LG":
		*IDmodelo = "32"
	case "Braviaa":
		*IDmodelo = "33"
	case "AOC":
		*IDmodelo = "34"
	case "Dell":
		*IDmodelo = "35"
	case "Flantron":
		*IDmodelo = "36"
	case "Samsung":
		*IDmodelo = "37"
	default:
		*IDmodelo = "11"
	}

	//Input Manual: Status
	fmt.Println("Digite o Status do Ativo (Padrão: Disponível): ")
	fmt.Scanf("%v", IDstatus)

	//Identificando o Status
	switch *IDstatus {
	case "Disponivel":
		*IDstatus = "5"
	case "Indisponivel":
		*IDstatus = "6"
	case "Ocupado":
		*IDstatus = "7"
	case "Descartado":
		*IDstatus = "4"
	default:
		*IDstatus = "5"
	}

	//Inputs Manuais: Asset Tag, Name, Modelo do Ativo
	fmt.Println("Digite uma Marcação do Ativo única (Exemplo: T008921): ")
	fmt.Scanf("%v", Assettag)
	fmt.Println("Digite o Nome do Ativo (Exemplo: Macbook2014/Mac/Macintosh): ")
	fmt.Scanf("%v", name)
	fmt.Println("Digite o Modelo do Ativo (Exemplo: Air): ")
	fmt.Scanf("%v", modeloAtivo)

	//Somente alguns prints para sinalização; Sem utilidade pratica para o código.
	fmt.Printf("\nNOME DO DISPOSITIVO: %v\n", mac.Name)
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

	fmt.Println("Você deseja enviar essas informações para o inventário Snipeit? (sim/nao)")
	var answer string
	fmt.Scanf("%v", &answer)

	switch answer {
	case "sim", "s":
		MacOS.Snipesending(mac)
	case "nao", "n":
		fmt.Println("Você deseja apagar os arquivos criados? (sim/nao)")
		var anotherAnswer string
		fmt.Scanf("%v", &anotherAnswer)

		switch anotherAnswer {
		case "sim", "s":
			MacOS.Clear()
		case "nao", "n":
			fmt.Println("Certo. Fique à Vontade!")
		}

	}
}

func main() {

	fmt.Println("Qual o seus sistema operacional? (MacOS/Linux/Windows)")
	var resposta string
	fmt.Scanf("%v\n", &resposta)

	switch resposta {
	case "MacOS", "Macos", "MacOs", "macOS", "macos", "MACOS":
		forMacOs()
	case "Linux", "linux", "LINUX":

	case "Windows", "WINDOWS", "windows":

	}
}