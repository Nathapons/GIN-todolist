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
                "description": "Return list of blog.",
                "tags": [
                    "Blog"
                ],
                "summary": "Create Blogs.",
                "parameters": [
                    {
                        "description": "Request body for creating a resource",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/forms.Blog"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/blog/{id}": {
            "put": {
                "description": "Return list of blog.",
                "tags": [
                    "Blog"
                ],
                "summary": "Update Blogs.",
                "operationId": "create-resource",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID of models.Blog",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Request body for update a resource",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/forms.Blog"
                        }
                    }
                ],
                "responses": {}
            },
            "delete": {
                "description": "Return list of blog.",
                "tags": [
                    "Blog"
                ],
                "summary": "Delete Blogs.",
                "responses": {}
            }
        }
    },
    "definitions": {
        "forms.Blog": {
            "type": "object",
            "properties": {
                "author": {
                    "type": "string",
                    "default": "admin"
                },
                "title": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:7000",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "Blog todolist API",
	Description:      "Blog todolist API in Go using GIN framework",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
