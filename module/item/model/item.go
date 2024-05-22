package Profile

type Profile struct {
	Id        int    `json:"id" gorm:"column:id"`
	Email     string `json:"email" gorm:"column:email"`
	FirstName string `json:"first_name" gorm:"column:first_name"`
	LastName  string `json:"last_name" gorm:"column:last_name"`
	Avatar    string `json:"avatar" gorm:"column:avatar"`
}

func (Profile) TableName() string { return "profile" }

func (item Profile) Validate() error {
	return nil
}

type DataPaging struct {
	Page  int   `json:"page" form:"page"`
	Limit int   `json:"limit" form:"limit"`
	Total int64 `json:"total" form:"-"`
}

func (p *DataPaging) Process() {
	if p.Page <= 0 {
		p.Page = 1
	}

	if p.Limit <= 0 {
		p.Limit = 5
	}
}
