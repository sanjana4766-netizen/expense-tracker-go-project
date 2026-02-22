package main

import "sort"

type Balance struct {
	UserID uint
	Amt    float64
}

type Txn struct {
	From uint
	To   uint
	Amt  float64
}

func CalculateSettlement(groupID uint) []Txn {
	var expenses []Expense
	var splits []Split

	DB.Where("group_id = ?", groupID).Find(&expenses)

	var ids []uint
	for _, e := range expenses {
		ids = append(ids, e.ID)
	}
	DB.Where("expense_id IN ?", ids).Find(&splits)

	bal := map[uint]float64{}

	for _, e := range expenses {
		bal[e.PaidBy] += e.Amount
	}
	for _, s := range splits {
		bal[s.UserID] -= s.Amount
	}

	var cr, db []Balance
	for u, a := range bal {
		if a > 0 {
			cr = append(cr, Balance{u, a})
		} else if a < 0 {
			db = append(db, Balance{u, -a})
		}
	}

	sort.Slice(cr, func(i, j int) bool { return cr[i].Amt > cr[j].Amt })
	sort.Slice(db, func(i, j int) bool { return db[i].Amt > db[j].Amt })

	var tx []Txn
	i, j := 0, 0
	for i < len(db) && j < len(cr) {
		m := min(db[i].Amt, cr[j].Amt)
		tx = append(tx, Txn{db[i].UserID, cr[j].UserID, m})
		db[i].Amt -= m
		cr[j].Amt -= m
		if db[i].Amt == 0 {
			i++
		}
		if cr[j].Amt == 0 {
			j++
		}
	}
	return tx
}

func min(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}