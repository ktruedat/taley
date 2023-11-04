package model

type Tale struct {
	ID         uint   `gorm:"primaryKey"`
	Title      string `gorm:"not null;unique" json:"title"`
	CategoryID uint   `gorm:"not null;" json:"category_id"` // Foreign key to Category
	Contents   string `gorm:"not null;" json:"contents"`
	UserID     uint   `gorm:"not null;" json:"user_id"`
	Likes      uint   `gorm:"not null;default:0" json:"likes"`
	User       User   `gorm:"foreignKey:UserID"`
}

type TaleRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Category uint   `json:"category"`
	Title    string `json:"title"`
	Contents string `json:"contents"`
}

type Category struct {
	ID    uint   `gorm:"primaryKey"`
	Title string `gorm:"not null;unique" json:"title"`
	Tales []Tale // One-to-Many relationship: a Category can have multiple Tales
}

var CategorySeedData = []Category{
	{Title: "Romantice"},
	{Title: "Triste"},
	{Title: "Haioase"},
	{Title: "Cu invățătură"},
}
