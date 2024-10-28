package util

import (
	"fmt"
	"regexp"
	"time"

	"gorm.io/gorm"
)

// YEar Info
var (
	yearInfoT string = "common" + "." + "yearinfo"
	yearcodeC string = "yearcode"
)

// Company Info
var (
	compinfoT string = "common" + "." + "compinfo"
	compcodeC string = "compcode"
)

// GetFinancialYear Returns the current financial year
func GetFinancialYear() string {
	y, m, d := time.Now().Date()
	if m >= time.July && d >= 1 {
		return fmt.Sprintf("%d-%d", y, y+1)
	}
	return fmt.Sprintf("%d-%d", y-1, y)
}

// GetFinancialYearFromDate returns financial year from input date
func GetFinancialYearFromDate(date string) (string, error) {
	dt, err := time.Parse("2006-01-02", date)
	if err != nil {
		return "", fmt.Errorf("invalid date provided")
	}
	y, m, d := dt.Date()
	if m >= time.July && d >= 1 {
		return fmt.Sprintf("%d-%d", y, y+1), nil
	}
	return fmt.Sprintf("%d-%d", y-1, y), nil
}

// GetFinancialYearFromDateTime returns financial year from input date
func GetFinancialYearFromDateTime(date *time.Time) string {
	y, m, d := date.Date()
	if m >= time.July && d >= 1 {
		return fmt.Sprintf("%d-%d", y, y+1)
	}
	return fmt.Sprintf("%d-%d", y-1, y)
}

func ValidateEmail(email string) bool {
	// Regular expression for a basic email validation
	// regex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	regex := `^[a-zA-Z0-9]+([._-][a-zA-Z0-9]+)*@[a-zA-Z]+([.-][a-zA-Z]+)*\.[a-zA-Z]{2,}$`
	//`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`
	//regex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.com$`

	// Compile the regex
	re := regexp.MustCompile(regex)

	// Use MatchString method to check if the email matches the pattern
	return re.MatchString(email)
}

func IsPhoneValid(e string) bool {
	// phoneno := regexp.MustCompile(`^01+\d{9}$`)
	phoneno := regexp.MustCompile(`^\+?(88)?0?1[3456789][0-9]{8}\b`)
	return phoneno.MatchString(e)
}

func Company_Year_into(db *gorm.DB) ([]int, []int) {

	var yearList []int
	var companyList []int

	queryYear := "SELECT " + yearcodeC + " FROM " + yearInfoT
	queryComp := "SELECT " + compcodeC + " FROM " + compinfoT
	_ = db.Raw(queryYear).Find(&yearList)
	_ = db.Raw(queryComp).Find(&companyList)

	return yearList, companyList
}

func GenerateArchiveDateTime() (string, string) {
	changeDate := time.Now().Format("2006-01-02")
	bdTimeZone, _ := time.LoadLocation("Asia/Dhaka")
	changeTime := time.Now().In(bdTimeZone).Format("15:04:05")
	return changeDate,changeTime
}

//archiveTable.ChangeDate,archiveTable.ChangeTime = util.GenerateArchiveDateTime()