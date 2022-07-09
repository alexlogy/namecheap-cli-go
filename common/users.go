package common

import (
	"encoding/xml"
	"log"
	"namecheap-cli/consts"
	"namecheap-cli/models"
	"namecheap-cli/utils"
)

func UsersGetBalance() {
	payload := make(map[string]string)
	payload["Command"] = "namecheap.users.getBalances"

	var usersGetBalanceResponse models.UsersGetBalanceResponse

	resp := utils.PostRequest(consts.NamecheapAPI, payload)
	
	err := xml.Unmarshal(resp, &usersGetBalanceResponse)
	if err != nil {
		log.Fatalln(err)
	}

	usersGetBalanceResults := *usersGetBalanceResponse.CommandResponse.UserGetBalancesResult
	utils.PrintUsersGetBalanceTable(usersGetBalanceResults)
}
