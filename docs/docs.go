// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/profile/add": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "profile"
                ],
                "summary": "New",
                "parameters": [
                    {
                        "description": "profile",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/repositories.Profile"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.DefaultResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.DefaultResponse"
                        }
                    }
                }
            }
        },
        "/profile/getById/{id}": {
            "post": {
                "description": "Getting by an profile.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "profile"
                ],
                "summary": "GetById.",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID of the profile",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "api response",
                        "schema": {
                            "$ref": "#/definitions/internal_http-server_handlers_profile.RespProfile"
                        }
                    },
                    "400": {
                        "description": "invalid request",
                        "schema": {
                            "$ref": "#/definitions/response.DefaultResponse"
                        }
                    },
                    "404": {
                        "description": "profile not found",
                        "schema": {
                            "$ref": "#/definitions/response.DefaultResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "github.com_VeryHappy2_GoApi_internal_http-server_handlers_profile.Profile": {
            "type": "object"
        },
        "github.com_VeryHappy2_GoApi_internal_http-server_handlers_profile.RespProfile": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/github.com_VeryHappy2_GoApi_internal_http-server_handlers_profile.Profile"
                },
                "error": {
                    "type": "string"
                },
                "status": {
                    "type": "string",
                    "example": "success"
                }
            }
        },
        "internal_http-server_handlers_profile.Profile": {
            "type": "object"
        },
        "internal_http-server_handlers_profile.RespProfile": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/internal_http-server_handlers_profile.Profile"
                },
                "error": {
                    "type": "string"
                },
                "status": {
                    "type": "string",
                    "example": "success"
                }
            }
        },
        "repositories.Profile": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "last_name": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "response.DefaultResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "integer"
                },
                "error": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "",
	Description: "",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
