/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package config

import (
	"time"
)

// HTTPVerb defines the restful api http verb
type HTTPVerb string

const (
	// MethodAny any method
	MethodAny HTTPVerb = "ANY"
	// MethodGet get
	MethodGet HTTPVerb = "GET"
	// MethodHead head
	MethodHead HTTPVerb = "HEAD"
	// MethodPost post
	MethodPost HTTPVerb = "POST"
	// MethodPut put
	MethodPut HTTPVerb = "PUT"
	// MethodPatch patch
	MethodPatch HTTPVerb = "PATCH" // RFC 5789
	// MethodDelete delete
	MethodDelete HTTPVerb = "DELETE"
	// MethodOptions options
	MethodOptions HTTPVerb = "OPTIONS"
)

// RequestType describes the type of the request. could be DUBBO/HTTP and others that we might implement in the future
type RequestType string

const (
	// DubboRequest represents the dubbo request
	DubboRequest RequestType = "dubbo"
	// HTTPRequest represents the http request
	HTTPRequest RequestType = "http"
)

// APIConfig defines the data structure of the api gateway configuration
type APIConfig struct {
	Name        string       `json:"name" yaml:"name"`
	Description string       `json:"description" yaml:"description"`
	Resources   []Resource   `json:"resources" yaml:"resources"`
	Definitions []Definition `json:"definitions" yaml:"definitions"`
}

// Resource defines the API path
type Resource struct {
	Type        string            `json:"type" yaml:"type"` // Restful, Dubbo
	Path        string            `json:"path" yaml:"path"`
	Timeout     time.Duration     `json:"timeout" yaml:"timeout"`
	Description string            `json:"description" yaml:"description"`
	Filters     []string          `json:"filters" yaml:"filters"`
	Methods     []Method          `json:"methods" yaml:"methods"`
	Resources   []Resource        `json:"resources,omitempty" yaml:"resources,omitempty"`
	Headers     map[string]string `json:"headers,omitempty" yaml:"headers,omitempty"`
}

// Method defines the method of the api
type Method struct {
	OnAir              bool          `json:"onAir" yaml:"onAir"` // true means the method is up and false means method is down
	Timeout            time.Duration `json:"timeout" yaml:"timeout"`
	Mock               bool          `json:"mock" yaml:"mock"`
	Filters            []string      `json:"filters" yaml:"filters"`
	HTTPVerb           `json:"httpVerb" yaml:"httpVerb"`
	InboundRequest     `json:"inboundRequest" yaml:"inboundRequest"`
	IntegrationRequest `json:"integrationRequest" yaml:"integrationRequest"`
}

// InboundRequest defines the details of the inbound
type InboundRequest struct {
	RequestType  `json:"requestType" yaml:"requestType"` //http, TO-DO: dubbo
	Headers      []Params                                `json:"headers" yaml:"headers"`
	QueryStrings []Params                                `json:"queryStrings" yaml:"queryStrings"`
	RequestBody  []BodyDefinition                        `json:"requestBody" yaml:"requestBody"`
}

// Params defines the simple parameter definition
type Params struct {
	Name     string `json:"name" yaml:"name"`
	Type     string `json:"type" yaml:"type"`
	Required bool   `json:"required" yaml:"required"`
}

// BodyDefinition connects the request body to the definitions
type BodyDefinition struct {
	DefinitionName string `json:"definitionName" yaml:"definitionName"`
}

// IntegrationRequest defines the backend request format and target
type IntegrationRequest struct {
	RequestType        `json:"requestType" yaml:"requestType"` // dubbo, TO-DO: http
	DubboBackendConfig `json:"dubboBackendConfig,inline,omitempty" yaml:"dubboBackendConfig,inline,omitempty"`
	HTTPBackendConfig  `json:"httpBackendConfig,inline,omitempty" yaml:"httpBackendConfig,inline,omitempty"`
	MappingParams      []MappingParam `json:"mappingParams,omitempty" yaml:"mappingParams,omitempty"`
}

// MappingParam defines the mapping rules of headers and queryStrings
type MappingParam struct {
	Name  string `json:"name,omitempty" yaml:"name"`
	MapTo string `json:"mapTo,omitempty" yaml:"mapTo"`
	// Opt some action.
	Opt Opt `json:"opt,omitempty" yaml:"opt,omitempty"`
}

// Opt option, action for compatibility.
type Opt struct {
	// Name match dubbo.DefaultMapOption key.
	Name string `json:"name,omitempty" yaml:"name"`
	// Open control opt create, only true will create a Opt.
	Open bool `json:"open,omitempty" yaml:"open"`
	// Usable setTarget condition, true can set, false not set.
	Usable bool `json:"usable,omitempty" yaml:"usable"`
}

// DubboBackendConfig defines the basic dubbo backend config
type DubboBackendConfig struct {
	ClusterName     string   `yaml:"clusterName" json:"clusterName"`
	ApplicationName string   `yaml:"applicationName" json:"applicationName"`
	Protocol        string   `yaml:"protocol" json:"protocol,omitempty" default:"dubbo"`
	Group           string   `yaml:"group" json:"group"`
	Version         string   `yaml:"version" json:"version"`
	Interface       string   `yaml:"interface" json:"interface"`
	Method          string   `yaml:"method" json:"method"`
	ParamTypes      []string `yaml:"paramTypes" json:"paramTypes"`
	ToParamTypes    []string `yaml:"toParamTypes" json:"toParamTypes"`
	Retries         string   `yaml:"retries" json:"retries,omitempty"`
}

// HTTPBackendConfig defines the basic dubbo backend config
type HTTPBackendConfig struct {
	URL string `yaml:"url" json:"url,omitempty"`
	// downstream host.
	Host string `yaml:"host" json:"host,omitempty"`
	// path to replace.
	Path string `yaml:"path" json:"path,omitempty"`
	// http protocol, http or https.
	Schema string `yaml:"schema" json:"scheme,omitempty"`
}

// Definition defines the complex json request body
type Definition struct {
	Name   string `json:"name" yaml:"name"`
	Schema string `json:"schema" yaml:"schema"` // use json schema
}
