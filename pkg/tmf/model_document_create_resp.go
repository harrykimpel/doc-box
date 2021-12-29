/*
Doc Box API

<- The intent of this API is to provide a consistent/standardized mechanism to create new documents, delete existing documents, retrieving information about uploaded documents. This Specification is based on TMF667 - Document Management Release 17.0.1 from November 2017.

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tmf

import (
	"encoding/json"
)

// DocumentCreateResp This is a document resource allowing meta-data and the information of the document model.
type DocumentCreateResp struct {
	Links *Links `json:"_links,omitempty"`
	// Identifier of the document.
	Id *string `json:"id,omitempty"`
	// The lifecycleState of the document, such as Active.
	LifecycleState *string `json:"lifecycleState,omitempty"`
	// The date and time the document was created.
	CreationDate *string `json:"creationDate,omitempty"`
	// Name of the document type
	Type *string `json:"type,omitempty"`
	// A string used to give a name to the document
	Name *string `json:"name,omitempty"`
}

// NewDocumentCreateResp instantiates a new DocumentCreateResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDocumentCreateResp() *DocumentCreateResp {
	this := DocumentCreateResp{}
	return &this
}

// NewDocumentCreateRespWithDefaults instantiates a new DocumentCreateResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDocumentCreateRespWithDefaults() *DocumentCreateResp {
	this := DocumentCreateResp{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *DocumentCreateResp) GetLinks() Links {
	if o == nil || o.Links == nil {
		var ret Links
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentCreateResp) GetLinksOk() (*Links, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *DocumentCreateResp) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given Links and assigns it to the Links field.
func (o *DocumentCreateResp) SetLinks(v Links) {
	o.Links = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DocumentCreateResp) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentCreateResp) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DocumentCreateResp) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *DocumentCreateResp) SetId(v string) {
	o.Id = &v
}

// GetLifecycleState returns the LifecycleState field value if set, zero value otherwise.
func (o *DocumentCreateResp) GetLifecycleState() string {
	if o == nil || o.LifecycleState == nil {
		var ret string
		return ret
	}
	return *o.LifecycleState
}

// GetLifecycleStateOk returns a tuple with the LifecycleState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentCreateResp) GetLifecycleStateOk() (*string, bool) {
	if o == nil || o.LifecycleState == nil {
		return nil, false
	}
	return o.LifecycleState, true
}

// HasLifecycleState returns a boolean if a field has been set.
func (o *DocumentCreateResp) HasLifecycleState() bool {
	if o != nil && o.LifecycleState != nil {
		return true
	}

	return false
}

// SetLifecycleState gets a reference to the given string and assigns it to the LifecycleState field.
func (o *DocumentCreateResp) SetLifecycleState(v string) {
	o.LifecycleState = &v
}

// GetCreationDate returns the CreationDate field value if set, zero value otherwise.
func (o *DocumentCreateResp) GetCreationDate() string {
	if o == nil || o.CreationDate == nil {
		var ret string
		return ret
	}
	return *o.CreationDate
}

// GetCreationDateOk returns a tuple with the CreationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentCreateResp) GetCreationDateOk() (*string, bool) {
	if o == nil || o.CreationDate == nil {
		return nil, false
	}
	return o.CreationDate, true
}

// HasCreationDate returns a boolean if a field has been set.
func (o *DocumentCreateResp) HasCreationDate() bool {
	if o != nil && o.CreationDate != nil {
		return true
	}

	return false
}

// SetCreationDate gets a reference to the given string and assigns it to the CreationDate field.
func (o *DocumentCreateResp) SetCreationDate(v string) {
	o.CreationDate = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *DocumentCreateResp) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentCreateResp) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *DocumentCreateResp) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *DocumentCreateResp) SetType(v string) {
	o.Type = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DocumentCreateResp) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentCreateResp) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DocumentCreateResp) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DocumentCreateResp) SetName(v string) {
	o.Name = &v
}

func (o DocumentCreateResp) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Links != nil {
		toSerialize["_links"] = o.Links
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.LifecycleState != nil {
		toSerialize["lifecycleState"] = o.LifecycleState
	}
	if o.CreationDate != nil {
		toSerialize["creationDate"] = o.CreationDate
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableDocumentCreateResp struct {
	value *DocumentCreateResp
	isSet bool
}

func (v NullableDocumentCreateResp) Get() *DocumentCreateResp {
	return v.value
}

func (v *NullableDocumentCreateResp) Set(val *DocumentCreateResp) {
	v.value = val
	v.isSet = true
}

func (v NullableDocumentCreateResp) IsSet() bool {
	return v.isSet
}

func (v *NullableDocumentCreateResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDocumentCreateResp(val *DocumentCreateResp) *NullableDocumentCreateResp {
	return &NullableDocumentCreateResp{value: val, isSet: true}
}

func (v NullableDocumentCreateResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDocumentCreateResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
