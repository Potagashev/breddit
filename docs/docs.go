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
        "/threads": {
            "get": {
                "tags": [
                    "threads"
                ],
                "summary": "List threads",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/threads.Thread"
                            }
                        }
                    }
                }
            },
            "post": {
                "tags": [
                    "threads"
                ],
                "summary": "Create thread",
                "parameters": [
                    {
                        "description": "Create data",
                        "name": "requestBody",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/threads.ThreadWrite"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/threads.ThreadWriteResponse"
                        }
                    }
                }
            }
        },
        "/threads/{id}": {
            "get": {
                "tags": [
                    "threads"
                ],
                "summary": "Get thread",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Thread ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/threads.Thread"
                        }
                    }
                }
            },
            "put": {
                "tags": [
                    "threads"
                ],
                "summary": "Update thread",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Thread ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Update data",
                        "name": "requestBody",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/threads.ThreadWrite"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/threads.ThreadWriteResponse"
                        }
                    }
                }
            },
            "delete": {
                "tags": [
                    "threads"
                ],
                "summary": "Delete thread",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Thread ID",
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
        "threads.Thread": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "text": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "threads.ThreadWrite": {
            "type": "object",
            "properties": {
                "text": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "threads.ThreadWriteResponse": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "Swagger Example API",
	Description:      "This is a sample server.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}