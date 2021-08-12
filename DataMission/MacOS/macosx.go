package MacOS

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"sync"
	"time"
)

//Definindo Tipo para popular com as informações do computador.
type MacOSt struct {
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

//Lista para leitura linha a linha
var Linhas []string

//Lista para Informações armazenadas
var Infos []string

//Cria arquivos com as informações retiradas do computador via Terminal
func Create(wg *sync.WaitGroup, command string, args string) {

	outFile, err := os.OpenFile(command+".out", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0755)
	defer outFile.Close()
	if err != nil {
		fmt.Println("error creating file", err)
	}
	cmd := exec.Command("bash", "-c", command+" "+args)

	out, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println("error attaching command stdout", err)
	}
	go io.Copy(outFile, out)

	err = cmd.Run()
	if err != nil {
		panic(err)
	}

	fmt.Println("Arquivo criado...")
	time.Sleep(time.Second)
	wg.Done()
}

//Lê os arquivos criados pela função Create
func Running() {

	file, err := os.Open("uname.out")
	if err != nil {
		log.Print(err)
	}
	fileScanner := bufio.NewScanner(file)
	Linhas = []string{}

	for fileScanner.Scan() {
		Linhas = append(Linhas, fileScanner.Text())
		if fileScanner.Err() != nil {
			log.Fatalf("Erro SCAN: %v", fileScanner.Err().Error())
		}
	}
	Infos = append(Infos, Linhas[0])
	file.Close()

	file, err = os.Open("sysctl.out")
	if err != nil {
		log.Print(err)
	}
	fileScanner = bufio.NewScanner(file)
	Linhas = []string{}
	for fileScanner.Scan() {
		Linhas = append(Linhas, fileScanner.Text())
	}
	Infos = append(Infos, Linhas[0])
	fmt.Println(Infos[1])
	file.Close()

	file, err = os.Open("hostinfo.out")
	if err != nil {
		log.Print(err)
	}
	fileScanner = bufio.NewScanner(file)
	Linhas = []string{}
	for fileScanner.Scan() {
		Linhas = append(Linhas, fileScanner.Text())
	}
	Infos = append(Infos, Linhas[0])
	fmt.Println(Infos[2])
	file.Close()

	file, err = os.Open("diskutil.out")
	if err != nil {
		log.Print(err)
	}
	fileScanner = bufio.NewScanner(file)
	Linhas = []string{}
	for fileScanner.Scan() {
		Linhas = append(Linhas, fileScanner.Text())
	}
	Infos = append(Infos, Linhas[0])
	fmt.Println(Infos[3])
	file.Close()

	file, err = os.Open("sw_vers.out")
	if err != nil {
		log.Print(err)
	}
	fileScanner = bufio.NewScanner(file)
	Linhas = []string{}
	for fileScanner.Scan() {
		Linhas = append(Linhas, fileScanner.Text())
	}
	Infos = append(Infos, Linhas[0])
	fmt.Println(Infos[4])
	file.Close()

	return
}

//Envia os dados do computador para o inventário no Snipeit. (Essa função recebe a variavel que recebe o tipo struct criado com os dados do computador)
func Snipesending(mac MacOSt) {

	//URL da API SnipeIt
	url := "http://10.20.1.79:8001/api/v1/hardware"

	// Token de autentiucação
	var bearer = "Bearer " + "eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9.eyJhdWQiOiIxIiwianRpIjoiM2NlMzRhNDM0OGNjMGRkMjczMWQyMDM0ZDQ4MzRkZTZiMTQ3MGI3ODE2YWQyM2RmMjRmMzg0YzE3ZjIzOWU1N2E5ZTg2N2E0ODhlMTg5YTEiLCJpYXQiOjE2MjY0MzU0MzYsIm5iZiI6MTYyNjQzNTQzNiwiZXhwIjoyMDk5ODIxMDM1LCJzdWIiOiIzIiwic2NvcGVzIjpbXX0.JtCQ_KStz4TluCkt_6JGJLmSGVhuY6dS_3OQ7KJicm8vSgYnfh2cwzrjjgoDU92u5RN2-fMHMji_ju6a4Lm33_nyj6_qclFV9SPRtO-UqMJe7EVkPhe0bP3co-9dVKyfUmSyi7GjVeHkUcD2OGG9m_zhu7krpwzQRBNiaNR9dJwCeBEbH1O13kKQItRl_V_-DDEtFF-bTnQ3DbnlEqZxtY4we6-qjpXmIrGmOU27pH5DUXZ8-cxqlAKP1ysBz_BJRBYGN0HZqYyL2AgrTG_k9sPds2CSyqPhbTvjm7yD5IxPOAcmasJbJoAPSyZecpNSecOL7JVsjB7UFcDPTdIy6zykIqJV6Zj-3qwkg4VrAt6iGvSIPCfSPzlydwk3o0znDHisp_9IDGuSTII49kAGnGb5Kw6WWsV9xQrXBtm6R41cwVAGc47r9j8tLux5PmlXdcrSxGS1uiiaMwZSx1ZdvZlC85f5LSpKiA0qP85acTX2R_Aav4oqsx_FN-UkBuBs8ADYC-sxMDVDuokr49IkkgVY9LUfkk8-pQX4IqKZKBOHuPAT1NsalgDPOZG9pFaIQ9kmt9Qm6TkkinNIPiwcBJ2mqHXziirtvQqylfrH2MBkXAofHK_-EEkOCAsARfFT41iw7wkJwW5ijliz5SC2ZiG6HTFS9WIG88WNiRzu9qc"

	//transformando em bytes a variável hw
	hardwarePostJSON, err := json.Marshal(mac)
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

	// Unmarshal do resultado do response
	response := SnipeITHardwareResponseT{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Printf("Reading body failed: %s", err)
		return
	}

	//Printando o Response
	fmt.Println("Response do POST:", response)

	//Animação de Apagando Junk Files
	fmt.Print("\nApagando Junk Files")
	for i := 0; i < 4; i++ {
		time.Sleep(time.Second * 1)
		fmt.Print(".")
	}

	//Apagando Junk Files
	cmd := exec.Command("rm", "uname.out", "sysctl.out", "hostinfo.out", "diskutil.out", "sw_vers.out")
	stdout, _ := cmd.Output()
	fmt.Println(string(stdout))
}

/*Apaga os arquivos criados no inicio do código.
Não é necessário utiliza-lo após a função Snipesending, já possui esse programa internamente.*/
func Clear(wg *sync.WaitGroup) {

	//Animação de Apagando Junk Files
	fmt.Print("\nApagando Junk Files")
	for i := 0; i < 4; i++ {
		time.Sleep(time.Second * 1)
		fmt.Print(".")
	}

	//Apagando Junk Files
	cmd := exec.Command("rm", "uname.out", "sysctl.out", "hostinfo.out", "diskutil.out", "sw_vers.out")
	stdout, _ := cmd.Output()
	fmt.Println(string(stdout))
	wg.Done()

}
