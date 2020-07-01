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
    "description": "Repository of recipe information",
    "title": "recipe-rolodex",
    "version": "0.0.1"
  },
  "basePath": "/rolodex",
  "paths": {
    "/ready": {
      "get": {
        "tags": [
          "ready"
        ],
        "summary": "verify that the API is running and ready to accept requests",
        "responses": {
          "200": {
            "description": "post successful responses",
            "schema": {
              "$ref": "#/definitions/ReturnCode"
            }
          },
          "500": {
            "description": "internal service error",
            "schema": {
              "$ref": "#/definitions/ReturnCode"
            }
          },
          "default": {
            "description": "unexpected error",
            "schema": {
              "$ref": "#/definitions/ReturnCode"
            }
          }
        }
      }
    },
    "/recipe": {
      "get": {
        "description": "Retrieve details about a recipe given a recipeID\n",
        "tags": [
          "recipe"
        ],
        "summary": "get details about a recipe",
        "parameters": [
          {
            "type": "integer",
            "description": "a recipe ID",
            "name": "recipeID",
            "in": "query",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "successful",
            "schema": {
              "$ref": "#/definitions/Recipe"
            }
          },
          "400": {
            "description": "bad request",
            "schema": {
              "$ref": "#/definitions/ReturnCode"
            }
          },
          "404": {
            "description": "not found",
            "schema": {
              "$ref": "#/definitions/ReturnCode"
            }
          },
          "500": {
            "description": "internal service error",
            "schema": {
              "$ref": "#/definitions/ReturnCode"
            }
          },
          "default": {
            "description": "unexpected error",
            "schema": {
              "$ref": "#/definitions/ReturnCode"
            }
          }
        }
      },
      "post": {
        "description": "Add a new set of recipe details to the database\n",
        "tags": [
          "recipe"
        ],
        "summary": "create a new recipe",
        "parameters": [
          {
            "description": "new recipe details",
            "name": "newRecipe",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/NewRecipe"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "successful",
            "schema": {
              "$ref": "#/definitions/ReturnCode"
            }
          },
          "400": {
            "description": "bad request",
            "schema": {
              "$ref": "#/definitions/ReturnCode"
            }
          },
          "500": {
            "description": "internal service error",
            "schema": {
              "$ref": "#/definitions/ReturnCode"
            }
          },
          "default": {
            "description": "unexpected error",
            "schema": {
              "$ref": "#/definitions/ReturnCode"
            }
          }
        }
      }
    },
    "/recipes": {
      "get": {
        "description": "Retrieve a list of recipes based on the user-generated query\n",
        "tags": [
          "recipes"
        ],
        "summary": "get a list of recipes",
        "parameters": [
          {
            "type": "string",
            "default": "",
            "description": "an ingredient to filter recipes by",
            "name": "ingredient1",
            "in": "query"
          },
          {
            "type": "string",
            "default": "",
            "description": "an ingredient to filter recipes by",
            "name": "ingredient2",
            "in": "query"
          },
          {
            "type": "string",
            "default": "",
            "description": "an ingredient to filter recipes by",
            "name": "ingredient3",
            "in": "query"
          },
          {
            "type": "string",
            "default": "",
            "description": "a season to filter recipes by",
            "name": "season",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "successful",
            "schema": {
              "$ref": "#/definitions/Recipes"
            }
          },
          "404": {
            "description": "not found",
            "schema": {
              "$ref": "#/definitions/ReturnCode"
            }
          },
          "500": {
            "description": "internal service error",
            "schema": {
              "$ref": "#/definitions/ReturnCode"
            }
          },
          "default": {
            "description": "unexpected error",
            "schema": {
              "$ref": "#/definitions/ReturnCode"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "NewRecipe": {
      "type": "object",
      "required": [
        "season",
        "title",
        "author",
        "link",
        "ingredientList"
      ],
      "properties": {
        "author": {
          "type": "string"
        },
        "ingredientList": {
          "type": "array",
          "items": {
            "type": "string"
          },
          "example": [
            "apples",
            "bananas",
            "peaches"
          ]
        },
        "link": {
          "type": "string"
        },
        "season": {
          "type": "string"
        },
        "title": {
          "type": "string"
        }
      }
    },
    "Recipe": {
      "type": "object",
      "properties": {
        "author": {
          "type": "string"
        },
        "id": {
          "type": "integer"
        },
        "link": {
          "type": "string"
        },
        "season": {
          "type": "string"
        },
        "title": {
          "type": "string"
        }
      }
    },
    "Recipes": {
      "type": "object",
      "properties": {
        "recipeList": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/Recipe"
          }
        }
      }
    },
    "ReturnCode": {
      "type": "object",
      "properties": {
        "code": {
          "type": "integer",
          "readOnly": true
        },
        "message": {
          "type": "string",
          "readOnly": true
        }
      }
    }
  }
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
    "description": "Repository of recipe information",
    "title": "recipe-rolodex",
    "version": "0.0.1"
  },
  "basePath": "/rolodex",
  "paths": {
    "/ready": {
      "get": {
        "tags": [
          "ready"
        ],
        "summary": "verify that the API is running and ready to accept requests",
        "responses": {
          "200": {
            "description": "post successful responses",
            "schema": {
              "$ref": "#/definitions/ReturnCode"
            }
          },
          "500": {
            "description": "internal service error",
            "schema": {
              "$ref": "#/definitions/ReturnCode"
            }
          },
          "default": {
            "description": "unexpected error",
            "schema": {
              "$ref": "#/definitions/ReturnCode"
            }
          }
        }
      }
    },
    "/recipe": {
      "get": {
        "description": "Retrieve details about a recipe given a recipeID\n",
        "tags": [
          "recipe"
        ],
        "summary": "get details about a recipe",
        "parameters": [
          {
            "type": "integer",
            "description": "a recipe ID",
            "name": "recipeID",
            "in": "query",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "successful",
            "schema": {
              "$ref": "#/definitions/Recipe"
            }
          },
          "400": {
            "description": "bad request",
            "schema": {
              "$ref": "#/definitions/ReturnCode"
            }
          },
          "404": {
            "description": "not found",
            "schema": {
              "$ref": "#/definitions/ReturnCode"
            }
          },
          "500": {
            "description": "internal service error",
            "schema": {
              "$ref": "#/definitions/ReturnCode"
            }
          },
          "default": {
            "description": "unexpected error",
            "schema": {
              "$ref": "#/definitions/ReturnCode"
            }
          }
        }
      },
      "post": {
        "description": "Add a new set of recipe details to the database\n",
        "tags": [
          "recipe"
        ],
        "summary": "create a new recipe",
        "parameters": [
          {
            "description": "new recipe details",
            "name": "newRecipe",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/NewRecipe"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "successful",
            "schema": {
              "$ref": "#/definitions/ReturnCode"
            }
          },
          "400": {
            "description": "bad request",
            "schema": {
              "$ref": "#/definitions/ReturnCode"
            }
          },
          "500": {
            "description": "internal service error",
            "schema": {
              "$ref": "#/definitions/ReturnCode"
            }
          },
          "default": {
            "description": "unexpected error",
            "schema": {
              "$ref": "#/definitions/ReturnCode"
            }
          }
        }
      }
    },
    "/recipes": {
      "get": {
        "description": "Retrieve a list of recipes based on the user-generated query\n",
        "tags": [
          "recipes"
        ],
        "summary": "get a list of recipes",
        "parameters": [
          {
            "type": "string",
            "default": "",
            "description": "an ingredient to filter recipes by",
            "name": "ingredient1",
            "in": "query"
          },
          {
            "type": "string",
            "default": "",
            "description": "an ingredient to filter recipes by",
            "name": "ingredient2",
            "in": "query"
          },
          {
            "type": "string",
            "default": "",
            "description": "an ingredient to filter recipes by",
            "name": "ingredient3",
            "in": "query"
          },
          {
            "type": "string",
            "default": "",
            "description": "a season to filter recipes by",
            "name": "season",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "successful",
            "schema": {
              "$ref": "#/definitions/Recipes"
            }
          },
          "404": {
            "description": "not found",
            "schema": {
              "$ref": "#/definitions/ReturnCode"
            }
          },
          "500": {
            "description": "internal service error",
            "schema": {
              "$ref": "#/definitions/ReturnCode"
            }
          },
          "default": {
            "description": "unexpected error",
            "schema": {
              "$ref": "#/definitions/ReturnCode"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "NewRecipe": {
      "type": "object",
      "required": [
        "season",
        "title",
        "author",
        "link",
        "ingredientList"
      ],
      "properties": {
        "author": {
          "type": "string"
        },
        "ingredientList": {
          "type": "array",
          "items": {
            "type": "string"
          },
          "example": [
            "apples",
            "bananas",
            "peaches"
          ]
        },
        "link": {
          "type": "string"
        },
        "season": {
          "type": "string"
        },
        "title": {
          "type": "string"
        }
      }
    },
    "Recipe": {
      "type": "object",
      "properties": {
        "author": {
          "type": "string"
        },
        "id": {
          "type": "integer"
        },
        "link": {
          "type": "string"
        },
        "season": {
          "type": "string"
        },
        "title": {
          "type": "string"
        }
      }
    },
    "Recipes": {
      "type": "object",
      "properties": {
        "recipeList": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/Recipe"
          }
        }
      }
    },
    "ReturnCode": {
      "type": "object",
      "properties": {
        "code": {
          "type": "integer",
          "readOnly": true
        },
        "message": {
          "type": "string",
          "readOnly": true
        }
      }
    }
  }
}`))
}
