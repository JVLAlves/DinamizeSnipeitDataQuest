package Linux

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
)

// Modelo POST/REQUEST
type snipeITHardwareRequestT struct {
	ModelID           string `json:"model_id"`
	StatusID          string `json:"status_id"`
	AssetTag          string `json:"asset_tag"`
	Name              string `json:"name"`
	SnipeitSo8        string `json:"_snipeit_so_8"`
	SnipeitModel12    string `json:"_snipeit_modelo_12"`
	SnipeitHostname10 string `json:"_snipeit_hostname_10"`
	SnipeitHd9        string `json:"_snipeit_hd_9"`
	SnipeitCPU11      string `json:"_snipeit_cpu_11"`
	SnipeitMema3Ria7  string `json:"_snipeit_mema3ria_7"`
}

//Modelo de RESPONSE
type snipeITHardwareResponseT struct {
	Status   string `json:"status"`
	Messages string `json:"messages"`
	Payload  struct {
		ModelID        int    `json:"model_id"`
		Name           string `json:"name"`
		Serial         string `json:"serial"`
		CompanyID      string `json:"company_id"`
		OrderNumber    string `json:"order_number"`
		Notes          string `json:"notes"`
		AssetTag       string `json:"asset_tag"`
		UserID         int    `json:"user_id"`
		Archived       string `json:"archived"`
		Physical       string `json:"physical"`
		Depreciate     string `json:"depreciate"`
		StatusID       int    `json:"status_id"`
		WarrantyMonths string `json:"warranty_months"`
		PurchaseCost   string `json:"purchase_cost"`
		PurchaseDate   string `json:"purchase_date"`
		AssignedTo     string `json:"assigned_to"`
		SupplierID     string `json:"supplier_id"`
		Requestable    int    `json:"requestable"`
		RtdLocationID  string `json:"rtd_location_id"`
		UpdatedAt      string `json:"updated_at"`
		CreatedAt      string `json:"created_at"`
		ID             int    `json:"id"`
		Model          struct {
			ID                   int    `json:"id"`
			Name                 string `json:"name"`
			ModelNumber          string `json:"model_number"`
			ManufacturerID       int    `json:"manufacturer_id"`
			CategoryID           int    `json:"category_id"`
			CreatedAt            string `json:"created_at"`
			UpdatedAt            string `json:"updated_at"`
			DepreciationID       int    `json:"depreciation_id"`
			Eol                  int    `json:"eol"`
			Image                string `json:"image"`
			DeprecatedMacAddress int    `json:"deprecated_mac_address"`
			FieldsetID           int    `json:"fieldset_id"`
			Notes                string `json:"notes"`
			Requestable          int    `json:"requestable"`
		} `json:"model"`
	} `json:"payload"`
}

//Variáveis de armazenamento dos dados da máquina
var linhas = []string{}
var infos = []string{}

func main() {
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
		linhas = append(linhas, fileScanner.Text())

	}
	// adicionando informação encontrada no arquivo CPU a variável

	infos = append(infos, linhas[4])

	//Tratando o ocasoional erro da leitura do arquivo
	if err := fileScanner.Err(); err != nil {
		log.Fatalf("Error while reading file: %s", err)
	}

	//fechando o arquivo lido
	file.Close()

	//Executa o comando Script para escrever a sessão do terminal em arquivo txt (Tamanho do disco)
	cmd := exec.Command("script", "-c", "free -h |grep Mem |awk '{print $2}'", "tamanhoDoHd.txt")
	stdout, _ := cmd.Output()
	fmt.Println(string(stdout))

	// abrindo o arquiuvo criado "tamanhoDoHd.txt"
	file, err = os.Open("tamanhoDoHd.txt")

	//Tratando o ocasoional erro da leitura do arquivo
	if err != nil {
		log.Fatalf("Error when opening file: %s", err)
	}

	//Lendo o arquivo "tamanhoDoHd.txt"
	fileScanner = bufio.NewScanner(file)

	//Limpa o Array das linhas
	linhas = []string{}

	//Lendo linha a linha
	for fileScanner.Scan() {
		fmt.Println(fileScanner.Text())
		linhas = append(linhas, fileScanner.Text())

	}

	// adicionando informação encontrada no arquivo "tamanhoDoHd.txt" a variável
	infos = append(infos, linhas[1])

	//Tratando o ocasoional erro da leitura do arquivo
	if err := fileScanner.Err(); err != nil {
		log.Fatalf("Error while reading file: %s", err)
	}

	//fechando o arquivo lido
	file.Close()

	//Executa o comando Script para escrever a sessão do terminal em arquivo txt (S.O.)
	cmd = exec.Command("script", "-c", "lsb_release -d |grep Description |awk '{print $2,$3,$4}'", "SO.txt")
	stdout, _ = cmd.Output()
	fmt.Println(string(stdout))

	// abrindo o arquiuvo criado "S0.txt"
	file, err = os.Open("SO.txt")

	//Tratando o ocasoional erro da leitura do arquivo
	if err != nil {
		log.Fatalf("Error when opening file: %s", err)
	}

	//Lendo o arquivo "SO.txt"
	fileScanner = bufio.NewScanner(file)

	//Limpa o Array das linhas
	linhas = []string{}

	//Lendo linha a linha
	for fileScanner.Scan() {
		fmt.Println(fileScanner.Text())
		linhas = append(linhas, fileScanner.Text())

	}

	// adicionando informação encontrada no arquivo "SO.txt" a variável
	infos = append(infos, linhas[1])

	//Tratando o ocasoional erro da leitura do arquivo
	if err := fileScanner.Err(); err != nil {
		log.Fatalf("Error while reading file: %s", err)
	}

	//fechando o arquivo lido
	file.Close()

	//Executa o comando Script para escrever a sessão do terminal em arquivo txt (Hostname)
	cmd = exec.Command("script", "-c", "hostname", "hostname.txt")
	stdout, _ = cmd.Output()
	fmt.Println(string(stdout))

	// abrindo o arquiuvo criado "Hostname.txt"
	file, err = os.Open("hostname.txt")

	//Tratando o ocasoional erro da leitura do arquivo
	if err != nil {
		log.Fatalf("Error when opening file: %s", err)
	}

	//Lendo o arquivo "Hostname.txt"
	fileScanner = bufio.NewScanner(file)

	//Limpa o Array das linhas
	linhas = []string{}

	//Lendo linha a linha
	for fileScanner.Scan() {
		fmt.Println(fileScanner.Text())
		linhas = append(linhas, fileScanner.Text())

	}

	infos = append(infos, linhas[1])

	// adicionando informação encontrada no arquivo "Hostname.txt" a variável
	if err := fileScanner.Err(); err != nil {
		log.Fatalf("Error while reading file: %s", err)
	}

	//fechando o arquivo lido
	file.Close()

	//Executa o comando Script para escrever a sessão do terminal em arquivo txt (Tamanho do Disco)
	cmd = exec.Command("script", "-c", "lsblk |grep disk |awk '{print $4}'", "tamanhoDoDisco.txt")
	stdout, _ = cmd.Output()
	fmt.Println(string(stdout))

	// abrindo o arquiuvo criado "tamanhoDoDisco.txt"
	file, err = os.Open("tamanhoDoDisco.txt")

	//Tratando o ocasoional erro da leitura do arquivo
	if err != nil {
		log.Fatalf("Error when opening file: %s", err)
	}

	fileScanner = bufio.NewScanner(file)
	linhas = []string{}

	//Lendo linha a linha
	for fileScanner.Scan() {
		fmt.Println(fileScanner.Text())
		linhas = append(linhas, fileScanner.Text())

	}

	// adicionando informação encontrada no arquivo "tamanhoDoDisco.txt" a variável
	infos = append(infos, linhas[1])

	//Tratando o ocasoional erro da leitura do arquivo
	if err := fileScanner.Err(); err != nil {
		log.Fatalf("Error while reading file: %s", err)
	}

	//fechando o arquivo lido
	file.Close()

	/*for i := 0; i != len(infos); i++ {

		fmt.Println(infos[i])
	}*/

	//Variável recebe o tipo REQUEST
	var hw snipeITHardwareRequestT = snipeITHardwareRequestT{}

	//Entrada Personalizada
	var IDmodelo *string = &hw.ModelID
	var IDstatus *string = &hw.StatusID
	var Assettag *string = &hw.AssetTag
	var name *string = &hw.Name
	var modeloAtivo *string = &hw.SnipeitModel12

	fmt.Println("Digite o Tipo de Ativo (Exemplo:Teclado/Desktop/MacBook): ")
	fmt.Scanf("%v", "%v", "%v", "%v", IDmodelo)

	//identificando o Modelo
	switch *IDmodelo {
	case "Leitor de Cartão Mesa":
		*IDmodelo = "1"
	case "Leitor de Cartão Porta":
		*IDmodelo = "2"
	case "Mouse Sem Fio":
		*IDmodelo = "3"
	case "Roteador":
		*IDmodelo = "4"
	case "Roteador Wireless":
		*IDmodelo = "5"
	case "Notebook":
		*IDmodelo = "6"
	case "Celular Samsumg":
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
	case "Telefone Grandstream":
		*IDmodelo = "13"
	case "NP350":
		*IDmodelo = "14"
	case "Samsung Default":
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
	case "Asus X":
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
	case "Multilaser Desk":
		*IDmodelo = "30"
	case "Zelman Desk":
		*IDmodelo = "31"
	case "TV LG":
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

	fmt.Println("Digite o Status do Ativo (Padrão: Disponível): ")
	fmt.Scanf("%v", IDstatus)

	//Identificando o Status
	switch *IDstatus {
	case "Disponível":
		*IDstatus = "5"
	case "Indisponível":
		*IDstatus = "6"
	case "Ocupado":
		*IDstatus = "7"
	case "Descartado":
		*IDstatus = "4"
	default:
		*IDstatus = "5"
	}
	fmt.Println("Digite a Marcação do Ativo (Precisa ser única): ")
	fmt.Scanf("%v", Assettag)
	fmt.Println("Digite o Nome do Ativo: ")
	fmt.Scanf("%v", name)
	fmt.Println("Digite o Modelo do Ativo: ")
	fmt.Scanf("%v", modeloAtivo)

	//Populando a variável
	hw.SnipeitCPU11 = infos[0]
	hw.SnipeitMema3Ria7 = infos[1]
	hw.SnipeitSo8 = infos[2]
	hw.SnipeitHostname10 = infos[3]
	hw.SnipeitHd9 = infos[4]

	//URL da API SnipeIt
	url := "http://10.20.1.79:8001/api/v1/hardware"

	// Token de autentiucação
	var bearer = "Bearer " + "eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9.eyJhdWQiOiIxIiwianRpIjoiM2NlMzRhNDM0OGNjMGRkMjczMWQyMDM0ZDQ4MzRkZTZiMTQ3MGI3ODE2YWQyM2RmMjRmMzg0YzE3ZjIzOWU1N2E5ZTg2N2E0ODhlMTg5YTEiLCJpYXQiOjE2MjY0MzU0MzYsIm5iZiI6MTYyNjQzNTQzNiwiZXhwIjoyMDk5ODIxMDM1LCJzdWIiOiIzIiwic2NvcGVzIjpbXX0.JtCQ_KStz4TluCkt_6JGJLmSGVhuY6dS_3OQ7KJicm8vSgYnfh2cwzrjjgoDU92u5RN2-fMHMji_ju6a4Lm33_nyj6_qclFV9SPRtO-UqMJe7EVkPhe0bP3co-9dVKyfUmSyi7GjVeHkUcD2OGG9m_zhu7krpwzQRBNiaNR9dJwCeBEbH1O13kKQItRl_V_-DDEtFF-bTnQ3DbnlEqZxtY4we6-qjpXmIrGmOU27pH5DUXZ8-cxqlAKP1ysBz_BJRBYGN0HZqYyL2AgrTG_k9sPds2CSyqPhbTvjm7yD5IxPOAcmasJbJoAPSyZecpNSecOL7JVsjB7UFcDPTdIy6zykIqJV6Zj-3qwkg4VrAt6iGvSIPCfSPzlydwk3o0znDHisp_9IDGuSTII49kAGnGb5Kw6WWsV9xQrXBtm6R41cwVAGc47r9j8tLux5PmlXdcrSxGS1uiiaMwZSx1ZdvZlC85f5LSpKiA0qP85acTX2R_Aav4oqsx_FN-UkBuBs8ADYC-sxMDVDuokr49IkkgVY9LUfkk8-pQX4IqKZKBOHuPAT1NsalgDPOZG9pFaIQ9kmt9Qm6TkkinNIPiwcBJ2mqHXziirtvQqylfrH2MBkXAofHK_-EEkOCAsARfFT41iw7wkJwW5ijliz5SC2ZiG6HTFS9WIG88WNiRzu9qc"

	//transformando em bytes a variável hw
	hardwarePostJSON, err := json.Marshal(hw)
	//Tratando o ocasoional erro transformação da variável em byte
	if err != nil {
		panic(err)
	}

	//POST REQUEST
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(hardwarePostJSON))

	//Tratando o ocasoional erro do POST/REQUEST
	if err != nil {
		panic(err)
	}

	//adicionando os headers a autorização
	req.Header.Add("Authorization", bearer)
	//definindo a formatação do REQUEST
	req.Header.Add("Content-type", "application/json")

	// Send req using http Client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error on request.\n[ERROR] -", err)
	}

	//fechando o Response após a conclusão do código
	defer resp.Body.Close()

	//lendo o RESQUEST
	body, err := ioutil.ReadAll(resp.Body)
	//Tratando o ocasoional erro do request
	if err != nil {
		log.Println("Error on parsing response.\n[ERROR] -", err)
	}

	//coventendo Body em bytes para Body em String
	bodyString := string(body)
	log.Print(bodyString)

	// Unmarshal do resultado do response
	response := snipeITHardwareResponseT{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Printf("Reading body failed: %s", err)
		return
	}

	//Printando o Response
	fmt.Println("Response do POST:", response)

	//Apagando Junk Files
	cmd = exec.Command("rm", "tamanhoDoHd.txt", "SO.txt", "hostname.txt", "tamanhoDoDisco.txt")
	stdout, _ = cmd.Output()
	fmt.Println(string(stdout))

}
