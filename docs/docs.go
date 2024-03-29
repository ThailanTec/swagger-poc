// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "https://example.com/terms/",
        "contact": {},
        "license": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/user": {
            "post": {
                "description": "Cria um novo usuário com base nos dados fornecidos",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Criar um novo usuário",
                "operationId": "create-resource",
                "parameters": [
                    {
                        "description": "Dados do novo usuário",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.UserInput"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/model.UserOutput"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/model.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.ErrorResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer",
                    "example": 400
                },
                "message": {
                    "type": "string",
                    "example": "Bad Request, erro ao criar usuário."
                }
            }
        },
        "model.Status": {
            "type": "string",
            "enum": [
                "active",
                "inactive"
            ],
            "x-enum-varnames": [
                "Active",
                "Inactive"
            ]
        },
        "model.UserInput": {
            "type": "object",
            "required": [
                "email",
                "name"
            ],
            "properties": {
                "email": {
                    "type": "string",
                    "example": "thailan@gmail.com"
                },
                "name": {
                    "type": "string",
                    "example": "Thailan Santos"
                }
            }
        },
        "model.UserOutput": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string",
                    "example": "thailan@gmail.com"
                },
                "id": {
                    "type": "string",
                    "example": "c087a2ed-6694-49ad-884a-eecade5a2f90"
                },
                "name": {
                    "type": "string",
                    "example": "Thailan Santos"
                },
                "status": {
                    "allOf": [
                        {
                            "$ref": "#/definitions/model.Status"
                        }
                    ],
                    "example": "active"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "2.0",
	Host:             "localhost:8080",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "Sua API Swagger Title",
	Description:      "Esta é uma API Swagger para um exemplo de projeto Go com Gin.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
