// Code generated by swaggo/swag. DO NOT EDIT.

package doc

import "github.com/swaggo/swag/v2"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "components": {"schemas":{"dto.CitizenCardResponse":{"properties":{"citizenId":{"type":"string"},"expireDate":{"type":"string"},"laserId":{"type":"string"},"picture":{"type":"string"}},"type":"object"},"dto.HttpError":{"properties":{"error":{"type":"string"}},"type":"object"},"dto.HttpResponse-dto_CitizenCardResponse":{"properties":{"result":{"$ref":"#/components/schemas/dto.CitizenCardResponse"}},"type":"object"},"dto.HttpResponse-dto_LoginResponse":{"properties":{"result":{"$ref":"#/components/schemas/dto.LoginResponse"}},"type":"object"},"dto.HttpResponse-dto_RegisterResponse":{"properties":{"result":{"$ref":"#/components/schemas/dto.RegisterResponse"}},"type":"object"},"dto.HttpResponse-dto_TokenResponse":{"properties":{"result":{"$ref":"#/components/schemas/dto.TokenResponse"}},"type":"object"},"dto.HttpResponse-dto_UserResponse":{"properties":{"result":{"$ref":"#/components/schemas/dto.UserResponse"}},"type":"object"},"dto.LoginRequest":{"properties":{"idToken":{"type":"string"},"provider":{"description":"GOOGLE","type":"string"}},"required":["idToken","provider"],"type":"object"},"dto.LoginResponse":{"properties":{"accessToken":{"type":"string"},"exp":{"type":"integer"},"refreshToken":{"type":"string"},"user":{"$ref":"#/components/schemas/dto.UserResponse"}},"type":"object"},"dto.RefreshTokenRequest":{"properties":{"refreshToken":{"type":"string"}},"required":["refreshToken"],"type":"object"},"dto.RegisterRequest":{"properties":{"idToken":{"type":"string"},"provider":{"description":"GOOGLE","type":"string"},"role":{"description":"CUSTOMER, PHOTOGRAPHER, ADMIN","type":"string"}},"required":["idToken","provider","role"],"type":"object"},"dto.RegisterResponse":{"properties":{"accessToken":{"type":"string"},"exp":{"type":"integer"},"refreshToken":{"type":"string"},"user":{"$ref":"#/components/schemas/dto.UserResponse"}},"type":"object"},"dto.TokenResponse":{"properties":{"accessToken":{"type":"string"},"exp":{"type":"integer"},"refreshToken":{"type":"string"}},"type":"object"},"dto.UserResponse":{"properties":{"accountNo":{"type":"string"},"bank":{"type":"string"},"bankBranch":{"type":"string"},"email":{"type":"string"},"facebook":{"type":"string"},"id":{"type":"integer"},"instagram":{"type":"string"},"name":{"type":"string"},"phoneNumber":{"type":"string"},"profilePictureUrl":{"type":"string"},"role":{"type":"string"}},"type":"object"}},"securitySchemes":{"@securitydefinitions.apikey\tApiKeyAuth":{"in":"header","name":"Authorization","type":"apiKey"}}},
    "info": {"description":"{{escape .Description}}","title":"{{.Title}}","version":"{{.Version}}"},
    "externalDocs": {"description":"","url":""},
    "paths": {"/api/v1/auth/login":{"post":{"description":"Login","requestBody":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/dto.LoginRequest"}}},"description":"request request","required":true},"responses":{"200":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/dto.HttpResponse-dto_LoginResponse"}}},"description":"OK"},"400":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/dto.HttpError"}}},"description":"Bad Request"},"500":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/dto.HttpError"}}},"description":"Internal Server Error"}},"summary":"Login","tags":["auth"]}},"/api/v1/auth/logout":{"post":{"description":"Logout","responses":{"204":{"description":"No Content"},"400":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/dto.HttpError"}}},"description":"Bad Request"},"500":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/dto.HttpError"}}},"description":"Internal Server Error"}},"security":[{"ApiKeyAuth":[]}],"summary":"Logout","tags":["auth"]}},"/api/v1/auth/refresh-token":{"post":{"description":"Refresh Token","requestBody":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/dto.RefreshTokenRequest"}}},"description":"request request","required":true},"responses":{"200":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/dto.HttpResponse-dto_TokenResponse"}}},"description":"OK"},"400":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/dto.HttpError"}}},"description":"Bad Request"},"500":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/dto.HttpError"}}},"description":"Internal Server Error"}},"summary":"Refresh Token","tags":["auth"]}},"/api/v1/auth/register":{"post":{"description":"register","requestBody":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/dto.RegisterRequest"}}},"description":"request request","required":true},"responses":{"200":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/dto.HttpResponse-dto_RegisterResponse"}}},"description":"OK"},"400":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/dto.HttpError"}}},"description":"Bad Request"},"500":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/dto.HttpError"}}},"description":"Internal Server Error"}},"summary":"register","tags":["auth"]}},"/api/v1/me":{"get":{"description":"Get me","responses":{"200":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/dto.HttpResponse-dto_UserResponse"}}},"description":"OK"},"400":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/dto.HttpError"}}},"description":"Bad Request"},"500":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/dto.HttpError"}}},"description":"Internal Server Error"}},"security":[{"ApiKeyAuth":[]}],"summary":"Get me","tags":["user"]},"patch":{"description":"Update user's profile","requestBody":{"content":{"application/x-www-form-urlencoded":{"schema":{"type":"string"}}},"description":"Bank Branch"},"responses":{"200":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/dto.HttpResponse-dto_UserResponse"}}},"description":"OK"},"400":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/dto.HttpError"}}},"description":"Bad Request"},"500":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/dto.HttpError"}}},"description":"Internal Server Error"}},"security":[{"ApiKeyAuth":[]}],"summary":"Update me","tags":["user"]}},"/api/v1/photographer/citizen-card":{"get":{"description":"Get Photographer Citizen Card","responses":{"200":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/dto.HttpResponse-dto_CitizenCardResponse"}}},"description":"OK"},"400":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/dto.HttpError"}}},"description":"Bad Request"},"500":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/dto.HttpError"}}},"description":"Internal Server Error"}},"security":[{"ApiKeyAuth":[]}],"summary":"Get Citizen Card","tags":["photographer"]}},"/api/v1/photographer/reverify":{"patch":{"description":"Reverify Photographer Citizen Card","requestBody":{"content":{"application/x-www-form-urlencoded":{"schema":{"type":"string"}}},"description":"Expire Date","required":true},"responses":{"200":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/dto.HttpResponse-dto_CitizenCardResponse"}}},"description":"OK"},"400":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/dto.HttpError"}}},"description":"Bad Request"},"500":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/dto.HttpError"}}},"description":"Internal Server Error"}},"security":[{"ApiKeyAuth":[]}],"summary":"Reverify Citizen Card","tags":["photographer"]}},"/api/v1/photographer/verify":{"post":{"description":"Verify Photographer Citizen Card","requestBody":{"content":{"application/x-www-form-urlencoded":{"schema":{"type":"string"}}},"description":"Expire Date","required":true},"responses":{"200":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/dto.HttpResponse-dto_CitizenCardResponse"}}},"description":"OK"},"400":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/dto.HttpError"}}},"description":"Bad Request"},"500":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/dto.HttpError"}}},"description":"Internal Server Error"}},"security":[{"ApiKeyAuth":[]}],"summary":"Verify Citizen Card","tags":["photographer"]}}},
    "openapi": "3.1.0"
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
