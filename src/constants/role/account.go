package role

var ROLE_ACCOUNT_WORDING = map[int]string{
	ROLE_ACCOUNT_ADMIN:    "admin",
	ROLE_ACCOUNT_OPERATOR: "operator",
}

var ROLE_ACCOUNT_COLOR = map[int]string{
	ROLE_ACCOUNT_ADMIN:    "badge-primary",
	ROLE_ACCOUNT_OPERATOR: "badge-danger",
}

const ROLE_ACCOUNT_ADMIN = 1
const ROLE_ACCOUNT_OPERATOR = 2

const ROLE_COMPANY_GOD = -999
const ROLE_COMPANY_ADMIN = 1
const ROLE_COMPANY_CLIENT = 2
