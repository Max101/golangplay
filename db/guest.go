package db

import (
	sq "github.com/Masterminds/squirrel"
	"log"
)

type GuestRepo struct {
}

// Guest asd
type Guest struct {
	GuestID         int    `db:"guest_id" json:"guest_id"`
	FirstName       string `db:"first_name" json:"first_name"`
	LastName        string `db:"last_name" json:"last_name"`
	CountryCode     string `db:"country" json:"country_code"`
	CreatedDatetime string `db:"created_datetime" json:"created_datetime"`
}

type GuestFilter struct {
	Limit   uint64
	Offset  uint64
	Country string
}

func NewGuestFilter() (guestFilter GuestFilter) {
	guestFilter = GuestFilter{10, 0, ""}
	return
}

func (gr GuestRepo) Get(filters GuestFilter) (guests []*Guest, total int, err error) {
	guests = make([]*Guest, 0)

	userQuery := getUserQuery(false).
		Limit(filters.Limit).
		Offset(filters.Offset)

	if filters.Country != "" {
		userQuery = userQuery.Where(sq.Eq{"country": filters.Country})
	}

	sql, args, err := userQuery.ToSql()

	if err != nil {
		log.Panic(err)
		return guests, 0, err
	}

	err = DB().Select(&guests, sql, args...)

	if err != nil {
		log.Panic(err)
	}

	sql, args, err = getUserQuery(true).Where(sq.Eq{"country": "it"}).ToSql()

	err = DB().Get(&total, sql, args...)

	if err != nil {
		log.Panic(err)
	}

	return guests, total, err
}

func getUserQuery(selectTotals bool) (userQuery sq.SelectBuilder) {

	var selectQuery1 string

	if selectTotals {
		selectQuery1 = "count(*)"
	} else {
		selectQuery1 = GetStructDBFields(Guest{})
	}

	userQuery = sq.
		Select(selectQuery1).
		From("Guest")

	//.Join("emails USING (email_id)")

	return
}
