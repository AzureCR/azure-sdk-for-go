package managementpartner

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"encoding/json"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/date"
	"github.com/Azure/go-autorest/autorest/to"
	"net/http"
)

// Code enumerates the values for code.
type Code string

const (
	// BadRequest ...
	BadRequest Code = "BadRequest"
	// Conflict ...
	Conflict Code = "Conflict"
	// NotFound ...
	NotFound Code = "NotFound"
)

// PossibleCodeValues returns an array of possible values for the Code const type.
func PossibleCodeValues() []Code {
	return []Code{BadRequest, Conflict, NotFound}
}

// State enumerates the values for state.
type State string

const (
	// Active ...
	Active State = "Active"
	// Deleted ...
	Deleted State = "Deleted"
)

// PossibleStateValues returns an array of possible values for the State const type.
func PossibleStateValues() []State {
	return []State{Active, Deleted}
}

// Error this is the management partner operations error
type Error struct {
	// Error - this is the ExtendedErrorInfo property
	Error *ExtendedErrorInfo `json:"error,omitempty"`
}

// ExtendedErrorInfo this is the extended error info
type ExtendedErrorInfo struct {
	// Code - this is the error response code. Possible values include: 'NotFound', 'Conflict', 'BadRequest'
	Code Code `json:"code,omitempty"`
	// Message - this is the extended error info message
	Message *string `json:"message,omitempty"`
}

// OperationDisplay this is the management partner operation
type OperationDisplay struct {
	// Provider - the is management partner provider
	Provider *string `json:"provider,omitempty"`
	// Resource - the is management partner resource
	Resource *string `json:"resource,omitempty"`
	// Operation - the is management partner operation
	Operation *string `json:"operation,omitempty"`
	// Description - the is management partner operation description
	Description *string `json:"description,omitempty"`
}

// OperationList this is the management partner operations list
type OperationList struct {
	autorest.Response `json:"-"`
	// Value - this is the operation response list
	Value *[]OperationResponse `json:"value,omitempty"`
	// NextLink - Url to get the next page of items.
	NextLink *string `json:"nextLink,omitempty"`
}

// OperationListIterator provides access to a complete listing of OperationResponse values.
type OperationListIterator struct {
	i    int
	page OperationListPage
}

// Next advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *OperationListIterator) Next() error {
	iter.i++
	if iter.i < len(iter.page.Values()) {
		return nil
	}
	err := iter.page.Next()
	if err != nil {
		iter.i--
		return err
	}
	iter.i = 0
	return nil
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter OperationListIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter OperationListIterator) Response() OperationList {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter OperationListIterator) Value() OperationResponse {
	if !iter.page.NotDone() {
		return OperationResponse{}
	}
	return iter.page.Values()[iter.i]
}

// IsEmpty returns true if the ListResult contains no values.
func (ol OperationList) IsEmpty() bool {
	return ol.Value == nil || len(*ol.Value) == 0
}

// operationListPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (ol OperationList) operationListPreparer() (*http.Request, error) {
	if ol.NextLink == nil || len(to.String(ol.NextLink)) < 1 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(ol.NextLink)))
}

// OperationListPage contains a page of OperationResponse values.
type OperationListPage struct {
	fn func(OperationList) (OperationList, error)
	ol OperationList
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *OperationListPage) Next() error {
	next, err := page.fn(page.ol)
	if err != nil {
		return err
	}
	page.ol = next
	return nil
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page OperationListPage) NotDone() bool {
	return !page.ol.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page OperationListPage) Response() OperationList {
	return page.ol
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page OperationListPage) Values() []OperationResponse {
	if page.ol.IsEmpty() {
		return nil
	}
	return *page.ol.Value
}

// OperationResponse this is the management partner operations response
type OperationResponse struct {
	// Name - this is the operation response name
	Name *string `json:"name,omitempty"`
	// Display - this is the operation display
	Display *OperationDisplay `json:"display,omitempty"`
	// Origin - the is operation response origin information
	Origin *string `json:"origin,omitempty"`
}

// PartnerProperties this is the management partner properties
type PartnerProperties struct {
	// PartnerID - This is the partner id
	PartnerID *string `json:"partnerId,omitempty"`
	// PartnerName - This is the partner name
	PartnerName *string `json:"partnerName,omitempty"`
	// TenantID - This is the tenant id.
	TenantID *string `json:"tenantId,omitempty"`
	// ObjectID - This is the object id.
	ObjectID *string `json:"objectId,omitempty"`
	// Version - This is the version.
	Version *string `json:"version,omitempty"`
	// UpdatedTime - This is the DateTime when the partner was updated.
	UpdatedTime *date.Time `json:"updatedTime,omitempty"`
	// CreatedTime - This is the DateTime when the partner was created.
	CreatedTime *date.Time `json:"createdTime,omitempty"`
	// State - This is the partner state. Possible values include: 'Active', 'Deleted'
	State State `json:"state,omitempty"`
}

// PartnerResponse this is the management partner operations response
type PartnerResponse struct {
	autorest.Response `json:"-"`
	// Etag - Type of the partner
	Etag *int32 `json:"etag,omitempty"`
	// ID - Identifier of the partner
	ID *string `json:"id,omitempty"`
	// Name - Name of the partner
	Name *string `json:"name,omitempty"`
	// PartnerProperties - Properties of the partner
	*PartnerProperties `json:"properties,omitempty"`
	// Type - Type of resource. "Microsoft.ManagementPartner/partners"
	Type *string `json:"type,omitempty"`
}

// MarshalJSON is the custom marshaler for PartnerResponse.
func (pr PartnerResponse) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if pr.Etag != nil {
		objectMap["etag"] = pr.Etag
	}
	if pr.ID != nil {
		objectMap["id"] = pr.ID
	}
	if pr.Name != nil {
		objectMap["name"] = pr.Name
	}
	if pr.PartnerProperties != nil {
		objectMap["properties"] = pr.PartnerProperties
	}
	if pr.Type != nil {
		objectMap["type"] = pr.Type
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for PartnerResponse struct.
func (pr *PartnerResponse) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "etag":
			if v != nil {
				var etag int32
				err = json.Unmarshal(*v, &etag)
				if err != nil {
					return err
				}
				pr.Etag = &etag
			}
		case "id":
			if v != nil {
				var ID string
				err = json.Unmarshal(*v, &ID)
				if err != nil {
					return err
				}
				pr.ID = &ID
			}
		case "name":
			if v != nil {
				var name string
				err = json.Unmarshal(*v, &name)
				if err != nil {
					return err
				}
				pr.Name = &name
			}
		case "properties":
			if v != nil {
				var partnerProperties PartnerProperties
				err = json.Unmarshal(*v, &partnerProperties)
				if err != nil {
					return err
				}
				pr.PartnerProperties = &partnerProperties
			}
		case "type":
			if v != nil {
				var typeVar string
				err = json.Unmarshal(*v, &typeVar)
				if err != nil {
					return err
				}
				pr.Type = &typeVar
			}
		}
	}

	return nil
}
