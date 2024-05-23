// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/delete": {
            "delete": {
                "description": "Delete a record from Database",
                "produces": [
                    "application/json"
                ],
                "summary": "Delete a record",
                "parameters": [
                    {
                        "type": "string",
                        "example": "patients",
                        "description": "The collection name",
                        "name": "collection",
                        "in": "query",
                        "required": true
                    },
                    {
                        "description": "Delete Request Body",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/app.requestID"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/record": {
            "put": {
                "description": "The endpoint inserts sent data to the database",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Insert record to DB",
                "parameters": [
                    {
                        "type": "string",
                        "example": "patients",
                        "description": "The collection name",
                        "name": "collection",
                        "in": "query",
                        "required": true
                    },
                    {
                        "description": "Insert Request Body",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/app.recordReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/records": {
            "get": {
                "description": "Get all records from selected collection in database",
                "produces": [
                    "application/json"
                ],
                "summary": "Get all records",
                "parameters": [
                    {
                        "type": "string",
                        "example": "patients",
                        "description": "The collection name",
                        "name": "collection",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/update": {
            "post": {
                "description": "Update a record from Database",
                "produces": [
                    "application/json"
                ],
                "summary": "Update a record",
                "parameters": [
                    {
                        "type": "string",
                        "example": "patients",
                        "description": "The collection name",
                        "name": "collection",
                        "in": "query",
                        "required": true
                    },
                    {
                        "description": "Delete Request Body",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/app.recordReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "app.recordReq": {
            "type": "object",
            "properties": {
                "dateOfBirth": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "patientName": {
                    "type": "string"
                }
            }
        },
        "app.requestID": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
