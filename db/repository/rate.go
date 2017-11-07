package repository

import  "github.com/jmoiron/sqlx"

// Rate asd
type Rate struct {
	RateID       int     `db:"rate_id"json:"rate_id"`
	PropertyID   int     `db:"property_id"json:"property_id"`
	Title        string  `db:"title"json:"title"`
	FromDate     string  `db:"from_date"json:"from_date"`
	ToDate       string  `db:"to_date"json:"to_date"`
	Currency     string  `db:"currency"json:"currency"`
	NightlyPrice string  `db:"nightly_price"json:"nightly_price"`
	WeeklyPrice  float32 `db:"weekly_price"json:"weekly_price"`
	MonthlyPrice float32 `db:"monthly_price"json:"monthly_price"`
	MinStayDays  int     `db:"min_stay_days"json:"min_stay_days"`
}

func GetAll() ([]*Rate, error) {
	rates := make([]*Rate, 0)

	err := sqlx.DB.Select(&rates, "SELECT rate_id FROM Rate")

	if err != nil {
		panic(err)
	}

	return rates, nil
}
