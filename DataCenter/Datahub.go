package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/JVLAlves/DinamizeSnipeitDataQuest/DataQuest/Linux"
	"github.com/JVLAlves/DinamizeSnipeitDataQuest/DataQuest/MacOS"
	"github.com/JVLAlves/DinamizeSnipeitDataQuest/DataQuest/Windows"
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

func Getidbytag(assettag string) int {
	url := "http://10.20.1.79:8001/api/v1/hardware/bytag/" + assettag
	var bearer = "Bearer " + "eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9.eyJhdWQiOiIxIiwianRpIjoiM2NlMzRhNDM0OGNjMGRkMjczMWQyMDM0ZDQ4MzRkZTZiMTQ3MGI3ODE2YWQyM2RmMjRmMzg0YzE3ZjIzOWU1N2E5ZTg2N2E0ODhlMTg5YTEiLCJpYXQiOjE2MjY0MzU0MzYsIm5iZiI6MTYyNjQzNTQzNiwiZXhwIjoyMDk5ODIxMDM1LCJzdWIiOiIzIiwic2NvcGVzIjpbXX0.JtCQ_KStz4TluCkt_6JGJLmSGVhuY6dS_3OQ7KJicm8vSgYnfh2cwzrjjgoDU92u5RN2-fMHMji_ju6a4Lm33_nyj6_qclFV9SPRtO-UqMJe7EVkPhe0bP3co-9dVKyfUmSyi7GjVeHkUcD2OGG9m_zhu7krpwzQRBNiaNR9dJwCeBEbH1O13kKQItRl_V_-DDEtFF-bTnQ3DbnlEqZxtY4we6-qjpXmIrGmOU27pH5DUXZ8-cxqlAKP1ysBz_BJRBYGN0HZqYyL2AgrTG_k9sPds2CSyqPhbTvjm7yD5IxPOAcmasJbJoAPSyZecpNSecOL7JVsjB7UFcDPTdIy6zykIqJV6Zj-3qwkg4VrAt6iGvSIPCfSPzlydwk3o0znDHisp_9IDGuSTII49kAGnGb5Kw6WWsV9xQrXBtm6R41cwVAGc47r9j8tLux5PmlXdcrSxGS1uiiaMwZSx1ZdvZlC85f5LSpKiA0qP85acTX2R_Aav4oqsx_FN-UkBuBs8ADYC-sxMDVDuokr49IkkgVY9LUfkk8-pQX4IqKZKBOHuPAT1NsalgDPOZG9pFaIQ9kmt9Qm6TkkinNIPiwcBJ2mqHXziirtvQqylfrH2MBkXAofHK_-EEkOCAsARfFT41iw7wkJwW5ijliz5SC2ZiG6HTFS9WIG88WNiRzu9qc"
	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", bearer)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("error (1): %s", err)
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Println("Error on parsing response.\n[ERROR] -", err)
	}

	//coventendo Body em bytes para Body em String
	bodyString := string(body)
	log.Print(bodyString)

	// Unmarshal do resultado do response
	response := IDT{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Printf("Reading body failed: %s", err)
	}

	Id := response.ID

	return Id

}

var Analyser MacOS.MacOSt = MacOS.MacOSt{}

func Getbytag(assettag string, ativo MacOS.MacOSt) (string, bool) {
	url := "http://10.20.1.79:8001/api/v1/hardware/bytag/" + assettag
	var bearer = "Bearer " + "eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9.eyJhdWQiOiIxIiwianRpIjoiM2NlMzRhNDM0OGNjMGRkMjczMWQyMDM0ZDQ4MzRkZTZiMTQ3MGI3ODE2YWQyM2RmMjRmMzg0YzE3ZjIzOWU1N2E5ZTg2N2E0ODhlMTg5YTEiLCJpYXQiOjE2MjY0MzU0MzYsIm5iZiI6MTYyNjQzNTQzNiwiZXhwIjoyMDk5ODIxMDM1LCJzdWIiOiIzIiwic2NvcGVzIjpbXX0.JtCQ_KStz4TluCkt_6JGJLmSGVhuY6dS_3OQ7KJicm8vSgYnfh2cwzrjjgoDU92u5RN2-fMHMji_ju6a4Lm33_nyj6_qclFV9SPRtO-UqMJe7EVkPhe0bP3co-9dVKyfUmSyi7GjVeHkUcD2OGG9m_zhu7krpwzQRBNiaNR9dJwCeBEbH1O13kKQItRl_V_-DDEtFF-bTnQ3DbnlEqZxtY4we6-qjpXmIrGmOU27pH5DUXZ8-cxqlAKP1ysBz_BJRBYGN0HZqYyL2AgrTG_k9sPds2CSyqPhbTvjm7yD5IxPOAcmasJbJoAPSyZecpNSecOL7JVsjB7UFcDPTdIy6zykIqJV6Zj-3qwkg4VrAt6iGvSIPCfSPzlydwk3o0znDHisp_9IDGuSTII49kAGnGb5Kw6WWsV9xQrXBtm6R41cwVAGc47r9j8tLux5PmlXdcrSxGS1uiiaMwZSx1ZdvZlC85f5LSpKiA0qP85acTX2R_Aav4oqsx_FN-UkBuBs8ADYC-sxMDVDuokr49IkkgVY9LUfkk8-pQX4IqKZKBOHuPAT1NsalgDPOZG9pFaIQ9kmt9Qm6TkkinNIPiwcBJ2mqHXziirtvQqylfrH2MBkXAofHK_-EEkOCAsARfFT41iw7wkJwW5ijliz5SC2ZiG6HTFS9WIG88WNiRzu9qc"
	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", bearer)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("error (1): %s", err)
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Println("Error on parsing response.\n[ERROR] -", err)
	}

	//coventendo Body em bytes para Body em String
	bodyString := string(body)
	log.Print("Billy Lowkão *easteregg*")
	log.Print(bodyString)

	// Unmarshal do resultado do response
	var responsevar UniversalGetT //= MacOS.SnipeITHardwareResponseT{}
	//responsevar =

	err = json.Unmarshal(body, &responsevar)
	if err != nil {
		log.Printf("Reading body failed: %s", err)
	}

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

	var AnalyserIndex = []string{Analyser.Name, Analyser.AssetTag, Analyser.ModelID, Analyser.StatusID, Analyser.SnipeitMema3Ria7, Analyser.SnipeitSo8, Analyser.SnipeitHd9, Analyser.SnipeitHostname10, Analyser.SnipeitCPU11, Analyser.SnipeitModel12}

	var AtivoIndex = []string{ativo.Name, ativo.AssetTag, ativo.ModelID, ativo.StatusID, ativo.SnipeitMema3Ria7, ativo.SnipeitSo8, ativo.SnipeitHd9, ativo.SnipeitHostname10, ativo.SnipeitCPU11, ativo.SnipeitModel12}

	var Pending []string

	var Patchresquest string = "{\"requestable\":false,\"archived\":false"

	if Analyser != ativo {
		fmt.Println("Disparidades encontradas!")
		fmt.Printf("Aprofundando análises")
		for i := 0; i < 4; i++ {
			time.Sleep(time.Second * 1)
			fmt.Print(".")
		}

		for i := 0; i < len(AnalyserIndex); i++ {
			if AnalyserIndex[i] != AtivoIndex[i] {
				fmt.Printf("\nDisparidade encontrada no Index[%v].\n", i)
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
		fmt.Println("err", err)
		log.Fatalf("Client Error")
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	//coventendo Body em bytes para Body em String
	bodyString := string(body)
	log.Print(bodyString)

	// Unmarshal do resultado do response
	response := PatchRespose{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Printf("Reading body failed: %s", err)
	}

}

func Verifybytag(assettag string) bool {

	url := "http://10.20.1.79:8001/api/v1/hardware/bytag/" + assettag
	var bearer = "Bearer " + "eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9.eyJhdWQiOiIxIiwianRpIjoiM2NlMzRhNDM0OGNjMGRkMjczMWQyMDM0ZDQ4MzRkZTZiMTQ3MGI3ODE2YWQyM2RmMjRmMzg0YzE3ZjIzOWU1N2E5ZTg2N2E0ODhlMTg5YTEiLCJpYXQiOjE2MjY0MzU0MzYsIm5iZiI6MTYyNjQzNTQzNiwiZXhwIjoyMDk5ODIxMDM1LCJzdWIiOiIzIiwic2NvcGVzIjpbXX0.JtCQ_KStz4TluCkt_6JGJLmSGVhuY6dS_3OQ7KJicm8vSgYnfh2cwzrjjgoDU92u5RN2-fMHMji_ju6a4Lm33_nyj6_qclFV9SPRtO-UqMJe7EVkPhe0bP3co-9dVKyfUmSyi7GjVeHkUcD2OGG9m_zhu7krpwzQRBNiaNR9dJwCeBEbH1O13kKQItRl_V_-DDEtFF-bTnQ3DbnlEqZxtY4we6-qjpXmIrGmOU27pH5DUXZ8-cxqlAKP1ysBz_BJRBYGN0HZqYyL2AgrTG_k9sPds2CSyqPhbTvjm7yD5IxPOAcmasJbJoAPSyZecpNSecOL7JVsjB7UFcDPTdIy6zykIqJV6Zj-3qwkg4VrAt6iGvSIPCfSPzlydwk3o0znDHisp_9IDGuSTII49kAGnGb5Kw6WWsV9xQrXBtm6R41cwVAGc47r9j8tLux5PmlXdcrSxGS1uiiaMwZSx1ZdvZlC85f5LSpKiA0qP85acTX2R_Aav4oqsx_FN-UkBuBs8ADYC-sxMDVDuokr49IkkgVY9LUfkk8-pQX4IqKZKBOHuPAT1NsalgDPOZG9pFaIQ9kmt9Qm6TkkinNIPiwcBJ2mqHXziirtvQqylfrH2MBkXAofHK_-EEkOCAsARfFT41iw7wkJwW5ijliz5SC2ZiG6HTFS9WIG88WNiRzu9qc"
	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", bearer)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("error (1): %s", err)
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	//coventendo Body em bytes para Body em String
	bodyString := string(body)
	log.Print(bodyString)

	// Unmarshal do resultado do response
	response := ErrorT{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Printf("Reading body failed: %s", err)
	}

	blankspace := ErrorT{}
	//Printando o Response
	if response == blankspace {
		return false

	} else {
		return true
	}
}

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
	fmt.Println("Digite o Tipo de Ativo (Exemplo:Desktop/MacBook/Notebook): ")
	fmt.Scanf("%v", IDmodelo)

	//identificando o Modelo
	switch *IDmodelo {
	case "Notebook":
		*IDmodelo = "6"
	case "Desktop":
		*IDmodelo = "8"
	case "MacBook":
		*IDmodelo = "22"
	default:
		*IDmodelo = "11"
	}

	//Status ID
	*IDstatus = "5"

	//Inputs Manuais: Modelo do Ativo
	*modeloAtivo = "DNZ-COMPUTER"

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

	if Verifybytag(mac.AssetTag) {
		fmt.Println("Esse ativo ainda não consta no inventário.\nVocê deseja enviar essas informações para o inventário Snipeit? (sim/nao)")
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

	} else {
		fmt.Print("Asset Tag idêntico encontrado. Iniciando análise de disparidades")
		for i := 0; i < 4; i++ {
			time.Sleep(time.Second * 1)
			fmt.Print(".")
		}
		patch, boolean := Getbytag(mac.AssetTag, mac)
		if boolean {
			fmt.Println("PATCH necessário.")
			fmt.Println("Executando PATCH RESQUEST.")
			time.Sleep(time.Second * 3)
			id := Getidbytag(mac.AssetTag)
			Patchbyid(id, patch)
		} else {
			fmt.Println("Sem alterações")
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
	fmt.Println("Digite o Tipo de Ativo (Exemplo:Desktop/MacBook/Notebook): ")
	fmt.Scanf("%v", IDmodelo)

	//identificando o Modelo
	switch *IDmodelo {
	case "Notebook":
		*IDmodelo = "6"
	case "Desktop":
		*IDmodelo = "8"
	case "MacBook":
		*IDmodelo = "22"
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

	if Verifybytag(win.AssetTag) {
		fmt.Println("Esse ativo ainda não consta no inventário.\nVocê deseja enviar essas informações para o inventário Snipeit? (sim/nao)")
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

	} else {
		fmt.Print("Asset Tag idêntico encontrado. Iniciando análise de disparidades")
		for i := 0; i < 4; i++ {
			time.Sleep(time.Second * 1)
			fmt.Print(".")
		}
		patch, boolean := Getbytag(win.AssetTag, win)
		if boolean {
			fmt.Println("PATCH necessário.")
			fmt.Println("Executando PATCH RESQUEST.")
			time.Sleep(time.Second * 3)
			id := Getidbytag(win.AssetTag)
			Patchbyid(id, patch)

		} else {
			fmt.Println("Sem alterações")
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
	fmt.Println("Digite o Tipo de Ativo (Exemplo:Desktop/MacBook/Notebook): ")
	fmt.Scanf("%v", IDmodelo)

	//identificando o Modelo
	switch *IDmodelo {
	case "Notebook":
		*IDmodelo = "6"
		*modeloAtivo = "DNZ-Notebook"
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

	if Verifybytag(lin.AssetTag) {
		fmt.Println("Esse ativo ainda não consta no inventário.\nVocê deseja enviar essas informações para o inventário Snipeit? (sim/nao)")
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

	} else {
		fmt.Print("Asset Tag idêntico encontrado. Iniciando análise de disparidades")
		for i := 0; i < 4; i++ {
			time.Sleep(time.Second * 1)
			fmt.Print(".")
		}
		patch, boolean := Getbytag(lin.AssetTag, lin)
		if boolean {
			fmt.Println("PATCH necessário.")
			fmt.Println("Executando PATCH RESQUEST.")
			time.Sleep(time.Second * 3)
			id := Getidbytag(lin.AssetTag)
			Patchbyid(id, patch)

		} else {
			fmt.Println("Sem alterações")
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
