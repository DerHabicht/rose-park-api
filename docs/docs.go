// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Robert Hawk",
            "email": "robert@the-hawk.us"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/health": {
            "get": {
                "description": "Healthcheck endpoint. Reports which statuses are currently\nrunning and the current API\\'s version number. If critical\nservices are running, it will return 200. If any of the\ncritical services are down, then the endpoint will return 503.",
                "summary": "Check to assure that the service is running.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers.ServiceStatus"
                        }
                    },
                    "503": {
                        "description": "Service Unavailable",
                        "schema": {
                            "$ref": "#/definitions/controllers.ServiceStatus"
                        }
                    }
                }
            }
        },
        "/sites": {
            "get": {
                "description": "Lists all blogs that are managed by this backend. Authentication is required to use this endpoint.",
                "summary": "List all blog sites managed by this backend.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Blog"
                            }
                        }
                    },
                    "204": {
                        "description": "No Content",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/controllers.ControllerError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/controllers.ControllerError"
                        }
                    }
                }
            },
            "post": {
                "description": "Creates a new blog site to be managed by this backend. Authentication is required to use this endpoint.",
                "summary": "Create a new blog site.",
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.Blog"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/controllers.ControllerError"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/controllers.ControllerError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/controllers.ControllerError"
                        }
                    }
                }
            }
        },
        "/sites/{domain}": {
            "get": {
                "description": "Retrieve the blog, all authors that write for this blog and their bios, and the ten most recent post\ntitles (with their slugs) published on this blog. Authentication is not required for this endpoint.",
                "summary": "Fetch data regarding this blog.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Blog"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/controllers.ControllerError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/controllers.ControllerError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/controllers.ControllerError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "controllers.ControllerError": {
            "type": "object",
            "additionalProperties": true
        },
        "controllers.ServiceStatus": {
            "type": "object",
            "properties": {
                "errors": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "status": {
                    "type": "object",
                    "additionalProperties": {
                        "type": "boolean"
                    }
                },
                "version": {
                    "type": "string"
                }
            }
        },
        "models.Author": {
            "type": "object",
            "properties": {
                "bio": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "posts": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Post"
                    }
                }
            }
        },
        "models.Blog": {
            "type": "object",
            "properties": {
                "authors": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Author"
                    }
                },
                "domain": {
                    "type": "string"
                },
                "name": {
                    "description": "The name of this blog.",
                    "type": "string"
                },
                "posts": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Post"
                    }
                }
            }
        },
        "models.Post": {
            "type": "object",
            "properties": {
                "authors": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Author"
                    }
                },
                "body": {
                    "type": "string"
                },
                "publish_date": {
                    "type": "string"
                },
                "slug": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "0.1.0+0",
	Host:        "https://the-hawk.us",
	BasePath:    "/apis/blogs/v1",
	Schemes:     []string{},
	Title:       "THUS Blogs Backend",
	Description: "UPDATE DESCRIPTION FIELD",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
