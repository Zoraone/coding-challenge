package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
  accountData, err := readFile()
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

  fmt.Println(revenue)
  fmt.Println(expenses)
  fmt.Println(float64(grossProfit) / float64(revenue))
  fmt.Println(float64(netProfit) / float64(revenue))
  fmt.Printf("%.1f%%", (float32(assets) / float32(liabilities) * 100))
}

func readFile() ([]AccountData, error) {
  jsonFile, err := os.Open("data.json")
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