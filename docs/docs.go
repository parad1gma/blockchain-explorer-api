// Code generated by swaggo/swag. DO NOT EDIT
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
        "/v1/example/helloworld": {
            "get": {
                "description": "do ping",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "example"
                ],
                "summary": "ping example",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/v1/transaction/hash/{txhash}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "transaction"
                ],
                "summary": "Get transaction by hash",
                "parameters": [
                    {
                        "type": "string",
                        "description": "transaction hash",
                        "name": "txhash",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Transaction"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/v1/transaction/txinblock/{blocknumber}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "transaction"
                ],
                "summary": "Get transactions in block",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "block number",
                        "name": "blocknumber",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Transaction"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/v2/example/helloworld": {
            "get": {
                "description": "do ping",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "example"
                ],
                "summary": "ping example",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Transaction": {
            "type": "object",
            "properties": {
                "blockHash": {
                    "type": "string"
                },
                "blockNUmber": {
                    "type": "integer"
                },
                "contractAddress": {
                    "type": "string"
                },
                "from": {
                    "type": "string"
                },
                "gas": {
                    "type": "integer"
                },
                "gasPrice": {
                    "type": "integer"
                },
                "gasUsed": {
                    "type": "integer"
                },
                "hash": {
                    "type": "string"
                },
                "input": {
                    "type": "string"
                },
                "nonce": {
                    "type": "integer"
                },
                "status": {
                    "type": "integer"
                },
                "timestamp": {
                    "type": "integer"
                },
                "to": {
                    "type": "string"
                },
                "transactionIndex": {
                    "type": "integer"
                },
                "value": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8888",
	BasePath:         "/api",
	Schemes:          []string{},
	Title:            "Block Explorer API",
	Description:      "This is a block explorer server. You can visit the GitHub repository at https://github.com/Ethernal-Tech/blockchain-explorer-api",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
