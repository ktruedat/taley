package model

var CountrySeedData = []Country{
	{Name: "Romania"},
	{Name: "Moldova"},
}

var RegionSeedData = []Region{
	{Name: "Sud", CountryID: 1},
	{Name: "Nord", CountryID: 2},
}

var TimeOfDaySeedData = []TimeOfTheDay{
	{Name: "ziua", WeatherID: 1},
	{Name: "noaptea", WeatherID: 2},
}

var WeatherSeedData = []Weather{
	{Name: "cu cer albastru"},
	{Name: "cu ploaie"},
}
