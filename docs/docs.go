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
            "name": "Dmitrii",
            "email": "ladovod@gmail.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/song": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Create a song",
                "parameters": [
                    {
                        "description": "body",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Song"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/song/{id}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "Retrieves song based on given ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Song ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "404": {
                        "description": "Not Found"
                    }
                }
            },
            "put": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Updates song based on given ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Song ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "body",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/main.Song"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            },
            "delete": {
                "summary": "Delete a song by id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "PAGE",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/song/{id}/verses": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "Retrieves song's verses based on given song ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Song ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "paginating results - ?page=1",
                        "name": "page",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "404": {
                        "description": "Not Found"
                    }
                }
            }
        },
        "/songs": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "Retrieves songs",
                "parameters": [
                    {
                        "type": "string",
                        "description": "paginating results - ?page=1",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "song search - ?title=Some title",
                        "name": "title",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "song search - ?description=Some descr",
                        "name": "description",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "song search - ?author=Some author",
                        "name": "author",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "song search - ?song_group=Some group",
                        "name": "song_group",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "song release_date - ?release_date=1995-07-16",
                        "name": "release_date",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Song"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error"
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
        "main.Song": {
            "type": "object",
            "properties": {
                "Author": {
                    "type": "string",
                    "example": "Song author"
                },
                "Description": {
                    "type": "string",
                    "example": "Song description text"
                },
                "Link": {
                    "type": "string",
                    "example": "https://www.youtube.com/watch?v=b_h8kh-PEfI9999"
                },
                "ReleaseDate": {
                    "type": "string",
                    "example": "2006-02-01T15:04:05Z"
                },
                "SongGroup": {
                    "type": "string",
                    "example": "Song group"
                },
                "Title": {
                    "type": "string",
                    "example": "Song title"
                }
            }
        },
        "model.Song": {
            "type": "object",
            "properties": {
                "Author": {
                    "type": "string",
                    "example": "Song author"
                },
                "Description": {
                    "type": "string",
                    "example": "Song description text"
                },
                "Link": {
                    "type": "string",
                    "example": "https://www.youtube.com/watch?v=b_h8kh-PEfI9999"
                },
                "ReleaseDate": {
                    "type": "string",
                    "example": "2006-02-01T15:04:05Z"
                },
                "SongGroup": {
                    "type": "string",
                    "example": "Song group"
                },
                "Title": {
                    "type": "string",
                    "example": "Song title"
                },
                "Verses": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Verses"
                    }
                },
                "createdAt": {
                    "type": "string"
                },
                "deletedAt": {
                    "$ref": "#/definitions/gorm.DeletedAt"
                },
                "id": {
                    "type": "integer"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "model.Verses": {
            "type": "object",
            "properties": {
                "Text": {
                    "type": "string",
                    "example": "Verse text"
                },
                "createdAt": {
                    "type": "string"
                },
                "deletedAt": {
                    "$ref": "#/definitions/gorm.DeletedAt"
                },
                "id": {
                    "type": "integer"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "Songs library API",
	Description:      "Swagger API for Songs library API.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
