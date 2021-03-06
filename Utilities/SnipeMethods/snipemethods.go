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

	"github.com/rodaine/table"
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

//Modelo de respose do m??todo PATCH
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

//Modelo de response do m??todo GET
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
		} `json:"Mem??ria"`
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
	//Define URL (link da API com IP do servidor + Assettag para localiza????o do Ativo)
	url := "http://" + IP + "/api/v1/hardware/bytag/" + assettag
	//C??digo de autentica????o
	var bearer = "Bearer " + "eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9.eyJhdWQiOiIxIiwianRpIjoiM2NlMzRhNDM0OGNjMGRkMjczMWQyMDM0ZDQ4MzRkZTZiMTQ3MGI3ODE2YWQyM2RmMjRmMzg0YzE3ZjIzOWU1N2E5ZTg2N2E0ODhlMTg5YTEiLCJpYXQiOjE2MjY0MzU0MzYsIm5iZiI6MTYyNjQzNTQzNiwiZXhwIjoyMDk5ODIxMDM1LCJzdWIiOiIzIiwic2NvcGVzIjpbXX0.JtCQ_KStz4TluCkt_6JGJLmSGVhuY6dS_3OQ7KJicm8vSgYnfh2cwzrjjgoDU92u5RN2-fMHMji_ju6a4Lm33_nyj6_qclFV9SPRtO-UqMJe7EVkPhe0bP3co-9dVKyfUmSyi7GjVeHkUcD2OGG9m_zhu7krpwzQRBNiaNR9dJwCeBEbH1O13kKQItRl_V_-DDEtFF-bTnQ3DbnlEqZxtY4we6-qjpXmIrGmOU27pH5DUXZ8-cxqlAKP1ysBz_BJRBYGN0HZqYyL2AgrTG_k9sPds2CSyqPhbTvjm7yD5IxPOAcmasJbJoAPSyZecpNSecOL7JVsjB7UFcDPTdIy6zykIqJV6Zj-3qwkg4VrAt6iGvSIPCfSPzlydwk3o0znDHisp_9IDGuSTII49kAGnGb5Kw6WWsV9xQrXBtm6R41cwVAGc47r9j8tLux5PmlXdcrSxGS1uiiaMwZSx1ZdvZlC85f5LSpKiA0qP85acTX2R_Aav4oqsx_FN-UkBuBs8ADYC-sxMDVDuokr49IkkgVY9LUfkk8-pQX4IqKZKBOHuPAT1NsalgDPOZG9pFaIQ9kmt9Qm6TkkinNIPiwcBJ2mqHXziirtvQqylfrH2MBkXAofHK_-EEkOCAsARfFT41iw7wkJwW5ijliz5SC2ZiG6HTFS9WIG88WNiRzu9qc"
	//REQUEST do GET
	req, _ := http.NewRequest("GET", url, nil)

	//HEADERs
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", bearer)

	//Comunica????o HTTP com o invent??rio
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("Falha de conex??o com o Host Snipeit.")
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

	//Recebimento do ID
	Id := response.ID

	return Id

}

/*GET

Busca as informa????es do Ativo existente no invent??rio Snipe it e compara com o Ativo criado ao executar o programa.
Ele recebe o Asset Tag ??nico do Ativo existente e a vari??vel que cont??m o tipo populado com as informa????es do Ativo criado.
Ao comparar ambos A. Existente e A. Criado ele destaca as disparidades e as retorna  em uma string Patchrequest, assim como um bool Needpatching que afirma se ?? necess??rio um PATCH ou n??o.

OBS: Patchrequest ?? um JSON padronizado especificamente para o envio atrav??s do m??todo PATCH.*/
func Getbytag(IP string, assettag string, ativo CollectionT, f io.Writer) (Patchrequest string, Needpatching bool) {

	//Define URL (link da API com IP do servidor + Assettag para localiza????o do Ativo)
	url := "http://" + IP + "/api/v1/hardware/bytag/" + assettag
	//C??digo de autentica????o
	var bearer = "Bearer " + "eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9.eyJhdWQiOiIxIiwianRpIjoiM2NlMzRhNDM0OGNjMGRkMjczMWQyMDM0ZDQ4MzRkZTZiMTQ3MGI3ODE2YWQyM2RmMjRmMzg0YzE3ZjIzOWU1N2E5ZTg2N2E0ODhlMTg5YTEiLCJpYXQiOjE2MjY0MzU0MzYsIm5iZiI6MTYyNjQzNTQzNiwiZXhwIjoyMDk5ODIxMDM1LCJzdWIiOiIzIiwic2NvcGVzIjpbXX0.JtCQ_KStz4TluCkt_6JGJLmSGVhuY6dS_3OQ7KJicm8vSgYnfh2cwzrjjgoDU92u5RN2-fMHMji_ju6a4Lm33_nyj6_qclFV9SPRtO-UqMJe7EVkPhe0bP3co-9dVKyfUmSyi7GjVeHkUcD2OGG9m_zhu7krpwzQRBNiaNR9dJwCeBEbH1O13kKQItRl_V_-DDEtFF-bTnQ3DbnlEqZxtY4we6-qjpXmIrGmOU27pH5DUXZ8-cxqlAKP1ysBz_BJRBYGN0HZqYyL2AgrTG_k9sPds2CSyqPhbTvjm7yD5IxPOAcmasJbJoAPSyZecpNSecOL7JVsjB7UFcDPTdIy6zykIqJV6Zj-3qwkg4VrAt6iGvSIPCfSPzlydwk3o0znDHisp_9IDGuSTII49kAGnGb5Kw6WWsV9xQrXBtm6R41cwVAGc47r9j8tLux5PmlXdcrSxGS1uiiaMwZSx1ZdvZlC85f5LSpKiA0qP85acTX2R_Aav4oqsx_FN-UkBuBs8ADYC-sxMDVDuokr49IkkgVY9LUfkk8-pQX4IqKZKBOHuPAT1NsalgDPOZG9pFaIQ9kmt9Qm6TkkinNIPiwcBJ2mqHXziirtvQqylfrH2MBkXAofHK_-EEkOCAsARfFT41iw7wkJwW5ijliz5SC2ZiG6HTFS9WIG88WNiRzu9qc"
	//REQUEST do GET
	req, _ := http.NewRequest("GET", url, nil)

	//HEADERs
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", bearer)

	//Comunica????o HTTP com o invent??rio
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("Falha de conex??o com o Host Snipeit.")
	}

	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)
	io.MultiReader()
	if err != nil {
		log.Println("Error on parsing response.\n[ERROR] -", err)
	}

	//Billy Lowk??o *easteregg*

	//Vari??vel Struct utilizada para a an??lise de disparidades entre Ativo Existente no invent??rio e Ativo Criado pela execu????o do programa
	var Analyser CollectionT = CollectionT{}

	//Variavel que cont??m os dados do Ativo Existente
	var responsevar UniversalGetT

	// Unmarshal do resultado do response
	err = json.Unmarshal(body, &responsevar)
	if err != nil {
		log.Printf("Reading body failed: %s", err)
	}

	//Armazena as informa????es selecioandas do Response na vari??vel Struct de an??lise
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

	//Vari??vel Array com as informa????es do Struct de an??lise
	var AnalyserIndex = []string{Analyser.Name, Analyser.AssetTag, Analyser.ModelID, Analyser.StatusID, Analyser.SnipeitMema3Ria7, Analyser.SnipeitSo8, Analyser.SnipeitHd9, Analyser.SnipeitHostname10, Analyser.SnipeitCPU11, Analyser.SnipeitModel12}

	//Vari??vel Array com as informa????es do Struct do Ativo Criado
	var AtivoIndex = []string{ativo.Name, ativo.AssetTag, ativo.ModelID, ativo.StatusID, ativo.SnipeitMema3Ria7, ativo.SnipeitSo8, ativo.SnipeitHd9, ativo.SnipeitHostname10, ativo.SnipeitCPU11, ativo.SnipeitModel12}

	//Variav??l Array que cont??m as altera????es pendentes
	var Pending []string

	//Vari??vel String que cont??m o pr??ncipio do Patchrequest
	var Patchresquest string = "{\"requestable\":false,\"archived\":false"

	//Verifica as disparidades, destacando-as e criando o Patchrequest.
	if Analyser != ativo {

		//Cria tabela com os Cabe??alhos "Ativo Existente", "Ativo Criado"
		tbl := table.New("Fieldname", "Ativo Existente", "Ativo Criado")

		//Implementa????o da formata????o
		tbl.WithWriter(f)

		fmt.Fprintln(f, "Disparidades encontradas.")

		//Analise de disparidades
		for i := 0; i < len(AnalyserIndex); i++ {
			if AnalyserIndex[i] != AtivoIndex[i] {
				var Fieldname string
				switch i {
				case 0:
					//Caso a disparidade seja encontrada no Index [0] do Array, ?? necess??rio PATCH no campo NAME
					Patchresquest += ",\"name\":\"" + AtivoIndex[i] + "\""
					Fieldname = "NAME"
				case 1:
					//Caso a disparidade seja encontrada no Index [1] do Array, ?? necess??rio PATCH no campo ASSET TAG
					Patchresquest += ",\"asset_tag\":\"" + AtivoIndex[i] + "\""
					Fieldname = "ASSET TAG"
				case 2:
					//Caso a disparidade seja encontrada no Index [2] do Array, ?? necess??rio PATCH no campo MODEL ID
					Patchresquest += ",\"model_id\":\"" + AtivoIndex[i] + "\""
					Fieldname = "MODEL ID"
				case 3:
					//Caso a disparidade seja encontrada no Index [3] do Array, ?? necess??rio PATCH no campo STATUS ID
					Patchresquest += ",\"status_id\":\"" + AtivoIndex[i] + "\""
					Fieldname = "STATUS ID"
				case 4:
					//Caso a disparidade seja encontrada no Index [4] do Array, ?? necess??rio PATCH no campo MEM??RIA
					Patchresquest += ",\"_snipeit_mema3ria_7\":\"" + AtivoIndex[i] + "\""
					Fieldname = "MEM??RIA"
				case 5:
					//Caso a disparidade seja encontrada no Index [5] do Array, ?? necess??rio PATCH no campo SISTEMA OPERACIONAL
					Patchresquest += ",\"_snipeit_so_8\":\"" + AtivoIndex[i] + "\""
					Fieldname = "SISTEMA OPERACIONAL"
				case 6:
					//Caso a disparidade seja encontrada no Index [6] do Array, ?? necess??rio PATCH no campo HD
					Patchresquest += ",\"_snipeit_hd_9\":\"" + AtivoIndex[i] + "\""
					Fieldname = "HD"
				case 7:
					//Caso a disparidade seja encontrada no Index [7] do Array, ?? necess??rio PATCH no campo HOSTNAME
					Patchresquest += ",\"_snipeit_hostname_10\":\"" + AtivoIndex[i] + "\""
					Fieldname = "HOSTNAME"
				case 8:
					//Caso a disparidade seja encontrada no Index [8] do Array, ?? necess??rio PATCH no campo CPU
					Patchresquest += ",\"_snipeit_cpu_11\":\"" + AtivoIndex[i] + "\""
					Fieldname = "CPU"
				case 9:
					//Caso a disparidade seja encontrada no Index [9] do Array, ?? necess??rio PATCH no campo MODELO
					Patchresquest += ",\"_snipeit_modelo_12\":\"" + AtivoIndex[i] + "\""
					Fieldname = "MODEL"
				}

				//Acrescenta informa????es a tabela
				tbl.AddRow(Fieldname, AnalyserIndex[i], AtivoIndex[i])

				//Acrescenta altera????es a uma lista de pend??ncias para expor visualmente depois
				Pending = append(Pending, AtivoIndex[i])
			} else {
				//Se n??o h?? disparidades, continue a an??lise
				continue
			}
		}

		//Fechamento do Patchresquest
		Patchresquest += "}"
		fmt.Printf("\nAltera????es pendentes:\n%v\n", Pending)
		//Caso haja altera????es,printe a tabela retorna true
		tbl.Print()
		return Patchresquest, true
	} else {
		//Caso n??o.. retorna false
		_, _ = fmt.Fprintf(f, "Nenhuma disparidade foi encontrada no Ativo...\n\n")
		_, _ = fmt.Fprintf(f, "")

		//Cria tabela com os Cabe??alhos "Fieldname" e "Ativo Existente"
		tbl := table.New("Fieldname", "Ativo Existente")

		//Implementa????o da formata????o
		tbl.WithWriter(f)

		for i := 0; i < len(AnalyserIndex); i++ {

			var Fieldname string
			switch i {
			case 0:
				Fieldname = "NAME"
			case 1:

				Fieldname = "ASSET TAG"
			case 2:

				Fieldname = "MODEL ID"
			case 3:

				Fieldname = "STATUS ID"
			case 4:

				Fieldname = "MEM??RIA"
			case 5:

				Fieldname = "SISTEMA OPERACIONAL"
			case 6:

				Fieldname = "HD"
			case 7:

				Fieldname = "HOSTNAME"
			case 8:

				Fieldname = "CPU"
			case 9:

				Fieldname = "MODEL"
			}

			//Acrescenta informa????es a tabela
			tbl.AddRow(Fieldname, AnalyserIndex[i])

		}

		//Exp??e tabela do Ativo Existente
		tbl.Print()

		return Patchresquest, false
	}
}

/*PATCH

Envia altera????es feitas no ativo existente no invent??rio atrav??s de seu ID.*/
func Patchbyid(id int, IP string, Patchresquest string) {
	//Converte ID de string para int
	strid := strconv.Itoa(id)
	//Define URL (link da API com IP do servidor + Assettag para localiza????o do Ativo)
	url := "http://" + IP + "/api/v1/hardware/" + strid

	payload := strings.NewReader(Patchresquest)
	//REQUEST do GET
	req, err := http.NewRequest("PATCH", url, payload)
	if err != nil {
		log.Fatalf("Request Error")
	}

	//C??digo de autentica????o
	var bearer = "Bearer " + "eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9.eyJhdWQiOiIxIiwianRpIjoiM2NlMzRhNDM0OGNjMGRkMjczMWQyMDM0ZDQ4MzRkZTZiMTQ3MGI3ODE2YWQyM2RmMjRmMzg0YzE3ZjIzOWU1N2E5ZTg2N2E0ODhlMTg5YTEiLCJpYXQiOjE2MjY0MzU0MzYsIm5iZiI6MTYyNjQzNTQzNiwiZXhwIjoyMDk5ODIxMDM1LCJzdWIiOiIzIiwic2NvcGVzIjpbXX0.JtCQ_KStz4TluCkt_6JGJLmSGVhuY6dS_3OQ7KJicm8vSgYnfh2cwzrjjgoDU92u5RN2-fMHMji_ju6a4Lm33_nyj6_qclFV9SPRtO-UqMJe7EVkPhe0bP3co-9dVKyfUmSyi7GjVeHkUcD2OGG9m_zhu7krpwzQRBNiaNR9dJwCeBEbH1O13kKQItRl_V_-DDEtFF-bTnQ3DbnlEqZxtY4we6-qjpXmIrGmOU27pH5DUXZ8-cxqlAKP1ysBz_BJRBYGN0HZqYyL2AgrTG_k9sPds2CSyqPhbTvjm7yD5IxPOAcmasJbJoAPSyZecpNSecOL7JVsjB7UFcDPTdIy6zykIqJV6Zj-3qwkg4VrAt6iGvSIPCfSPzlydwk3o0znDHisp_9IDGuSTII49kAGnGb5Kw6WWsV9xQrXBtm6R41cwVAGc47r9j8tLux5PmlXdcrSxGS1uiiaMwZSx1ZdvZlC85f5LSpKiA0qP85acTX2R_Aav4oqsx_FN-UkBuBs8ADYC-sxMDVDuokr49IkkgVY9LUfkk8-pQX4IqKZKBOHuPAT1NsalgDPOZG9pFaIQ9kmt9Qm6TkkinNIPiwcBJ2mqHXziirtvQqylfrH2MBkXAofHK_-EEkOCAsARfFT41iw7wkJwW5ijliz5SC2ZiG6HTFS9WIG88WNiRzu9qc"

	//HEADERs
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", bearer)
	req.Header.Add("Content-Type", "application/json")

	//Comunica????o HTTP com o invent??rio
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("Falha de conex??o com o Host Snipeit.")
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

Verifica se ativo existe procurando-o (GET) no invent??rio atrav??s do seu Asset Tag ??nico.*/
func Verifybytag(assettag string, IP string) bool {
	//Define URL (link da API com IP do servidor + Assettag para localiza????o do Ativo)
	url := "http://" + IP + "/api/v1/hardware/bytag/" + assettag

	//C??digo de autentica????o
	var bearer = "Bearer " + "eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9.eyJhdWQiOiIxIiwianRpIjoiM2NlMzRhNDM0OGNjMGRkMjczMWQyMDM0ZDQ4MzRkZTZiMTQ3MGI3ODE2YWQyM2RmMjRmMzg0YzE3ZjIzOWU1N2E5ZTg2N2E0ODhlMTg5YTEiLCJpYXQiOjE2MjY0MzU0MzYsIm5iZiI6MTYyNjQzNTQzNiwiZXhwIjoyMDk5ODIxMDM1LCJzdWIiOiIzIiwic2NvcGVzIjpbXX0.JtCQ_KStz4TluCkt_6JGJLmSGVhuY6dS_3OQ7KJicm8vSgYnfh2cwzrjjgoDU92u5RN2-fMHMji_ju6a4Lm33_nyj6_qclFV9SPRtO-UqMJe7EVkPhe0bP3co-9dVKyfUmSyi7GjVeHkUcD2OGG9m_zhu7krpwzQRBNiaNR9dJwCeBEbH1O13kKQItRl_V_-DDEtFF-bTnQ3DbnlEqZxtY4we6-qjpXmIrGmOU27pH5DUXZ8-cxqlAKP1ysBz_BJRBYGN0HZqYyL2AgrTG_k9sPds2CSyqPhbTvjm7yD5IxPOAcmasJbJoAPSyZecpNSecOL7JVsjB7UFcDPTdIy6zykIqJV6Zj-3qwkg4VrAt6iGvSIPCfSPzlydwk3o0znDHisp_9IDGuSTII49kAGnGb5Kw6WWsV9xQrXBtm6R41cwVAGc47r9j8tLux5PmlXdcrSxGS1uiiaMwZSx1ZdvZlC85f5LSpKiA0qP85acTX2R_Aav4oqsx_FN-UkBuBs8ADYC-sxMDVDuokr49IkkgVY9LUfkk8-pQX4IqKZKBOHuPAT1NsalgDPOZG9pFaIQ9kmt9Qm6TkkinNIPiwcBJ2mqHXziirtvQqylfrH2MBkXAofHK_-EEkOCAsARfFT41iw7wkJwW5ijliz5SC2ZiG6HTFS9WIG88WNiRzu9qc"

	//REQUEST do GET
	req, _ := http.NewRequest("GET", url, nil)

	//HEADERs
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", bearer)

	//Comunica????o HTTP com o invent??rio
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("Falha de conex??o com o Host Snipeit.")
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	// Unmarshal do resultado do response
	response := ErrorT{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Printf("Reading body failed: %s", err)
	}

	//tipo vazio para a compara????o. (Se a Reponse for igual a ele, isto ??, vazio, ent??o ele retorna um false significando que n??o h?? erro)
	blankspace := ErrorT{}
	//Printando o Response
	if response == blankspace {
		return false

	} else {
		return true
	}
}

/* POST

Envia os dados do computador para o invent??rio no Snipeit. (Essa fun????o recebe a variavel que recebe o tipo struct criado com os dados do computador)*/
func PostSnipe(Ativo CollectionT, IP string, f io.Writer) {

	var AtivoIndex = []string{Ativo.Name, Ativo.AssetTag, Ativo.ModelID, Ativo.StatusID, Ativo.SnipeitMema3Ria7, Ativo.SnipeitSo8, Ativo.SnipeitHd9, Ativo.SnipeitHostname10, Ativo.SnipeitCPU11, Ativo.SnipeitModel12}

	//URL da API SnipeIt
	url := "http://" + IP + "/api/v1/hardware"

	// Token de autentiuca????o
	var bearer = "Bearer " + "eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9.eyJhdWQiOiIxIiwianRpIjoiM2NlMzRhNDM0OGNjMGRkMjczMWQyMDM0ZDQ4MzRkZTZiMTQ3MGI3ODE2YWQyM2RmMjRmMzg0YzE3ZjIzOWU1N2E5ZTg2N2E0ODhlMTg5YTEiLCJpYXQiOjE2MjY0MzU0MzYsIm5iZiI6MTYyNjQzNTQzNiwiZXhwIjoyMDk5ODIxMDM1LCJzdWIiOiIzIiwic2NvcGVzIjpbXX0.JtCQ_KStz4TluCkt_6JGJLmSGVhuY6dS_3OQ7KJicm8vSgYnfh2cwzrjjgoDU92u5RN2-fMHMji_ju6a4Lm33_nyj6_qclFV9SPRtO-UqMJe7EVkPhe0bP3co-9dVKyfUmSyi7GjVeHkUcD2OGG9m_zhu7krpwzQRBNiaNR9dJwCeBEbH1O13kKQItRl_V_-DDEtFF-bTnQ3DbnlEqZxtY4we6-qjpXmIrGmOU27pH5DUXZ8-cxqlAKP1ysBz_BJRBYGN0HZqYyL2AgrTG_k9sPds2CSyqPhbTvjm7yD5IxPOAcmasJbJoAPSyZecpNSecOL7JVsjB7UFcDPTdIy6zykIqJV6Zj-3qwkg4VrAt6iGvSIPCfSPzlydwk3o0znDHisp_9IDGuSTII49kAGnGb5Kw6WWsV9xQrXBtm6R41cwVAGc47r9j8tLux5PmlXdcrSxGS1uiiaMwZSx1ZdvZlC85f5LSpKiA0qP85acTX2R_Aav4oqsx_FN-UkBuBs8ADYC-sxMDVDuokr49IkkgVY9LUfkk8-pQX4IqKZKBOHuPAT1NsalgDPOZG9pFaIQ9kmt9Qm6TkkinNIPiwcBJ2mqHXziirtvQqylfrH2MBkXAofHK_-EEkOCAsARfFT41iw7wkJwW5ijliz5SC2ZiG6HTFS9WIG88WNiRzu9qc"

	//transformando em bytes a vari??vel hw
	hardwarePostJSON, err := json.Marshal(Ativo)
	//Tratando o ocasoional erro transforma????o da vari??vel em byte
	if err != nil {
		panic(err)
	}

	//POST REQUEST
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(hardwarePostJSON))

	//Tratando o ocasoional erro do POST/REQUEST
	if err != nil {
		panic(err)
	}

	//adicionando os headers a autoriza????o
	req.Header.Add("Authorization", bearer)
	//definindo a formata????o do REQUEST
	req.Header.Add("Content-type", "application/json")

	// Send req using http Client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error on request.\n[ERROR] -", err)
	}

	//fechando o Response ap??s a conclus??o do c??digo
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

	//Cria tabela com os Cabe??alhos "Fieldname" e "Ativo Existente"
	tbl := table.New("Fieldname", "Novo Ativo")

	//Implementa????o da formata????o
	tbl.WithWriter(f)

	for i := 0; i < len(AtivoIndex); i++ {

		var Fieldname string
		switch i {
		case 0:
			Fieldname = "NAME"
		case 1:

			Fieldname = "ASSET TAG"
		case 2:

			Fieldname = "MODEL ID"
		case 3:

			Fieldname = "STATUS ID"
		case 4:

			Fieldname = "MEM??RIA"
		case 5:

			Fieldname = "SISTEMA OPERACIONAL"
		case 6:

			Fieldname = "HD"
		case 7:

			Fieldname = "HOSTNAME"
		case 8:

			Fieldname = "CPU"
		case 9:

			Fieldname = "MODEL"
		}

		//Acrescenta informa????es a tabela
		tbl.AddRow(Fieldname, AtivoIndex[i])

	}

	//Exp??e tabela do Ativo Existente
	tbl.Print()
	//Printando o Response
	fmt.Println("Response do POST:", response)
}
