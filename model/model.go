package model

type Web struct {
	StatusWater string `json: "statuswater"`
	StatusWind  string `json: "statuswind"`
	Water       int    `json: "water"`
	Wind        int    `json: "wind"`
}
