package dto

type Customization struct{
	Type string `json:"type"`
	Value string `json:"value"`
}

type ImageGenerateRequest struct {
	Template string `json:"template"`
	Customization []Customization `json:"customization"`
}

type ImageGenerateReponse struct{
	URL string `json:"url"`
}