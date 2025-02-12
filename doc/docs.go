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
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
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
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Update user's profile",
                "tags": [
                    "user"
                ],
                "summary": "Update me",
                "parameters": [
                    {
                        "type": "file",
                        "description": "Profile picture (optional)",
                        "name": "profilePicture",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "Name",
                        "name": "name",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "Phone Number",
                        "name": "phoneNumber",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "Facebook",
                        "name": "facebook",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "Instagram",
                        "name": "instagram",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "Bank",
                        "name": "bank",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "Account No",
                        "name": "accountNo",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "Bank Branch",
                        "name": "bankBranch",
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
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
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
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Reverify Photographer Citizen Card",
                "tags": [
                    "photographer"
                ],
                "summary": "Reverify Citizen Card",
                "parameters": [
                    {
                        "type": "file",
                        "description": "Card picture (optional)",
                        "name": "cardPicture",
                        "in": "formData"
                    },
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
                        "type": "string",
                        "description": "Expire Date",
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
        },
        "/api/v1/photographer/verify": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Verify Photographer Citizen Card",
                "tags": [
                    "photographer"
                ],
                "summary": "Verify Citizen Card",
                "parameters": [
                    {
                        "type": "file",
                        "description": "Card picture (optional)",
                        "name": "cardPicture",
                        "in": "formData"
                    },
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
                        "type": "string",
                        "description": "Expire Date",
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
                "picture": {
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
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "0.1",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Pic Me Pls API",
	Description:      "Pic Me Pls API Documentation",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
