package main

import (
	"fmt"
	"log"
	"runtime"
	"sync"
	"time"

	"github.com/JVLAlves/DinamizeSnipeitDataQuest/DataQuest/Linux"
	"github.com/JVLAlves/DinamizeSnipeitDataQuest/DataQuest/MacOS"
	"github.com/JVLAlves/DinamizeSnipeitDataQuest/DataQuest/Windows"
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
	mac.AssetTag = MacOS.Infos[0]
	mac.Name = MacOS.Infos[0]

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

	//Status ID
	*IDstatus = "5"

	//Inputs Manuais: Modelo do Ativo
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
			wg := &sync.WaitGroup{}
			wg.Add(1)
			go MacOS.Clear(wg)
			wg.Wait()
		case "nao", "n":
			fmt.Println("Certo. Fique à Vontade!")
		}

	}
}

func forWindows() {

	Windows.MainProgram()

	//This variable receive a type MacOSt even being Windows. This is because this type is universal, but was named in the MacOS Program.
	var win MacOS.MacOSt = MacOS.MacOSt{}

	//Populando Struct MacOSt
	win.SnipeitCPU11 = Windows.Infos[3]
	win.SnipeitMema3Ria7 = Windows.Infos[2]
	win.SnipeitSo8 = Windows.Infos[1]
	win.SnipeitHostname10 = Windows.Infos[0]
	win.SnipeitHd9 = Windows.Infos[4]
	win.AssetTag = Windows.Infos[0]
	win.Name = Windows.Infos[0]

	//Entrada Personalizada
	var IDmodelo *string = &win.ModelID
	var IDstatus *string = &win.StatusID
	var modeloAtivo *string = &win.SnipeitModel12

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

	//Status ID
	*IDstatus = "5"

	//Inputs Manuais: Asset Tag, Name, Modelo do Ativo
	fmt.Println("Digite o Modelo do Ativo (Exemplo: Air): ")
	fmt.Scanf("%v", modeloAtivo)

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

	fmt.Println("Você deseja enviar essas informações para o inventário Snipeit? (sim/nao)")
	var answer string
	fmt.Scanf("%v", &answer)

	switch answer {
	case "sim", "s":
		MacOS.Snipesending(win)
	case "nao", "n":
		fmt.Println("Você deseja apagar os arquivos criados? (sim/nao)")
		var anotherAnswer string
		fmt.Scanf("%v", &anotherAnswer)

		switch anotherAnswer {
		case "sim", "s":
			wg := &sync.WaitGroup{}
			wg.Add(1)
			go MacOS.Clear(wg)
			wg.Wait()
		case "nao", "n":
			fmt.Println("Certo. Fique à Vontade!")
		}

	}

}

func forLinux() {

	Linux.MainProgram()

	//This variable receive a type MacOSt even being Windows. This is because this type is universal, but was named in the MacOS Program.
	var lin MacOS.MacOSt = MacOS.MacOSt{}

	//Populando Struct MacOSt
	lin.SnipeitCPU11 = Linux.Infos[0]
	lin.SnipeitMema3Ria7 = Linux.Infos[1]
	lin.SnipeitSo8 = Linux.Infos[2]
	lin.SnipeitHostname10 = Linux.Infos[3]
	lin.SnipeitHd9 = Linux.Infos[4]
	lin.AssetTag = Linux.Infos[3]
	lin.Name = Linux.Infos[3]

	//Entrada Personalizada
	var IDmodelo *string = &lin.ModelID
	var IDstatus *string = &lin.StatusID
	var modeloAtivo *string = &lin.SnipeitModel12

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

	//Status ID
	*IDstatus = "5"

	//Inputs Manuais: Asset Tag, Name, Modelo do Ativo
	fmt.Println("Digite o Modelo do Ativo (Exemplo: Air): ")
	fmt.Scanf("%v", modeloAtivo)

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

	fmt.Println("Você deseja enviar essas informações para o inventário Snipeit? (sim/nao)")
	var answer string
	fmt.Scanf("%v", &answer)

	switch answer {
	case "sim", "s":
		MacOS.Snipesending(lin)
	case "nao", "n":
		fmt.Println("Você deseja apagar os arquivos criados? (sim/nao)")
		var anotherAnswer string
		fmt.Scanf("%v", &anotherAnswer)

		switch anotherAnswer {
		case "sim", "s":
			wg := &sync.WaitGroup{}
			wg.Add(1)
			go MacOS.Clear(wg)
			wg.Wait()
		case "nao", "n":
			fmt.Println("Certo. Fique à Vontade!")
		}

	}

}

func main() {

	fmt.Print("Dectecting your Operating System")
	for i := 0; i < 4; i++ {
		time.Sleep(time.Second * 1)
		fmt.Print(".")
	}

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

}
