package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/leekchan/accounting"
)

func main() {
  accountData, err := readFile("data.json")
  if err != nil {
    fmt.Println(err)
  }

  revenue := 0
  expenses := 0
  grossProfit := 0
  assets := 0
  liabilities := 0
  for _, ad := range accountData {
    totalAsInt := int(ad.TotalValue * 100)

    switch ad.AccountCategory {
    case "revenue":
      revenue += totalAsInt
    case "expense":
      expenses += totalAsInt
    case "assets":
      switch ad.ValueType {
      case "debit":
        switch ad.AccountType {
        case "current", "bank", "current_accounts_receivable":
          assets += totalAsInt
        }
      case "credit":
        switch ad.AccountType {
        case "current", "bank", "current_accounts_receivable":
          assets -= totalAsInt
        }
      }
    case "liability":
      switch ad.ValueType {
      case "debit":
        switch ad.AccountType {
        case "current", "current_accounts_payable":
          liabilities += totalAsInt
        }
      case "credit":
        switch ad.AccountType {
        case "current", "current_accounts_payable":
          liabilities -= totalAsInt
        }
      }
    }

    if ad.AccountType == "sales" && ad.ValueType == "debit" {
      grossProfit += totalAsInt
    }
  }

  netProfit := revenue - expenses

  ac := accounting.Accounting{Symbol: "$"}
  fmt.Printf("Revenue: %s\n", ac.FormatMoney(float32(revenue)/100.0))
  fmt.Printf("Expenses: %s\n", ac.FormatMoney(float32(expenses)/100.0))
  fmt.Printf("Gross Profit Margin: %.1f%%\n", (float64(grossProfit) / float64(revenue) * 100))
  fmt.Printf("Net Profit Margin: %.1f%%\n", (float64(netProfit) / float64(revenue) * 100))
  fmt.Printf("Working Capital Ratio: %.1f%%\n", (float32(assets) / float32(liabilities) * 100))
}

func readFile(filename string) ([]AccountData, error) {
  jsonFile, err := os.Open(filename)
  if err != nil {
    return nil, err
  }
  defer jsonFile.Close()

  jsonBytes, err := ioutil.ReadAll(jsonFile)
  if err != nil {
    return nil, err
  }

  var gl GeneralLedger
  json.Unmarshal(jsonBytes, &gl)

  return gl.Data, nil
}