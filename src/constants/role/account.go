package role

var ROLE_ACCOUNT_WORDING = map[int]string{
	ROLE_ACCOUNT_ADMIN:  "admin",
	ROLE_ACCOUNT_CLIENT: "operator",
}

var ROLE_ACCOUNT_COLOR = map[int]string{
	ROLE_ACCOUNT_ADMIN:  "badge-primary",
	ROLE_ACCOUNT_CLIENT: "badge-danger",
}

const ROLE_ACCOUNT_ADMIN = 1
const ROLE_ACCOUNT_CLIENT = 2