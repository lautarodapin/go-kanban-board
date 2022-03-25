// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
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
        "/columns": {
            "get": {
                "description": "Returns all columns",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "GetColumns"
                ],
                "summary": "Get all columns",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "array",
                                "items": {
                                    "$ref": "#/definitions/models.Column"
                                }
                            }
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new column",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "CreateColumn"
                ],
                "summary": "Create a new column",
                "parameters": [
                    {
                        "description": "Column",
                        "name": "column",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Column"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Column"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/columns/:id": {
            "get": {
                "description": "Returns a column by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "GetColumn"
                ],
                "summary": "Get a column by id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Column ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Column"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "put": {
                "description": "Update a column",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "UpdateColumn"
                ],
                "summary": "Update a column",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Column ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Column",
                        "name": "column",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Column"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Column"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/dropzones": {
            "get": {
                "description": "Returns all dropzones",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Dropzones list"
                ],
                "summary": "Get all dropzones",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "array",
                                "items": {
                                    "$ref": "#/definitions/models.Dropzone"
                                }
                            }
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new dropzone",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Create Dropzone"
                ],
                "summary": "Create a new dropzone",
                "parameters": [
                    {
                        "description": "Dropzone",
                        "name": "dropzone",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.BaseDropzone"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Dropzone"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/dropzones/:id": {
            "get": {
                "description": "Returns a dropzone by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Get Dropzone"
                ],
                "summary": "Get a dropzone by id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Dropzone ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Dropzone"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "put": {
                "description": "Update a dropzone",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Update Dropzone"
                ],
                "summary": "Update a dropzone",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Dropzone ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Dropzone",
                        "name": "dropzone",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.BaseDropzone"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Dropzone"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/kanbans": {
            "get": {
                "description": "Returns all kanbans",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "GetKanbans"
                ],
                "summary": "get all kanbans",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "array",
                                "items": {
                                    "$ref": "#/definitions/models.Kanban"
                                }
                            }
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new kanban",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "CreateKanban"
                ],
                "summary": "Create a new kanban",
                "parameters": [
                    {
                        "description": "Kanban",
                        "name": "kanban",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.CreateKanban"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Kanban"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/kanbans/:id": {
            "get": {
                "description": "Return a kanban by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "GetKanban"
                ],
                "summary": "Return a kanban by id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Kanban ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Kanban"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "put": {
                "description": "Update a kanban",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "UpdateKanban"
                ],
                "summary": "Update a kanban",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Kanban ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Kanban",
                        "name": "kanban",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.UpdateKanban"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Kanban"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/tickets": {
            "get": {
                "description": "Returns all tickets",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Tickets list"
                ],
                "summary": "Get all tickets",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "array",
                                "items": {
                                    "$ref": "#/definitions/models.Ticket"
                                }
                            }
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new ticket",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Create Ticket"
                ],
                "summary": "Create a new ticket",
                "parameters": [
                    {
                        "description": "Ticket",
                        "name": "ticket",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.BaseTicket"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Ticket"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/tickets/:id": {
            "get": {
                "description": "Returns a ticket by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Get Ticket"
                ],
                "summary": "Get a ticket by id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Ticket ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Ticket"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "put": {
                "description": "Update a ticket",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Update Ticket"
                ],
                "summary": "Update a ticket",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Ticket ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Ticket",
                        "name": "ticket",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.BaseTicket"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Ticket"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "gorm.DeletedAt": {
            "type": "object",
            "properties": {
                "time": {
                    "type": "string"
                },
                "valid": {
                    "description": "Valid is true if Time is not NULL",
                    "type": "boolean"
                }
            }
        },
        "models.BaseDropzone": {
            "type": "object",
            "properties": {
                "column_id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "order": {
                    "type": "integer"
                }
            }
        },
        "models.BaseTicket": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "dropzone_id": {
                    "type": "integer"
                },
                "kanban_id": {
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "models.Column": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "deleted_at": {
                    "$ref": "#/definitions/gorm.DeletedAt"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "order": {
                    "type": "integer"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "models.CreateKanban": {
            "type": "object",
            "properties": {
                "end_date": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "start_date": {
                    "type": "string"
                }
            }
        },
        "models.Dropzone": {
            "type": "object",
            "properties": {
                "column": {
                    "$ref": "#/definitions/models.Column"
                },
                "column_id": {
                    "type": "integer"
                },
                "created_at": {
                    "type": "string"
                },
                "deleted_at": {
                    "$ref": "#/definitions/gorm.DeletedAt"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "order": {
                    "type": "integer"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "models.Kanban": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "deleted_at": {
                    "$ref": "#/definitions/gorm.DeletedAt"
                },
                "end_date": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "start_date": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "models.Ticket": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "deleted_at": {
                    "$ref": "#/definitions/gorm.DeletedAt"
                },
                "description": {
                    "type": "string"
                },
                "dropzone": {
                    "$ref": "#/definitions/models.Dropzone"
                },
                "dropzone_id": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "kanban": {
                    "$ref": "#/definitions/models.Kanban"
                },
                "kanban_id": {
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "models.UpdateKanban": {
            "type": "object",
            "properties": {
                "end_date": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
