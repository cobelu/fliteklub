package model

type Membership struct {
	ID     int  `json:"id"`
	UserID int  `gorm:"column:user_id"`
	User   User `json:"user"`
	ClubID int  `gorm:"column:club_id"`
	Club   Club `json:"club"`
}

func (m Membership) TableName() string {
	return "membership"
}

func ExampleMembership() Membership {
	return Membership{
		ID:   1,
		User: ExampleUser(),
		Club: ExampleClub(),
	}
}
