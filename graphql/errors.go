package graphql

import (
	"encoding/json"
	"errors"
	"strings"

	"github.com/graphql-go/graphql/gqlerrors"
	"github.com/graphql-go/graphql/language/location"
)

// ContextKey is self explanatory
type ContextKey string

// MBError represents an application specific error
type MBError struct {
	Subject string `json:"subject"`
	Context string `json:"context"`
	Code    string `json:"code"`
	Error   string `json:"error"`
}

// JSONError format for GraphQL
type JSONError struct {
	Data   interface{} `json:"data"`
	Errors []GQLErrors `json:"errors"`
}

// GQLErrors struct
type GQLErrors struct {
	Message string `json:"message"`
}

const (
	//////////////////////////////////////
	// Error messages
	//////////////////////////////////////

	// Unauthorized is returned when no valid access objects have been supplied
	Unauthorized = "unauthorized"
	// MarformedID is returned when the UUID is malformed
	MarformedID = "malformed_id"
	// EmptyFieldError occurs when a field in the JSON body was empty
	EmptyFieldError = "empty_field_error"
	// InternalError occurs when we have a problem writing to the database
	InternalError = "internal_error"
	// JSONValidationError self explanatory
	JSONValidationError = "json_validation_error"
	// UserNotFound is returned if the UUID is not in the database
	UserNotFound = "user_not_found"
	// MissingUserID is self explanatory
	MissingUserID = "missing_user_id"
	// InvalidCredentials when email is malformed
	InvalidCredentials = "invalid_credentials"
	// InvalidEmailFormat when email is malformed
	InvalidEmailFormat = "invalid_email_format"
	// EmailAlreadyExists is self explanatory
	EmailAlreadyExists = "email_already_exists"
	// PublisherAlreadyExists is self explanatory
	PublisherAlreadyExists = "publisher_already_exists"
	// EmailNotFound is self explanatory
	EmailNotFound = "email_not_found"
	// RequestBodyEmpty is self explanatory
	RequestBodyEmpty = "request_body_empty"
	// MethodNotImplemented means exactly that
	MethodNotImplemented = "method_not_implemented"
	// MissingPublisherID is self explanatory
	MissingPublisherID = "missing_publisher_id"
	// PublisherNotFound is self explanatory
	PublisherNotFound = "publisher_not_found"
	// PhoneAlreadyExists is self explanatory
	PhoneAlreadyExists = "phone_already_exists"
	// PhoneNotFound is self explanatory
	PhoneNotFound = "phone_not_found"
	// AddressAlreadyExists is self explanatory
	AddressAlreadyExists = "address_already_exists"
	// AddressNameAlreadyExists is self explanatory
	AddressNameAlreadyExists = "address_name_already_exists"
	// PhoneNameAlreadyExists is self explanatory
	PhoneNameAlreadyExists = "phone_name_already_exists"
	// AddressNotFound is self explanatory
	AddressNotFound = "address_not_found"
	// RecordMismatch is self explanatory
	RecordMismatch = "record_mismatch"
	// CategoryAlreadyExists is self explanatory
	CategoryAlreadyExists = "category_already_exists"
	// PolicyAlreadyExists is self explanatory
	PolicyAlreadyExists = "policy_already_exists"
	// PolicyNameAlreadyExists is self explanatory
	PolicyNameAlreadyExists = "policy_name_already_exists"
	// PolicyNotFound is self explanatory
	PolicyNotFound = "policy_not_found"
	// CategoryNotFound is self explanatory
	CategoryNotFound = "category_not_found"
	// AdNotFound is self explanatory
	AdNotFound = "ad_not_found"
	// CampaignNotFound is self explanatory
	CampaignNotFound = "campaign_not_found"
	// CampaignAlreadyExists is self explanatory
	CampaignAlreadyExists = "campaign_already_exists"
	// CampaignAdAlreadyExists is self explanatory
	CampaignAdAlreadyExists = "campaign_ad_already_exists"
	// CampaignAdNotFound is self explanatory
	CampaignAdNotFound = "campaign_ad_not_found"
	// DemographicNotFound is self explanatory
	DemographicNotFound = "demographic_not_found"
	// CampaignDemographicAlreadyExists is self explanatory
	CampaignDemographicAlreadyExists = "campaign_demographic_already_exists"
	// CampaignDemographicNotFound is self explanatory
	CampaignDemographicNotFound = "campaign_demographic_not_found"
	// CharacteristicNotFound is self explanatory
	CharacteristicNotFound = "characteristic_not_found"
	// CharacteristicAlreadyExists is self explanatory
	CharacteristicAlreadyExists = "characteristic_already_exists"
	// DemographicCharacteristicNotFound is self explanatory
	DemographicCharacteristicNotFound = "demographic_characteristic_not_found"
	// DemographicCharacteristicAlreadyExists is self explanatory
	DemographicCharacteristicAlreadyExists = "demographic_characteristic_already_exists"
	// QueryParameterError is self explanatory
	QueryParameterError = "query_parameter_error"
	// AttributeNotFound is self explanatory
	AttributeNotFound = "attribute_not_found"
	// AttributeAlreadyExists is self explanatory
	AttributeAlreadyExists = "attribute_already_exists"
	// WeakPasswordError is self explanatory
	WeakPasswordError = "weak_password_error"
	// InvalidPasswordError is self explanatory
	InvalidPasswordError = "invalid_password_error"
	// InvalidJSONFormatInVariables is self explanatory
	InvalidJSONFormatInVariables = "invalid_json_format_variables"
	// UsernameExists is self explanatory
	UsernameExists = "username_exists"
	// UnableToCreateUser is self explanatory
	UnableToCreateUser = "unable_to_create_user"
	// AccountLocked is self explanatory
	AccountLocked = "account_locked"
	// ProductAlreadyExists is self explanatory
	ProductAlreadyExists = "product_already_exists"
	// ProductNotFound is self explanatory
	ProductNotFound = "product_not_found"
	// ProductAttributeAlreadyExists is self explanatory
	ProductAttributeAlreadyExists = "product_attribute_already_exists"
	// ProductAttributeNotFound is self explanatory
	ProductAttributeNotFound = "product_attribute_not_found"
	// TypeRecursionViolation is self explanatory
	TypeRecursionViolation = "type_recursion_violoation"
	// AdAlreadyExists is self explanatory
	AdAlreadyExists = "ad_already_exists"
	// SubscriptionNotFound is self explanatory
	SubscriptionNotFound = "subscription_not_found"
	// SubscriptionAlreadyExists is self explanatory
	SubscriptionAlreadyExists = "subscription_already_exists"

	//////////////////////////////////////
	// Response messages
	//////////////////////////////////////

	// PasswordUpdated is self explanatory
	PasswordUpdated = "password_updated"
)

var messageMap = map[string]string{
	"password_updated":                          "password has been updated successfully",
	"unauthorized":                              "Access to this resource is restricted",
	"user_not_found":                            "the requested user was not found",
	"attribute_not_found":                       "the requested attribute was not found",
	"attribute_already_exists":                  "the attribute already exists",
	"query_parameter_error":                     "Some parameters in the query are invalid or not found",
	"malformed_id":                              "the id provided has an invalid format",
	"empty_field_error":                         "One or more empty fields were found in JSON message",
	"internal_error":                            "something went wrong.  please contact us @ 1-800-555-1212",
	"json_validation_error":                     "JSON schema validation failed in one or more ",
	"invalid_email_format":                      "Email format is invalid",
	"invalid_credentials":                       "Invalid user id or password",
	"email_already_exists":                      "the provided email has already been registered",
	"publisher_already_exists":                  "the provided marketer name has already been registered",
	"request_body_empty":                        "the request did not contain a boy or it's empty",
	"method_not_implemented":                    "the requested method is not being used",
	"missing_user_id":                           "You are missing a 'user_id' field",
	"missing_publisher_id":                      "You are missing a 'publisher_id' field",
	"email_not_found":                           "the requested email was not found",
	"publisher_not_found":                       "the requested marketer was not found",
	"phone_already_exists":                      "the provided phone number has already been registered",
	"phone_not_found":                           "the requested phone was not found",
	"address_already_exists":                    "the provided address has already been registered",
	"address_name_already_exists":               "the provided address label already exists",
	"phone_name_already_exists":                 "the provided phone label already exists",
	"address_not_found":                         "the requested address was not found",
	"record_mismatch":                           "the record and id pair are not associated",
	"category_already_exists":                   "the provided category has already been registered",
	"policy_already_exists":                     "a policy with that name has already been created",
	"policy_name_already_exists":                "a policy with that name already exists. Please use a different name",
	"policy_not_found":                          "the provided policy was not found",
	"category_not_found":                        "the requested category was not found",
	"campaign_not_found":                        "the requested campaign was not found",
	"campaign_ad_not_found":                     "the Campaign Ad realtionship was not found",
	"campaign_already_exists":                   "the Campaign already exists",
	"campaign_ad_already_exists":                "the Campaign Ad realtionship already exists",
	"demographic_not_found":                     "the requested demographic was not found",
	"campaign_demographic_already_exists":       "the Campaign Demographic realtionship already exists",
	"campaign_demographic_not_found":            "the Campaign Demographic realtionship not found",
	"product_attribute_already_exists":          "the product attribute already exists",
	"product_attribute_not_found":               "the product attribute was not found",
	"product_already_exists":                    "the product already exists",
	"product_not_found":                         "the product was not found",
	"demographic_characteristic_already_exists": "the Demographic Characteristic already exists",
	"demographic_characteristic_not_found":      "the Demographic Characteristic was not found",
	"weak_password_error":                       "password is weak",
	"invalid_password_error":                    "Invalid password",
	"invalid_json_format_variables":             "Invalid JSON format in variables",
	"username_exists":                           "username already exists",
	"unable_to_create_user":                     "We are unable to create the user account.  Please contact customer support.",
	"account_locked":                            "the user account is locked",
	"type_recursion_violoation":                 "your query contains a type within the same type. e.g. {product {maketer{product}}}",
	"ad_already_exists":                         "the ad already exists",
	"ad_not_found":                              "the requested ad was not found",
	"subscription_already_exists":               "the subscription already exists",
	"subscription_not_found":                    "the requested subscription was not found",
}

// Error function
func Error(code string) error {
	return errors.New(code)
}

// ErrorMessage function
func ErrorMessage(code string) error {
	return errors.New(messageMap[code])
}

// FormatError function creates a graphql error with location information
func FormatError(field string, argument string, code string, locations []location.SourceLocation) gqlerrors.FormattedError {
	e := gqlerrors.FormattedError{
		Message:   strings.Join([]string{field, argument, code, messageMap[code]}, ":"),
		Locations: locations,
	}
	return e
}

// SimpleFormattedError function creates a graphql error with no location
func SimpleFormattedError(field string, argument string, code string) error {
	return errors.New(strings.Join([]string{field, argument, code, messageMap[code]}, ":"))
}

// SimpleJSONFormattedError function creates a graphql error with no location
func SimpleJSONFormattedError(field string, argument string, code string) error {

	mbe := MBError{
		Subject: field,
		Context: argument,
		Code:    code,
		Error:   messageMap[code],
	}

	d, _ := json.Marshal(&mbe)

	m := JSONError{
		Data: nil,
		Errors: []GQLErrors{
			GQLErrors{
				Message: string(d),
			},
		},
	}

	msg, err := json.Marshal(m)
	if err == nil {
		return errors.New(string(msg))
		// return fmt.Errorf("%s", msg)
	}

	return err
}
