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

// Modelo POST/REQUEST
type SnipeITHardwareRequestT struct {
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
type SnipeITHardwareResponseT struct {
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
	//fmt.Println("\n\n\n", Infos)
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
	stdout, _ = cmd.Output()
	fmt.Println(string(stdout))

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
	stdout, _ = cmd.Output()
	fmt.Println(string(stdout))

	fmt.Println(Infos)

	cmd = exec.Command("rm", "tamanhoDoHd.txt", "SO.txt", "hostname.txt", "tamanhoDoDisco.txt")
	stdout, _ = cmd.Output()
	fmt.Println(string(stdout))
}
