package status

var COMPANY_IS_ENABLED_WORDING = map[bool]string{
	true:  COMPANY_IS_ENABLED_STATUS_TRUE,
	false: COMPANY_IS_ENABLED_STATUS_FALSE,
}

const COMPANY_IS_ENABLED_STATUS_TRUE = "company account active"
const COMPANY_IS_ENABLED_STATUS_FALSE = "company account deactive"
