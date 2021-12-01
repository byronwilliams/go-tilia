package libraries

import "time"

type Stock struct {
	Name        string          `json:"name"`
	ID          string          `json:"id,omitempty"`
	ExternalID  string          `json:"external-id"`
	Description string          `json:"description"`
	StockType   CreateStockType `json:"stock-type"`
	Vendor      string          `json:"vendor"`
	Grades      []Grades        `json:"grades"`
}
type CreateStockType struct {
	Name        string  `json:"name,omitempty"`
	ID          *string `json:"id,omitempty"`
	ExternalID  *string `json:"external-id,omitempty"`
	Description *string `json:"description,omitempty"`
}
type Sheets struct {
	Name        string  `json:"name"`
	ID          string  `json:"id,omitempty"`
	ExternalID  string  `json:"external-id"`
	Description string  `json:"description"`
	StockID     string  `json:"stock-id"`
	GradeID     string  `json:"grade-id"`
	Dimension1  string  `json:"dimension1"`
	Dimension2  string  `json:"dimension2"`
	Cost        float64 `json:"cost"`
	CostUnits   string  `json:"cost-units"`
	Grain       string  `json:"grain"`
}
type Rolls struct {
	Name        *string `json:"name,omitempty"`
	ID          string  `json:"id,omitempty"`
	ExternalID  string  `json:"external-id"`
	Description string  `json:"description,omitempty"`
	StockID     string  `json:"stock-id,omitempty"`
	GradeID     string  `json:"grade-id,omitempty"`
	Width       string  `json:"width"`
	Cost        float64 `json:"cost"`
	CostUnits   string  `json:"cost-units"`
	Grain       string  `json:"grain"`
}
type Grades struct {
	Name         *string  `json:"name,omitempty"`
	ID           *string  `json:"id,omitempty"`
	ExternalID   string   `json:"external-id"`
	Description  *string  `json:"description,omitempty"`
	GradeDisplay string   `json:"grade-display"`
	Weight       *int     `json:"weight,omitempty"`
	WeightUnits  *string  `json:"weight-units,omitempty"`
	WeightType   *string  `json:"weight-type,omitempty"`
	Caliper      string   `json:"caliper,omitempty"`
	Cost         float64  `json:"cost"`
	CostUnits    string   `json:"cost-units"`
	AnySheetSize *bool    `json:"any-sheet-size,omitempty"`
	Sheets       []Sheets `json:"sheets,omitempty"`
	Rolls        []Rolls  `json:"rolls,omitempty"`
}

type StockType struct {
	ID          string                `json:"id"`
	Name        string                `json:"name"`
	CreatedOn   time.Time             `json:"created-on"`
	ModifiedOn  time.Time             `json:"modified-on"`
	Version     string                `json:"version"`
	Description string                `json:"description"`
	Notes       string                `json:"notes"`
	ExternalID  string                `json:"external-id"`
	Properties  []StockTypeProperties `json:"properties"`
	Path        string                `json:"path"`
}
type StockTypeProperties struct {
	Name  string `json:"name"`
	Value bool   `json:"value"`
	Type  string `json:"type"`
}

type StockV2 struct {
	Type       string        `json:"type"`
	ID         string        `json:"id"`
	Name       string        `json:"name"`
	CreatedOn  time.Time     `json:"created-on"`
	ModifiedOn time.Time     `json:"modified-on"`
	Path       string        `json:"path"`
	Properties []interface{} `json:"properties"`
	Version    string        `json:"version"`
	ExternalID string        `json:"external-id"`
	StockType  StockTypeV2   `json:"stock-type"`
	Vendor     string        `json:"vendor"`
	Grades     []GradeV2     `json:"grades"`
}
type StockTypeV2 struct {
	ID string `json:"id"`
}
type RollV2 struct {
	ID         string        `json:"id"`
	Name       string        `json:"name"`
	CreatedOn  time.Time     `json:"created-on"`
	ModifiedOn time.Time     `json:"modified-on"`
	Path       string        `json:"path"`
	Properties []interface{} `json:"properties"`
	Version    string        `json:"version"`
	ExternalID string        `json:"external-id"`
	Dimension1 string        `json:"dimension1"`
	Cost       float64       `json:"cost"`
	CostUnits  string        `json:"cost-units"`
	Grain      string        `json:"grain"`
}
type GradeV2 struct {
	ID           string        `json:"id"`
	Name         string        `json:"name"`
	CreatedOn    time.Time     `json:"created-on"`
	ModifiedOn   time.Time     `json:"modified-on"`
	Path         string        `json:"path"`
	Properties   []interface{} `json:"properties"`
	Version      string        `json:"version"`
	ExternalID   string        `json:"external-id"`
	GradeDisplay string        `json:"grade-display"`
	Weight       float64       `json:"weight"`
	Caliper      string        `json:"caliper"`
	Cost         float64       `json:"cost"`
	CostUnits    string        `json:"cost-units"`
	AnySheetSize bool          `json:"any-sheet-size"`
	Sheets       []interface{} `json:"sheets"`
	Rolls        []RollV2      `json:"rolls"`
}
