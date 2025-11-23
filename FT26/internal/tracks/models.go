package tracks

type Track struct {
    ID          string  `json:"id"`
    Name        string  `json:"name"`
    Country     string  `json:"country"`
    LengthKm    float64 `json:"length_km"`
    Turns       int     `json:"turns"`
    CreatedYear int     `json:"created_year"`
}