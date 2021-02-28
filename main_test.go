package main

import "testing"

func TestCalcRevenue(t *testing.T) {
	var ad []AccountData
	ad = append(ad, AccountData{AccountCategory: "revenue", TotalValue: 32.40})
	ad = append(ad, AccountData{AccountCategory: "revenue", TotalValue: 122.23})

	totals := calcTotals(ad)

	exp := float32(154.63)
	if totals.Revenue != exp {
		t.Errorf("TestCalcRevenue test failed, got %f, expected %f\n", totals.Revenue, exp)
	} else {
		t.Logf("TestCalcRevenue test succeeded!")
	}
}

func TestCalcExpense(t *testing.T) {
	var ad []AccountData
	ad = append(ad, AccountData{AccountCategory: "expense", TotalValue: 55.40})
	ad = append(ad, AccountData{AccountCategory: "expense", TotalValue: 144.55})

	totals := calcTotals(ad)

	exp := float32(199.95)
	if totals.Expenses != exp {
		t.Errorf("TestCalcExpense test failed, got %f, expected %f\n", totals.Expenses, exp)
	} else {
		t.Logf("TestCalcExpense test succeeded!")
	}
}

func TestCalcGrossProfit(t *testing.T) {
	var ad []AccountData
	ad = append(ad, AccountData{AccountType: "sales", ValueType: "debit", TotalValue: 11.40})
	ad = append(ad, AccountData{AccountType: "sales", ValueType: "debit", TotalValue: 22.40})
	ad = append(ad, AccountData{AccountType: "sales", ValueType: "credit", TotalValue: 33.85})
	ad = append(ad, AccountData{AccountCategory: "revenue", TotalValue: 10.00})

	totals := calcTotals(ad)

	exp := float32(338.00)
	if totals.GrossProfitMargin != exp {
		t.Errorf("TestCalcGrossProfit test failed, got %f, expected %f\n", totals.GrossProfitMargin, exp)
	} else {
		t.Logf("TestCalcGrossProfit test succeeded!")
	}
}

func TestCalcNetProfitMargin(t *testing.T) {
	var ad []AccountData
	ad = append(ad, AccountData{AccountCategory: "expense", TotalValue: 20.00})
	ad = append(ad, AccountData{AccountCategory: "revenue", TotalValue: 80.00})
	ad = append(ad, AccountData{AccountCategory: "revenue", TotalValue: 20.00})

	totals := calcTotals(ad)

	exp := float32(80.00)
	if totals.NetProfitMargin != exp {
		t.Errorf("TestCalcNetProfitMargin test failed, got %f, expected %f\n", totals.NetProfitMargin, exp)
	} else {
		t.Logf("TestCalcNetProfitMargin test succeeded!")
	}
}

func TestCalcWorkingCapitalRatio(t *testing.T) {
	var ad []AccountData
	ad = append(ad, AccountData{AccountCategory: "assets", AccountType: "current", ValueType: "debit", TotalValue: 10.00})
	ad = append(ad, AccountData{AccountCategory: "liability", AccountType: "current", ValueType: "debit", TotalValue: 10.00})
	ad = append(ad, AccountData{AccountCategory: "assets", AccountType: "bank", ValueType: "debit", TotalValue: 15.00})
	ad = append(ad, AccountData{AccountCategory: "liability", AccountType: "bank", ValueType: "debit", TotalValue: 15.00})
	ad = append(ad, AccountData{AccountCategory: "assets", AccountType: "current_accounts_receivable", ValueType: "credit", TotalValue: 15.00})
	ad = append(ad, AccountData{AccountCategory: "liability", AccountType: "current_accounts_payable", ValueType: "credit", TotalValue: 15.00})

	totals := calcTotals(ad)

	exp := float32(-200.00)
	if totals.WorkingCaptialRatio != exp {
		t.Errorf("TestCalcNetProfitMargin test failed, got %f, expected %f\n", totals.WorkingCaptialRatio, exp)
	} else {
		t.Logf("TestCalcNetProfitMargin test succeeded!")
	}
}
