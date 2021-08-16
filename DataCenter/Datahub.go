package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"

	functions "github.com/JVLAlves/Dinamize-DataHub/Utilities/Functions"
	snipe "github.com/JVLAlves/Dinamize-DataHub/Utilities/SnipeMethods"
	"github.com/JVLAlves/DinamizeSnipeitDataQuest/DataMission/Linux"
	"github.com/JVLAlves/DinamizeSnipeitDataQuest/DataMission/MacOS"
	"github.com/JVLAlves/DinamizeSnipeitDataQuest/DataMission/Windows"
)

//IP do inventário Snipeit
var IP string = "10.20.1.79:8001"

//Função de execução do programa em MacOS
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

	var mac snipe.CollectionT = snipe.CollectionT{}

	//Populando Struct MacOSt
	mac.SnipeitCPU11 = MacOS.Infos[1]
	mac.SnipeitMema3Ria7 = MacOS.Infos[2]
	mac.SnipeitHostname10 = MacOS.Infos[0]
	mac.SnipeitHd9 = MacOS.Infos[3]
	mac.Name = MacOS.Infos[0]

	mac.SnipeitMema3Ria7 = functions.RegexThis(`(^\d{1,3})`, MacOS.Infos[2]) + "GB"

	mac.AssetTag = functions.RegexThis(`\d`, MacOS.Infos[0])
	if mac.AssetTag == "" {
		mac.AssetTag = "No Asset Tag"
		log.Printf("Nenhum Asset Tag foi colocado, pois nenhuma sequência numérica foi encontrada no HOSTNAME: %v", MacOS.Infos[0])

	}

	SOregexed := functions.RegexThis(`(^\d{2}\.\d+)`, MacOS.Infos[4])
	numSO, err := strconv.ParseFloat(SOregexed, 64)
	if err != nil {
		log.Fatalf("Erro na conversão do S.O. para float")
	}

	if numSO >= 11.4 && numSO < 12.0 {
		mac.SnipeitSo8 = "11.4"
	}

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
	var modeloAtivo *string = &mac.SnipeitModel12

	//Input Manual: Tipo de Ativo
	fmt.Println("Digite o Tipo de Ativo (Exemplo:Desktop/MacBook): ")
	fmt.Scanf("%v", IDmodelo)

	//identificando o Modelo
	switch *IDmodelo {

	case "Desktop":
		*IDmodelo = "8"
		*modeloAtivo = "DNZ-Desktop"
	case "MacBook":
		*IDmodelo = "22"
		*modeloAtivo = "DNZ-Macbook"
	default:
		*IDmodelo = "11"
	}

	//Status ID
	*IDstatus = "5"

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

	//Verificando a existência de um ativo semelhante no inventário Snipe it
	if snipe.Verifybytag(mac.AssetTag, IP) {
		log.Println("Os dados do Ativo Criado não constam no sistema.")
		fmt.Println("Enviando Ativo para o Snipeit ")

		snipe.PostSnipe(mac, IP)
		log.Printf("NOVO ATIVO: %v", MacOS.Infos)
		log.Println("Ativo Criado enviado para o sistema.")

	} else {
		log.Println("Um Ativo semelhante foi encontrado no sistema.")
		fmt.Print("Asset Tag idêntico encontrado. Iniciando análise de disparidades")
		for i := 0; i < 4; i++ {
			time.Sleep(time.Second * 1)
			fmt.Print(".")
		}
		patch, boolean := snipe.Getbytag(IP, mac.AssetTag, mac)
		if boolean {
			fmt.Println("\nPATCH necessário.")
			fmt.Println("\nExecutando PATCH RESQUEST.")
			time.Sleep(time.Second * 3)
			id := snipe.Getidbytag(mac.AssetTag, IP)
			snipe.Patchbyid(id, IP, patch)

		} else {
			log.Println("Não foram encontradas disparidades entre o Ativo Existente no sistema e o Ativo Criado.")
			fmt.Println("\nSem alterações")
		}
	}
}

//Função de execução do programa em Windows
func forWindows() {

	Windows.MainProgram()

	//Essa variavel recebe um Tipo MacOSt, pois é o contrato padrão para a execução do programa
	var win snipe.CollectionT = snipe.CollectionT{}

	//Populando Struct MacOSt
	win.SnipeitCPU11 = Windows.Infos[3]
	win.SnipeitMema3Ria7 = Windows.Infos[2]
	win.SnipeitSo8 = Windows.Infos[1]
	win.SnipeitHostname10 = Windows.Infos[0]
	win.Name = Windows.Infos[0]

	win.AssetTag = functions.RegexThis(`\d`, Windows.Infos[0])
	if win.AssetTag == "" {
		win.AssetTag = "No Asset Tag"
		log.Printf("Nenhum Asset Tag foi colocado, pois nenhuma sequência numérica foi encontrada no HOSTNAME: %v", Windows.Infos[0])

	}

	win.SnipeitHd9 = functions.RegexThis(`(^\d{3})`, Windows.Infos[4]) + "GB"

	//Entrada Personalizada
	var IDmodelo *string = &win.ModelID
	var IDstatus *string = &win.StatusID
	var modeloAtivo *string = &win.SnipeitModel12

	//Input Manual: Tipo de Ativo
	fmt.Println("Digite o Tipo de Ativo (Exemplo:Desktop/Notebook): ")
	fmt.Scanf("%v", IDmodelo)

	//identificando o Modelo
	switch *IDmodelo {
	case "Notebook":
		*IDmodelo = "6"
		*modeloAtivo = "DNZ-Notebook"
	case "Desktop":
		*IDmodelo = "8"
		*modeloAtivo = "DNZ-Desktop"
	default:
		*IDmodelo = "11"
	}

	//Status ID
	*IDstatus = "5"

	//Somente alguns prints para sinalização; Sem utilidade pratica para o código.
	fmt.Printf("\nNOME DO DISPOSITIVO: %v\n", win.Name)
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
		log.Println("Os dados do Ativo Criado não constam no sistema.")
		fmt.Println("Enviando Ativo para o Snipeit ")

		snipe.PostSnipe(win, IP)
		log.Printf("NOVO ATIVO: %v", Windows.Infos)
		log.Println("Ativo Criado enviado para o sistema.")

	} else {
		log.Println("Um Ativo semelhante foi encontrado no sistema.")
		fmt.Print("Asset Tag idêntico encontrado. Iniciando análise de disparidades")
		for i := 0; i < 4; i++ {
			time.Sleep(time.Second * 1)
			fmt.Print(".")
		}
		patch, boolean := snipe.Getbytag(IP, win.AssetTag, win)
		if boolean {
			fmt.Println("\nPATCH necessário.")
			fmt.Println("\nExecutando PATCH RESQUEST.")
			time.Sleep(time.Second * 3)
			id := snipe.Getidbytag(win.AssetTag, IP)
			snipe.Patchbyid(id, IP, patch)

		} else {
			log.Println("Não foram encontradas disparidades entre o Ativo Existente no sistema e o Ativo Criado.")
			fmt.Println("\nSem alterações")
		}
	}

}

//Função de execução do programa em Linux
func forLinux() {

	//programa principal para a coleta de informações em Linux
	Linux.MainProgram()

	//Essa variavel recebe um Tipo MacOSt, pois é o contrato padrão para a execução do programa
	var lin snipe.CollectionT = snipe.CollectionT{}

	//Populando Struct MacOSt
	lin.SnipeitCPU11 = Linux.Infos[0]
	lin.SnipeitSo8 = Linux.Infos[2]
	lin.SnipeitHostname10 = Linux.Infos[3]

	lin.Name = Linux.Infos[3]
	interHD := functions.RegexThis(`(^\d{3}[,.]\d?)`, Linux.Infos[4])
	indexHD := strings.Split(interHD, ",")
	lin.SnipeitHd9 = strings.Join(indexHD, ".") + "GB"

	intermem := functions.RegexThis(`\d`, Linux.Infos[1])
	lin.SnipeitMema3Ria7 = intermem + "GB"

	lin.AssetTag = functions.RegexThis(`\d`, Linux.Infos[3])
	if lin.AssetTag == "" {
		lin.AssetTag = "No Asset Tag"
		log.Printf("Nenhum Asset Tag foi colocado, pois nenhuma sequência numérica foi encontrada no HOSTNAME: %v", Linux.Infos[0])

	}

	//Entrada Personalizada
	var IDmodelo *string = &lin.ModelID
	var IDstatus *string = &lin.StatusID
	var modeloAtivo *string = &lin.SnipeitModel12

	//Input Manual: Tipo de Ativo
	fmt.Println("Digite o Tipo de Ativo (Exemplo:Desktop/Notebook): ")
	fmt.Scanf("%v", IDmodelo)

	//identificando o Modelo
	switch *IDmodelo {
	case "Notebook":
		*IDmodelo = "6"
		*modeloAtivo = "DNZ-Notebook"
	case "Desktop":
		*IDmodelo = "8"
		*modeloAtivo = "DNZ-Desktop"
	default:
		*IDmodelo = "11"
	}

	//Status ID
	*IDstatus = "5"

	//Somente alguns prints para sinalização; Sem utilidade pratica para o código.
	fmt.Printf("\nNOME DO DISPOSITIVO: %v\n", lin.Name)
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
		log.Println("Os dados do Ativo Criado não constam no sistema.")
		fmt.Println("Enviando Ativo para o Snipeit ")

		snipe.PostSnipe(lin, IP)
		log.Printf("NOVO ATIVO: %v", Linux.Infos)
		log.Println("Ativo Criado enviado para o sistema.")

	} else {
		log.Println("Um Ativo semelhante foi encontrado no sistema.")
		fmt.Print("Asset Tag idêntico encontrado. Iniciando análise de disparidades")
		for i := 0; i < 4; i++ {
			time.Sleep(time.Second * 1)
			fmt.Print(".")
		}
		patch, boolean := snipe.Getbytag(IP, lin.AssetTag, lin)
		if boolean {
			fmt.Println("\nPATCH necessário.")
			fmt.Println("\nExecutando PATCH RESQUEST.")
			time.Sleep(time.Second * 3)
			id := snipe.Getidbytag(lin.AssetTag, IP)
			snipe.Patchbyid(id, IP, patch)

		} else {
			log.Println("Não foram encontradas disparidades entre o Ativo Existente no sistema e o Ativo Criado.")
			fmt.Println("\nSem alterações")
		}
	}

}

//função principal
func main() {
	logname := "logs" + functions.Today() + ".txt"
	file, err := os.OpenFile(logname, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(file)

	//mensagem de abertura
	fmt.Print("Dectecting your Operating System")
	for i := 0; i < 4; i++ {
		time.Sleep(time.Second * 1)
		fmt.Print(".")
	}

	log.Printf("\nInicio de execução.\n")

	//Identificando sistema operacional
	switch runtime.GOOS {
	case "darwin":
		forMacOs()
	case "linux":
		forLinux()

	case "windows":
		forWindows()
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

	//mensagem de encerramento
	fmt.Println("\n\nObrigado pela paciência! (FIM)")
	log.Printf("\nFim de execução.\n")
}
