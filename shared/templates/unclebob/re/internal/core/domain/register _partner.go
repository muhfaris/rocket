package domain

type DetailPartner struct {
	PartnerID string `params:"partner_id"`
}
type UpdatePartner struct {
	Email    string `json:"email"`
	Fullname string `json:"fullname"`
}
