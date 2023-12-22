package util

import (
	"fmt"
	"strconv"
	"strings"

	"gorm.io/gorm"
)

func Paginate(q map[string]string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		var (
			page, size, offset int
			sortBy             string
			sortAsDesc         bool
		)

		if v, exist := q["page"]; exist {
			page, _ = strconv.Atoi(v)
		}
		if page <= 0 {
			page = 1
		}

		if v, exist := q["size"]; exist {
			size, _ = strconv.Atoi(v)
		}
		switch {
		case size > 100:
			size = 100
		case size <= 0:
			size = 10
		}

		offset = (page - 1) * size

		if v, exist := q["sort_by"]; exist {
			sortBy = v
		}
		if v, exist := q["sort_as"]; exist && strings.ToLower(v) == "desc" || v == "descending" || v == "0" {
			sortAsDesc = true
		}

		if sortBy != "" {
			var sort string
			if sortAsDesc {
				sort = fmt.Sprintf("%s %s", sortBy, "DESC")
			} else {
				sort = fmt.Sprintf("%s %s", sortBy, "ASC")
			}
			return db.Order(sort).Offset(offset).Limit(size)
		}

		return db.Offset(offset).Limit(size)
	}
}
