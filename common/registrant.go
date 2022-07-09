package common

import (
	"namecheap-cli/consts"
	"namecheap-cli/models"
)

var (
	registrant *models.Registrant
)

func InitRegistrant() {

	registrant = new(models.Registrant)

	registrant.RegistrantFirstName = consts.FirstName
	registrant.RegistrantLastName = consts.LastName
	registrant.RegistrantAddress1 = consts.Address1
	registrant.RegistrantCity = consts.City
	registrant.RegistrantStateProvince = consts.StateProvince
	registrant.RegistrantPostalCode = consts.PostalCode
	registrant.RegistrantCountry = consts.Country
	registrant.RegistrantPhone = consts.Phone
	registrant.RegistrantEmailAddress = consts.EmailAddress
	registrant.TechFirstName = consts.FirstName
	registrant.TechLastName = consts.LastName
	registrant.TechAddress1 = consts.Address1
	registrant.TechCity = consts.City
	registrant.TechStateProvince = consts.StateProvince
	registrant.TechPostalCode = consts.PostalCode
	registrant.TechCountry = consts.Country
	registrant.TechPhone = consts.Phone
	registrant.TechEmailAddress = consts.EmailAddress
	registrant.AdminFirstName = consts.FirstName
	registrant.AdminLastName = consts.LastName
	registrant.AdminAddress1 = consts.Address1
	registrant.AdminCity = consts.City
	registrant.AdminStateProvince = consts.StateProvince
	registrant.AdminPostalCode = consts.PostalCode
	registrant.AdminCountry = consts.Country
	registrant.AdminPhone = consts.Phone
	registrant.AdminEmailAddress = consts.EmailAddress
	registrant.AuxBillingFirstName = consts.FirstName
	registrant.AuxBillingLastName = consts.LastName
	registrant.AuxBillingAddress1 = consts.Address1
	registrant.AuxBillingCity = consts.City
	registrant.AuxBillingStateProvince = consts.StateProvince
	registrant.AuxBillingPostalCode = consts.PostalCode
	registrant.AuxBillingCountry = consts.Country
	registrant.AuxBillingPhone = consts.Phone
	registrant.AuxBillingEmailAddress = consts.EmailAddress
}
