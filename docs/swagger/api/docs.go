// Package api GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package api

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/account/check_account": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Account"
                ],
                "summary": "アカウント確認",
                "parameters": [
                    {
                        "description": "アカウント確認",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/github.com_game-core_gocrafter_api_presentation_request_account.CheckAccount"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/github.com_game-core_gocrafter_api_presentation_response_account.CheckAccount"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/error.Error"
                        }
                    }
                }
            }
        },
        "/account/login_account": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Account"
                ],
                "summary": "アカウントログイン",
                "parameters": [
                    {
                        "description": "アカウントログイン",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/github.com_game-core_gocrafter_api_presentation_request_account.LoginAccount"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/github.com_game-core_gocrafter_api_presentation_response_account.LoginAccount"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/error.Error"
                        }
                    }
                }
            }
        },
        "/account/register_account": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Account"
                ],
                "summary": "アカウント登録",
                "parameters": [
                    {
                        "description": "アカウント登録",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/github.com_game-core_gocrafter_api_presentation_request_account.RegisterAccount"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/github.com_game-core_gocrafter_api_presentation_response_account.RegisterAccount"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/error.Error"
                        }
                    }
                }
            }
        },
        "/login_eward/receive_login_reward": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "LoginReward"
                ],
                "summary": "ログイン報酬受け取り",
                "parameters": [
                    {
                        "description": "ログイン報酬受け取り",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/github.com_game-core_gocrafter_api_presentation_request_loginReward.ReceiveLoginReward"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/github.com_game-core_gocrafter_api_presentation_response_loginReward.ReceiveLoginReward"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/error.Error"
                        }
                    }
                }
            }
        },
        "/login_reward/get_login_reward_model": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "LoginReward"
                ],
                "summary": "ログイン報酬モデル取得",
                "parameters": [
                    {
                        "description": "ログイン報酬モデル取得",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/github.com_game-core_gocrafter_api_presentation_request_loginReward.GetLoginRewardModel"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/github.com_game-core_gocrafter_api_presentation_response_loginReward.GetLoginRewardModel"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/error.Error"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "account.Account": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "shard_key": {
                    "type": "string"
                },
                "token": {
                    "type": "string"
                },
                "uuid": {
                    "type": "string"
                }
            }
        },
        "error.Error": {
            "type": "object",
            "properties": {
                "error_message": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
                }
            }
        },
        "github.com_game-core_gocrafter_api_presentation_request_account.CheckAccount": {
            "type": "object",
            "properties": {
                "shard_key": {
                    "type": "string"
                },
                "uuid": {
                    "type": "string"
                }
            }
        },
        "github.com_game-core_gocrafter_api_presentation_request_account.LoginAccount": {
            "type": "object",
            "properties": {
                "password": {
                    "type": "string"
                },
                "shard_key": {
                    "type": "string"
                },
                "uuid": {
                    "type": "string"
                }
            }
        },
        "github.com_game-core_gocrafter_api_presentation_request_account.RegisterAccount": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                }
            }
        },
        "github.com_game-core_gocrafter_api_presentation_request_loginReward.GetLoginRewardModel": {
            "type": "object",
            "properties": {
                "login_reward_model_name": {
                    "type": "string"
                }
            }
        },
        "github.com_game-core_gocrafter_api_presentation_request_loginReward.ReceiveLoginReward": {
            "type": "object",
            "properties": {
                "account_id": {
                    "type": "integer"
                },
                "login_reward_model_name": {
                    "type": "string"
                },
                "shard_key": {
                    "type": "string"
                },
                "uuid": {
                    "type": "string"
                }
            }
        },
        "github.com_game-core_gocrafter_api_presentation_response_account.CheckAccount": {
            "type": "object",
            "properties": {
                "account": {
                    "$ref": "#/definitions/account.Account"
                },
                "status": {
                    "type": "integer"
                }
            }
        },
        "github.com_game-core_gocrafter_api_presentation_response_account.LoginAccount": {
            "type": "object",
            "properties": {
                "account": {
                    "$ref": "#/definitions/account.Account"
                },
                "status": {
                    "type": "integer"
                }
            }
        },
        "github.com_game-core_gocrafter_api_presentation_response_account.RegisterAccount": {
            "type": "object",
            "properties": {
                "account": {
                    "$ref": "#/definitions/account.Account"
                },
                "status": {
                    "type": "integer"
                }
            }
        },
        "github.com_game-core_gocrafter_api_presentation_response_loginReward.GetLoginRewardModel": {
            "type": "object",
            "properties": {
                "login_reward_model": {
                    "$ref": "#/definitions/loginReward.LoginRewardModel"
                },
                "status": {
                    "type": "integer"
                }
            }
        },
        "github.com_game-core_gocrafter_api_presentation_response_loginReward.ReceiveLoginReward": {
            "type": "object",
            "properties": {
                "login_reward_status": {
                    "$ref": "#/definitions/loginReward.LoginRewardStatus"
                },
                "status": {
                    "type": "integer"
                }
            }
        },
        "loginReward.Event": {
            "type": "object",
            "properties": {
                "end_at": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "repeat_setting": {
                    "type": "boolean"
                },
                "repeat_start_at": {
                    "type": "string"
                },
                "reset_hour": {
                    "type": "integer"
                },
                "start_at": {
                    "type": "string"
                }
            }
        },
        "loginReward.Item": {
            "type": "object",
            "properties": {
                "count": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "loginReward.LoginRewardModel": {
            "type": "object",
            "properties": {
                "event": {
                    "$ref": "#/definitions/loginReward.Event"
                },
                "id": {
                    "type": "integer"
                },
                "login_reward_rewards": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/loginReward.LoginRewardReward"
                    }
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "loginReward.LoginRewardReward": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "items": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/loginReward.Item"
                    }
                },
                "name": {
                    "type": "string"
                },
                "step_number": {
                    "type": "integer"
                }
            }
        },
        "loginReward.LoginRewardStatus": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "items": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/loginReward.Item"
                    }
                },
                "last_received_at": {
                    "type": "string"
                },
                "login_reward_model": {
                    "$ref": "#/definitions/loginReward.LoginRewardModel"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8001",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "github.com/game-core/gocrafter",
	Description:      "This is a sample swagger server.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
