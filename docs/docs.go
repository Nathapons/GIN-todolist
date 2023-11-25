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
        "/blog": {
            "get": {
                "description": "Return list of blog.",
                "tags": [
                    "Blog"
                ],
                "summary": "Get Blogs.",
                "responses": {}
            },
            "post": {
                "description": "Return blog data.",
                "tags": [
                    "Blog"
                ],
                "summary": "Create Blogs.",
                "parameters": [
                    {
                        "type": "file",
                        "description": "File to upload",
                        "name": "file",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Title data",
                        "name": "title",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Author data",
                        "name": "author",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/blog/{id}": {
            "put": {
                "description": "Return new blog after update data.",
                "tags": [
                    "Blog"
                ],
                "summary": "Update Blogs.",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID of models.Blog",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "file",
                        "description": "File",
                        "name": "file",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Title",
                        "name": "title",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Author",
                        "name": "author",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {}
            },
            "delete": {
                "description": "Return no context status.",
                "tags": [
                    "Blog"
                ],
                "summary": "Delete Blogs.",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID of models.Blog",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/user": {
            "get": {
                "description": "Return users.",
                "tags": [
                    "User"
                ],
                "summary": "Get Users.",
                "responses": {}
            },
            "post": {
                "description": "Return user.",
                "tags": [
                    "User"
                ],
                "summary": "Create Users.",
                "parameters": [
                    {
                        "description": "User data in JSON format",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/forms.User"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/user/{id}": {
            "put": {
                "description": "Return new user.",
                "tags": [
                    "User"
                ],
                "summary": "Update Users.",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID of models.User",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "User data in JSON format",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/forms.User"
                        }
                    }
                ],
                "responses": {}
            },
            "delete": {
                "description": "Return new user.",
                "tags": [
                    "User"
                ],
                "summary": "Update Users.",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID of models.User",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        }
    },
    "definitions": {
        "forms.User": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "first_name": {
                    "type": "string"
                },
                "last_name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8000",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "Blog Service API",
	Description:      "Blog Service API in Go using GIN framework",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
