package dpfm_api_input_reader

type SDC struct {
	ConnectionKey     string   `json:"connection_key"`
	Result            bool     `json:"result"`
	RedisKey          string   `json:"redis_key"`
	Filepath          string   `json:"filepath"`
	APIStatusCode     int      `json:"api_status_code"`
	RuntimeSessionID  string   `json:"runtime_session_id"`
	BusinessPartnerID *int     `json:"business_partner"`
	ServiceLabel      string   `json:"service_label"`
	APIType           string   `json:"APIType"`
	Message           Message  `json:"message"`
	APISchema         string   `json:"api_schema"`
	Accepter          []string `json:"accepter"`
	Deleted           bool     `json:"deleted"`
}

type Message struct {
	Partner *[]Partner `json:"Partner"`
}

type Partner struct {
	PartnerFunction         string `json:"PartnerFunction"`
	BusinessPartner         int    `json:"BusinessPartner"`
	BusinessPartnerFullName string `json:"BusinessPartnerFullName"`
	BusinessPartnerName     string `json:"BusinessPartnerName"`
	Organization            string `json:"Organization"`
	Country                 string `json:"Country"`
	Language                string `json:"Language"`
	Currency                string `json:"Currency"`
	ExternalDocumentID      string `json:"ExternalDocumentID"`
	AddressID               int    `json:"AddressID"`
	EmailAddress            string `json:"EmailAddress"`
}
