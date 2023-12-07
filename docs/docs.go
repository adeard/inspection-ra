// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "ICT INDOAGRI",
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/v1/attachment": {
            "post": {
                "description": "Upload File",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Attachment"
                ],
                "summary": "Upload File",
                "parameters": [
                    {
                        "type": "string",
                        "default": "Bearer \u003cAdd access token here\u003e",
                        "description": "Insert your access token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "file",
                        "description": " AttachmentRequest Schema ",
                        "name": "file",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": " AttachmentRequest Schema ",
                        "name": "no_inspec",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": " AttachmentRequest Schema ",
                        "name": "image_category",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.AttachmentResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/auth/logged": {
            "get": {
                "description": "Get Logged User",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Get Logged User",
                "parameters": [
                    {
                        "type": "string",
                        "default": "Bearer \u003cAdd access token here\u003e",
                        "description": "Insert your access token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/domain.AuthResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/domain.AuthLoggedData"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/api/v1/auth/sign_in": {
            "post": {
                "description": "Sign In",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Sign In",
                "parameters": [
                    {
                        "description": " AuthRequest Schema ",
                        "name": "AuthRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.AuthRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/domain.AuthResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/domain.AuthData"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/api/v1/mob": {
            "get": {
                "description": "Get Mob",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Mob"
                ],
                "summary": "Get Mob",
                "parameters": [
                    {
                        "type": "string",
                        "default": "Bearer \u003cAdd access token here\u003e",
                        "description": "Insert your access token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/domain.MobResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/domain.MobData"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/api/v1/mob/batch": {
            "post": {
                "description": "Insert Inspection",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Mob"
                ],
                "summary": "Insert Inspection",
                "parameters": [
                    {
                        "type": "string",
                        "default": "Bearer \u003cAdd access token here\u003e",
                        "description": "Insert your access token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": " MobRequest Schema ",
                        "name": "MobRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/domain.MobRequest"
                            }
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/domain.MobResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/domain.MobData"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/api/v1/objpart": {
            "get": {
                "description": "Get Object Part",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Object Part"
                ],
                "summary": "Get Object Part",
                "parameters": [
                    {
                        "type": "string",
                        "default": "Bearer \u003cAdd access token here\u003e",
                        "description": "Insert your access token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/domain.ObjPartResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/domain.ObjPartData"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/api/v1/plan": {
            "get": {
                "description": "Get Plan",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Plan"
                ],
                "summary": "Get Plan",
                "parameters": [
                    {
                        "type": "string",
                        "default": "Bearer \u003cAdd access token here\u003e",
                        "description": "Insert your access token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/domain.PlanResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/domain.PlanData"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            },
            "post": {
                "description": "Store Plan",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Plan"
                ],
                "summary": "Store Plan",
                "parameters": [
                    {
                        "type": "string",
                        "default": "Bearer \u003cAdd access token here\u003e",
                        "description": "Insert your access token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": " PlanRequest Schema ",
                        "name": "PlanRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.PlanRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/domain.PlanResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/domain.PlanData"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/api/v1/plan/batch": {
            "post": {
                "description": "Insert Plan",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Plan"
                ],
                "summary": "Insert Plan",
                "parameters": [
                    {
                        "type": "string",
                        "default": "Bearer \u003cAdd access token here\u003e",
                        "description": "Insert your access token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": " PlanRequest Schema ",
                        "name": "PlanRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/domain.PlanRequest"
                            }
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/domain.PlanResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/domain.PlanData"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/api/v1/runacct": {
            "get": {
                "description": "Get Running Account",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Running Account"
                ],
                "summary": "Get Running Account",
                "parameters": [
                    {
                        "type": "string",
                        "default": "Bearer \u003cAdd access token here\u003e",
                        "description": "Insert your access token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "name": "anln1",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "class_group",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "class_ra",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "company_code",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "equnr",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "estate",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "kostl",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "license_plate",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "running_account",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "running_text",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "status_flag",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/domain.RunAcctResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/domain.RunAcctData"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/api/v1/vehicletype": {
            "get": {
                "description": "Get Vehicle Type",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Vehicle Type"
                ],
                "summary": "Get Vehicle Type",
                "parameters": [
                    {
                        "type": "string",
                        "default": "Bearer \u003cAdd access token here\u003e",
                        "description": "Insert your access token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/domain.VehicleTypeResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/domain.VehicleTypeData"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "domain.AttachmentResponse": {
            "type": "object",
            "properties": {
                "data": {},
                "elapsed_time": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "domain.AuthCompanyData": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "estates": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/domain.AuthEstateData"
                    }
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "domain.AuthData": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "token": {
                    "type": "string"
                }
            }
        },
        "domain.AuthEstateData": {
            "type": "object",
            "properties": {
                "code": {
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
                }
            }
        },
        "domain.AuthLoggedData": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string"
                },
                "company": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/domain.AuthCompanyData"
                    }
                },
                "email": {
                    "type": "string"
                },
                "fullname": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "domain.AuthRequest": {
            "type": "object",
            "required": [
                "password",
                "username"
            ],
            "properties": {
                "domain": {
                    "type": "string"
                },
                "ldap": {
                    "type": "boolean"
                },
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "domain.AuthResponse": {
            "type": "object",
            "properties": {
                "data": {},
                "elapsed_time": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "domain.MobData": {
            "type": "object",
            "properties": {
                "ba": {
                    "type": "string"
                },
                "coordinate": {
                    "type": "string"
                },
                "created_by": {
                    "type": "string"
                },
                "created_date": {
                    "type": "string"
                },
                "created_time": {
                    "type": "string"
                },
                "dam_date": {
                    "type": "string"
                },
                "dam_detail": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "is_gps": {
                    "type": "integer"
                },
                "loc": {
                    "type": "string"
                },
                "no_inspec": {
                    "type": "string"
                },
                "note": {
                    "type": "string"
                },
                "obj_part": {
                    "type": "string"
                },
                "photo": {
                    "type": "string"
                },
                "plan_date": {
                    "type": "string"
                },
                "running_account": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
                },
                "vehicle_type": {
                    "type": "string"
                },
                "zinspec_vehicletype_id": {
                    "type": "integer"
                },
                "ztuagri_runacct_id": {
                    "type": "integer"
                }
            }
        },
        "domain.MobRequest": {
            "type": "object",
            "properties": {
                "ba": {
                    "type": "string"
                },
                "coordinate": {
                    "type": "string"
                },
                "created_by": {
                    "type": "string"
                },
                "created_date": {
                    "type": "string"
                },
                "created_time": {
                    "type": "string"
                },
                "dam_date": {
                    "type": "string"
                },
                "dam_detail": {
                    "type": "string"
                },
                "is_gps": {
                    "type": "integer"
                },
                "loc": {
                    "type": "string"
                },
                "no_inspec": {
                    "type": "string"
                },
                "note": {
                    "type": "string"
                },
                "obj_part": {
                    "type": "string"
                },
                "photo": {
                    "type": "string"
                },
                "plan_date": {
                    "type": "string"
                },
                "running_account": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
                },
                "vehicle_type": {
                    "type": "string"
                },
                "zinspec_vehicletype_id": {
                    "type": "integer"
                },
                "ztuagri_runacct_id": {
                    "type": "integer"
                }
            }
        },
        "domain.MobResponse": {
            "type": "object",
            "properties": {
                "data": {},
                "elapsed_time": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "domain.ObjPartData": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "obj_part": {
                    "type": "string"
                },
                "vehicle_type": {
                    "type": "string"
                },
                "zinspec_vehicletype_id": {
                    "type": "integer"
                }
            }
        },
        "domain.ObjPartResponse": {
            "type": "object",
            "properties": {
                "data": {},
                "elapsed_time": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "domain.PlanData": {
            "type": "object",
            "properties": {
                "ba": {
                    "type": "string"
                },
                "company_code": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "inspect_date": {
                    "type": "string"
                },
                "inspect_time": {
                    "type": "string"
                },
                "plan_date": {
                    "type": "string"
                },
                "running_account": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                },
                "week": {
                    "type": "integer"
                },
                "ztuagri_runacct_id": {
                    "type": "integer"
                }
            }
        },
        "domain.PlanRequest": {
            "type": "object",
            "properties": {
                "ba": {
                    "type": "string"
                },
                "company_code": {
                    "type": "string"
                },
                "inspect_date": {
                    "type": "string"
                },
                "inspect_time": {
                    "type": "string"
                },
                "plan_date": {
                    "type": "string"
                },
                "running_account": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                },
                "week": {
                    "type": "integer"
                },
                "ztuagri_runacct_id": {
                    "type": "integer"
                }
            }
        },
        "domain.PlanResponse": {
            "type": "object",
            "properties": {
                "data": {},
                "elapsed_time": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "domain.RunAcctData": {
            "type": "object",
            "properties": {
                "anln1": {
                    "type": "integer"
                },
                "class_group": {
                    "type": "string"
                },
                "class_ra": {
                    "type": "string"
                },
                "company_code": {
                    "type": "string"
                },
                "equnr": {
                    "type": "string"
                },
                "estate": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "kostl": {
                    "type": "string"
                },
                "license_plate": {
                    "type": "string"
                },
                "running_account": {
                    "type": "string"
                },
                "running_text": {
                    "type": "string"
                },
                "status_flag": {
                    "type": "string"
                }
            }
        },
        "domain.RunAcctResponse": {
            "type": "object",
            "properties": {
                "data": {},
                "elapsed_time": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "domain.VehicleTypeData": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "obj_part": {
                    "type": "string"
                },
                "vehicle_type": {
                    "type": "string"
                }
            }
        },
        "domain.VehicleTypeResponse": {
            "type": "object",
            "properties": {
                "data": {},
                "elapsed_time": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0.0",
	Host:             "dev.indoagri.co.id",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "API SWAGGER FOR INSPECTION RA API SERVICE",
	Description:      "INSPECTION RA API SERVICE",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
