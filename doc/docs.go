// Package doc Code generated by swaggo/swag. DO NOT EDIT
package doc

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
        "/api/v1/auth/login": {
            "post": {
                "description": "Login",
                "tags": [
                    "auth"
                ],
                "summary": "Login",
                "parameters": [
                    {
                        "description": "request request",
                        "name": "RequestBody",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.LoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/dto.HttpResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "result": {
                                            "$ref": "#/definitions/dto.LoginResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dto.HttpResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/dto.HttpResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/auth/logout": {
            "post": {
                "description": "Logout",
                "tags": [
                    "auth"
                ],
                "summary": "Logout",
                "responses": {
                    "204": {
                        "description": "No Content"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dto.HttpResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/dto.HttpResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/auth/refresh-token": {
            "post": {
                "description": "Refresh Token",
                "tags": [
                    "auth"
                ],
                "summary": "Refresh Token",
                "parameters": [
                    {
                        "description": "request request",
                        "name": "RequestBody",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.RefreshTokenRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/dto.HttpResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "result": {
                                            "$ref": "#/definitions/dto.TokenResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dto.HttpResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/dto.HttpResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/me": {
            "get": {
                "description": "Get me",
                "tags": [
                    "user"
                ],
                "summary": "Get me",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/dto.HttpResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "result": {
                                            "$ref": "#/definitions/dto.UserResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dto.HttpResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/dto.HttpResponse"
                        }
                    }
                }
            },
            "patch": {
                "description": "Update user's profile, Send only the fields that need to be changed.",
                "consumes": [
                    "multipart/form-data"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Update user's profile",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User's name",
                        "name": "name",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "User's phone number",
                        "name": "phone_number",
                        "in": "formData"
                    },
                    {
                        "type": "file",
                        "description": "User's profile picture",
                        "name": "profile_picture",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "User's Facebook link",
                        "name": "facebook",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "User's Instagram link",
                        "name": "instagram",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "User's bank name",
                        "name": "bank",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "User's bank account number",
                        "name": "account_no",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "User's bank branch",
                        "name": "bank_branch",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/dto.HttpResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "result": {
                                            "$ref": "#/definitions/dto.UserResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dto.HttpResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/dto.HttpResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/photographer/citizen-card": {
            "get": {
                "description": "Get Photographer Citizen Card",
                "tags": [
                    "photographer"
                ],
                "summary": "Get Citizen Card",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/dto.HttpResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "result": {
                                            "$ref": "#/definitions/dto.CitizenCardResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dto.HttpResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/dto.HttpResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/photographer/reverify": {
            "patch": {
                "description": "Allows photographers to update their citizen card details. Send only the fields that need to be changed.",
                "consumes": [
                    "multipart/form-data"
                ],
                "tags": [
                    "photographer"
                ],
                "summary": "Reverify Citizen Card",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Citizen ID",
                        "name": "citizenId",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "Laser ID",
                        "name": "laserId",
                        "in": "formData"
                    },
                    {
                        "type": "file",
                        "description": "Citizen card picture",
                        "name": "picture",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "Expire date (YYYY-MM-DD)",
                        "name": "expireDate",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/dto.HttpResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "result": {
                                            "$ref": "#/definitions/dto.CitizenCardResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dto.HttpResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/dto.HttpResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/photographer/verify": {
            "post": {
                "description": "Verify Photographer Citizen Card",
                "consumes": [
                    "multipart/form-data"
                ],
                "tags": [
                    "photographer"
                ],
                "summary": "Verify Citizen Card",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Citizen ID",
                        "name": "citizenId",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Laser ID",
                        "name": "laserId",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "file",
                        "description": "Citizen card picture",
                        "name": "picture",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Expire date (YYYY-MM-DD)",
                        "name": "expireDate",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/dto.HttpResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "result": {
                                            "$ref": "#/definitions/dto.CitizenCardResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dto.HttpResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/dto.HttpResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.CitizenCardResponse": {
            "type": "object",
            "properties": {
                "citizenId": {
                    "type": "string"
                },
                "expireDate": {
                    "type": "string"
                },
                "laserId": {
                    "type": "string"
                },
                "picture_url": {
                    "type": "string"
                }
            }
        },
        "dto.HttpResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                },
                "result": {}
            }
        },
        "dto.LoginRequest": {
            "type": "object",
            "required": [
                "idToken",
                "provider",
                "role"
            ],
            "properties": {
                "idToken": {
                    "type": "string"
                },
                "provider": {
                    "description": "GOOGLE",
                    "type": "string"
                },
                "role": {
                    "description": "CUSTOMER, PHOTOGRAPHER, ADMIN",
                    "type": "string"
                }
            }
        },
        "dto.LoginResponse": {
            "type": "object",
            "properties": {
                "accessToken": {
                    "type": "string"
                },
                "exp": {
                    "type": "integer"
                },
                "refreshToken": {
                    "type": "string"
                },
                "user": {
                    "$ref": "#/definitions/dto.UserResponse"
                }
            }
        },
        "dto.RefreshTokenRequest": {
            "type": "object",
            "required": [
                "refreshToken"
            ],
            "properties": {
                "refreshToken": {
                    "type": "string"
                }
            }
        },
        "dto.TokenResponse": {
            "type": "object",
            "properties": {
                "accessToken": {
                    "type": "string"
                },
                "exp": {
                    "type": "integer"
                },
                "refreshToken": {
                    "type": "string"
                }
            }
        },
        "dto.UserResponse": {
            "type": "object",
            "properties": {
                "account_no": {
                    "type": "string"
                },
                "bank": {
                    "type": "string"
                },
                "bank_branch": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "facebook": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "instagram": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "phone_number": {
                    "type": "string"
                },
                "profile_picture_url": {
                    "type": "string"
                },
                "role": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "Bearer": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{"https", "http"},
	Title:            "pic-me-pls API",
	Description:      "pic-me-pls API documentation",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
