package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/JVLAlves/DinamizeSnipeitDataQuest/DataQuest/Linux"
	"github.com/JVLAlves/DinamizeSnipeitDataQuest/DataQuest/MacOS"
	"github.com/JVLAlves/DinamizeSnipeitDataQuest/DataQuest/Windows"
	functions "github.com/JVLAlves/DinamizeSnipeitDataQuest/Utilities"
)

type PatchRespose struct {
	Status   string `json:"status"`
	Messages string `json:"messages"`
	Payload  struct {
		ID                  int    `json:"id"`
		Name                string `json:"name"`
		AssetTag            string `json:"asset_tag"`
		ModelID             int    `json:"model_id"`
		Serial              string `json:"serial"`
		PurchaseDate        string `json:"purchase_date"`
		PurchaseCost        string `json:"purchase_cost"`
		OrderNumber         string `json:"order_number"`
		AssignedTo          string `json:"assigned_to"`
		Notes               string `json:"notes"`
		Image               string `json:"image"`
		UserID              int    `json:"user_id"`
		CreatedAt           string `json:"created_at"`
		UpdatedAt           string `json:"updated_at"`
		Physical            int    `json:"physical"`
		DeletedAt           string `json:"deleted_at"`
		StatusID            int    `json:"status_id"`
		Archived            int    `json:"archived"`
		WarrantyMonths      string `json:"warranty_months"`
		Depreciate          int    `json:"depreciate"`
		SupplierID          int    `json:"supplier_id"`
		Requestable         bool   `json:"requestable"`
		RtdLocationID       int    `json:"rtd_location_id"`
		Accepted            string `json:"accepted"`
		LastCheckout        string `json:"last_checkout"`
		ExpectedCheckin     string `json:"expected_checkin"`
		CompanyID           string `json:"company_id"`
		AssignedType        string `json:"assigned_type"`
		LastAuditDate       string `json:"last_audit_date"`
		NextAuditDate       string `json:"next_audit_date"`
		LocationID          int    `json:"location_id"`
		CheckinCounter      int    `json:"checkin_counter"`
		CheckoutCounter     int    `json:"checkout_counter"`
		RequestsCounter     int    `json:"requests_counter"`
		SnipeitImei1        string `json:"_snipeit_imei_1"`
		SnipeitPhoneNumber2 string `json:"_snipeit_phone_number_2"`
		SnipeitRAM3         string `json:"_snipeit_ram_3"`
		SnipeitCPU4         string `json:"_snipeit_cpu_4"`
		SnipeitMacAddress5  string `json:"_snipeit_mac_address_5"`
	} `json:"payload"`
}

type UniversalGetT struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	AssetTag string `json:"asset_tag"`
	Serial   string `json:"serial"`
	Model    struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"model"`
	ModelNumber string `json:"model_number"`
	Eol         string `json:"eol"`
	StatusLabel struct {
		ID         int    `json:"id"`
		Name       string `json:"name"`
		StatusType string `json:"status_type"`
		StatusMeta string `json:"status_meta"`
	} `json:"status_label"`
	Category struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"category"`
	Manufacturer struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"manufacturer"`
	Supplier        string `json:"supplier"`
	Notes           string `json:"notes"`
	OrderNumber     string `json:"order_number"`
	Company         string `json:"company"`
	Location        string `json:"location"`
	RtdLocation     string `json:"rtd_location"`
	Image           string `json:"image"`
	AssignedTo      string `json:"assigned_to"`
	WarrantyMonths  string `json:"warranty_months"`
	WarrantyExpires string `json:"warranty_expires"`
	CreatedAt       struct {
		Datetime  string `json:"datetime"`
		Formatted string `json:"formatted"`
	} `json:"created_at"`
	UpdatedAt struct {
		Datetime  string `json:"datetime"`
		Formatted string `json:"formatted"`
	} `json:"updated_at"`
	LastAuditDate   string `json:"last_audit_date"`
	NextAuditDate   string `json:"next_audit_date"`
	DeletedAt       string `json:"deleted_at"`
	PurchaseDate    string `json:"purchase_date"`
	LastCheckout    string `json:"last_checkout"`
	ExpectedCheckin string `json:"expected_checkin"`
	PurchaseCost    string `json:"purchase_cost"`
	CheckinCounter  int    `json:"checkin_counter"`
	CheckoutCounter int    `json:"checkout_counter"`
	RequestsCounter int    `json:"requests_counter"`
	UserCanCheckout bool   `json:"user_can_checkout"`
	CustomFields    struct {
		Modelo struct {
			Field       string `json:"field"`
			Value       string `json:"value"`
			FieldFormat string `json:"field_format"`
		} `json:"Modelo"`
		Hostname struct {
			Field       string `json:"field"`
			Value       string `json:"value"`
			FieldFormat string `json:"field_format"`
		} `json:"Hostname"`
		Hd struct {
			Field       string `json:"field"`
			Value       string `json:"value"`
			FieldFormat string `json:"field_format"`
		} `json:"HD"`
		CPU struct {
			Field       string `json:"field"`
			Value       string `json:"value"`
			FieldFormat string `json:"field_format"`
		} `json:"CPU"`
		MemRia struct {
			Field       string `json:"field"`
			Value       string `json:"value"`
			FieldFormat string `json:"field_format"`
		} `json:"Memória"`
		SO struct {
			Field       string `json:"field"`
			Value       string `json:"value"`
			FieldFormat string `json:"field_format"`
		} `json:"S.O."`
		Office struct {
			Field       string `json:"field"`
			Value       string `json:"value"`
			FieldFormat string `json:"field_format"`
		} `json:"Office"`
		Setor struct {
			Field       string `json:"field"`
			Value       string `json:"value"`
			FieldFormat string `json:"field_format"`
		} `json:"Setor"`
	} `json:"custom_fields"`
	AvailableActions struct {
		Checkout bool `json:"checkout"`
		Checkin  bool `json:"checkin"`
		Clone    bool `json:"clone"`
		Restore  bool `json:"restore"`
		Update   bool `json:"update"`
		Delete   bool `json:"delete"`
	} `json:"available_actions"`
}
type ErrorT struct {
	Status   string `json:"status"`
	Messages string `json:"messages"`
	Payload  string `json:"payload"`
}

type IDT struct {
	ID int `json:"id"`
}

/*Busca o ID do Ativo Existente.
Ele recebe o Asset Tag do Ativo Existente e retorna o ID int*/
func Getidbytag(assettag string) (ID int) {
	url := "http://10.20.1.79:8001/api/v1/hardware/bytag/" + assettag
	var bearer = "Bearer " + "eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9.eyJhdWQiOiIxIiwianRpIjoiM2NlMzRhNDM0OGNjMGRkMjczMWQyMDM0ZDQ4MzRkZTZiMTQ3MGI3ODE2YWQyM2RmMjRmMzg0YzE3ZjIzOWU1N2E5ZTg2N2E0ODhlMTg5YTEiLCJpYXQiOjE2MjY0MzU0MzYsIm5iZiI6MTYyNjQzNTQzNiwiZXhwIjoyMDk5ODIxMDM1LCJzdWIiOiIzIiwic2NvcGVzIjpbXX0.JtCQ_KStz4TluCkt_6JGJLmSGVhuY6dS_3OQ7KJicm8vSgYnfh2cwzrjjgoDU92u5RN2-fMHMji_ju6a4Lm33_nyj6_qclFV9SPRtO-UqMJe7EVkPhe0bP3co-9dVKyfUmSyi7GjVeHkUcD2OGG9m_zhu7krpwzQRBNiaNR9dJwCeBEbH1O13kKQItRl_V_-DDEtFF-bTnQ3DbnlEqZxtY4we6-qjpXmIrGmOU27pH5DUXZ8-cxqlAKP1ysBz_BJRBYGN0HZqYyL2AgrTG_k9sPds2CSyqPhbTvjm7yD5IxPOAcmasJbJoAPSyZecpNSecOL7JVsjB7UFcDPTdIy6zykIqJV6Zj-3qwkg4VrAt6iGvSIPCfSPzlydwk3o0znDHisp_9IDGuSTII49kAGnGb5Kw6WWsV9xQrXBtm6R41cwVAGc47r9j8tLux5PmlXdcrSxGS1uiiaMwZSx1ZdvZlC85f5LSpKiA0qP85acTX2R_Aav4oqsx_FN-UkBuBs8ADYC-sxMDVDuokr49IkkgVY9LUfkk8-pQX4IqKZKBOHuPAT1NsalgDPOZG9pFaIQ9kmt9Qm6TkkinNIPiwcBJ2mqHXziirtvQqylfrH2MBkXAofHK_-EEkOCAsARfFT41iw7wkJwW5ijliz5SC2ZiG6HTFS9WIG88WNiRzu9qc"
	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", bearer)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("Falha de conexão com o Host Snipeit.")
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Println("Error on parsing response.\n[ERROR] -", err)
	}

	// Unmarshal do resultado do response
	response := IDT{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Printf("Reading body failed: %s", err)
	}

	Id := response.ID

	return Id

}

//Variável Struct utilizada para a análise de disparidades entre Ativo Existente no inventário e Ativo Criado pela execução do programa
var Analyser MacOS.MacOSt = MacOS.MacOSt{}

/*Busca as informações do Ativo existente no inventário Snipe it e compara com o Ativo criado ao executar o programa.
Ele recebe o Asset Tag único do Ativo existente e a variável que contém o tipo populado com as informações do Ativo criado.
Ao comparar ambos A. Existente e A. Criado ele destaca as disparidades e as retorna  em uma string Patchrequest, assim como um bool Needpatching que afirma se é necessário um PATCH ou não.

OBS: Patchrequest é um JSON padronizado especificamente para o envio através do método PATCH.*/
func Getbytag(assettag string, ativo MacOS.MacOSt) (Patchrequest string, Needpatching bool) {
	url := "http://10.20.1.79:8001/api/v1/hardware/bytag/" + assettag
	var bearer = "Bearer " + "eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9.eyJhdWQiOiIxIiwianRpIjoiM2NlMzRhNDM0OGNjMGRkMjczMWQyMDM0ZDQ4MzRkZTZiMTQ3MGI3ODE2YWQyM2RmMjRmMzg0YzE3ZjIzOWU1N2E5ZTg2N2E0ODhlMTg5YTEiLCJpYXQiOjE2MjY0MzU0MzYsIm5iZiI6MTYyNjQzNTQzNiwiZXhwIjoyMDk5ODIxMDM1LCJzdWIiOiIzIiwic2NvcGVzIjpbXX0.JtCQ_KStz4TluCkt_6JGJLmSGVhuY6dS_3OQ7KJicm8vSgYnfh2cwzrjjgoDU92u5RN2-fMHMji_ju6a4Lm33_nyj6_qclFV9SPRtO-UqMJe7EVkPhe0bP3co-9dVKyfUmSyi7GjVeHkUcD2OGG9m_zhu7krpwzQRBNiaNR9dJwCeBEbH1O13kKQItRl_V_-DDEtFF-bTnQ3DbnlEqZxtY4we6-qjpXmIrGmOU27pH5DUXZ8-cxqlAKP1ysBz_BJRBYGN0HZqYyL2AgrTG_k9sPds2CSyqPhbTvjm7yD5IxPOAcmasJbJoAPSyZecpNSecOL7JVsjB7UFcDPTdIy6zykIqJV6Zj-3qwkg4VrAt6iGvSIPCfSPzlydwk3o0znDHisp_9IDGuSTII49kAGnGb5Kw6WWsV9xQrXBtm6R41cwVAGc47r9j8tLux5PmlXdcrSxGS1uiiaMwZSx1ZdvZlC85f5LSpKiA0qP85acTX2R_Aav4oqsx_FN-UkBuBs8ADYC-sxMDVDuokr49IkkgVY9LUfkk8-pQX4IqKZKBOHuPAT1NsalgDPOZG9pFaIQ9kmt9Qm6TkkinNIPiwcBJ2mqHXziirtvQqylfrH2MBkXAofHK_-EEkOCAsARfFT41iw7wkJwW5ijliz5SC2ZiG6HTFS9WIG88WNiRzu9qc"
	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", bearer)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("Falha de conexão com o Host Snipeit.")
	}

	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)
	io.MultiReader()
	if err != nil {
		log.Println("Error on parsing response.\n[ERROR] -", err)
	}

	//log.Println("\nBilly Lowkão *easteregg*")

	//Variavel que contém os dados do Ativo Existente
	var responsevar UniversalGetT

	// Unmarshal do resultado do response
	err = json.Unmarshal(body, &responsevar)
	if err != nil {
		log.Printf("Reading body failed: %s", err)
	}

	//Armazena as informações selecioandas do Response na variável Struct de análise
	Analyser.Name = responsevar.Name
	Analyser.AssetTag = responsevar.AssetTag
	Analyser.ModelID = strconv.Itoa(responsevar.Model.ID)
	Analyser.StatusID = strconv.Itoa(responsevar.StatusLabel.ID)
	Analyser.SnipeitMema3Ria7 = responsevar.CustomFields.MemRia.Value
	Analyser.SnipeitSo8 = responsevar.CustomFields.SO.Value
	Analyser.SnipeitHd9 = responsevar.CustomFields.Hd.Value
	Analyser.SnipeitHostname10 = responsevar.CustomFields.Hostname.Value
	Analyser.SnipeitCPU11 = responsevar.CustomFields.CPU.Value
	Analyser.SnipeitModel12 = responsevar.CustomFields.Modelo.Value

	//Variável Array com as informacões da Variável Struct de análise
	var AnalyserIndex = []string{Analyser.Name, Analyser.AssetTag, Analyser.ModelID, Analyser.StatusID, Analyser.SnipeitMema3Ria7, Analyser.SnipeitSo8, Analyser.SnipeitHd9, Analyser.SnipeitHostname10, Analyser.SnipeitCPU11, Analyser.SnipeitModel12}

	//Variável Array com as informacões da Variável Struct do Ativo Criado
	var AtivoIndex = []string{ativo.Name, ativo.AssetTag, ativo.ModelID, ativo.StatusID, ativo.SnipeitMema3Ria7, ativo.SnipeitSo8, ativo.SnipeitHd9, ativo.SnipeitHostname10, ativo.SnipeitCPU11, ativo.SnipeitModel12}

	//Variavél Array que contém as alterações pendentes
	var Pending []string

	//Variável String que contém o príncipio do Patchrequest
	var Patchresquest string = "{\"requestable\":false,\"archived\":false"

	//Verifica as disparidades, destacando-as e criando o Patchrequest.
	if Analyser != ativo {
		log.Println("Disparidades encontradas.")
		fmt.Println("Disparidades encontradas!")
		fmt.Printf("Aprofundando análises")
		for i := 0; i < 4; i++ {
			time.Sleep(time.Second * 1)
			fmt.Print(".")
		}

		for i := 0; i < len(AnalyserIndex); i++ {
			if AnalyserIndex[i] != AtivoIndex[i] {
				fmt.Printf("\nDisparidade encontrada no Index[%v].\n", i)
				log.Printf("\nAtivo no invetário apresenta: %v\nEnquanto, novo ativo apresenta:%v\n", AnalyserIndex[i], AtivoIndex[i])
				fmt.Printf("\nAtivo no invetário apresenta: %v\nEnquanto, novo ativo apresenta:%v\n", AnalyserIndex[i], AtivoIndex[i])
				switch i {
				case 0:
					Patchresquest += ",\"name\":\"" + AtivoIndex[i] + "\""
				case 1:
					Patchresquest += ",\"asset_tag\":\"" + AtivoIndex[i] + "\""
				case 2:
					Patchresquest += ",\"model_id\":\"" + AtivoIndex[i] + "\""
				case 3:
					Patchresquest += ",\"status_id\":\"" + AtivoIndex[i] + "\""
				case 4:
					Patchresquest += ",\"_snipeit_mema3ria_7\":\"" + AtivoIndex[i] + "\""
				case 5:
					Patchresquest += ",\"_snipeit_so_8\":\"" + AtivoIndex[i] + "\""
				case 6:
					Patchresquest += ",\"_snipeit_hd_9\":\"" + AtivoIndex[i] + "\""
				case 7:
					Patchresquest += ",\"_snipeit_hostname_10\":\"" + AtivoIndex[i] + "\""
				case 8:
					Patchresquest += ",\"_snipeit_cpu_11\":\"" + AtivoIndex[i] + "\""
				case 9:
					Patchresquest += ",\"_snipeit_modelo_12\":\"" + AtivoIndex[i] + "\""

				}
				Pending = append(Pending, AtivoIndex[i])
				time.Sleep(time.Second * 3)
			} else {
				continue
			}
		}
		Patchresquest += "}"
		fmt.Printf("\nAlterações pendentes:\n%v\n", Pending)
		return Patchresquest, true
	} else {
		return Patchresquest, false
	}
}

//Envia alterações feitas no ativo existente no inventário através de seu ID.
func Patchbyid(id int, Patchresquest string) {
	strid := strconv.Itoa(id)
	url := "http://10.20.1.79:8001/api/v1/hardware/" + strid

	payload := strings.NewReader(Patchresquest)
	req, err := http.NewRequest("PATCH", url, payload)
	if err != nil {
		log.Fatalf("Request Error")
	}
	var bearer = "Bearer " + "eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9.eyJhdWQiOiIxIiwianRpIjoiM2NlMzRhNDM0OGNjMGRkMjczMWQyMDM0ZDQ4MzRkZTZiMTQ3MGI3ODE2YWQyM2RmMjRmMzg0YzE3ZjIzOWU1N2E5ZTg2N2E0ODhlMTg5YTEiLCJpYXQiOjE2MjY0MzU0MzYsIm5iZiI6MTYyNjQzNTQzNiwiZXhwIjoyMDk5ODIxMDM1LCJzdWIiOiIzIiwic2NvcGVzIjpbXX0.JtCQ_KStz4TluCkt_6JGJLmSGVhuY6dS_3OQ7KJicm8vSgYnfh2cwzrjjgoDU92u5RN2-fMHMji_ju6a4Lm33_nyj6_qclFV9SPRtO-UqMJe7EVkPhe0bP3co-9dVKyfUmSyi7GjVeHkUcD2OGG9m_zhu7krpwzQRBNiaNR9dJwCeBEbH1O13kKQItRl_V_-DDEtFF-bTnQ3DbnlEqZxtY4we6-qjpXmIrGmOU27pH5DUXZ8-cxqlAKP1ysBz_BJRBYGN0HZqYyL2AgrTG_k9sPds2CSyqPhbTvjm7yD5IxPOAcmasJbJoAPSyZecpNSecOL7JVsjB7UFcDPTdIy6zykIqJV6Zj-3qwkg4VrAt6iGvSIPCfSPzlydwk3o0znDHisp_9IDGuSTII49kAGnGb5Kw6WWsV9xQrXBtm6R41cwVAGc47r9j8tLux5PmlXdcrSxGS1uiiaMwZSx1ZdvZlC85f5LSpKiA0qP85acTX2R_Aav4oqsx_FN-UkBuBs8ADYC-sxMDVDuokr49IkkgVY9LUfkk8-pQX4IqKZKBOHuPAT1NsalgDPOZG9pFaIQ9kmt9Qm6TkkinNIPiwcBJ2mqHXziirtvQqylfrH2MBkXAofHK_-EEkOCAsARfFT41iw7wkJwW5ijliz5SC2ZiG6HTFS9WIG88WNiRzu9qc"
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", bearer)
	req.Header.Add("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("Falha de conexão com o Host Snipeit.")
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	// Unmarshal do resultado do response
	response := PatchRespose{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Printf("Reading body failed: %s", err)
	}

}

//Verifica se ativo existe procurando-o (GET) no inventário através do seu Asset Tag único.
func Verifybytag(assettag string) bool {

	url := "http://10.20.1.79:8001/api/v1/hardware/bytag/" + assettag
	var bearer = "Bearer " + "eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9.eyJhdWQiOiIxIiwianRpIjoiM2NlMzRhNDM0OGNjMGRkMjczMWQyMDM0ZDQ4MzRkZTZiMTQ3MGI3ODE2YWQyM2RmMjRmMzg0YzE3ZjIzOWU1N2E5ZTg2N2E0ODhlMTg5YTEiLCJpYXQiOjE2MjY0MzU0MzYsIm5iZiI6MTYyNjQzNTQzNiwiZXhwIjoyMDk5ODIxMDM1LCJzdWIiOiIzIiwic2NvcGVzIjpbXX0.JtCQ_KStz4TluCkt_6JGJLmSGVhuY6dS_3OQ7KJicm8vSgYnfh2cwzrjjgoDU92u5RN2-fMHMji_ju6a4Lm33_nyj6_qclFV9SPRtO-UqMJe7EVkPhe0bP3co-9dVKyfUmSyi7GjVeHkUcD2OGG9m_zhu7krpwzQRBNiaNR9dJwCeBEbH1O13kKQItRl_V_-DDEtFF-bTnQ3DbnlEqZxtY4we6-qjpXmIrGmOU27pH5DUXZ8-cxqlAKP1ysBz_BJRBYGN0HZqYyL2AgrTG_k9sPds2CSyqPhbTvjm7yD5IxPOAcmasJbJoAPSyZecpNSecOL7JVsjB7UFcDPTdIy6zykIqJV6Zj-3qwkg4VrAt6iGvSIPCfSPzlydwk3o0znDHisp_9IDGuSTII49kAGnGb5Kw6WWsV9xQrXBtm6R41cwVAGc47r9j8tLux5PmlXdcrSxGS1uiiaMwZSx1ZdvZlC85f5LSpKiA0qP85acTX2R_Aav4oqsx_FN-UkBuBs8ADYC-sxMDVDuokr49IkkgVY9LUfkk8-pQX4IqKZKBOHuPAT1NsalgDPOZG9pFaIQ9kmt9Qm6TkkinNIPiwcBJ2mqHXziirtvQqylfrH2MBkXAofHK_-EEkOCAsARfFT41iw7wkJwW5ijliz5SC2ZiG6HTFS9WIG88WNiRzu9qc"
	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", bearer)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("Falha de conexão com o Host Snipeit.")
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	// Unmarshal do resultado do response
	response := ErrorT{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Printf("Reading body failed: %s", err)
	}

	//tipo vazio para a comparação. (Se a Reponse for igual a ele, isto é, vazio, então ele retorna um false significando que não há erro)
	blankspace := ErrorT{}
	//Printando o Response
	if response == blankspace {
		return false

	} else {
		return true
	}
}

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
	if Verifybytag(mac.AssetTag) {
		log.Println("Os dados do Ativo Criado não constam no sistema.")
		fmt.Println("Enviando Ativo para o Snipeit ")

		MacOS.Snipesending(mac)
		log.Printf("NOVO ATIVO: %v", MacOS.Infos)
		log.Println("Ativo Criado enviado para o sistema.")

	} else {
		log.Println("Um Ativo semelhante foi encontrado no sistema.")
		fmt.Print("Asset Tag idêntico encontrado. Iniciando análise de disparidades")
		for i := 0; i < 4; i++ {
			time.Sleep(time.Second * 1)
			fmt.Print(".")
		}
		patch, boolean := Getbytag(mac.AssetTag, mac)
		if boolean {
			fmt.Println("\nPATCH necessário.")
			fmt.Println("\nExecutando PATCH RESQUEST.")
			time.Sleep(time.Second * 3)
			id := Getidbytag(mac.AssetTag)
			Patchbyid(id, patch)

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
	if Verifybytag(win.AssetTag) {
		log.Println("Os dados do Ativo Criado não constam no sistema.")
		fmt.Println("Enviando Ativo para o Snipeit ")

		MacOS.Snipesending(win)
		log.Printf("NOVO ATIVO: %v", Windows.Infos)
		log.Println("Ativo Criado enviado para o sistema.")

	} else {
		log.Println("Um Ativo semelhante foi encontrado no sistema.")
		fmt.Print("Asset Tag idêntico encontrado. Iniciando análise de disparidades")
		for i := 0; i < 4; i++ {
			time.Sleep(time.Second * 1)
			fmt.Print(".")
		}
		patch, boolean := Getbytag(win.AssetTag, win)
		if boolean {
			fmt.Println("\nPATCH necessário.")
			fmt.Println("\nExecutando PATCH RESQUEST.")
			time.Sleep(time.Second * 3)
			id := Getidbytag(win.AssetTag)
			Patchbyid(id, patch)

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
	var lin MacOS.MacOSt = MacOS.MacOSt{}

	//Populando Struct MacOSt
	lin.SnipeitCPU11 = Linux.Infos[0]
	//lin.SnipeitMema3Ria7 = Linux.Infos[1]
	lin.SnipeitSo8 = Linux.Infos[2]
	lin.SnipeitHostname10 = Linux.Infos[3]
	//lin.SnipeitHd9 = Linux.Infos[4]
	//lin.AssetTag = Linux.Infos[3]
	lin.Name = Linux.Infos[3]
	re1 := regexp.MustCompile(`(^\d{3}[,.]\d?)`)
	regexHD := re1.FindAllString(Linux.Infos[4], -1)
	interHD := strings.Join(regexHD, "")
	indexHD := strings.Split(interHD, ",")
	lin.SnipeitHd9 = strings.Join(indexHD, ".") + "GB"

	re2 := regexp.MustCompile(`\d`)
	regexmem := re2.FindAllString(Linux.Infos[1], -1)
	intermem := strings.Join(regexmem, "")
	lin.SnipeitMema3Ria7 = intermem + "GB"

	regextag := re2.FindAllString(Linux.Infos[3], -1)
	lin.AssetTag = strings.Join(regextag, "")

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
	if Verifybytag(lin.AssetTag) {
		log.Println("Os dados do Ativo Criado não constam no sistema.")
		fmt.Println("Enviando Ativo para o Snipeit ")

		MacOS.Snipesending(lin)
		log.Printf("NOVO ATIVO: %v", Linux.Infos)
		log.Println("Ativo Criado enviado para o sistema.")

	} else {
		log.Println("Um Ativo semelhante foi encontrado no sistema.")
		fmt.Print("Asset Tag idêntico encontrado. Iniciando análise de disparidades")
		for i := 0; i < 4; i++ {
			time.Sleep(time.Second * 1)
			fmt.Print(".")
		}
		patch, boolean := Getbytag(lin.AssetTag, lin)
		if boolean {
			fmt.Println("\nPATCH necessário.")
			fmt.Println("\nExecutando PATCH RESQUEST.")
			time.Sleep(time.Second * 3)
			id := Getidbytag(lin.AssetTag)
			Patchbyid(id, patch)

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
