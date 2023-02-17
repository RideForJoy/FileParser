package product

type Products []Product

type Product struct {
	Barcodes []string `json:"barcodes"`
	CaseItem struct {
		CaseCount  int `json:"case-count"`
		Dimensions struct {
			Height        float64 `json:"height"`
			Length        float64 `json:"length"`
			UnitOfMeasure string  `json:"unit-of-measure"`
			Width         float64 `json:"width"`
		} `json:"dimensions"`
		Weight struct {
			UnitOfMeasure string  `json:"unit-of-measure"`
			Weight        float64 `json:"weight"`
		} `json:"weight"`
	} `json:"case-item"`
	CategoriesHierarchy []interface{} `json:"categories-hierarchy"`
	CorpIds             []string      `json:"corp-ids"`
	Description         string        `json:"description"`
	Discontinued        bool          `json:"discontinued"`
	EcomIds             []string      `json:"ecom-ids"`
	FeatureAttributes   struct {
		FoodSafety                string `json:"food-safety"`
		IsBulk                    bool   `json:"is-bulk"`
		IsCaseStorage             bool   `json:"is-case-storage"`
		IsChemical                bool   `json:"is-chemical"`
		IsCrushable               bool   `json:"is-crushable"`
		IsWeightVariableOnReceipt bool   `json:"is-weight-variable-on-receipt"`
		IsEgg                     bool   `json:"is-egg"`
		IsGlassPackaged           bool   `json:"is-glass-packaged"`
		IsHazardous               bool   `json:"is-hazardous"`
		IsHeavy                   bool   `json:"is-heavy"`
		IsRaw                     bool   `json:"is-raw"`
		IsSpecificChilled         bool   `json:"is-specific-chilled"`
		IsStackable               bool   `json:"is-stackable"`
		IsVertical                bool   `json:"is-vertical"`
	} `json:"feature-attributes"`
	Image       string `json:"image"`
	ItemAddress struct {
		Aisle    string `json:"aisle"`
		Area     string `json:"area"`
		Bay      string `json:"bay"`
		Position string `json:"position"`
		Shelf    string `json:"shelf"`
	} `json:"item-address"`
	ItemType               string `json:"item-type"`
	MfcID                  string `json:"mfc-id"`
	MfcStopBuy             bool   `json:"mfc-stop-buy"`
	MfcStopFulfill         bool   `json:"mfc-stop-fulfill"`
	MinRemainingShelfLife  int    `json:"min-remaining-shelf-life"`
	Name                   string `json:"name"`
	PrimaryCorpID          string `json:"primary-corp-id"`
	RequiresExpirationDate bool   `json:"requires-expiration-date"`
	RetailItem             struct {
		Dimensions struct {
			Height        float64 `json:"height"`
			Length        float64 `json:"length"`
			UnitOfMeasure string  `json:"unit-of-measure"`
			Width         float64 `json:"width"`
		} `json:"dimensions"`
		Weight struct {
			UnitOfMeasure string  `json:"unit-of-measure"`
			Weight        float64 `json:"weight"`
		} `json:"weight"`
	} `json:"retail-item"`
	Seasonal         bool `json:"seasonal"`
	ShelfLife        int  `json:"shelf-life"`
	SlottingEligible bool `json:"slotting-eligible"`
	SupplierInfo     []struct {
		SupplierID           string `json:"supplier-id"`
		SupplierKeepBuy      bool   `json:"supplier-keep-buy"`
		SupplierName         string `json:"supplier-name"`
		SupplierPerformCycle int    `json:"supplier-perform-cycle"`
		SupplierProductID    string `json:"supplier-product-id"`
		SupplierSubAccount   string `json:"supplier-sub-account"`
		SupplierType         string `json:"supplier-type"`
	} `json:"supplier-info"`
	TemperatureZone []string `json:"temperature-zone"`
	TomID           string   `json:"tom-id"`
}
