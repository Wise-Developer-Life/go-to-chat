package exception

import (
	"fmt"
	"net/http"
)

type ResourceConflictError struct {
	resourceName   string
	conflictReason string
}

func NewResourceConflictError(resourceName string, conflictReason string) *ResourceConflictError {
	return &ResourceConflictError{
		resourceName:   resourceName,
		conflictReason: conflictReason,
	}
}

func (r *ResourceConflictError) Error() string {
	return fmt.Sprintf("Resource %s has conflict: %s", r.resourceName, r.conflictReason)
}

func (r *ResourceConflictError) Code() string {
	return "resource_conflict"
}

func (r *ResourceConflictError) HttpStatusCode() int {
	return http.StatusConflict
}
