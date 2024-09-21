package store

import (
	"eps-backend/db"
	"eps-backend/model"
	"eps-backend/utils"
)

type RevenueConstruct struct {
	db db.DBConnection
}

func NewRevenueStore(db db.DBConnection) *RevenueConstruct {
	return &RevenueConstruct{db}
}

func (c *RevenueConstruct) GetRevenueByHour(path string) ([][]model.RevenuePerHour, error) {
	query := "SELECT * FROM v_revenue_3h where MONTH(tgl_entri) = MONTH(GETDATE()) AND YEAR(tgl_entri) = YEAR(GETDATE()) ORDER BY tgl_entri ASC"
	// execute query
	var result []model.RevenuePerHour
	if err := utils.SelectConn(path, c.db).Debug().Raw(query).Scan(&result).Error; err != nil {
		return nil, err
	}

	var resultByDate [][]model.RevenuePerHour
	for _, entry := range result {
		// Check if we already have an array for the current date
		found := false
		for i := range resultByDate {
			if resultByDate[i][0].TglEntri == entry.TglEntri {
				// Append the entry to the existing group
				resultByDate[i] = append(resultByDate[i], entry)
				found = true
				break
			}
		}
		// If no group found for the date, create a new group
		if !found {
			resultByDate = append(resultByDate, []model.RevenuePerHour{entry})
		}
	}
	return resultByDate, nil
}
