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
        "/dec": {
            "get": {
                "description": "Read one dec",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Dec"
                ],
                "summary": "Read one dec",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Dec identification",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handler.ReadDecResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handler.ErrorDecResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/handler.ErrorDecResponse"
                        }
                    }
                }
            },
            "put": {
                "description": "Update dec",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Dec"
                ],
                "summary": "Update dec",
                "parameters": [
                    {
                        "description": "Request body",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handler.UpdateDecRequest"
                        }
                    },
                    {
                        "type": "string",
                        "description": "Dec identification",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handler.UpdateDecResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handler.ErrorDecResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/handler.ErrorDecResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new dec",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Dec"
                ],
                "summary": "Create Dec",
                "parameters": [
                    {
                        "description": "Request body",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handler.CreateDecRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/handler.CreateDecResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handler.ErrorDecResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handler.ErrorDecResponse"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete dec",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Dec"
                ],
                "summary": "Delete dec",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Dec identification",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handler.DeleteDecResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handler.ErrorDecResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/handler.ErrorDecResponse"
                        }
                    }
                }
            }
        },
        "/decs": {
            "get": {
                "description": "List dec",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Dec"
                ],
                "summary": "List dec",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handler.ListDecResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handler.ErrorDecResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/handler.ErrorDecResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "handler.CreateDecRequest": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                },
                "userId": {
                    "type": "integer"
                }
            }
        },
        "handler.CreateDecResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/handler.CreateDecResponseData"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "handler.CreateDecResponseData": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "userId": {
                    "type": "integer"
                }
            }
        },
        "handler.DecReadResponseData": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "userId": {
                    "type": "integer"
                }
            }
        },
        "handler.DecUpdateResponseData": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "userId": {
                    "type": "integer"
                }
            }
        },
        "handler.DeleteDecResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/handler.DeleteDecResponseData"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "handler.DeleteDecResponseData": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "userId": {
                    "type": "integer"
                }
            }
        },
        "handler.ErrorDecResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "handler.ListDecResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/handler.ListDecResponseData"
                    }
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "handler.ListDecResponseData": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "userId": {
                    "type": "integer"
                }
            }
        },
        "handler.ReadDecResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/handler.DecReadResponseData"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "handler.UpdateDecRequest": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                },
                "userId": {
                    "type": "integer"
                }
            }
        },
        "handler.UpdateDecResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/handler.DecUpdateResponseData"
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
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
