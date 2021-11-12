package projects

import "time"

type StandardResponse struct {
	Type       string     `json:"type"`
	Success    bool       `json:"success"`
	StatusCode int        `json:"status-code"`
	Errors     []Errors   `json:"errors"`
	Warnings   []Warnings `json:"warnings"`
	Resource   []string   `json:"resource"`
}

type Errors struct {
	ID     int    `json:"id"`
	Text   string `json:"text"`
	Action string `json:"action"`
}
type Warnings struct {
	ID     int    `json:"id"`
	Text   string `json:"text"`
	Action string `json:"action"`
}

type CreateProjectRequest struct {
	ID           string  `json:"id"`
	Name         *string `json:"name"`
	Contact      *string `json:"contact"`
	Phone        *string `json:"phone"`
	Client       *string `json:"client"`
	Notes        *string `json:"notes"`
	Units        *string `json:"units"`
	Idref        *int    `json:"idref"`
	TemplatePath *string `json:"template-path"`
}

type PlanProjectRequest struct {
	Idref          int      `json:"idref"`
	Products       []string `json:"products"`
	Profiles       []string `json:"profiles"`
	ProfilesInline []struct {
		Name        string `json:"name"`
		ID          string `json:"id"`
		ExternalID  string `json:"external-id"`
		Description string `json:"description"`
		Strategies  struct {
			HorizontalCut   bool `json:"horizontal-cut"`
			VerticalCut     bool `json:"vertical-cut"`
			Nesting         bool `json:"nesting"`
			FreeNesting     bool `json:"free-nesting"`
			GridNesting     bool `json:"grid-nesting"`
			StripNesting    bool `json:"strip-nesting"`
			HorizontalStrip bool `json:"horizontal-strip"`
			VerticalStrip   bool `json:"vertical-strip"`
			Templates       bool `json:"templates"`
		} `json:"strategies"`
		StripOptions struct {
			StripRule    string `json:"strip-rule"`
			TemplateRule string `json:"template-rule"`
			Property     string `json:"property"`
			Alignment    string `json:"alignment"`
			Gutter       string `json:"gutter"`
			GutterRule   string `json:"gutter-rule"`
		} `json:"strip-options"`
		LayoutOptions struct {
			SheetFill           string `json:"sheet-fill"`
			LimitUniqueProducts int    `json:"limit-unique-products"`
			AllowBleedInGripper bool   `json:"allow-bleed-in-gripper"`
			UseDerivedSheets    bool   `json:"use-derived-sheets"`
			OrderedPlacement    struct {
				FavorOrderedPlacement bool   `json:"favor-ordered-placement"`
				StartCorner           string `json:"start-corner"`
				OrderMethod           string `json:"order-method"`
			} `json:"ordered-placement"`
		} `json:"layout-options"`
		PlanOptions struct {
			PlanMode                 string `json:"plan-mode"`
			StackSize                int    `json:"stack-size"`
			Finishing                string `json:"finishing"`
			StackingOrder            string `json:"stacking-order"`
			AllowProductSpanning     bool   `json:"allow-product-spanning"`
			AllowMultiplePressPasses bool   `json:"allow-multiple-press-passes"`
			FixedRunLength           int    `json:"fixed-run-length"`
		} `json:"plan-options"`
		WebOptions struct {
			AllowSignatureFrameSpanning bool `json:"allow-signature-frame-spanning"`
		} `json:"web-options"`
		ApplyingOptions struct {
			SplitOverlapsOnApply   bool `json:"split-overlaps-on-apply"`
			GroupProductsOnApply   bool `json:"group-products-on-apply"`
			EnsureMarginsPlacement bool `json:"ensure-margins-placement"`
		} `json:"applying-options"`
		PlanRules struct {
			Logic string   `json:"logic"`
			Rules []string `json:"rules"`
			Type  string   `json:"type"`
		} `json:"plan-rules"`
		Scripts []struct {
			Name        string `json:"name"`
			ID          string `json:"id"`
			ExternalID  string `json:"external-id"`
			Description string `json:"description"`
		} `json:"scripts"`
	} `json:"profiles-inline"`
	StopMinutes int      `json:"stop-minutes"`
	Things      []string `json:"things"`
	Sheets      []struct {
		Stock string `json:"stock"`
		Grade string `json:"grade"`
		Name  string `json:"name"`
	} `json:"sheets"`
	Rolls       []Roll   `json:"rolls"`
	Templates   []string `json:"templates"`
	ApplyResult bool     `json:"apply-result"`
	Presses     []string `json:"presses"`
}

type Roll struct {
	Stock string `json:"stock"`
	Grade string `json:"grade"`
	Name  string `json:"name"`
}

type AddProductToProjectRequest struct {
	// Name is unique to a Project
	Name    string  `json:"name"`
	Type    *string `json:"type"`
	Color   *string `json:"color"`
	Ordered int     `json:"ordered"`
	Stock   *string `json:"stock"`
	Grade   *string `json:"grade"`
	Colors  *[]struct {
		Name     string    `json:"name"`
		Type     string    `json:"type"`
		Values   []float64 `json:"values"`
		Coverage float64   `json:"coverage"`
		Process  string    `json:"process"`
	} `json:"colors"`
	BackColors *[]struct {
		Name     string    `json:"name"`
		Type     string    `json:"type"`
		Values   []float64 `json:"values"`
		Coverage float64   `json:"coverage"`
		Process  string    `json:"process"`
	} `json:"back-colors"`
	ColorSource         *string    `json:"color-source"`
	Grain               *string    `json:"grain"`
	Width               string     `json:"width"`
	Height              string     `json:"height"`
	Rotation            *Rotation  `json:"rotation"`
	AllowedRotations    *string    `json:"allowed-rotations"`
	Templates           *[]string  `json:"templates"`
	ScaleProportionally *bool      `json:"scale-proportionally"`
	Artwork             string     `json:"artwork"`
	Page                *int       `json:"page"`
	BackArtwork         *string    `json:"back-artwork"`
	BackPage            *int       `json:"back-page"`
	DieshapeSource      *string    `json:"dieshape-source"`
	DieDesign           *string    `json:"die-design"`
	AutosnapInk         *string    `json:"autosnap-ink"`
	BackAutosnapInk     *string    `json:"back-autosnap-ink"`
	AutosnapLayer       *string    `json:"autosnap-layer"`
	BackAutosnapLayer   *string    `json:"back-autosnap-layer"`
	ShapeHandling       *string    `json:"shape-handling"`
	PageHandling        *string    `json:"page-handling"`
	FrontToBack         *string    `json:"front-to-back"`
	CadFile             *string    `json:"cad-file"`
	CadDesign           *string    `json:"cad-design"`
	Group               *string    `json:"group"`
	Priority            int        `json:"priority"`
	DueDate             *string    `json:"due-date"`
	BleedType           *BleedType `json:"bleed-type"`
	BleedMargin         string     `json:"bleed-margin"`
	BleedMargins        *struct {
		Type   string `json:"type"`
		Left   string `json:"left"`
		Top    string `json:"top"`
		Right  string `json:"right"`
		Bottom string `json:"bottom"`
		Linked bool   `json:"linked"`
	} `json:"bleed-margins"`
	PageBleed      *string `json:"page-bleed"`
	SpacingType    *string `json:"spacing-type"`
	SpacingMargin  *string `json:"spacing-margin"`
	SpacingMargins *struct {
		Type   string `json:"type"`
		Left   string `json:"left"`
		Top    string `json:"top"`
		Right  string `json:"right"`
		Bottom string `json:"bottom"`
		Linked bool   `json:"linked"`
	} `json:"spacing-margins"`
	OffcutMargins *struct {
		Type   string `json:"type"`
		Left   string `json:"left"`
		Top    string `json:"top"`
		Right  string `json:"right"`
		Bottom string `json:"bottom"`
		Linked bool   `json:"linked"`
	} `json:"offcut-margins"`
	MinOverruns     string `json:"min-overruns"`
	MaxOverruns     string `json:"max-overruns"`
	Description     string `json:"description"`
	Notes           string `json:"notes"`
	ProcessSettings *[]struct {
		Process      string   `json:"process"`
		Mode         string   `json:"mode"`
		ModeValue    float64  `json:"mode-value"`
		Things       []string `json:"things"`
		ProcessTypes []string `json:"process-types"`
	} `json:"process-settings"`
	Marks           []string `json:"marks"`
	BackMarks       []string `json:"back-marks"`
	FoldingPatterns []string `json:"folding-patterns"`
	BindingMethod   *string  `json:"binding-method"`
	BindingEdge     *string  `json:"binding-edge"`
	JogEdge         *string  `json:"jog-edge"`
	ReadingOrder    *string  `json:"reading-order"`
	Pages           *int     `json:"pages"`
	PagesPerSection *int     `json:"pages-per-section"`
	SelfCover       *bool    `json:"self-cover"`
	Trim            *struct {
		SpineTrim  string `json:"spine-trim"`
		JogTrim    string `json:"jog-trim"`
		FaceTrim   string `json:"face-trim"`
		NonJogTrim string `json:"non-jog-trim"`
		LipType    string `json:"lip-type"`
		Lip        string `json:"lip"`
	} `json:"trim"`
	NUp *struct {
		Number int    `json:"number"`
		Gap    string `json:"gap"`
	} `json:"n-up"`
	Creep *struct {
		Type        string `json:"type"`
		Transition  string `json:"transition"`
		Method      string `json:"method"`
		Calculation string `json:"calculation"`
		Amount      string `json:"amount"`
	} `json:"creep"`
	BundleSize  *int               `json:"bundle-size"`
	Properties  *[]ProductProperty `json:"properties"`
	Tiling      *Tiling            `json:"tiling"`
	FrontInks   []string           `json:"front-inks"`
	BackInks    []string           `json:"back-inks"`
	CutInk      *string            `json:"cut-ink"`
	CreaseInk   *string            `json:"crease-ink"`
	BleedInk    *string            `json:"bleed-ink"`
	CutLayer    *string            `json:"cut-layer"`
	CreaseLayer *string            `json:"crease-layer"`
	BleedLayer  *string            `json:"bleed-layer"`
}

type ProductPropertyType string

const (
	Text     ProductPropertyType = "Text"
	Integer  ProductPropertyType = "Integer"
	Boolean  ProductPropertyType = "Boolean"
	Number   ProductPropertyType = "Number"
	Date     ProductPropertyType = "Date"
	TextList ProductPropertyType = "TextList"
)

type ProductProperty struct {
	Name  string              `json:"name"`
	Value interface{}         `json:"value"`
	Type  ProductPropertyType `json:"type"`
}

type Tiling struct {
	Name           string `json:"name"`
	Description    string `json:"description"`
	Notes          string `json:"notes"`
	Start          string `json:"start"`
	Order          string `json:"order"`
	HorizontalRule struct {
		Number           int    `json:"number"`
		UniformFinalSize bool   `json:"uniform-final-size"`
		Type             string `json:"type"`
	} `json:"horizontal-rule"`
	HorizontalMethod struct {
		Gap           float64 `json:"gap"`
		ExtensionRule string  `json:"extension-rule"`
		Extension     float64 `json:"extension"`
		Type          string  `json:"type"`
	} `json:"horizontal-method"`
	VerticalRule struct {
		Number           int    `json:"number"`
		UniformFinalSize bool   `json:"uniform-final-size"`
		Type             string `json:"type"`
	} `json:"vertical-rule"`
	VerticalMethod struct {
		Gap           float64 `json:"gap"`
		ExtensionRule string  `json:"extension-rule"`
		Extension     float64 `json:"extension"`
		Type          string  `json:"type"`
	} `json:"vertical-method"`
	Type string `json:"type"`
	Path string `json:"path"`
}

type ApplyMarkRequest struct {
	Name  string `json:"name"`
	Idref *int   `json:"idref"`
}

type ExportType string

const (
	ExportCoverSheet   ExportType = "cover-sheet"
	ExportDieCff2      ExportType = "die/cff2"
	ExportDieDxf       ExportType = "die/dxf"
	ExportDiePdf       ExportType = "die/pdf"
	ExportDieZcc       ExportType = "die/zcc"
	ExportHPJDF        ExportType = "hp-jdf"
	ExportJDFCutting   ExportType = "jdf-cutting"
	ExportPDF          ExportType = "pdf"
	ExportPDFVector    ExportType = "pdf-vector"
	ExportReportJSON   ExportType = "report/json"
	ExportReportPDF    ExportType = "report/pdf"
	ExportReportXML    ExportType = "report/xml"
	ExportReportTiling ExportType = "tiling-report"
)

type ExportRequest struct {
	Idref  *int    `json:"idref"`
	Path   *string `json:"path"`
	Preset *string `json:"preset"`
}

type PlanStatus string

const (
	PlanNotStarted PlanStatus = "NotStarted"
	PlanRunning    PlanStatus = "Running"
	PlanFinished   PlanStatus = "Finished"

	// Dont know
	PlanError PlanStatus = "Error"
)

type PlanStatusResponse struct {
	State         PlanStatus `json:"state"`
	RunningTime   float64    `json:"running-time"`
	Errors        []Errors   `json:"errors"`
	Warnings      []Warnings `json:"warnings"`
	Results       int        `json:"results"`
	LowestCost    float64    `json:"lowest-cost"`
	LowestWaste   float64    `json:"lowest-waste"`
	LowestLayouts int        `json:"lowest-layouts"`
}

type Sorting string

func SortByCost() *Sorting    { var s Sorting = "Cost"; return &s }
func SortByWaste() *Sorting   { var s Sorting = "Waste"; return &s }
func SortByTime() *Sorting    { var s Sorting = "Time"; return &s }
func SortByLayouts() *Sorting { var s Sorting = "Layouts"; return &s }

type RenderMode string

const (
	RenderArtwork  RenderMode = "Artwork"
	RenderColors   RenderMode = "Colors"
	RenderDielines RenderMode = "Dielines"
)

type ListPlanResultsRequestOpts struct {
	Limit       *int     `qs:"limit"`
	Start       *int     `qs:"start"`
	Sorting     *Sorting `qs:"sorting"`
	Layouts     *bool    `qs:"layouts"`
	Thumb       *bool    `qs:"thumb"`
	PlanThumb   *bool    `qs:"plan-thumb"`
	ThumbWidth  *int     `qs:"thumb-width"`
	ThumbHeight *int     `qs:"thumb-height"`
	RenderMode  *string  `qs:"render-mode"`
}

type ListPlanResultsResponse struct {
	ID                 int        `json:"id"`
	Timestamp          float64    `json:"timestamp"`
	PendingEstimations bool       `json:"pending-estimations"`
	PressMinutes       float64    `json:"press-minutes"`
	CuttingTime        float64    `json:"cutting-time"`
	RunLength          int        `json:"run-length"`
	TotalCost          float64    `json:"total-cost"`
	PlateCost          float64    `json:"plate-cost"`
	StockCost          float64    `json:"stock-cost"`
	PressCost          float64    `json:"press-cost"`
	CuttingCost        float64    `json:"cutting-cost"`
	Waste              float64    `json:"waste"`
	SheetUsage         float64    `json:"sheet-usage"`
	TotalUnderruns     int        `json:"total-underruns"`
	TotalOverruns      int        `json:"total-overruns"`
	LayoutCount        int        `json:"layout-count"`
	Layouts            *[]Layouts `json:"layouts"`
	Thumbnail          Thumbnail  `json:"thumbnail"`
}
type Bounds struct {
	X      string `json:"x"`
	Y      string `json:"y"`
	Width  string `json:"width"`
	Height string `json:"height"`
}
type Items struct {
	Count  int    `json:"count"`
	Bounds Bounds `json:"bounds"`
}
type Templates struct {
	Name   string `json:"name"`
	Source string `json:"source"`
	Items  int    `json:"items"`
	Placed int    `json:"placed"`
}
type Thumbnail struct {
	Data      string `json:"data"`
	MediaType string `json:"media-type"`
	Width     int    `json:"width"`
	Height    int    `json:"height"`
}
type Layouts struct {
	ID                 int         `json:"id"`
	Timestamp          float64     `json:"timestamp"`
	PendingEstimations bool        `json:"pending-estimations"`
	LayoutType         string      `json:"layout-type"`
	Sheet              string      `json:"sheet"`
	Things             []string    `json:"things"`
	RunLength          int         `json:"run-length"`
	PressMinutes       float64     `json:"press-minutes"`
	CuttingTime        float64     `json:"cutting-time"`
	TotalCost          float64     `json:"total-cost"`
	PlateCost          float64     `json:"plate-cost"`
	StockCost          float64     `json:"stock-cost"`
	PressCost          float64     `json:"press-cost"`
	CuttingCost        float64     `json:"cutting-cost"`
	Waste              float64     `json:"waste"`
	SheetUsage         float64     `json:"sheet-usage"`
	MaxOverrun         float64     `json:"max-overrun"`
	TotalOverruns      int         `json:"total-overruns"`
	Placed             int         `json:"placed"`
	Items              Items       `json:"items"`
	CutComplexity      int         `json:"cut-complexity"`
	Templates          []Templates `json:"templates"`
	Thumbnail          Thumbnail   `json:"thumbnail"`
	Press              string      `json:"press"`
}

func intPtr(n int) *int {
	return &n
}

func Limit(n int) *int {
	return intPtr(n)
}

func IsTerminalStatus(s PlanStatus) bool {
	return s == PlanError || s == PlanFinished
}

type AutoGenerated struct {
	Type                      string        `json:"type"`
	ID                        string        `json:"id"`
	Name                      string        `json:"name"`
	Description               string        `json:"description"`
	CreatedOn                 time.Time     `json:"created-on"`
	ModifiedOn                time.Time     `json:"modified-on"`
	Path                      string        `json:"path"`
	Properties                []interface{} `json:"properties"`
	Version                   string        `json:"version"`
	AllowPassThrough          bool          `json:"allow-pass-through"`
	Connections               []interface{} `json:"connections"`
	Costing                   Costing       `json:"costing"`
	Capabilities              Capabilities  `json:"capabilities"`
	Script                    Script        `json:"script"`
	Manufacturer              string        `json:"manufacturer"`
	MediaRules                MediaRules    `json:"media-rules"`
	Placement                 Placement     `json:"placement"`
	FeedType                  string        `json:"feed-type"`
	ProcessType               ProcessType   `json:"process-type"`
	Operation                 Operation     `json:"operation"`
	FrontUnits                FrontUnits    `json:"front-units"`
	SinglePassDoubleSided     bool          `json:"single-pass-double-sided"`
	BackUnits                 BackUnits     `json:"back-units"`
	DoubleSidedSpeedReduction string        `json:"double-sided-speed-reduction"`
	Gripper                   string        `json:"gripper"`
}
type Rate struct {
	Value float64 `json:"value"`
	Units string  `json:"units"`
}
type Time struct {
	Units string `json:"units"`
	Value int    `json:"value"`
}
type Setup struct {
	Type    string `json:"type"`
	Time    Time   `json:"time"`
	Layouts int    `json:"layouts"`
}
type Costing struct {
	Type         string `json:"type"`
	Currency     string `json:"currency"`
	Rate         Rate   `json:"rate"`
	Setup        Setup  `json:"setup"`
	RunningWaste int    `json:"running-waste"`
}
type WidthRange struct {
	Type  string `json:"type"`
	End   string `json:"end"`
	Start string `json:"start"`
}
type HeightRange struct {
	Type string `json:"type"`
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
type MediaDefaultRules struct {
	Type           string        `json:"type"`
	Marks          []interface{} `json:"marks"`
	SpeedReduction string        `json:"speed-reduction"`
	Content        Content       `json:"content"`
	Image          Image         `json:"image"`
	Regions        []interface{} `json:"regions"`
	InkAdjustment  string        `json:"ink-adjustment"`
}
type MediaRules struct {
	Type    string            `json:"type"`
	Default MediaDefaultRules `json:"default"`
	Map     []interface{}     `json:"map"`
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
type InkCost struct {
	Type     string `json:"type"`
	Units    string `json:"units"`
	Coverage string `json:"coverage"`
	Cyan     int    `json:"cyan"`
	Magenta  int    `json:"magenta"`
	Yellow   int    `json:"yellow"`
	Black    int    `json:"black"`
	Spots    int    `json:"spots"`
	Coatings int    `json:"coatings"`
	Foils    int    `json:"foils"`
}
type SpeedUnits struct {
	Units string `json:"units"`
	Time  string `json:"time"`
}
type Speed struct {
	Type       string     `json:"type"`
	SpeedUnits SpeedUnits `json:"speed-units"`
	FixedSpeed int        `json:"fixed-speed"`
}
type ThingDefault struct {
	Type    string  `json:"type"`
	InkCost InkCost `json:"ink-cost"`
	Speed   Speed   `json:"speed"`
}
type Operation struct {
	Type    string        `json:"type"`
	Default ThingDefault  `json:"default"`
	Map     []interface{} `json:"map"`
}
type FrontUnits struct {
	Units    int `json:"units"`
	Coatings int `json:"coatings"`
	Foils    int `json:"foils"`
}
type BackUnits struct {
	Units    int `json:"units"`
	Coatings int `json:"coatings"`
	Foils    int `json:"foils"`
}

type Rotation string

const (
	RotationNone Rotation = "None"
	RotationAny  Rotation = "Any"
)

type BleedType string

const (
	BleedTypeNone    BleedType = "None"
	BleedTypeMargins BleedType = "Margins"
	BleedTypeContour BleedType = "Contour"
)
