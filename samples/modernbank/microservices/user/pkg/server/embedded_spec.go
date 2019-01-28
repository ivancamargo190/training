// Code generated by go-swagger; DO NOT EDIT.

package server

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
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "This is the User Microservice, responsible for managing Users in Modern Bank.",
    "title": "User",
    "version": "1.0.0"
  },
  "host": "users",
  "basePath": "/v1",
  "paths": {
    "/users": {
      "post": {
        "description": "Creates ",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "users"
        ],
        "summary": "Create a new user identity for a customer",
        "operationId": "createUser",
        "parameters": [
          {
            "description": "Created user object",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/User"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Created",
            "schema": {
              "$ref": "#/definitions/User"
            }
          },
          "409": {
            "description": "User alreadys exists"
          },
          "500": {
            "description": "Internal server error"
          }
        }
      }
    },
    "/users/{username}": {
      "get": {
        "produces": [
          "application/json"
        ],
        "tags": [
          "users"
        ],
        "summary": "Get user by user name",
        "operationId": "getUserByUserName",
        "parameters": [
          {
            "type": "string",
            "description": "The name that needs to be fetched.",
            "name": "username",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Success!",
            "schema": {
              "$ref": "#/definitions/User"
            }
          },
          "404": {
            "description": "User not found"
          },
          "500": {
            "description": "Internal server error"
          }
        }
      },
      "put": {
        "description": "Update user by username.",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "users"
        ],
        "summary": "Update user",
        "operationId": "updateUser",
        "parameters": [
          {
            "type": "string",
            "description": "name that need to be updated",
            "name": "username",
            "in": "path",
            "required": true
          },
          {
            "description": "Updated user object",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/User"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Success!",
            "schema": {
              "$ref": "#/definitions/User"
            }
          },
          "404": {
            "description": "User not found"
          },
          "500": {
            "description": "Internal server error"
          }
        }
      },
      "delete": {
        "description": "Delete user by username.",
        "tags": [
          "users"
        ],
        "summary": "Delete user",
        "operationId": "deleteUser",
        "parameters": [
          {
            "type": "string",
            "description": "The name that needs to be deleted",
            "name": "username",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Success!"
          },
          "404": {
            "description": "User not found"
          },
          "500": {
            "description": "Internal server error"
          }
        }
      }
    }
  },
  "definitions": {
    "User": {
      "type": "object",
      "required": [
        "username",
        "firstName",
        "lastName",
        "email",
        "password"
      ],
      "properties": {
        "email": {
          "type": "string"
        },
        "firstName": {
          "type": "string"
        },
        "lastName": {
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
  },
  "tags": [
    {
      "description": "Operations about users",
      "name": "user"
    }
  ]
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "This is the User Microservice, responsible for managing Users in Modern Bank.",
    "title": "User",
    "version": "1.0.0"
  },
  "host": "users",
  "basePath": "/v1",
  "paths": {
    "/users": {
      "post": {
        "description": "Creates ",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "users"
        ],
        "summary": "Create a new user identity for a customer",
        "operationId": "createUser",
        "parameters": [
          {
            "description": "Created user object",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/User"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Created",
            "schema": {
              "$ref": "#/definitions/User"
            }
          },
          "409": {
            "description": "User alreadys exists"
          },
          "500": {
            "description": "Internal server error"
          }
        }
      }
    },
    "/users/{username}": {
      "get": {
        "produces": [
          "application/json"
        ],
        "tags": [
          "users"
        ],
        "summary": "Get user by user name",
        "operationId": "getUserByUserName",
        "parameters": [
          {
            "type": "string",
            "description": "The name that needs to be fetched.",
            "name": "username",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Success!",
            "schema": {
              "$ref": "#/definitions/User"
            }
          },
          "404": {
            "description": "User not found"
          },
          "500": {
            "description": "Internal server error"
          }
        }
      },
      "put": {
        "description": "Update user by username.",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "users"
        ],
        "summary": "Update user",
        "operationId": "updateUser",
        "parameters": [
          {
            "type": "string",
            "description": "name that need to be updated",
            "name": "username",
            "in": "path",
            "required": true
          },
          {
            "description": "Updated user object",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/User"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Success!",
            "schema": {
              "$ref": "#/definitions/User"
            }
          },
          "404": {
            "description": "User not found"
          },
          "500": {
            "description": "Internal server error"
          }
        }
      },
      "delete": {
        "description": "Delete user by username.",
        "tags": [
          "users"
        ],
        "summary": "Delete user",
        "operationId": "deleteUser",
        "parameters": [
          {
            "type": "string",
            "description": "The name that needs to be deleted",
            "name": "username",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Success!"
          },
          "404": {
            "description": "User not found"
          },
          "500": {
            "description": "Internal server error"
          }
        }
      }
    }
  },
  "definitions": {
    "User": {
      "type": "object",
      "required": [
        "username",
        "firstName",
        "lastName",
        "email",
        "password"
      ],
      "properties": {
        "email": {
          "type": "string"
        },
        "firstName": {
          "type": "string"
        },
        "lastName": {
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
  },
  "tags": [
    {
      "description": "Operations about users",
      "name": "user"
    }
  ]
}`))
}
