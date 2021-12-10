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

type Thing struct {
	Type                  string                 `json:"type"`
	ID                    string                 `json:"id"`
	Name                  string                 `json:"name"`
	CreatedOn             time.Time              `json:"created-on"`
	ModifiedOn            time.Time              `json:"modified-on"`
	Path                  string                 `json:"path"`
	Properties            []Properties           `json:"properties"`
	Version               string                 `json:"version"`
	AllowPassThrough      bool                   `json:"allow-pass-through"`
	Connections           []Connections          `json:"connections"`
	Costing               Costing                `json:"costing"`
	Capabilities          Capabilities           `json:"capabilities"`
	Script                *Script                `json:"script"`
	Manufacturer          *string                `json:"manufacturer"`
	MediaRules            *MediaRules            `json:"media-rules"`
	Placement             *Placement             `json:"placement"`
	FeedType              *string                `json:"feed-type"`
	ProcessType           *ProcessType           `json:"process-type"`
	Operation             *Operation             `json:"operation"`
	EstimatingEngine      *string                `json:"estimating-engine"`
	ZccConnectionSettings *ZccConnectionSettings `json:"zcc-connection-settings"`

	FrontUnits                FrontUnits `json:"front-units"`
	SinglePassDoubleSided     bool       `json:"single-pass-double-sided"`
	BackUnits                 BackUnits  `json:"back-units"`
	DoubleSidedSpeedReduction string     `json:"double-sided-speed-reduction"`
	Gripper                   string     `json:"gripper"`
}
type Rate struct {
	Value float64 `json:"value"`
	Units string  `json:"units"`
}
type SetupTime struct {
	Units string  `json:"units"`
	Value float64 `json:"value"`
}
type Setup struct {
	Type string    `json:"type"`
	Time SetupTime `json:"time"`
}
type Costing struct {
	Type     string `json:"type"`
	Currency string `json:"currency"`
	Rate     Rate   `json:"rate"`
	Setup    Setup  `json:"setup"`
}
type WidthRange struct {
	Type  string `json:"type"`
	End   string `json:"end"`
	Start string `json:"start"`
}
type HeightRange struct {
	Type  string `json:"type"`
	End   string `json:"end"`
	Start string `json:"start"`
}
type CaliperRange struct {
	Type string `json:"type"`
}
type WeightRange struct {
	Type string `json:"type"`
}
type Capabilities struct {
	Type           string        `json:"type"`
	WidthRange     WidthRange    `json:"width-range"`
	HeightRange    HeightRange   `json:"height-range"`
	Limit          bool          `json:"limit"`
	LimitLogic     string        `json:"limit-logic"`
	CaliperRange   CaliperRange  `json:"caliper-range"`
	WeightRange    WeightRange   `json:"weight-range"`
	SheetHandling  string        `json:"sheet-handling"`
	StockTypes     []interface{} `json:"stock-types"`
	SpecificStocks []interface{} `json:"specific-stocks"`
}
type Script struct {
	ID string `json:"id"`
}
type Content struct {
	Type   string `json:"type"`
	Left   string `json:"left"`
	Top    string `json:"top"`
	Right  string `json:"right"`
	Bottom string `json:"bottom"`
	Linked bool   `json:"linked"`
}
type Image struct {
	Type   string `json:"type"`
	Left   string `json:"left"`
	Top    string `json:"top"`
	Right  string `json:"right"`
	Bottom string `json:"bottom"`
	Linked bool   `json:"linked"`
}
type DefaultMediaRule struct {
	Type           string       `json:"type"`
	Marks          []DeviceMark `json:"marks"`
	SpeedReduction string       `json:"speed-reduction"`
	Content        Content      `json:"content"`
	Image          Image        `json:"image"`
}
type MediaRules struct {
	Type    string           `json:"type"`
	Default DefaultMediaRule `json:"default"`
	Map     []interface{}    `json:"map"`
}

func ID(id string) IDWrap {
	return IDWrap{ID: id}
}

type IDWrap struct {
	ID string `json:"id"`
}

type DeviceMark struct {
	Mark IDWrap `json:"mark"`
	Side string `json:"side"`
}
type Rule struct {
	Anchor           string `json:"anchor"`
	ReferencePoint   string `json:"reference-point"`
	Placement        string `json:"placement"`
	HorizontalOffset string `json:"horizontal-offset"`
	VerticalOffset   string `json:"vertical-offset"`
	ResizeSheet      bool   `json:"resize-sheet"`
}
type Placement struct {
	Type string `json:"type"`
	Rule Rule   `json:"rule"`
}
type Process struct {
	ID string `json:"id"`
}
type ProcessType struct {
	ID         string        `json:"id"`
	Name       string        `json:"name"`
	CreatedOn  time.Time     `json:"created-on"`
	ModifiedOn time.Time     `json:"modified-on"`
	Path       string        `json:"path"`
	Properties []interface{} `json:"properties"`
	Version    string        `json:"version"`
	Process    Process       `json:"process"`
}
type Velocity struct {
}
type Acceleration struct {
	Value float64 `json:"value"`
	Units string  `json:"units"`
}
type MotionLowered struct {
	Velocity     Velocity     `json:"velocity"`
	Acceleration Acceleration `json:"acceleration"`
}
type MotionLifted struct {
	Velocity     Velocity     `json:"velocity"`
	Acceleration Acceleration `json:"acceleration"`
}
type MotionLowering struct {
	Velocity     Velocity     `json:"velocity"`
	Acceleration Acceleration `json:"acceleration"`
}
type MotionLifting struct {
	Velocity     Velocity     `json:"velocity"`
	Acceleration Acceleration `json:"acceleration"`
}
type DefaultOperation struct {
	Type             string         `json:"type"`
	MotionLowered    MotionLowered  `json:"motion-lowered"`
	MotionLifted     MotionLifted   `json:"motion-lifted"`
	MotionLowering   MotionLowering `json:"motion-lowering"`
	MotionLifting    MotionLifting  `json:"motion-lifting"`
	ClearingDistance string         `json:"clearing-distance"`
	AutoLiftAngle    float64        `json:"auto-lift-angle"`
	AutoLiftDistance string         `json:"auto-lift-distance"`

	InkCost InkCost `json:"ink-cost"`
	Speed   Speed   `json:"speed"`
}

type InkCost struct {
	Type     string  `json:"type"`
	Units    string  `json:"units"`
	Coverage string  `json:"coverage"`
	Cyan     float64 `json:"cyan"`
	Magenta  float64 `json:"magenta"`
	Yellow   float64 `json:"yellow"`
	Black    float64 `json:"black"`
	Spots    float64 `json:"spots"`
	Coatings float64 `json:"coatings"`
	Foils    float64 `json:"foils"`
}
type SpeedUnits struct {
	Units string `json:"units"`
	Time  string `json:"time"`
}
type Speed struct {
	Type       string     `json:"type"`
	SpeedUnits SpeedUnits `json:"speed-units"`
	FixedSpeed float64    `json:"fixed-speed"`
}

type Operation struct {
	Type    string           `json:"type"`
	Default DefaultOperation `json:"default"`
	Map     []interface{}    `json:"map"`
}
type ZccConnectionSettings struct {
	Hostname string `json:"hostname"`
	Port     int    `json:"port"`
	Timeout  int    `json:"timeout"`
}

type ThingConnection struct {
	ID string `json:"id"`
}
type Connections struct {
	Thing ThingConnection `json:"thing"`
}
type Properties struct {
	Name  string `json:"name"`
	Value bool   `json:"value"`
	Type  string `json:"type"`
}

type CreateThing struct {
	Name                  string                 `json:"name"`
	Description           string                 `json:"description"`
	Notes                 string                 `json:"notes"`
	ExternalID            string                 `json:"external-id"`
	AllowPassThrough      bool                   `json:"allow-pass-through"`
	Connections           []Connections          `json:"connections"`
	Type                  string                 `json:"type"`
	Properties            []Properties           `json:"properties"`
	Path                  string                 `json:"path"`
	Script                *Script                `json:"script"`
	Manufacturer          *string                `json:"manufacturer"`
	MediaRules            *MediaRules            `json:"media-rules"`
	Placement             *Placement             `json:"placement"`
	FeedType              *string                `json:"feed-type"`
	ProcessType           *ProcessType           `json:"process-type"`
	Operation             *Operation             `json:"operation"`
	EstimatingEngine      *string                `json:"estimating-engine"`
	ZccConnectionSettings *ZccConnectionSettings `json:"zcc-connection-settings"`
}

type UpdateThing struct {
	Name                  string                 `json:"name"`
	Description           string                 `json:"description"`
	Notes                 string                 `json:"notes"`
	ExternalID            string                 `json:"external-id"`
	AllowPassThrough      bool                   `json:"allow-pass-through"`
	Connections           []Connections          `json:"connections"`
	Type                  string                 `json:"type"`
	Properties            []Properties           `json:"properties"`
	Path                  string                 `json:"path"`
	Script                *Script                `json:"script"`
	Manufacturer          *string                `json:"manufacturer"`
	MediaRules            *MediaRules            `json:"media-rules"`
	Placement             *Placement             `json:"placement"`
	Capabilities          *Capabilities          `json:"capabilities"`
	FeedType              *string                `json:"feed-type"`
	ProcessType           *ProcessType           `json:"process-type"`
	Operation             *Operation             `json:"operation"`
	EstimatingEngine      *string                `json:"estimating-engine"`
	ZccConnectionSettings *ZccConnectionSettings `json:"zcc-connection-settings"`

	FrontUnits                FrontUnits `json:"front-units"`
	SinglePassDoubleSided     bool       `json:"single-pass-double-sided"`
	BackUnits                 BackUnits  `json:"back-units"`
	DoubleSidedSpeedReduction string     `json:"double-sided-speed-reduction"`
	Gripper                   string     `json:"gripper"`
}

type Mark struct {
	Type   string `json:"type"`
	Name   string `json:"name"`
	ID     string `json:"id"`
	Smart  bool   `json:"smart"`
	Anchor string `json:"anchor"`
}

type FrontUnits struct {
	Units    float64 `json:"units"`
	Coatings float64 `json:"coatings"`
	Foils    float64 `json:"foils"`
}
type BackUnits struct {
	Units    float64 `json:"units"`
	Coatings float64 `json:"coatings"`
	Foils    float64 `json:"foils"`
}
