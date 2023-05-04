package response

import "time"

type Form struct {
	ID                   int       `json:"id"`
	First_name           string    `json:"firstname"`
	Last_name            string    `json:"lastname"`
	Email                string    `json:"email"`
	Gender               string    `json:"gender" gorm:"type:char(1)"`
	Marital_status       string    `json:"maritalstatus"  gorm:"type:char(1)"`
	Phone                int64     `json:"phone"`
	Date_of_birth        time.Time `json:"dateofbirth"`
	P_address            string    `json:"paddress"`
	C_address            string    `json:"caddress"`
	Account_no           string    `json:"accno"`
	Ifsc_code            string    `json:"ifsccode"`
	Name_as_per_passbokk string    `json:"nameaspass"`
	Pan_number           string    `json:"pannumber"`
	Adhaar_no            string    `json:"adhaarnumber"`
	Photo                string    `json:"photo"`
}
