package backend

type F1 struct {
	MRData struct {
		RaceTable struct {
			Season string `json:"season"`
			Round  string `json:"round"`
			Races  []struct {
				Season   string `json:"season"`
				Round    string `json:"round"`
				RaceName string `json:"raceName"`
				Circuit  struct {
					CircuitID   string `json:"circuitId"`
					URL         string `json:"url"`
					CircuitName string `json:"circuitName"`
					Location    struct {
						Lat      string `json:"lat"`
						Long     string `json:"long"`
						Locality string `json:"locality"`
						Country  string `json:"country"`
					} `json:"Location"`
				} `json:"Circuit"`
				Date    string `json:"date"`
				Time    string `json:"time"`
				Results []struct {
					Number   string `json:"number"`
					Position string `json:"position"`
					Points   string `json:"points"`
					Driver   struct {
						DriverID        string `json:"driverId"`
						PermanentNumber string `json:"permanentNumber"`
						Code            string `json:"code"`
						GivenName       string `json:"givenName"`
						FamilyName      string `json:"familyName"`
						DateOfBirth     string `json:"dateOfBirth"`
						Nationality     string `json:"nationality"`
					} `json:"Driver"`
					Constructor struct {
						ConstructorID string `json:"constructorId"`
						URL           string `json:"url"`
						Name          string `json:"name"`
						Nationality   string `json:"nationality"`
					} `json:"Constructor"`
					Grid   string `json:"grid"`
					Laps   string `json:"laps"`
					Status string `json:"status"`
					Time   struct {
						Millis string `json:"millis"`
						Time   string `json:"time"`
					} `json:"Time,omitempty"`
					FastestLap struct {
						Rank string `json:"rank"`
						Lap  string `json:"lap"`
						Time struct {
							Time string `json:"time"`
						} `json:"Time"`
						AverageSpeed struct {
							Units string `json:"units"`
							Speed string `json:"speed"`
						} `json:"AverageSpeed"`
					} `json:"FastestLap"`
				} `json:"Results"`
			} `json:"Races"`
		} `json:"RaceTable"`
	} `json:"MRData"`
}

type InfoPilotes struct {
	MRData struct {
		RaceTable struct {
			Races []struct {
				Results []struct {
					Driver struct {
						DriverID        string `json:"driverId"`
						PermanentNumber string `json:"permanentNumber"`
						Code            string `json:"code"`
						GivenName       string `json:"givenName"`
						FamilyName      string `json:"familyName"`
						DateOfBirth     string `json:"dateOfBirth"`
						Nationality     string `json:"nationality"`
					} `json:"Driver"`
					Constructor struct {
						ConstructorID string `json:"constructorId"`
						Name          string `json:"name"`
						Nationality   string `json:"nationality"`
					} `json:"Constructor"`
				} `json:"Results"`
			} `json:"Races"`
		} `json:"RaceTable"`
	} `json:"MRData"`
}

type Pilote struct {
	DriverID      string
	Name          string
	FamilyName    string
	DateOfBirth   string
	Code          int
	Number        int
	Nationality   string
	Constructor   string
	ConstructorID string
}

type Circuit struct{
	IDCircuit string
	Name string
}
