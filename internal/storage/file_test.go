package storage_test

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"

	"github.com/AndriiShevchun/parse-pc-files/internal/storage"
)

func TestMovePath(t *testing.T) {
	tests := []struct {
		name         string
		fileDate     time.Time
		expectedPath string
	}{
		{
			name:         "Parse date ",
			fileDate:     time.Date(2023, time.February, 23, 10, 11, 12, 0, time.UTC),
			expectedPath: "2023/20230223",
		},
		{
			name:         "Parse any date ",
			fileDate:     time.Date(1900, time.September, 05, 0, 0, 0, 0, time.UTC),
			expectedPath: "1900/19000905",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expectedPath, storage.MovePath(test.fileDate))
		})
	}

}

func TestDateFromName(t *testing.T) {
	tests := []struct {
		name          string
		fileName      string
		expectedTime  time.Time
		expectedError error
	}{
		{
			name:          "Correct file with h/m/s",
			fileName:      "Takeoff_product_catalog_diff_2006_20230223101112.json",
			expectedTime:  time.Date(2023, time.February, 23, 10, 11, 12, 0, time.UTC),
			expectedError: nil,
		},
		{
			name:          "Correct file with 0",
			fileName:      "Takeoff_product_catalog_diff_2006_20230223000000.json",
			expectedTime:  time.Date(2023, time.February, 23, 0, 0, 0, 0, time.UTC),
			expectedError: nil,
		},
		{
			name:          "Incorrect file name without date by regexp",
			fileName:      "Takeoff_product_catalog_diff_2006_2023.json",
			expectedTime:  time.Time{},
			expectedError: errors.New("the file fileName: Takeoff_product_catalog_diff_2006_2023.json doesn't contains a string in the regexp: \\d{14,}"),
		},
		{
			name:          "Incorrect file without date",
			fileName:      "Takeoff_product_catalog_diff_2006_20239999999999.json",
			expectedTime:  time.Time{},
			expectedError: errors.New("the file fileName: Takeoff_product_catalog_diff_2006_20239999999999.json doesn't contains a date by layout: 20060102150405"),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := storage.DateFromName(test.fileName)
			assert.Equal(t, result, test.expectedTime)
			assert.Equal(t, err, test.expectedError)
		})
	}
}

func TestIsFresh(t *testing.T) {
	tests := []struct {
		name           string
		fileDate       time.Time
		expectedResult bool
	}{
		{
			name:           "Now",
			fileDate:       time.Now(),
			expectedResult: true,
		},
		{
			name:           "Last month",
			fileDate:       time.Now().AddDate(0, 0, -31),
			expectedResult: true,
		},
		{
			name:           "More that last month",
			fileDate:       time.Now().AddDate(0, 0, -32),
			expectedResult: false,
		},
		{
			name:           "Tomorrow",
			fileDate:       time.Now().AddDate(0, 0, 1),
			expectedResult: true,
		},
		{
			name:           "Day after tomorrow",
			fileDate:       time.Now().AddDate(0, 0, 2),
			expectedResult: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, storage.IsFresh(test.fileDate), test.expectedResult)
		})
	}
}

func TestIsPCv6(t *testing.T) {
	tests := []struct {
		name           string
		fileName       string
		expectedResult bool
	}{
		{
			name:           "UpperCase",
			fileName:       "Takeoff_product_catalog_diff_2006_20230223010010.json",
			expectedResult: true,
		},
		{
			name:           "LoverCase",
			fileName:       "takeoff_product_catalog_diff_2006_20230223010010.json",
			expectedResult: true,
		},
		{
			name:           "With folder",
			fileName:       "data/takeoff_product_catalog_diff_.json",
			expectedResult: true,
		},
		{
			name:           "Not PCv6 file name",
			fileName:       "another_file.json",
			expectedResult: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, storage.IsPCv6(test.fileName), test.expectedResult)
		})
	}
}
