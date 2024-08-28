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
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/login": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Log in a user",
                "parameters": [
                    {
                        "description": "Request Body",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/user_models.Login"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/user_models.LoginResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid request format",
                        "schema": {
                            "$ref": "#/definitions/fiber.Map"
                        }
                    },
                    "401": {
                        "description": "Unauthorized (Invalid email or password)",
                        "schema": {
                            "$ref": "#/definitions/fiber.Map"
                        }
                    },
                    "404": {
                        "description": "User not found",
                        "schema": {
                            "$ref": "#/definitions/fiber.Map"
                        }
                    },
                    "default": {
                        "description": "An unexpected error response.",
                        "schema": {
                            "$ref": "#/definitions/controller_common.SingleErrorMessage"
                        }
                    }
                }
            }
        },
        "/register": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Create a new user",
                "parameters": [
                    {
                        "description": "Request Body",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/user_models.CreateUser"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created user",
                        "schema": {
                            "$ref": "#/definitions/user_models.User"
                        }
                    },
                    "400": {
                        "description": "Invalid request format or validation error",
                        "schema": {
                            "$ref": "#/definitions/fiber.Map"
                        }
                    },
                    "500": {
                        "description": "Failed to hash password or create user",
                        "schema": {
                            "$ref": "#/definitions/fiber.Map"
                        }
                    },
                    "default": {
                        "description": "An unexpected error response.",
                        "schema": {
                            "$ref": "#/definitions/controller_common.SingleErrorMessage"
                        }
                    }
                }
            }
        },
        "/users": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "List users",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Filter users by discount percentage",
                        "name": "discount_percentage",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "default": 1,
                        "description": "Page number for pagination",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "default": 10,
                        "description": "Limit of results per page",
                        "name": "limit",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Paginated list of users",
                        "schema": {
                            "$ref": "#/definitions/utils.Pagination"
                        }
                    },
                    "500": {
                        "description": "Failed to fetch users",
                        "schema": {
                            "$ref": "#/definitions/fiber.Map"
                        }
                    },
                    "default": {
                        "description": "An unexpected error response.",
                        "schema": {
                            "$ref": "#/definitions/controller_common.SingleErrorMessage"
                        }
                    }
                }
            }
        },
        "/users/{id}": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Get user by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "User details",
                        "schema": {
                            "$ref": "#/definitions/user_models.User"
                        }
                    },
                    "400": {
                        "description": "User ID is required",
                        "schema": {
                            "$ref": "#/definitions/fiber.Map"
                        }
                    },
                    "404": {
                        "description": "User not found",
                        "schema": {
                            "$ref": "#/definitions/fiber.Map"
                        }
                    },
                    "default": {
                        "description": "An unexpected error response.",
                        "schema": {
                            "$ref": "#/definitions/controller_common.SingleErrorMessage"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "controller_common.SingleErrorMessage": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "fiber.Map": {
            "type": "object",
            "additionalProperties": true
        },
        "user_models.CreateUser": {
            "type": "object",
            "required": [
                "address",
                "email",
                "firstName",
                "lastName",
                "password",
                "phone"
            ],
            "properties": {
                "address": {
                    "type": "string",
                    "minLength": 6
                },
                "discountPercentage": {
                    "type": "integer"
                },
                "email": {
                    "type": "string",
                    "maxLength": 64,
                    "minLength": 6
                },
                "firstName": {
                    "type": "string",
                    "maxLength": 64,
                    "minLength": 3
                },
                "isAdmin": {
                    "type": "boolean"
                },
                "lastName": {
                    "type": "string",
                    "maxLength": 64,
                    "minLength": 3
                },
                "password": {
                    "type": "string",
                    "minLength": 6
                },
                "phone": {
                    "type": "string",
                    "maxLength": 64,
                    "minLength": 6
                }
            }
        },
        "user_models.Login": {
            "type": "object",
            "required": [
                "email",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string",
                    "maxLength": 64,
                    "minLength": 6
                },
                "password": {
                    "type": "string",
                    "minLength": 6
                }
            }
        },
        "user_models.LoginResponse": {
            "type": "object",
            "properties": {
                "token": {
                    "type": "string"
                }
            }
        },
        "user_models.User": {
            "type": "object",
            "required": [
                "address",
                "email",
                "firstName",
                "lastName",
                "password",
                "phone"
            ],
            "properties": {
                "address": {
                    "type": "string",
                    "maxLength": 64,
                    "minLength": 6
                },
                "discountPercentage": {
                    "type": "integer"
                },
                "email": {
                    "type": "string",
                    "maxLength": 64,
                    "minLength": 6
                },
                "firstName": {
                    "type": "string",
                    "maxLength": 64,
                    "minLength": 3
                },
                "isAdmin": {
                    "type": "boolean"
                },
                "lastName": {
                    "type": "string",
                    "maxLength": 64,
                    "minLength": 3
                },
                "password": {
                    "type": "string",
                    "maxLength": 64,
                    "minLength": 6
                },
                "phone": {
                    "type": "string",
                    "maxLength": 64,
                    "minLength": 6
                }
            }
        },
        "utils.Pagination": {
            "type": "object",
            "properties": {
                "data": {},
                "limit": {
                    "type": "integer"
                },
                "next": {
                    "type": "integer"
                },
                "page": {
                    "type": "integer"
                },
                "prev": {
                    "type": "integer"
                },
                "total": {
                    "type": "integer"
                },
                "totalPages": {
                    "type": "integer"
                }
            }
        }
    },
    "securityDefinitions": {
        "BearerAuth": {
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
	BasePath:         "/v1",
	Schemes:          []string{},
	Title:            "Evolution Fitness Store API",
	Description:      "This is Evolution Fitness Store API.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
