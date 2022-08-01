package models

type User struct {
	ID           uint   `json:"id" gorm:"primary_key"`
	FirstName    string `json:"FirstName"`
	LastName     string `json:"LastName"`
	UserName     string `json:"UserName"`
	Email        string `json:"Email"`
	PasswordHash string `json:"PasswordHash"`
}
type Kuisioner struct {
	ID             uint   `json:"id" gorm:"primary_key"`
	JudulKuisioner string `json:"JudulKuisioner"`
	IsiKuisioner   string `json:"IsiKuisioner"`
	//UserName string `gorm:"foreignKey:UserRefer"`
}
