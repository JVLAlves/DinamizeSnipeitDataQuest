package snipe

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

//Modelo para coleta e envio de dados do computador.
type CollectionT struct {
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

//Modelo geral de RESPONSE
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
			ID                     int    `json:"id"`
			Name                   string `json:"name"`
			ModelNumber            string `json:"model_number"`
			ManufacturerID         int    `json:"manufacturer_id"`
			CategoryID             int    `json:"category_id"`
			CreatedAt              string `json:"created_at"`
			UpdatedAt              string `json:"updated_at"`
			DepreciationID         int    `json:"depreciation_id"`
			Eol                    int    `json:"eol"`
			Image                  string `json:"image"`
			DeprecatedAtivoAddress int    `json:"deprecated_Ativo_address"`
			FieldsetID             int    `json:"fieldset_id"`
			Notes                  string `json:"notes"`
			Requestable            int    `json:"requestable"`
		} `json:"model"`
	} `json:"payload"`
}

//Modelo de respose do método PATCH
type PatchRespose struct {
	Status   string `json:"status"`
	Messages string `json:"messages"`
	Payload  struct {
		ID                   int    `json:"id"`
		Name                 string `json:"name"`
		AssetTag             string `json:"asset_tag"`
		ModelID              int    `json:"model_id"`
		Serial               string `json:"serial"`
		PurchaseDate         string `json:"purchase_date"`
		PurchaseCost         string `json:"purchase_cost"`
		OrderNumber          string `json:"order_number"`
		AssignedTo           string `json:"assigned_to"`
		Notes                string `json:"notes"`
		Image                string `json:"image"`
		UserID               int    `json:"user_id"`
		CreatedAt            string `json:"created_at"`
		UpdatedAt            string `json:"updated_at"`
		Physical             int    `json:"physical"`
		DeletedAt            string `json:"deleted_at"`
		StatusID             int    `json:"status_id"`
		Archived             int    `json:"archived"`
		WarrantyMonths       string `json:"warranty_months"`
		Depreciate           int    `json:"depreciate"`
		SupplierID           int    `json:"supplier_id"`
		Requestable          bool   `json:"requestable"`
		RtdLocationID        int    `json:"rtd_location_id"`
		Accepted             string `json:"accepted"`
		LastCheckout         string `json:"last_checkout"`
		ExpectedCheckin      string `json:"expected_checkin"`
		CompanyID            string `json:"company_id"`
		AssignedType         string `json:"assigned_type"`
		LastAuditDate        string `json:"last_audit_date"`
		NextAuditDate        string `json:"next_audit_date"`
		LocationID           int    `json:"location_id"`
		CheckinCounter       int    `json:"checkin_counter"`
		CheckoutCounter      int    `json:"checkout_counter"`
		RequestsCounter      int    `json:"requests_counter"`
		SnipeitImei1         string `json:"_snipeit_imei_1"`
		SnipeitPhoneNumber2  string `json:"_snipeit_phone_number_2"`
		SnipeitRAM3          string `json:"_snipeit_ram_3"`
		SnipeitCPU4          string `json:"_snipeit_cpu_4"`
		SnipeitAtivoAddress5 string `json:"_snipeit_Ativo_address_5"`
	} `json:"payload"`
}

//Modelo de response do método GET
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

//Modelo de reponse de ERRO
type ErrorT struct {
	Status   string `json:"status"`
	Messages string `json:"messages"`
	Payload  string `json:"payload"`
}

//Modelo exclusivo para ID
type IDT struct {
	ID int `json:"id"`
}

/*GET

Busca o ID do Ativo Existente.*/
func Getidbytag(assettag string, IP string) (ID int) {
	url := "http://" + IP + "/api/v1/hardware/bytag/" + assettag
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

/*GET

Busca as informações do Ativo existente no inventário Snipe it e compara com o Ativo criado ao executar o programa.
Ele recebe o Asset Tag único do Ativo existente e a variável que contém o tipo populado com as informações do Ativo criado.
Ao comparar ambos A. Existente e A. Criado ele destaca as disparidades e as retorna  em uma string Patchrequest, assim como um bool Needpatching que afirma se é necessário um PATCH ou não.

OBS: Patchrequest é um JSON padronizado especificamente para o envio através do método PATCH.*/
func Getbytag(IP string, assettag string, ativo CollectionT) (Patchrequest string, Needpatching bool) {
	url := "http://" + IP + "/api/v1/hardware/bytag/" + assettag
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

	//Variável Struct utilizada para a análise de disparidades entre Ativo Existente no inventário e Ativo Criado pela execução do programa
	var Analyser CollectionT = CollectionT{}

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

	//Variável Array com as inforAtivoões da Variável Struct de análise
	var AnalyserIndex = []string{Analyser.Name, Analyser.AssetTag, Analyser.ModelID, Analyser.StatusID, Analyser.SnipeitMema3Ria7, Analyser.SnipeitSo8, Analyser.SnipeitHd9, Analyser.SnipeitHostname10, Analyser.SnipeitCPU11, Analyser.SnipeitModel12}

	//Variável Array com as inforAtivoões da Variável Struct do Ativo Criado
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

/*PATCH

Envia alterações feitas no ativo existente no inventário através de seu ID.*/
func Patchbyid(id int, IP string, Patchresquest string) {
	strid := strconv.Itoa(id)
	url := "http://" + IP + "/api/v1/hardware/" + strid

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

/*GET

Verifica se ativo existe procurando-o (GET) no inventário através do seu Asset Tag único.*/
func Verifybytag(assettag string, IP string) bool {

	url := "http://" + IP + "/api/v1/hardware/bytag/" + assettag
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

/* POST

Envia os dados do computador para o inventário no Snipeit. (Essa função recebe a variavel que recebe o tipo struct criado com os dados do computador)*/
func PostSnipe(Ativo CollectionT, IP string) {

	//URL da API SnipeIt
	url := "http://" + IP + "/api/v1/hardware"

	// Token de autentiucação
	var bearer = "Bearer " + "eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9.eyJhdWQiOiIxIiwianRpIjoiM2NlMzRhNDM0OGNjMGRkMjczMWQyMDM0ZDQ4MzRkZTZiMTQ3MGI3ODE2YWQyM2RmMjRmMzg0YzE3ZjIzOWU1N2E5ZTg2N2E0ODhlMTg5YTEiLCJpYXQiOjE2MjY0MzU0MzYsIm5iZiI6MTYyNjQzNTQzNiwiZXhwIjoyMDk5ODIxMDM1LCJzdWIiOiIzIiwic2NvcGVzIjpbXX0.JtCQ_KStz4TluCkt_6JGJLmSGVhuY6dS_3OQ7KJicm8vSgYnfh2cwzrjjgoDU92u5RN2-fMHMji_ju6a4Lm33_nyj6_qclFV9SPRtO-UqMJe7EVkPhe0bP3co-9dVKyfUmSyi7GjVeHkUcD2OGG9m_zhu7krpwzQRBNiaNR9dJwCeBEbH1O13kKQItRl_V_-DDEtFF-bTnQ3DbnlEqZxtY4we6-qjpXmIrGmOU27pH5DUXZ8-cxqlAKP1ysBz_BJRBYGN0HZqYyL2AgrTG_k9sPds2CSyqPhbTvjm7yD5IxPOAcmasJbJoAPSyZecpNSecOL7JVsjB7UFcDPTdIy6zykIqJV6Zj-3qwkg4VrAt6iGvSIPCfSPzlydwk3o0znDHisp_9IDGuSTII49kAGnGb5Kw6WWsV9xQrXBtm6R41cwVAGc47r9j8tLux5PmlXdcrSxGS1uiiaMwZSx1ZdvZlC85f5LSpKiA0qP85acTX2R_Aav4oqsx_FN-UkBuBs8ADYC-sxMDVDuokr49IkkgVY9LUfkk8-pQX4IqKZKBOHuPAT1NsalgDPOZG9pFaIQ9kmt9Qm6TkkinNIPiwcBJ2mqHXziirtvQqylfrH2MBkXAofHK_-EEkOCAsARfFT41iw7wkJwW5ijliz5SC2ZiG6HTFS9WIG88WNiRzu9qc"

	//transformando em bytes a variável hw
	hardwarePostJSON, err := json.Marshal(Ativo)
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
}
