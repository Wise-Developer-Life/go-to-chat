package exception

import "net/http"

type ResourceNotFoundError struct {
	ResourceName string
	ResourceID   string
}

func NewResourceNotFoundError(resourceName string, resourceID string) *ResourceNotFoundError {
	return &ResourceNotFoundError{ResourceName: resourceName, ResourceID: resourceID}
}

func (e *ResourceNotFoundError) Error() string {
	return e.ResourceName + " with ID " + e.ResourceID + " not found"
}

func (e *ResourceNotFoundError) Code() string {
	return "resource_not_found"
}

func (e *ResourceNotFoundError) HttpStatusCode() int {
	return http.StatusNotFound
}
