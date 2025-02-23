package doc

import "github.com/swaggo/swag/v2"


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
                            "$ref": "#/definitions/dto.HttpResponse-dto_LoginResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dto.HttpError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/dto.HttpError"
                        }
                    }
                }
            }
        },
        "/api/v1/auth/logout": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
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
                            "$ref": "#/definitions/dto.HttpError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/dto.HttpError"
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
                            "$ref": "#/definitions/dto.HttpResponse-dto_TokenResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dto.HttpError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/dto.HttpError"
                        }
                    }
                }
            }
        },
        "/api/v1/auth/register": {
            "post": {
                "description": "Register",
                "tags": [
                    "auth"
                ],
                "summary": "Register",
                "parameters": [
                    {
                        "description": "request request",
                        "name": "RequestBody",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.RegisterRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.HttpResponse-dto_RegisterResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dto.HttpError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/dto.HttpError"
                        }
                    }
                }
            }
        },
        "/api/v1/gallery": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Create gallery by photographer",
                "tags": [
                    "gallery"
                ],
                "summary": "Create gallery",
                "parameters": [
                    {
                        "description": "Gallery details",
                        "name": "RequestBody",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.CreateGalleryRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dto.HttpError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/dto.HttpError"
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
                            "$ref": "#/definitions/dto.HttpResponse-dto_UserResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dto.HttpError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/dto.HttpError"
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
                "consumes": [
                    "multipart/form-data"
                ],
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
                            "$ref": "#/definitions/dto.HttpResponse-dto_UserResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dto.HttpError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/dto.HttpError"
                        }
                    }
                }
            }
        },
        "/api/v1/objects": {
            "post": {
                "description": "receive formData body, path (string, folder path, don't include \"..\" or prefix with \"/\") and file",
                "tags": [
                    "objects"
                ],
                "summary": "Upload image",
                "parameters": [
                    {
                        "type": "file",
                        "description": "picture (optional)",
                        "name": "file",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "folder enum (GALLERY, VERIFY_CITIZENCARD, PROFILE_IMAGE)",
                        "name": "folder",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.HttpResponse-dto_ObjectUploadResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dto.HttpError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/dto.HttpError"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete image",
                "tags": [
                    "objects"
                ],
                "summary": "Delete image",
                "parameters": [
                    {
                        "type": "string",
                        "description": "image url",
                        "name": "URL",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dto.HttpError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/dto.HttpError"
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
                            "$ref": "#/definitions/dto.HttpResponse-dto_CitizenCardResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dto.HttpError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/dto.HttpError"
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
                "consumes": [
                    "multipart/form-data"
                ],
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
                            "$ref": "#/definitions/dto.HttpResponse-dto_CitizenCardResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dto.HttpError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/dto.HttpError"
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
                "consumes": [
                    "multipart/form-data"
                ],
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
                            "$ref": "#/definitions/dto.HttpResponse-dto_CitizenCardResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dto.HttpError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/dto.HttpError"
                        }
                    }
                }
            }
        },
        "/api/v1/quotations": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Creates a new quotation for a customer and gallery",
                "tags": [
                    "quotation"
                ],
                "summary": "Create a quotation",
                "parameters": [
                    {
                        "description": "Quotation details",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.CreateQuotationRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/dto.HttpResponse-dto_CreateQuotationResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dto.HttpError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/dto.HttpError"
                        }
                    }
                }
            }
        },
        "/api/v1/quotations/{id}": {
            "patch": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Updates an existing quotation",
                "tags": [
                    "quotation"
                ],
                "summary": "Update a quotation",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Quotation ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Quotation update details",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.UpdateQuotationRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dto.HttpError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/dto.HttpError"
                        }
                    }
                }
            }
        },
        "/api/v1/quotations/{id}/accept": {
            "patch": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Acceot quotaion",
                "tags": [
                    "quotation"
                ],
                "summary": "Accept quotation",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "quotaion id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/dto.HttpError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/dto.HttpError"
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
        "dto.CreateGalleryRequest": {
            "type": "object",
            "required": [
                "description",
                "media",
                "name",
                "price"
            ],
            "properties": {
                "description": {
                    "type": "string"
                },
                "media": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dto.MediaGalleryRequest"
                    }
                },
                "name": {
                    "type": "string"
                },
                "price": {
                    "type": "number"
                }
            }
        },
        "dto.CreateQuotationRequest": {
            "type": "object",
            "required": [
                "customer_id",
                "from_date",
                "gallery_id",
                "price",
                "to_date"
            ],
            "properties": {
                "customer_id": {
                    "type": "integer"
                },
                "description": {
                    "type": "string"
                },
                "from_date": {
                    "type": "string"
                },
                "gallery_id": {
                    "type": "integer"
                },
                "price": {
                    "type": "number"
                },
                "to_date": {
                    "type": "string"
                }
            }
        },
        "dto.CreateQuotationResponse": {
            "type": "object",
            "properties": {
                "quotation_id": {
                    "type": "integer"
                }
            }
        },
        "dto.HttpError": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                }
            }
        },
        "dto.HttpResponse-dto_CitizenCardResponse": {
            "type": "object",
            "properties": {
                "result": {
                    "$ref": "#/definitions/dto.CitizenCardResponse"
                }
            }
        },
        "dto.HttpResponse-dto_CreateQuotationResponse": {
            "type": "object",
            "properties": {
                "result": {
                    "$ref": "#/definitions/dto.CreateQuotationResponse"
                }
            }
        },
        "dto.HttpResponse-dto_LoginResponse": {
            "type": "object",
            "properties": {
                "result": {
                    "$ref": "#/definitions/dto.LoginResponse"
                }
            }
        },
        "dto.HttpResponse-dto_ObjectUploadResponse": {
            "type": "object",
            "properties": {
                "result": {
                    "$ref": "#/definitions/dto.ObjectUploadResponse"
                }
            }
        },
        "dto.HttpResponse-dto_RegisterResponse": {
            "type": "object",
            "properties": {
                "result": {
                    "$ref": "#/definitions/dto.RegisterResponse"
                }
            }
        },
        "dto.HttpResponse-dto_TokenResponse": {
            "type": "object",
            "properties": {
                "result": {
                    "$ref": "#/definitions/dto.TokenResponse"
                }
            }
        },
        "dto.HttpResponse-dto_UserResponse": {
            "type": "object",
            "properties": {
                "result": {
                    "$ref": "#/definitions/dto.UserResponse"
                }
            }
        },
        "dto.LoginRequest": {
            "type": "object",
            "required": [
                "idToken",
                "provider"
            ],
            "properties": {
                "idToken": {
                    "type": "string"
                },
                "provider": {
                    "description": "GOOGLE",
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
        "dto.MediaGalleryRequest": {
            "type": "object",
            "required": [
                "pictureUrl"
            ],
            "properties": {
                "description": {
                    "type": "string"
                },
                "pictureUrl": {
                    "type": "string"
                }
            }
        },
        "dto.ObjectUploadResponse": {
            "type": "object",
            "properties": {
                "url": {
                    "type": "string"
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
        "dto.RegisterRequest": {
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
        "dto.RegisterResponse": {
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
        "dto.UpdateQuotationRequest": {
            "type": "object",
            "required": [
                "customer_id",
                "from_date",
                "gallery_id",
                "price",
                "status",
                "to_date"
            ],
            "properties": {
                "customer_id": {
                    "type": "integer"
                },
                "description": {
                    "type": "string"
                },
                "from_date": {
                    "type": "string"
                },
                "gallery_id": {
                    "type": "integer"
                },
                "price": {
                    "type": "number"
                },
                "status": {
                    "type": "string"
                },
                "to_date": {
                    "type": "string"
                }
            }
        },
        "dto.UserResponse": {
            "type": "object",
            "properties": {
                "accountNo": {
                    "type": "string"
                },
                "bank": {
                    "type": "string"
                },
                "bankBranch": {
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
                "phoneNumber": {
                    "type": "string"
                },
                "profilePictureUrl": {
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