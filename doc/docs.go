// Code generated by swaggo/swag. DO NOT EDIT.

package doc

import "github.com/swaggo/swag/v2"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "components": {"schemas":{"dto.CategoryResponse":{"properties":{"description":{"type":"string"},"id":{"type":"integer"},"name":{"type":"string"}},"type":"object"},"dto.CitizenCardResponse":{"properties":{"citizenId":{"type":"string"},"expireDate":{"type":"string"},"laserId":{"type":"string"},"picture":{"type":"string"}},"type":"object"},"dto.CreateCategoryRequest":{"properties":{"description":{"type":"string"},"name":{"type":"string"}},"required":["description","name"],"type":"object"},"dto.CreateMediaRequest":{"properties":{"description":{"type":"string"},"packageId":{"minimum":1,"type":"integer"},"pictureUrl":{"type":"string"}},"required":["packageId","pictureUrl"],"type":"object"},"dto.CreatePackageRequest":{"properties":{"description":{"type":"string"},"media":{"items":{"$ref":"#/components/schemas/dto.MediaPackageRequest"},"type":"array","uniqueItems":false},"name":{"type":"string"},"price":{"type":"number"}},"required":["description","media","name","price"],"type":"object"},"dto.DeleteMediaRequest":{"properties":{"mediaId":{"type":"integer"}},"type":"object"},"dto.HttpError":{"properties":{"error":{"type":"string"}},"type":"object"},"dto.HttpResponse-PaginationResponse[dto_CategoryResponse]":{"properties":{"result":{"$ref":"#/components/schemas/dto.PaginationResponse-dto_CategoryResponse"}},"type":"object"},"dto.HttpResponse-dto_CitizenCardResponse":{"properties":{"result":{"$ref":"#/components/schemas/dto.CitizenCardResponse"}},"type":"object"},"dto.HttpResponse-dto_LoginResponse":{"properties":{"result":{"$ref":"#/components/schemas/dto.LoginResponse"}},"type":"object"},"dto.HttpResponse-dto_ObjectUploadResponse":{"properties":{"result":{"$ref":"#/components/schemas/dto.ObjectUploadResponse"}},"type":"object"},"dto.HttpResponse-dto_QuotationResponse":{"properties":{"result":{"$ref":"#/components/schemas/dto.QuotationResponse"}},"type":"object"},"dto.HttpResponse-dto_RegisterResponse":{"properties":{"result":{"$ref":"#/components/schemas/dto.RegisterResponse"}},"type":"object"},"dto.HttpResponse-dto_TokenResponse":{"properties":{"result":{"$ref":"#/components/schemas/dto.TokenResponse"}},"type":"object"},"dto.HttpResponse-dto_UserResponse":{"properties":{"result":{"$ref":"#/components/schemas/dto.UserResponse"}},"type":"object"},"dto.LoginRequest":{"properties":{"idToken":{"type":"string"},"provider":{"description":"GOOGLE","type":"string"}},"required":["idToken","provider"],"type":"object"},"dto.LoginResponse":{"properties":{"accessToken":{"type":"string"},"exp":{"type":"integer"},"refreshToken":{"type":"string"},"user":{"$ref":"#/components/schemas/dto.UserResponse"}},"type":"object"},"dto.MediaPackageRequest":{"properties":{"description":{"type":"string"},"pictureUrl":{"type":"string"}},"required":["pictureUrl"],"type":"object"},"dto.MediaResponse":{"properties":{"id":{"type":"integer"},"pictureUrl":{"type":"string"}},"type":"object"},"dto.ObjectUploadResponse":{"properties":{"url":{"type":"string"}},"type":"object"},"dto.PackageResponse":{"properties":{"categories":{"items":{"$ref":"#/components/schemas/dto.CategoryResponse"},"type":"array","uniqueItems":false},"description":{"type":"string"},"id":{"type":"integer"},"media":{"items":{"$ref":"#/components/schemas/dto.MediaResponse"},"type":"array","uniqueItems":false},"name":{"type":"string"},"photographer":{"$ref":"#/components/schemas/dto.PhotographerResponse"},"price":{"type":"number"},"reviews":{"items":{"$ref":"#/components/schemas/dto.ReviewResponse"},"type":"array","uniqueItems":false},"tags":{"items":{"$ref":"#/components/schemas/dto.TagResponse"},"type":"array","uniqueItems":false}},"type":"object"},"dto.PaginationResponse-QuotationResponse":{"properties":{"data":{"items":{"$ref":"#/components/schemas/dto.QuotationResponse"},"type":"array","uniqueItems":false},"page":{"type":"integer"},"pageSize":{"type":"integer"},"totalPage":{"type":"integer"}},"type":"object"},"dto.PaginationResponse-dto_CategoryResponse":{"properties":{"data":{"items":{"$ref":"#/components/schemas/dto.CategoryResponse"},"type":"array","uniqueItems":false},"page":{"type":"integer"},"pageSize":{"type":"integer"},"totalPage":{"type":"integer"}},"type":"object"},"dto.PaginationResponse-dto_PackageResponse":{"properties":{"data":{"items":{"$ref":"#/components/schemas/dto.PackageResponse"},"type":"array","uniqueItems":false},"page":{"type":"integer"},"pageSize":{"type":"integer"},"totalPage":{"type":"integer"}},"type":"object"},"dto.PaginationResponse-dto_PhotographerResponse":{"properties":{"data":{"items":{"$ref":"#/components/schemas/dto.PhotographerResponse"},"type":"array","uniqueItems":false},"page":{"type":"integer"},"pageSize":{"type":"integer"},"totalPage":{"type":"integer"}},"type":"object"},"dto.PhotographerResponse":{"properties":{"activeStatus":{"type":"boolean"},"email":{"type":"string"},"id":{"type":"integer"},"isVerified":{"type":"boolean"},"name":{"type":"string"},"phoneNumber":{"type":"string"},"profilePictureUrl":{"type":"string"}},"type":"object"},"dto.QuotationResponse":{"properties":{"customer":{"$ref":"#/components/schemas/dto.UserResponse"},"description":{"type":"string"},"fromDate":{"type":"string"},"id":{"type":"integer"},"package":{"$ref":"#/components/schemas/dto.PackageResponse"},"photographer":{"$ref":"#/components/schemas/dto.PhotographerResponse"},"price":{"type":"number"},"status":{"$ref":"#/components/schemas/model.QuotationStatus"},"toDate":{"type":"string"}},"type":"object"},"dto.ReVerifyCitizenCardRequest":{"properties":{"citizenId":{"type":"string"},"expireDate":{"type":"string"},"imageUrl":{"type":"string"},"laserId":{"type":"string"}},"type":"object"},"dto.RefreshTokenRequest":{"properties":{"refreshToken":{"type":"string"}},"required":["refreshToken"],"type":"object"},"dto.RegisterRequest":{"properties":{"idToken":{"type":"string"},"provider":{"description":"GOOGLE","type":"string"},"role":{"description":"CUSTOMER, PHOTOGRAPHER, ADMIN","type":"string"}},"required":["idToken","provider","role"],"type":"object"},"dto.RegisterResponse":{"properties":{"accessToken":{"type":"string"},"exp":{"type":"integer"},"refreshToken":{"type":"string"},"user":{"$ref":"#/components/schemas/dto.UserResponse"}},"type":"object"},"dto.ReviewResponse":{"properties":{"comment":{"type":"string"},"customer":{"type":"string"},"id":{"type":"integer"},"rating":{"type":"number"}},"type":"object"},"dto.TagResponse":{"properties":{"id":{"type":"integer"},"name":{"type":"string"}},"type":"object"},"dto.TokenResponse":{"properties":{"accessToken":{"type":"string"},"exp":{"type":"integer"},"refreshToken":{"type":"string"}},"type":"object"},"dto.UpdateCategoryRequest":{"properties":{"description":{"type":"string"},"id":{"type":"integer"},"name":{"type":"string"}},"required":["id"],"type":"object"},"dto.UpdateMediaRequest":{"properties":{"description":{"type":"string"},"mediaId":{"minimum":1,"type":"integer"},"pictureUrl":{"type":"string"}},"required":["mediaId"],"type":"object"},"dto.UpdatePackageRequest":{"properties":{"description":{"minLength":0,"type":"string"},"name":{"type":"string"},"packageId":{"type":"integer"},"price":{"minimum":0,"type":"number"}},"type":"object"},"dto.UserResponse":{"properties":{"accountNo":{"type":"string"},"bank":{"type":"string"},"bankBranch":{"type":"string"},"email":{"type":"string"},"facebook":{"type":"string"},"id":{"type":"integer"},"instagram":{"type":"string"},"name":{"type":"string"},"phoneNumber":{"type":"string"},"profilePictureUrl":{"type":"string"},"role":{"$ref":"#/components/schemas/model.UserRole"}},"type":"object"},"dto.VerifyCitizenCardRequest":{"properties":{"citizenId":{"type":"string"},"expireDate":{"type":"string"},"imageUrl":{"type":"string"},"laserId":{"type":"string"}},"required":["citizenId","expireDate","imageUrl","laserId"],"type":"object"},"model.QuotationStatus":{"type":"string","x-enum-varnames":["QuotationPending","QuotationConfirm","QuotationCancelled","QuotationPaid"]},"model.UserRole":{"type":"string","x-enum-varnames":["UserRoleUnknown","UserRoleAdmin","UserRolePhotographer","UserRoleCustomer"]}},"securitySchemes":{"@securitydefinitions.apikey\tApiKeyAuth":{"in":"header","name":"Authorization","type":"apiKey"}}},
    "info": {"description":"{{escape .Description}}","title":"{{.Title}}","version":"{{.Version}}"},
    "externalDocs": {"description":"","url":""},
    "paths": {"/api/v1/admin/categories":{"post":{"description":"create category","requestBody":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/dto.CreateCategoryRequest"}}},"description":"request body","required":true},"responses":{"201":{"description":"Created"},"400":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/dto.HttpError"}}},"description":"Bad Request"},"401":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/dto.HttpError"}}},"description":"Unauthorized"},"500":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/dto.HttpError"}}},"description":"Internal Server Error"}},"security":[{"ApiKeyAuth":[]}],"summary":"create category","tags":["categories"]}},"/api/v1/admin/categories/{id}":{"delete":{"description":"delete category","parameters":[{"description":"category id","in":"path","name":"id","required":true,"schema":{"type":"integer"}}],"responses":{"204":{"description":"No Content"},"400":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/dto.HttpError"}}},"description":"Bad Request"},"401":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/dto.HttpError"}}},"description":"Unauthorized"},"500":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/dto.HttpError"}}},"description":"Internal Server Error"}},"security":[{"ApiKeyAuth":[]}],"summary":"delete category","tags":["categories"]},"patch":{"description":"update category","parameters":[{"description":"category id","in":"path","name":"id","required":true,"schema":{"type":"integer"}}],"requestBody":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/dto.UpdateCategoryRequest"}}},"description":"request body (don't need to include id)","required":true},"responses":{"204":{"description":"No Content"},"400":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/dto.HttpError"}}},"description":"Bad Request"},"401":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/dto.HttpError"}}},"description":"Unauthorized"},"500":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/dto.HttpError"}}},"description":"Internal Server Error"}},"security":[{"ApiKeyAuth":[]}],"summary":"update category","tags":["categories"]}},"/api/v1/auth/login":{"post":{"description":"Login","requestBody":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/dto.LoginRequest"}}},"description":"request request","required":true},"responses":{"200":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/dto.HttpResponse-dto_LoginResponse"}}},"description":"OK"},"400":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/dto.HttpError"}}},"description":"Bad Request"},"500":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/dto.HttpError"}}},"description":"Internal Server Error"}},"summary":"Login","tags":["auth"]}},"/api/v1/auth/logout":{"post":{"description":"Logout","responses":{"204":{"description":"No Content"},"400":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/dto.HttpError"}}},"description":"Bad Request"},"500":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/dto.HttpError"}}},"description":"Internal Server Error"}},"security":[{"ApiKeyAuth":[]}],"summary":"Logout","tags":["auth"]}},"/api/v1/auth/refresh-token":{"post":{"description":"Refresh Token","requestBody":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/dto.RefreshTokenRequest"}}},"description":"request request","required":true},"responses":{"200":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/dto.HttpResponse-dto_TokenResponse"}}},"description":"OK"},"400":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/dto.HttpError"}}},"description":"Bad Request"},"500":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/dto.HttpError"}}},"description":"Internal Server Error"}},"summary":"Refresh Token","tags":["auth"]}},"/api/v1/auth/register":{"post":{"description":"Register","requestBody":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/dto.RegisterRequest"}}},"description":"request request","required":true},"responses":{"200":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/dto.HttpResponse-dto_RegisterResponse"}}},"description":"OK"},"400":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/dto.HttpError"}}},"description":"Bad Request"},"500":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/dto.HttpError"}}},"description":"Internal Server Error"}},"summary":"Register","tags":["auth"]}},"/api/v1/categories":{"get":{"description":"list category","parameters":[{"description":"Page number for pagination (default: 1)","in":"query","name":"page","schema":{"type":"integer"}},{"description":"Number of records per page (default: 20)","in":"query","name":"pageSize","schema":{"type":"integer"}}],"responses":{"200":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/dto.HttpResponse-PaginationResponse[dto_CategoryResponse]"}}},"description":"OK"},"500":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/dto.HttpError"}}},"description":"Internal Server Error"}},"summary":"list category","tags":["categories"]}},"/api/v1/customer/quotations/{id}/cancel":{"patch":{"description":"cancelled quotaion","parameters":[{"description":"quotaion id","in":"path","name":"id","required":true,"schema":{"type":"integer"}}],"responses":{"204":{"description":"No Content"},"401":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/dto.HttpError"}}},"description":"Unauthorized"},"403":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/dto.HttpError"}}},"description":"Forbidden"},"404":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/dto.HttpError"}}},"description":"Not Found"},"500":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/dto.HttpError"}}},"description":"Internal Server Error"}},"security":[{"ApiKeyAuth":[]}],"summary":"cancelled quotation","tags":["quotations"]}},"/api/v1/customer/quotations/{id}/confirm":{"patch":{"description":"confirm quotaion","parameters":[{"description":"quotaion id","in":"path","name":"id","required":true,"schema":{"type":"integer"}}],"responses":{"204":{"description":"No Content"},"401":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/dto.HttpError"}}},"description":"Unauthorized"},"403":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/dto.HttpError"}}},"description":"Forbidden"},"404":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/dto.HttpError"}}},"description":"Not Found"},"500":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/dto.HttpError"}}},"description":"Internal Server Error"}},"security":[{"ApiKeyAuth":[]}],"summary":"confirm quotation","tags":["quotations"]}},"/api/v1/me":{"get":{"description":"Get me","responses":{"200":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/dto.HttpResponse-dto_UserResponse"}}},"description":"OK"},"400":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/dto.HttpError"}}},"description":"Bad Request"},"500":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/dto.HttpError"}}},"description":"Internal Server Error"}},"security":[{"ApiKeyAuth":[]}],"summary":"Get me","tags":["user"]},"patch":{"description":"Update user's profile","requestBody":{"content":{"multipart/form-data":{"schema":{"type":"string"}}},"description":"Bank Branch"},"responses":{"200":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/dto.HttpResponse-dto_UserResponse"}}},"description":"OK"},"400":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/dto.HttpError"}}},"description":"Bad Request"},"500":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/dto.HttpError"}}},"description":"Internal Server Error"}},"security":[{"ApiKeyAuth":[]}],"summary":"Update me","tags":["user"]}},"/api/v1/objects":{"delete":{"description":"Delete image","parameters":[{"description":"image url","in":"query","name":"URL","required":true,"schema":{"type":"string"}}],"responses":{"204":{"description":"No Content"},"400":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/dto.HttpError"}}},"description":"Bad Request"},"500":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/dto.HttpError"}}},"description":"Internal Server Error"}},"summary":"Delete image","tags":["objects"]},"post":{"description":"receive formData body, path (string, folder path, don't include \"..\" or prefix with \"/\") and file","requestBody":{"content":{"application/x-www-form-urlencoded":{"schema":{"type":"string"}}},"description":"folder enum (PACKAGE, VERIFY_CITIZENCARD, PROFILE_IMAGE)"},"responses":{"200":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/dto.HttpResponse-dto_ObjectUploadResponse"}}},"description":"OK"},"400":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/dto.HttpError"}}},"description":"Bad Request"},"500":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/dto.HttpError"}}},"description":"Internal Server Error"}},"summary":"Upload image","tags":["objects"]}},"/api/v1/packages":{"get":{"description":"Show all available packages with pagination","parameters":[{"description":"Page number","in":"query","name":"page","schema":{"type":"integer"}},{"description":"Page size","in":"query","name":"pageSize","schema":{"type":"integer"}},{"description":"Minimum price","in":"query","name":"minPrice","schema":{"type":"number"}},{"description":"Maximum price","in":"query","name":"maxPrice","schema":{"type":"number"}},{"description":"Photographer ID","in":"query","name":"photographerId","schema":{"type":"integer"}}],"responses":{"200":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/dto.PaginationResponse-dto_PackageResponse"}}},"description":"OK"},"400":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/dto.HttpError"}}},"description":"Bad Request"},"500":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/dto.HttpError"}}},"description":"Internal Server Error"}},"summary":"Get all packages","tags":["packages"]}},"/api/v1/photographer/citizen-card":{"get":{"description":"Get Photographer Citizen Card","responses":{"200":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/dto.HttpResponse-dto_CitizenCardResponse"}}},"description":"OK"},"400":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/dto.HttpError"}}},"description":"Bad Request"},"500":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/dto.HttpError"}}},"description":"Internal Server Error"}},"security":[{"ApiKeyAuth":[]}],"summary":"Get Citizen Card","tags":["citizencard"]}},"/api/v1/photographer/citizen-card/reverify":{"patch":{"description":"Reverify Photographer Citizen Card","requestBody":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/dto.ReVerifyCitizenCardRequest"}}},"description":"request request","required":true},"responses":{"200":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/dto.HttpResponse-dto_CitizenCardResponse"}}},"description":"OK"},"400":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/dto.HttpError"}}},"description":"Bad Request"},"500":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/dto.HttpError"}}},"description":"Internal Server Error"}},"security":[{"ApiKeyAuth":[]}],"summary":"Reverify Citizen Card","tags":["citizencard"]}},"/api/v1/photographer/citizen-card/verify":{"post":{"description":"Verify Photographer Citizen Card","requestBody":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/dto.VerifyCitizenCardRequest"}}},"description":"request request","required":true},"responses":{"200":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/dto.HttpResponse-dto_CitizenCardResponse"}}},"description":"OK"},"400":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/dto.HttpError"}}},"description":"Bad Request"},"500":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/dto.HttpError"}}},"description":"Internal Server Error"}},"security":[{"ApiKeyAuth":[]}],"summary":"Verify Citizen Card","tags":["citizencard"]}},"/api/v1/photographer/media":{"post":{"description":"Create media by photographer","requestBody":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/dto.CreateMediaRequest"}}},"description":"Media details","required":true},"responses":{"201":{"description":"Created"},"400":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/dto.HttpError"}}},"description":"Bad Request"},"500":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/dto.HttpError"}}},"description":"Internal Server Error"}},"security":[{"ApiKeyAuth":[]}],"summary":"Create Media","tags":["media"]}},"/api/v1/photographer/media/{mediaId}":{"delete":{"description":"Delete media","requestBody":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/dto.DeleteMediaRequest"}}},"description":"Media details","required":true},"responses":{"204":{"description":"No Content"},"400":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/dto.HttpError"}}},"description":"Bad Request"},"500":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/dto.HttpError"}}},"description":"Internal Server Error"}},"security":[{"ApiKeyAuth":[]}],"summary":"Delete Media","tags":["media"]},"patch":{"description":"Update media","requestBody":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/dto.UpdateMediaRequest"}}},"description":"Media details","required":true},"responses":{"204":{"description":"No Content"},"400":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/dto.HttpError"}}},"description":"Bad Request"},"500":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/dto.HttpError"}}},"description":"Internal Server Error"}},"security":[{"ApiKeyAuth":[]}],"summary":"Update Media","tags":["media"]}},"/api/v1/photographer/packages":{"post":{"description":"Create Package by photographer","requestBody":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/dto.CreatePackageRequest"}}},"description":"Package details","required":true},"responses":{"201":{"description":"Created"},"400":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/dto.HttpError"}}},"description":"Bad Request"},"500":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/dto.HttpError"}}},"description":"Internal Server Error"}},"security":[{"ApiKeyAuth":[]}],"summary":"Create Package","tags":["packages"]}},"/api/v1/photographer/packages/{packageId}":{"patch":{"description":"Update","requestBody":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/dto.UpdatePackageRequest"}}},"description":"Package details","required":true},"responses":{"204":{"description":"No Content"},"400":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/dto.HttpError"}}},"description":"Bad Request"},"500":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/dto.HttpError"}}},"description":"Internal Server Error"}},"security":[{"ApiKeyAuth":[]}],"summary":"Update package","tags":["packages"]}},"/api/v1/photographers":{"get":{"description":"Retrieve a paginated list of photographers, optionally filtered by name.","parameters":[{"description":"Page number for pagination (default: 1)","in":"query","name":"page","schema":{"type":"integer"}},{"description":"Number of records per page (default: 5, max: 20)","in":"query","name":"pageSize","schema":{"type":"integer"}},{"description":"Filter by photographer's name (case-insensitive)","in":"query","name":"name","schema":{"type":"string"}}],"responses":{"200":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/dto.PaginationResponse-dto_PhotographerResponse"}}},"description":"OK"},"400":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/dto.HttpError"}}},"description":"Bad Request"},"500":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/dto.HttpError"}}},"description":"Internal Server Error"}},"summary":"Get All Photographers","tags":["photographers"]}},"/api/v1/quotations":{"get":{"description":"list quotations","parameters":[{"description":"Page number","in":"query","name":"page","schema":{"type":"integer"}},{"description":"Page size","in":"query","name":"page_size","schema":{"type":"integer"}}],"responses":{"200":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/dto.PaginationResponse-QuotationResponse"}}},"description":"OK"},"401":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/dto.PaginationResponse-QuotationResponse"}}},"description":"Unauthorized"},"500":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/dto.HttpError"}}},"description":"Internal Server Error"}},"security":[{"ApiKeyAuth":[]}],"summary":"list quotations","tags":["quotations"]}},"/api/v1/quotations/{id}":{"get":{"description":"Get Quotation By ID","responses":{"200":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/dto.HttpResponse-dto_QuotationResponse"}}},"description":"OK"},"400":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/dto.HttpError"}}},"description":"Bad Request"},"401":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/dto.HttpError"}}},"description":"Unauthorized"},"403":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/dto.HttpError"}}},"description":"Forbidden"},"500":{"content":{"application/json":{"schema":{"$ref":"#/components/schemas/dto.HttpError"}}},"description":"Internal Server Error"}},"security":[{"ApiKeyAuth":[]}],"summary":"Get Quotation By ID","tags":["quotations"]}}},
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
