// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package servicequotas

import (
	"github.com/aws/aws-sdk-go/private/protocol"
)

const (

	// ErrCodeAWSServiceAccessNotEnabledException for service response error code
	// "AWSServiceAccessNotEnabledException".
	//
	// The action you attempted is not allowed unless Service Access with Service
	// Quotas is enabled in your organization. To enable, call AssociateServiceQuotaTemplate.
	ErrCodeAWSServiceAccessNotEnabledException = "AWSServiceAccessNotEnabledException"

	// ErrCodeAccessDeniedException for service response error code
	// "AccessDeniedException".
	//
	// You do not have sufficient access to perform this action.
	ErrCodeAccessDeniedException = "AccessDeniedException"

	// ErrCodeDependencyAccessDeniedException for service response error code
	// "DependencyAccessDeniedException".
	//
	// You can't perform this action because a dependency does not have access.
	ErrCodeDependencyAccessDeniedException = "DependencyAccessDeniedException"

	// ErrCodeIllegalArgumentException for service response error code
	// "IllegalArgumentException".
	//
	// Invalid input was provided.
	ErrCodeIllegalArgumentException = "IllegalArgumentException"

	// ErrCodeInvalidPaginationTokenException for service response error code
	// "InvalidPaginationTokenException".
	//
	// Invalid input was provided.
	ErrCodeInvalidPaginationTokenException = "InvalidPaginationTokenException"

	// ErrCodeInvalidResourceStateException for service response error code
	// "InvalidResourceStateException".
	//
	// Invalid input was provided for the .
	ErrCodeInvalidResourceStateException = "InvalidResourceStateException"

	// ErrCodeNoAvailableOrganizationException for service response error code
	// "NoAvailableOrganizationException".
	//
	// The account making this call is not a member of an organization.
	ErrCodeNoAvailableOrganizationException = "NoAvailableOrganizationException"

	// ErrCodeNoSuchResourceException for service response error code
	// "NoSuchResourceException".
	//
	// The specified resource does not exist.
	ErrCodeNoSuchResourceException = "NoSuchResourceException"

	// ErrCodeOrganizationNotInAllFeaturesModeException for service response error code
	// "OrganizationNotInAllFeaturesModeException".
	//
	// The organization that your account belongs to, is not in All Features mode.
	// To enable all features mode, see EnableAllFeatures (https://docs.aws.amazon.com/organizations/latest/APIReference/API_EnableAllFeatures.html).
	ErrCodeOrganizationNotInAllFeaturesModeException = "OrganizationNotInAllFeaturesModeException"

	// ErrCodeQuotaExceededException for service response error code
	// "QuotaExceededException".
	//
	// You have exceeded your service quota. To perform the requested action, remove
	// some of the relevant resources, or use Service Quotas to request a service
	// quota increase.
	ErrCodeQuotaExceededException = "QuotaExceededException"

	// ErrCodeResourceAlreadyExistsException for service response error code
	// "ResourceAlreadyExistsException".
	//
	// The specified resource already exists.
	ErrCodeResourceAlreadyExistsException = "ResourceAlreadyExistsException"

	// ErrCodeServiceException for service response error code
	// "ServiceException".
	//
	// Something went wrong.
	ErrCodeServiceException = "ServiceException"

	// ErrCodeServiceQuotaTemplateNotInUseException for service response error code
	// "ServiceQuotaTemplateNotInUseException".
	//
	// The quota request template is not associated with your organization.
	//
	// To use the template, call AssociateServiceQuotaTemplate.
	ErrCodeServiceQuotaTemplateNotInUseException = "ServiceQuotaTemplateNotInUseException"

	// ErrCodeTemplatesNotAvailableInRegionException for service response error code
	// "TemplatesNotAvailableInRegionException".
	//
	// The Service Quotas template is not available in the Region where you are
	// making the request. Please make the request in us-east-1.
	ErrCodeTemplatesNotAvailableInRegionException = "TemplatesNotAvailableInRegionException"

	// ErrCodeTooManyRequestsException for service response error code
	// "TooManyRequestsException".
	//
	// Due to throttling, the request was denied. Slow down the rate of request
	// calls, or request an increase for this quota.
	ErrCodeTooManyRequestsException = "TooManyRequestsException"
)

var exceptionFromCode = map[string]func(protocol.ResponseMetadata) error{
	"AWSServiceAccessNotEnabledException":       newErrorAWSServiceAccessNotEnabledException,
	"AccessDeniedException":                     newErrorAccessDeniedException,
	"DependencyAccessDeniedException":           newErrorDependencyAccessDeniedException,
	"IllegalArgumentException":                  newErrorIllegalArgumentException,
	"InvalidPaginationTokenException":           newErrorInvalidPaginationTokenException,
	"InvalidResourceStateException":             newErrorInvalidResourceStateException,
	"NoAvailableOrganizationException":          newErrorNoAvailableOrganizationException,
	"NoSuchResourceException":                   newErrorNoSuchResourceException,
	"OrganizationNotInAllFeaturesModeException": newErrorOrganizationNotInAllFeaturesModeException,
	"QuotaExceededException":                    newErrorQuotaExceededException,
	"ResourceAlreadyExistsException":            newErrorResourceAlreadyExistsException,
	"ServiceException":                          newErrorServiceException,
	"ServiceQuotaTemplateNotInUseException":     newErrorServiceQuotaTemplateNotInUseException,
	"TemplatesNotAvailableInRegionException":    newErrorTemplatesNotAvailableInRegionException,
	"TooManyRequestsException":                  newErrorTooManyRequestsException,
}
