package model

type Country struct {
	ID         uint   `gorm:"primaryKey"`
	Name       string `gorm:"not null;unique" json:"name"`
	Regions    []Region
	References []References
}

type Region struct {
	ID         uint   `gorm:"primaryKey"`
	Name       string `gorm:"not null;unique" json:"name"`
	CountryID  uint
	References []References
}

type TimeOfTheDay struct {
	ID         uint   `gorm:"primaryKey"`
	Name       string `gorm:"not null;unique" json:"name"`
	WeatherID  uint
	References []References
}

type Weather struct {
	ID           uint   `gorm:"primaryKey"`
	Name         string `gorm:"not null;unique" json:"name"`
	TimeOfTheDay []TimeOfTheDay
	References   []References
}

type References struct {
	ID             uint         `gorm:"primaryKey"`
	RegionID       uint         `json:"regionID"`
	Region         Region       `gorm:"foreignKey:RegionID" json:"region"`
	CountryID      uint         `json:"countryID"`
	Country        Country      `gorm:"foreignKey:CountryID" json:"country"`
	TimeOfTheDayID uint         `json:"timeOfTheDayID"`
	TimeOfTheDay   TimeOfTheDay `gorm:"foreignKey:TimeOfTheDayID" json:"timeOfTheDay"`
	WeatherID      uint         `json:"WeatherID"`
	Weather        Weather      `gorm:"foreignKey:WeatherID" json:"weather"`
}
