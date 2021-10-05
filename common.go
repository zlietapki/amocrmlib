package amocrmlib

const (
	PhoneWork   = "WORK"
	PhoneWorkDD = "WORKDD"
	PhoneMob    = "MOB"
	PhoneFax    = "FAX"
	PhoneHome   = "HOME"
	PhoneOther  = "OTHER"
)

const EntityTypeLeads = "leads"
const EntityTypeContacts = "contacts"
const EntityTypeCompanies = "companies"

//func addPhoneToCustomFields(fields []*apimodel.CustomFieldsValue, phone string, phoneType string) []*apimodel.CustomFieldsValue {
//	var phonesField *apimodel.CustomFieldsValue
//	for _, f := range fields {
//		if f.FieldCode == "PHONE" {
//			phonesField = f
//			break
//		}
//	}
//	if phonesField == nil {
//		phonesField = &apimodel.CustomFieldsValue{
//			FieldCode: "PHONE",
//			Values:    []*apimodel.CustomFieldsValueItem{},
//		}
//		fields = append(fields, phonesField)
//	}
//
//	phonesField.Values = append(phonesField.Values, &apimodel.CustomFieldsValueItem{
//		Value:    phone,
//		EnumCode: phoneType,
//	})
//	return fields
//}
//
//func phonesByType(fields []*apimodel.CustomFieldsValue, phoneType string) []string {
//	phonesList := make([]string, 0)
//	for _, f := range fields {
//		if f.FieldCode == "PHONE" {
//			for _, v := range f.Values {
//				if v.EnumCode == phoneType {
//					phonesList = append(phonesList, v.Value)
//				}
//			}
//		}
//	}
//
//	if len(phonesList) == 0 {
//		return nil
//	}
//	return phonesList
//}
//
//func clearPhoneInCustomFields(fields []*apimodel.CustomFieldsValue, phone string, phoneType string) error {
//	var phonesField *apimodel.CustomFieldsValue
//	for _, f := range fields {
//		if f.FieldCode == "PHONE" {
//			phonesField = f
//			break
//		}
//	}
//	if phonesField == nil {
//		return errors.Errorf("phone not found: %s phoneType: %s", phone, phoneType)
//	}
//
//	found := false
//	for i := range phonesField.Values {
//		if phonesField.Values[i].Value == phone && phonesField.Values[i].EnumCode == phoneType {
//			found = true
//			phonesField.Values[i].Value = ""
//		}
//	}
//	if !found {
//		return errors.Errorf("phone not found: %s phoneType: %s", phone, phoneType)
//	}
//	return nil
//}
