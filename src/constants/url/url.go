package constants

// PUBLIC or NON-ADMIN ENDPOINT
const URL_REQUEST_LOGIN string = "/auth/v1/login"
const URL_REQUEST_LOGOUT string = "/auth/v1/logout"
const URL_GET_USER_PROFILE string = "/auth/v1/get-profile"
const URL_CHECK_FEATURE string = "/auth/v1/validate-feature"
const URL_CHANGE_PASSWORD string = "/account/v1/change-password"

// ADMIN ACCOUNT ENDPOINT
const URL_ADMIN_GET_DETAIL_ACCOUNT string = "/admin/v1/get-account"
const URL_ADMIN_GET_DETAIL_COMPANY string = "/admin/v1/get-company"
const URL_ADMIN_GET_LIST_USER string = "/admin/v1/get-list-account"
const URL_ADMIN_GET_LIST_COMPANY string = "/admin/v1/get-list-company"
const URL_ADMIN_CREATE_ACCOUNT string = "/admin/v1/add-new-account"
const URL_ADMIN_UPDATE_ACCOUNT string = "/admin/v1/update-account"
const URL_ADMIN_CREATE_COMPANY string = "/admin/v1/add-new-company"
const URL_ADMIN_UPDATE_COMPANY string = "/admin/v1/update-company"
const URL_ADMIN_CHANGE_PASSWORD string = "/admin/v1/change-password"
const URL_ADMIN_CHANGE_CHANGE_STATUS_ACTIVATION string = "/admin/v1/change-status-activation"

// ORDER ENDPOINT
const URL_ADMIN_CREATE_ORDER_FOR_ADMIN string = "/order/v1/add"
const URL_ADMIN_GET_ORDER_DETAIL string = "/order/internal/v1/get"
const URL_ADMIN_GET_ORDER string = "/order/internal/v1/get"

// ADMIN FEATURE ENDPOINT
const URL_ADMIN_GET_LIST_MODULE string = "/admin/v1/account-module"
const URL_ADMIN_GET_LIST_FEATURE string = "/admin/v1/account-feature"

// Utility Endpoint (1 hit and re-use)
const URL_DESC_GET_LIST_ORDER_STATUS string = "/order/v1/get/order-status"
const URL_DESC_GET_LIST_PICKUP_STATUS string = "/order/v1/get/pickup-status"
const URL_ADMIN_GET_UNIT_ORDER string = "order/internal/v1/get/unit"
