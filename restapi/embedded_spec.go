// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "プロラボの図書館理システムです。",
    "title": "Toshokan",
    "version": "1.0.0"
  },
  "paths": {
    "/sessions": {
      "post": {
        "description": "メールアドレスとパスワードでログインする。",
        "consumes": [
          "application/json"
        ],
        "tags": [
          "sessions"
        ],
        "summary": "ログイン",
        "operationId": "login",
        "parameters": [
          {
            "description": "メールアドレスとパスワード",
            "name": "email and password",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/EmailAndPassword"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Created",
            "schema": {
              "type": "object",
              "properties": {
                "token": {
                  "type": "string",
                  "example": "1234_abcdefghijklmn"
                },
                "user": {
                  "$ref": "#/definitions/User"
                }
              }
            }
          },
          "401": {
            "description": "Unauthorized"
          }
        }
      },
      "delete": {
        "description": "ログアウトする。",
        "tags": [
          "sessions"
        ],
        "summary": "ログアウト",
        "operationId": "logout",
        "responses": {
          "204": {
            "description": "No Content"
          },
          "401": {
            "description": "Unauthorized"
          }
        }
      }
    }
  },
  "definitions": {
    "EmailAndPassword": {
      "type": "object",
      "required": [
        "email",
        "password"
      ],
      "properties": {
        "email": {
          "type": "string",
          "example": "hoge@example.com"
        },
        "password": {
          "type": "string"
        }
      }
    },
    "User": {
      "type": "object",
      "properties": {
        "displayName": {
          "type": "string",
          "example": "あいうえお"
        },
        "email": {
          "type": "string",
          "example": "hoge@example.com"
        },
        "isAdmin": {
          "type": "boolean"
        },
        "name": {
          "type": "string"
        },
        "userID": {
          "type": "integer",
          "format": "int32"
        }
      }
    }
  },
  "tags": [
    {
      "description": "ログインとログアウト",
      "name": "sessions"
    }
  ]
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "プロラボの図書館理システムです。",
    "title": "Toshokan",
    "version": "1.0.0"
  },
  "paths": {
    "/sessions": {
      "post": {
        "description": "メールアドレスとパスワードでログインする。",
        "consumes": [
          "application/json"
        ],
        "tags": [
          "sessions"
        ],
        "summary": "ログイン",
        "operationId": "login",
        "parameters": [
          {
            "description": "メールアドレスとパスワード",
            "name": "email and password",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/EmailAndPassword"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Created",
            "schema": {
              "type": "object",
              "properties": {
                "token": {
                  "type": "string",
                  "example": "1234_abcdefghijklmn"
                },
                "user": {
                  "$ref": "#/definitions/User"
                }
              }
            }
          },
          "401": {
            "description": "Unauthorized"
          }
        }
      },
      "delete": {
        "description": "ログアウトする。",
        "tags": [
          "sessions"
        ],
        "summary": "ログアウト",
        "operationId": "logout",
        "responses": {
          "204": {
            "description": "No Content"
          },
          "401": {
            "description": "Unauthorized"
          }
        }
      }
    }
  },
  "definitions": {
    "EmailAndPassword": {
      "type": "object",
      "required": [
        "email",
        "password"
      ],
      "properties": {
        "email": {
          "type": "string",
          "example": "hoge@example.com"
        },
        "password": {
          "type": "string"
        }
      }
    },
    "User": {
      "type": "object",
      "properties": {
        "displayName": {
          "type": "string",
          "example": "あいうえお"
        },
        "email": {
          "type": "string",
          "example": "hoge@example.com"
        },
        "isAdmin": {
          "type": "boolean"
        },
        "name": {
          "type": "string"
        },
        "userID": {
          "type": "integer",
          "format": "int32"
        }
      }
    }
  },
  "tags": [
    {
      "description": "ログインとログアウト",
      "name": "sessions"
    }
  ]
}`))
}