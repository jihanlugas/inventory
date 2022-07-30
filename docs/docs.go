// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
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
            "name": "API Support",
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
        "/": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Ping"
                ],
                "summary": "Ping",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            }
        },
        "/item": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Item"
                ],
                "summary": "Create Item",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Item Name",
                        "name": "itemName",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Item Description",
                        "name": "itemDescription",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "boolean",
                        "description": "Active",
                        "name": "isActive",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "file",
                        "description": "Photo",
                        "name": "photo",
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
                                    "$ref": "#/definitions/response.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "payload": {
                                            "$ref": "#/definitions/controller.itemRes"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            }
        },
        "/item/page": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Item"
                ],
                "summary": "Page Item",
                "parameters": [
                    {
                        "description": "payload",
                        "name": "req",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controller.pageItemReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "payload": {
                                            "$ref": "#/definitions/response.Pagination"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            }
        },
        "/item/{item_id}": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Item"
                ],
                "summary": "Get Item",
                "parameters": [
                    {
                        "type": "number",
                        "description": "item_id",
                        "name": "item_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "payload": {
                                            "$ref": "#/definitions/controller.itemRes"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            },
            "put": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Item"
                ],
                "summary": "Update Item",
                "parameters": [
                    {
                        "type": "number",
                        "description": "item_id",
                        "name": "item_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Item Name",
                        "name": "itemName",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Item Description",
                        "name": "itemDescription",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "boolean",
                        "description": "Active",
                        "name": "isActive",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "file",
                        "description": "Photo",
                        "name": "photo",
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
                                    "$ref": "#/definitions/response.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "payload": {
                                            "$ref": "#/definitions/controller.itemRes"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            },
            "delete": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Item"
                ],
                "summary": "Delete Item",
                "parameters": [
                    {
                        "type": "number",
                        "description": "item_id",
                        "name": "item_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            }
        },
        "/itemvariant": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Itemvariant"
                ],
                "summary": "Create Itemvariant",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Itemvariant Name",
                        "name": "itemvariantName",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Itemvariant Description",
                        "name": "itemvariantDescription",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Price",
                        "name": "price",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "boolean",
                        "description": "Active",
                        "name": "isActive",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "file",
                        "description": "Photo",
                        "name": "photo",
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
                                    "$ref": "#/definitions/response.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "payload": {
                                            "$ref": "#/definitions/controller.itemvariantRes"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            }
        },
        "/itemvariant/page": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Itemvariant"
                ],
                "summary": "Page Itemvariant",
                "parameters": [
                    {
                        "description": "payload",
                        "name": "req",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controller.pageItemvariantReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "payload": {
                                            "$ref": "#/definitions/response.Pagination"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            }
        },
        "/itemvariant/{itemvariant_id}": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Itemvariant"
                ],
                "summary": "Get Itemvariant",
                "parameters": [
                    {
                        "type": "number",
                        "description": "itemvariant_id",
                        "name": "itemvariant_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "payload": {
                                            "$ref": "#/definitions/controller.itemvariantRes"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            },
            "put": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Itemvariant"
                ],
                "summary": "Update Itemvariant",
                "parameters": [
                    {
                        "type": "number",
                        "description": "itemvariant_id",
                        "name": "itemvariant_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Itemvariant Name",
                        "name": "itemvariantName",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Itemvariant Description",
                        "name": "itemvariantDescription",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Price",
                        "name": "price",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "boolean",
                        "description": "Active",
                        "name": "isActive",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "file",
                        "description": "Photo",
                        "name": "photo",
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
                                    "$ref": "#/definitions/response.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "payload": {
                                            "$ref": "#/definitions/controller.itemvariantRes"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            },
            "delete": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Itemvariant"
                ],
                "summary": "Delete Itemvariant",
                "parameters": [
                    {
                        "type": "number",
                        "description": "itemvariant_id",
                        "name": "itemvariant_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            }
        },
        "/sign-in": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Authentication"
                ],
                "summary": "Sign in a user",
                "parameters": [
                    {
                        "description": "json req body",
                        "name": "req",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controller.signinReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            }
        },
        "/sign-out": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Authentication"
                ],
                "summary": "Sign out a user",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            }
        },
        "/sign-up": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Authentication"
                ],
                "summary": "Sign up a user",
                "parameters": [
                    {
                        "description": "json req body",
                        "name": "req",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controller.signupReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            }
        },
        "/user/me": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "To do get current active user",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "controller.itemRes": {
            "type": "object",
            "properties": {
                "createBy": {
                    "type": "integer"
                },
                "createDt": {
                    "type": "string"
                },
                "isActive": {
                    "type": "boolean"
                },
                "itemDescription": {
                    "type": "string"
                },
                "itemId": {
                    "type": "integer"
                },
                "itemName": {
                    "type": "string"
                },
                "photoUrl": {
                    "type": "string"
                },
                "propertyId": {
                    "type": "integer"
                },
                "updateBy": {
                    "type": "integer"
                },
                "updateDt": {
                    "type": "string"
                }
            }
        },
        "controller.itemvariantRes": {
            "type": "object",
            "properties": {
                "createBy": {
                    "type": "integer"
                },
                "createDt": {
                    "type": "string"
                },
                "isActive": {
                    "type": "boolean"
                },
                "itemId": {
                    "type": "integer"
                },
                "itemvariantDescription": {
                    "type": "string"
                },
                "itemvariantId": {
                    "type": "integer"
                },
                "itemvariantName": {
                    "type": "string"
                },
                "photoUrl": {
                    "type": "string"
                },
                "price": {
                    "type": "integer"
                },
                "updateBy": {
                    "type": "integer"
                },
                "updateDt": {
                    "type": "string"
                }
            }
        },
        "controller.pageItemReq": {
            "type": "object",
            "properties": {
                "itemDescription": {
                    "type": "string"
                },
                "itemName": {
                    "type": "string"
                },
                "limit": {
                    "type": "integer"
                },
                "page": {
                    "type": "integer"
                }
            }
        },
        "controller.pageItemvariantReq": {
            "type": "object",
            "required": [
                "itemId"
            ],
            "properties": {
                "itemId": {
                    "type": "integer"
                },
                "itemvariantDescription": {
                    "type": "string"
                },
                "itemvariantName": {
                    "type": "string"
                },
                "limit": {
                    "type": "integer"
                },
                "page": {
                    "type": "integer"
                }
            }
        },
        "controller.signinReq": {
            "type": "object",
            "required": [
                "passwd",
                "username"
            ],
            "properties": {
                "passwd": {
                    "type": "string",
                    "maxLength": 200
                },
                "username": {
                    "type": "string",
                    "maxLength": 20
                }
            }
        },
        "controller.signupReq": {
            "type": "object",
            "required": [
                "confirmPasswd",
                "email",
                "fullname",
                "noHp",
                "passwd",
                "propertyName",
                "username"
            ],
            "properties": {
                "confirmPasswd": {
                    "type": "string",
                    "maxLength": 200
                },
                "email": {
                    "type": "string",
                    "maxLength": 200
                },
                "fullname": {
                    "type": "string",
                    "maxLength": 80
                },
                "noHp": {
                    "type": "string",
                    "maxLength": 20
                },
                "passwd": {
                    "type": "string",
                    "maxLength": 200
                },
                "propertyName": {
                    "type": "string",
                    "maxLength": 200
                },
                "username": {
                    "type": "string",
                    "maxLength": 20
                }
            }
        },
        "response.Pagination": {
            "type": "object",
            "properties": {
                "dataPerPage": {
                    "type": "integer"
                },
                "list": {
                    "type": "array",
                    "items": {
                        "type": "object"
                    }
                },
                "page": {
                    "type": "integer"
                },
                "totalData": {
                    "type": "integer"
                },
                "totalPage": {
                    "type": "integer"
                }
            }
        },
        "response.Response": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "message": {
                    "type": "string"
                },
                "payload": {
                    "type": "object"
                },
                "success": {
                    "type": "boolean"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Swagger Example API",
	Description:      "This is a sample server Petstore server.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}