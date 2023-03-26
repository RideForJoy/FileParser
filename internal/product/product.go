package product

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

type Products []Product

type Product struct {
	Barcodes []string `json:"barcodes,omitempty"`
	CaseItem struct {
		CaseCount  int `json:"case-count,omitempty"`
		Dimensions struct {
			Height        float64 `json:"height,omitempty"`
			Length        float64 `json:"length,omitempty"`
			UnitOfMeasure string  `json:"unit-of-measure,omitempty"`
			Width         float64 `json:"width,omitempty"`
		} `json:"dimensions,omitempty"`
		Weight struct {
			UnitOfMeasure string  `json:"unit-of-measure,omitempty"`
			Weight        float64 `json:"weight,omitempty"`
		} `json:"weight,omitempty"`
	} `json:"case-item,omitempty"`
	CategoriesHierarchy []interface{} `json:"categories-hierarchy,omitempty"`
	CorpIds             []string      `json:"corp-ids,omitempty"`
	Description         string        `json:"description,omitempty"`
	Discontinued        bool          `json:"discontinued,omitempty"`
	EcomIds             []string      `json:"ecom-ids,omitempty"`
	FeatureAttributes   struct {
		FoodSafety                string `json:"food-safety,omitempty"`
		IsBulk                    bool   `json:"is-bulk,omitempty"`
		IsCaseStorage             bool   `json:"is-case-storage,omitempty"`
		IsChemical                bool   `json:"is-chemical,omitempty"`
		IsCrushable               bool   `json:"is-crushable,omitempty"`
		IsWeightVariableOnReceipt bool   `json:"is-weight-variable-on-receipt,omitempty"`
		IsEgg                     bool   `json:"is-egg,omitempty"`
		IsGlassPackaged           bool   `json:"is-glass-packaged,omitempty"`
		IsHazardous               bool   `json:"is-hazardous,omitempty"`
		IsHeavy                   bool   `json:"is-heavy,omitempty"`
		IsRaw                     bool   `json:"is-raw,omitempty"`
		IsSpecificChilled         bool   `json:"is-specific-chilled,omitempty"`
		IsStackable               bool   `json:"is-stackable,omitempty"`
		IsVertical                bool   `json:"is-vertical,omitempty"`
	} `json:"feature-attributes,omitempty"`
	Image       string `json:"image,omitempty"`
	ItemAddress struct {
		Aisle    string `json:"aisle,omitempty"`
		Area     string `json:"area,omitempty"`
		Bay      string `json:"bay,omitempty"`
		Position string `json:"position,omitempty"`
		Shelf    string `json:"shelf,omitempty"`
	} `json:"item-address,omitempty"`
	ItemType               string `json:"item-type,omitempty"`
	MfcID                  string `json:"mfc-id,omitempty"`
	MfcStopBuy             bool   `json:"mfc-stop-buy,omitempty"`
	MfcStopFulfill         bool   `json:"mfc-stop-fulfill,omitempty"`
	MinRemainingShelfLife  int    `json:"min-remaining-shelf-life,omitempty"`
	Name                   string `json:"name,omitempty"`
	PrimaryCorpID          string `json:"primary-corp-id,omitempty"`
	RequiresExpirationDate bool   `json:"requires-expiration-date,omitempty"`
	RetailItem             struct {
		Dimensions struct {
			Height        float64 `json:"height,omitempty"`
			Length        float64 `json:"length,omitempty"`
			UnitOfMeasure string  `json:"unit-of-measure,omitempty"`
			Width         float64 `json:"width,omitempty"`
		} `json:"dimensions,omitempty"`
		Weight struct {
			UnitOfMeasure string  `json:"unit-of-measure,omitempty"`
			Weight        float64 `json:"weight,omitempty"`
		} `json:"weight,omitempty"`
	} `json:"retail-item,omitempty"`
	Seasonal         bool `json:"seasonal,omitempty"`
	ShelfLife        int  `json:"shelf-life,omitempty"`
	SlottingEligible bool `json:"slotting-eligible,omitempty"`
	SupplierInfo     []struct {
		SupplierID           string `json:"supplier-id,omitempty"`
		SupplierKeepBuy      bool   `json:"supplier-keep-buy,omitempty"`
		SupplierName         string `json:"supplier-name,omitempty"`
		SupplierPerformCycle int    `json:"supplier-perform-cycle,omitempty"`
		SupplierProductID    string `json:"supplier-product-id,omitempty"`
		SupplierSubAccount   string `json:"supplier-sub-account,omitempty"`
		SupplierType         string `json:"supplier-type,omitempty"`
	} `json:"supplier-info,omitempty"`
	TemperatureZone []string `json:"temperature-zone,omitempty"`
	TomID           string   `json:"tom-id,omitempty"`

	UnknownFields map[string]json.RawMessage `json:"-"`
}

func UnmarshalWithUnknownFields(v interface{}, data []byte, unknownFields *map[string]json.RawMessage) error {
	allFields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(data, &allFields); err != nil {
		return err
	}

	reflectType := reflect.TypeOf(v).Elem()
	reflectValue := reflect.ValueOf(v).Elem()
	for i := 0; i < reflectType.NumField(); i++ {
		fieldType := reflectType.Field(i)
		fieldValue := reflectValue.Field(i)
		tagValue, ok := fieldType.Tag.Lookup("json")
		if ok && tagValue != "-" {
			jsonFieldName := strings.Split(tagValue, ",")[0]
			rawValue, ok := allFields[jsonFieldName]
			if !ok {
				continue
			}
			if err := json.Unmarshal(rawValue, fieldValue.Addr().Interface()); err != nil {
				return fmt.Errorf("field: %v parsed with error: %v", jsonFieldName, err)
			}
			delete(allFields, jsonFieldName)
		}
	}

	*unknownFields = allFields
	return nil
}

func (p *Product) UnmarshalJSON(data []byte) error {
	return UnmarshalWithUnknownFields(p, data, &p.UnknownFields)
}
