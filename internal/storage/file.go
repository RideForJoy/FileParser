package storage

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const (
	dateRegEpp      = "\\d{14,}"
	layoutTimeStamp = "20060102150405" //yyyyMMddHHmmss
	layoutDate      = "20060102"       //yyyyMMdd
)

func MovePath(fileDate time.Time) string {
	yearString := strconv.Itoa(fileDate.Year())
	folderDatePart := yearString + "/" + fileDate.Format(layoutDate)
	return folderDatePart
}

func DateFromName(fileName string) (time.Time, error) {
	r := regexp.MustCompile(dateRegEpp)
	stringDates := r.FindStringSubmatch(fileName)
	if len(stringDates) == 0 {
		errorString := "the file fileName: " + fileName + " doesn't contains a string in the regexp: " + dateRegEpp
		return time.Time{}, errors.New(errorString)
	}

	fileDate, err := time.Parse(layoutTimeStamp, stringDates[0])
	if err != nil {
		errorString := "the file fileName: " + fileName + " doesn't contains a date by layout: " + layoutTimeStamp
		return time.Time{}, errors.New(errorString)
	}
	return fileDate, nil
}

func IsFresh(fileDate time.Time) bool {
	fileAge := int(fileDate.Sub(time.Now()).Hours()) / 24 //in days
	if fileAge > -32 && fileAge < 1 {
		return true
	} else {
		return false
	}
}

func IsPCv6(name string) bool {
	if strings.Contains(name, "takeoff_product_catalog_") ||
		strings.Contains(name, "Takeoff_product_catalog_") {
		return true
	} else {
		return false
	}
}
