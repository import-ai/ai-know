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
        "license": {
            "name": "GPLv3"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/sidebar/list/entries": {
            "post": {
                "description": "Create an entry with the specified properties.",
                "tags": [
                    "Sidebar"
                ],
                "summary": "Create Entry",
                "parameters": [
                    {
                        "description": "Request Body",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handlers.CreateEntry.Req"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handlers.CreateEntry.Resp"
                        }
                    }
                }
            }
        },
        "/api/sidebar/list/entries/{entry_id}": {
            "get": {
                "description": "Get properties of an entry.",
                "tags": [
                    "Sidebar"
                ],
                "summary": "Get Entry",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Entry ID",
                        "name": "entry_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handlers.GetEntry.Resp"
                        }
                    }
                }
            },
            "put": {
                "description": "Update properties of an entry.",
                "tags": [
                    "Sidebar"
                ],
                "summary": "Update Entry",
                "parameters": [
                    {
                        "description": "Request Body",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handlers.PutEntry.Req"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handlers.PutEntry.Resp"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete an entry and all its sub-entries.",
                "tags": [
                    "Sidebar"
                ],
                "summary": "Delete Entry",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Entry ID",
                        "name": "entry_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/api/sidebar/list/entries/{entry_id}/duplicate": {
            "post": {
                "description": "Duplicate an entry.",
                "tags": [
                    "Sidebar"
                ],
                "summary": "Duplicate Entry",
                "parameters": [
                    {
                        "description": "Request Body",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handlers.DuplicateEntry.Req"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handlers.DuplicateEntry.Resp"
                        }
                    }
                }
            }
        },
        "/api/sidebar/list/entries/{entry_id}/sub_entries": {
            "get": {
                "description": "Get sub-entries of an entry.",
                "tags": [
                    "Sidebar"
                ],
                "summary": "Get Sub-Entries",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Entry ID",
                        "name": "entry_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handlers.GetSubEntries.Resp"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "handlers.CreateEntry.Req": {
            "type": "object",
            "properties": {
                "parent": {
                    "type": "string"
                },
                "position_after": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                },
                "type": {
                    "type": "string"
                }
            }
        },
        "handlers.CreateEntry.Resp": {
            "type": "object",
            "properties": {
                "entry": {
                    "$ref": "#/definitions/handlers.Entry"
                }
            }
        },
        "handlers.DuplicateEntry.Req": {
            "type": "object",
            "properties": {
                "parent": {
                    "type": "string"
                },
                "position_after": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "handlers.DuplicateEntry.Resp": {
            "type": "object",
            "properties": {
                "entry": {
                    "$ref": "#/definitions/handlers.Entry"
                }
            }
        },
        "handlers.Entry": {
            "type": "object",
            "properties": {
                "has_sub_entries": {
                    "type": "boolean",
                    "example": false
                },
                "id": {
                    "type": "string",
                    "example": "1000001"
                },
                "title": {
                    "type": "string",
                    "example": "Note Title"
                },
                "type": {
                    "type": "string",
                    "enum": [
                        "note",
                        "group",
                        "link"
                    ],
                    "example": "note"
                }
            }
        },
        "handlers.GetEntry.Resp": {
            "type": "object",
            "properties": {
                "entry": {
                    "$ref": "#/definitions/handlers.Entry"
                }
            }
        },
        "handlers.GetSubEntries.Resp": {
            "type": "object",
            "properties": {
                "sub_entries": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/handlers.Entry"
                    }
                }
            }
        },
        "handlers.PutEntry.Req": {
            "type": "object",
            "properties": {
                "parent": {
                    "type": "string"
                },
                "position_after": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "handlers.PutEntry.Resp": {
            "type": "object",
            "properties": {
                "entry": {
                    "$ref": "#/definitions/handlers.Entry"
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
	Title:            "AIKnow API",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
