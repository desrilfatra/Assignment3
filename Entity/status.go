package entity

type Status struct {
	Water       int    `json:"water"`
	Wind        int    `json:"wind"`
	StatusWater string `json:"statuswater"`
	StatusWind  string `json:"statuswind"`
}
