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
        "license": {
            "name": "MIT"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/task": {
            "get": {
                "description": "Get task from database",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "tasks"
                ],
                "summary": "Get task",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Task identifer",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/tasks.TaskResponse"
                        }
                    },
                    "502": {
                        "description": "Get task failed",
                        "schema": {
                            "$ref": "#/definitions/tasks.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/task/all": {
            "get": {
                "description": "Get all tasks from database",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "tasks"
                ],
                "summary": "Get all tasks",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/tasks.TaskListResponse"
                        }
                    },
                    "502": {
                        "description": "Get all tasks failed",
                        "schema": {
                            "$ref": "#/definitions/tasks.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/task/create": {
            "post": {
                "description": "Add task to database with data which contains content",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "tasks"
                ],
                "summary": "Add task",
                "parameters": [
                    {
                        "description": "Task content",
                        "name": "content",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/tasks.TaskResponse"
                        }
                    },
                    "400": {
                        "description": "Non-valid data",
                        "schema": {
                            "$ref": "#/definitions/tasks.ErrorResponse"
                        }
                    },
                    "502": {
                        "description": "Add task failed",
                        "schema": {
                            "$ref": "#/definitions/tasks.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/task/delete": {
            "delete": {
                "description": "Delete task from database with data which contains id",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "tasks"
                ],
                "summary": "Delete task",
                "parameters": [
                    {
                        "description": "Task identifer",
                        "name": "id",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "integer"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/tasks.TaskResponse"
                        }
                    },
                    "400": {
                        "description": "Non-valid data",
                        "schema": {
                            "$ref": "#/definitions/tasks.ErrorResponse"
                        }
                    },
                    "502": {
                        "description": "Delete task failed",
                        "schema": {
                            "$ref": "#/definitions/tasks.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/task/done": {
            "get": {
                "description": "Get tasks which marked as done",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "tasks"
                ],
                "summary": "Get done tasks",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/tasks.TaskListResponse"
                        }
                    },
                    "502": {
                        "description": "Get done tasks failed",
                        "schema": {
                            "$ref": "#/definitions/tasks.ErrorResponse"
                        }
                    }
                }
            },
            "put": {
                "description": "Mark task as done with data which contains id",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "tasks"
                ],
                "summary": "Done task",
                "parameters": [
                    {
                        "description": "Task indentifer",
                        "name": "id",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "integer"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/tasks.TaskResponse"
                        }
                    },
                    "400": {
                        "description": "Non-valid data",
                        "schema": {
                            "$ref": "#/definitions/tasks.ErrorResponse"
                        }
                    },
                    "502": {
                        "description": "Mark task as done failed",
                        "schema": {
                            "$ref": "#/definitions/tasks.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/task/notDone": {
            "get": {
                "description": "Get tasks which not marked as done",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "tasks"
                ],
                "summary": "Get not done tasks",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/tasks.TaskListResponse"
                        }
                    },
                    "502": {
                        "description": "Get not done tasks failed",
                        "schema": {
                            "$ref": "#/definitions/tasks.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/user/create": {
            "post": {
                "description": "Create user to database with data which contains username and password",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Create user",
                "parameters": [
                    {
                        "description": "Your future username",
                        "name": "username",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "Your future password",
                        "name": "password",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/users.UserResponse"
                        }
                    },
                    "400": {
                        "description": "Non-valid password",
                        "schema": {
                            "$ref": "#/definitions/users.ErrorResponse"
                        }
                    },
                    "502": {
                        "description": "Create user failed",
                        "schema": {
                            "$ref": "#/definitions/users.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "tasks.ErrorResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "tasks.Task": {
            "type": "object",
            "properties": {
                "content": {
                    "type": "string"
                },
                "done": {
                    "type": "boolean"
                },
                "id": {
                    "type": "integer"
                }
            }
        },
        "tasks.TaskListResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/tasks.Task"
                    }
                }
            }
        },
        "tasks.TaskResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/tasks.Task"
                }
            }
        },
        "users.ErrorResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "users.User": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "updated_at": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "users.UserResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/users.User"
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
	Title:            "Todo API",
	Description:      "Simple todo API",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
