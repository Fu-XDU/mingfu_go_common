package datetime

import "time"

const yearMonthLayout = "2006-01"

// YearMonthStringToDatetime 将格式为YYYY-MM的日期字符串转为*time.Time
// 如果日期字符串为空或不符合格式，返回nil
func YearMonthStringToDatetime(dateStr string) (birthdate *time.Time) {
	if len(dateStr) == 0 {
		return
	}

	parsedDate, err := time.Parse(yearMonthLayout, dateStr)
	if err != nil {
		// 处理解析错误（如格式不正确）
		return nil
	}
	birthdate = &parsedDate
	return
}

// BirthdateToYearMonthString 将日期转为YYYY-MM格式的字符串，如果参数为nil，返回空字符串
func BirthdateToYearMonthString(birthdate *time.Time) (output string) {
	if birthdate == nil {
		return
	}

	output = birthdate.Format(yearMonthLayout)
	return
}

func StartEarlierThanEnd(startDateStr, endDateStr string) bool {
	return !YearMonthStringToDatetime(startDateStr).After(*YearMonthStringToDatetime(endDateStr))
}
