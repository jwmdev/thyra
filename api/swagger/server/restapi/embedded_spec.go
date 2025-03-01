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
    "http",
    "https"
  ],
  "swagger": "2.0",
  "info": {
    "description": "Thyra HTTP server API.",
    "title": "thyra-server",
    "version": "0.0.0"
  },
  "paths": {
    "/cmd/executeFunction": {
      "post": {
        "produces": [
          "application/json"
        ],
        "operationId": "cmdExecuteFunction",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "type": "object",
              "default": {
                "args": "",
                "at": "A1MrqLgWq5XXDpTBH6fzXHUg7E8M5U2fYDAF3E1xnUSzyZuKpMh",
                "coins": {
                  "parallel": 0,
                  "sequential": 0
                },
                "expiry": 3,
                "fee": 0,
                "gaz": {
                  "limit": 700000000,
                  "price": 0
                },
                "keyId": "default",
                "name": "test"
              },
              "required": [
                "name",
                "at"
              ],
              "properties": {
                "args": {
                  "description": "Arguments to pass to the function.",
                  "type": "string",
                  "default": ""
                },
                "at": {
                  "description": "Smart contract address exporting the function to call.",
                  "type": "string"
                },
                "coins": {
                  "description": "Coins to be send from caller to smart contract address.",
                  "type": "object",
                  "properties": {
                    "parallel": {
                      "description": "Number of parallel coins to transfer from the caller to the smart contract address.",
                      "type": "number",
                      "default": 0
                    },
                    "sequential": {
                      "description": "Number of sequential coins to transfer from the caller to the smart contract address.",
                      "type": "number",
                      "default": 0
                    }
                  }
                },
                "expiry": {
                  "description": "Set the expiry duration (in number of slots) of the transaction.",
                  "type": "integer",
                  "default": 3
                },
                "fee": {
                  "description": "Set the fee amount (in massa) that will be given to the block creator.",
                  "type": "number",
                  "default": 0
                },
                "gaz": {
                  "description": "Gaz attibutes. Gaz is a virtual resource consumed by node while running smart contract.",
                  "type": "object",
                  "properties": {
                    "limit": {
                      "description": "Maximum number of gaz unit that a node will be able consume.",
                      "type": "integer",
                      "default": 700000000
                    },
                    "price": {
                      "description": "Price of a gaz unit.",
                      "type": "number",
                      "default": 0
                    }
                  }
                },
                "keyId": {
                  "description": "Defines the key to used to sign the transaction.",
                  "type": "string",
                  "default": "default"
                },
                "name": {
                  "description": "Function name to call.",
                  "type": "string"
                }
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK.",
            "schema": {
              "description": "Operation id.",
              "type": "string"
            }
          },
          "422": {
            "description": "Unprocessable Entity - syntax is correct, but the server was unable to process the contained instructions.",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "500": {
            "description": "Internal Server Error - The server has encountered a situation it does not know how to handle.",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/kpi": {
      "get": {
        "produces": [
          "application/json"
        ],
        "operationId": "kpi",
        "parameters": [
          {
            "uniqueItems": true,
            "type": "array",
            "items": {
              "enum": [
                "wallet",
                "node",
                "stacking",
                "blockchain"
              ],
              "type": "string"
            },
            "collectionFormat": "csv",
            "name": "scope",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "kpi message.",
            "schema": {
              "type": "object",
              "properties": {
                "node": {
                  "type": "array",
                  "items": {
                    "type": "object",
                    "properties": {
                      "cpu": {
                        "type": "array",
                        "items": {
                          "type": "number"
                        }
                      },
                      "memory": {
                        "type": "array",
                        "items": {
                          "type": "number"
                        }
                      },
                      "network": {
                        "type": "array",
                        "items": {
                          "type": "number"
                        }
                      },
                      "storage": {
                        "type": "array",
                        "items": {
                          "type": "number"
                        }
                      }
                    }
                  }
                },
                "stacking": {
                  "type": "array",
                  "items": {
                    "type": "object",
                    "properties": {
                      "address": {
                        "type": "string"
                      },
                      "gains": {
                        "type": "number"
                      },
                      "rolls": {
                        "type": "integer"
                      },
                      "slashing": {
                        "type": "integer"
                      }
                    }
                  }
                },
                "wallet": {
                  "type": "array",
                  "items": {
                    "type": "object",
                    "properties": {
                      "balance": {
                        "type": "number"
                      },
                      "coin": {
                        "type": "string"
                      }
                    }
                  }
                }
              }
            }
          }
        }
      }
    },
    "/mgmt/wallet": {
      "get": {
        "produces": [
          "application/json"
        ],
        "operationId": "mgmtWalletGet",
        "responses": {
          "200": {
            "description": "Wallets retrieved",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/Wallet"
              }
            }
          },
          "400": {
            "description": "Bad request.",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "500": {
            "description": "Internal Server Error - The server has encountered a situation it does not know how to handle.",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      },
      "put": {
        "produces": [
          "application/json"
        ],
        "operationId": "mgmtWalletImport",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/Wallet"
            }
          }
        ],
        "responses": {
          "204": {
            "description": "Wallet imported."
          },
          "400": {
            "description": "Bad request.",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "422": {
            "description": "Unprocessable Entity - syntax is correct, but the server was unable to process the contained instructions.",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "500": {
            "description": "Internal Server Error - The server has encountered a situation it does not know how to handle.",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      },
      "post": {
        "produces": [
          "application/json"
        ],
        "operationId": "mgmtWalletCreate",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "type": "object",
              "required": [
                "nickname",
                "password"
              ],
              "properties": {
                "nickname": {
                  "description": "Wallet's short name.",
                  "type": "string"
                },
                "password": {
                  "description": "Private key password.",
                  "type": "string"
                }
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "New wallet created.",
            "schema": {
              "$ref": "#/definitions/Wallet"
            }
          },
          "400": {
            "description": "Bad request.",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "422": {
            "description": "Unprocessable Entity - syntax is correct, but the server was unable to process the contained instructions.",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "500": {
            "description": "Internal Server Error - The server has encountered a situation it does not know how to handle.",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/mgmt/wallet/{nickname}": {
      "delete": {
        "produces": [
          "application/json"
        ],
        "operationId": "mgmtWalletDelete",
        "parameters": [
          {
            "type": "string",
            "description": "Wallet's short name.",
            "name": "nickname",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "204": {
            "description": "Wallet removed.",
            "schema": {
              "$ref": "#/definitions/Wallet"
            }
          },
          "400": {
            "description": "Bad request.",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "404": {
            "description": "Bad request.",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "422": {
            "description": "Unprocessable Entity - syntax is correct, but the server was unable to process the contained instructions.",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "500": {
            "description": "Internal Server Error - The server has encountered a situation it does not know how to handle.",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/website/{address}/{resource}": {
      "get": {
        "produces": [
          "media type"
        ],
        "operationId": "websiteGet",
        "parameters": [
          {
            "type": "string",
            "description": "Address containing the website.",
            "name": "address",
            "in": "path",
            "required": true
          },
          {
            "type": "string",
            "default": "index.html",
            "description": "Website resource.",
            "name": "resource",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Resource retrieved."
          },
          "400": {
            "description": "Bad request.",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "404": {
            "description": "Resource not found.",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "500": {
            "description": "Internal Server Error - The server has encountered a situation it does not know how to handle.",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "Error": {
      "description": "Error object.",
      "type": "object",
      "required": [
        "code",
        "message"
      ],
      "properties": {
        "code": {
          "description": "error code.",
          "type": "string"
        },
        "message": {
          "description": "error message.",
          "type": "string"
        }
      }
    },
    "Wallet": {
      "description": "Wallet object (V0).",
      "type": "object",
      "required": [
        "nickname",
        "address",
        "keyPairs"
      ],
      "properties": {
        "address": {
          "description": "wallet's address.",
          "type": "string"
        },
        "keyPairs": {
          "description": "wallet's key pairs.",
          "type": "array",
          "items": {
            "type": "object",
            "required": [
              "privateKey",
              "publicKey",
              "salt",
              "nonce"
            ],
            "properties": {
              "nonce": {
                "description": "Nonce used by the AES-GCM algorithm used to protect the key pair's private key.",
                "type": "string",
                "format": "base58check"
              },
              "privateKey": {
                "description": "Key pair's private key.",
                "type": "string",
                "format": "base58check"
              },
              "publicKey": {
                "description": "Key pair's public key.",
                "type": "string",
                "format": "base58check"
              },
              "salt": {
                "description": "Salt used by the PBKDF that generates the secret key used to protect the key pair's private key.",
                "type": "string",
                "format": "base58check"
              }
            }
          }
        },
        "nickname": {
          "description": "wallet's nickname.",
          "type": "string"
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
    "http",
    "https"
  ],
  "swagger": "2.0",
  "info": {
    "description": "Thyra HTTP server API.",
    "title": "thyra-server",
    "version": "0.0.0"
  },
  "paths": {
    "/cmd/executeFunction": {
      "post": {
        "produces": [
          "application/json"
        ],
        "operationId": "cmdExecuteFunction",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "type": "object",
              "default": {
                "args": "",
                "at": "A1MrqLgWq5XXDpTBH6fzXHUg7E8M5U2fYDAF3E1xnUSzyZuKpMh",
                "coins": {
                  "parallel": 0,
                  "sequential": 0
                },
                "expiry": 3,
                "fee": 0,
                "gaz": {
                  "limit": 700000000,
                  "price": 0
                },
                "keyId": "default",
                "name": "test"
              },
              "required": [
                "name",
                "at"
              ],
              "properties": {
                "args": {
                  "description": "Arguments to pass to the function.",
                  "type": "string",
                  "default": ""
                },
                "at": {
                  "description": "Smart contract address exporting the function to call.",
                  "type": "string"
                },
                "coins": {
                  "description": "Coins to be send from caller to smart contract address.",
                  "type": "object",
                  "properties": {
                    "parallel": {
                      "description": "Number of parallel coins to transfer from the caller to the smart contract address.",
                      "type": "number",
                      "default": 0
                    },
                    "sequential": {
                      "description": "Number of sequential coins to transfer from the caller to the smart contract address.",
                      "type": "number",
                      "default": 0
                    }
                  }
                },
                "expiry": {
                  "description": "Set the expiry duration (in number of slots) of the transaction.",
                  "type": "integer",
                  "default": 3
                },
                "fee": {
                  "description": "Set the fee amount (in massa) that will be given to the block creator.",
                  "type": "number",
                  "default": 0
                },
                "gaz": {
                  "description": "Gaz attibutes. Gaz is a virtual resource consumed by node while running smart contract.",
                  "type": "object",
                  "properties": {
                    "limit": {
                      "description": "Maximum number of gaz unit that a node will be able consume.",
                      "type": "integer",
                      "default": 700000000
                    },
                    "price": {
                      "description": "Price of a gaz unit.",
                      "type": "number",
                      "default": 0
                    }
                  }
                },
                "keyId": {
                  "description": "Defines the key to used to sign the transaction.",
                  "type": "string",
                  "default": "default"
                },
                "name": {
                  "description": "Function name to call.",
                  "type": "string"
                }
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK.",
            "schema": {
              "description": "Operation id.",
              "type": "string"
            }
          },
          "422": {
            "description": "Unprocessable Entity - syntax is correct, but the server was unable to process the contained instructions.",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "500": {
            "description": "Internal Server Error - The server has encountered a situation it does not know how to handle.",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/kpi": {
      "get": {
        "produces": [
          "application/json"
        ],
        "operationId": "kpi",
        "parameters": [
          {
            "minItems": 0,
            "uniqueItems": true,
            "type": "array",
            "items": {
              "enum": [
                "wallet",
                "node",
                "stacking",
                "blockchain"
              ],
              "type": "string"
            },
            "collectionFormat": "csv",
            "name": "scope",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "kpi message.",
            "schema": {
              "type": "object",
              "properties": {
                "node": {
                  "type": "array",
                  "items": {
                    "$ref": "#/definitions/NodeItems0"
                  }
                },
                "stacking": {
                  "type": "array",
                  "items": {
                    "$ref": "#/definitions/StackingItems0"
                  }
                },
                "wallet": {
                  "type": "array",
                  "items": {
                    "$ref": "#/definitions/WalletItems0"
                  }
                }
              }
            }
          }
        }
      }
    },
    "/mgmt/wallet": {
      "get": {
        "produces": [
          "application/json"
        ],
        "operationId": "mgmtWalletGet",
        "responses": {
          "200": {
            "description": "Wallets retrieved",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/Wallet"
              }
            }
          },
          "400": {
            "description": "Bad request.",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "500": {
            "description": "Internal Server Error - The server has encountered a situation it does not know how to handle.",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      },
      "put": {
        "produces": [
          "application/json"
        ],
        "operationId": "mgmtWalletImport",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/Wallet"
            }
          }
        ],
        "responses": {
          "204": {
            "description": "Wallet imported."
          },
          "400": {
            "description": "Bad request.",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "422": {
            "description": "Unprocessable Entity - syntax is correct, but the server was unable to process the contained instructions.",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "500": {
            "description": "Internal Server Error - The server has encountered a situation it does not know how to handle.",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      },
      "post": {
        "produces": [
          "application/json"
        ],
        "operationId": "mgmtWalletCreate",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "type": "object",
              "required": [
                "nickname",
                "password"
              ],
              "properties": {
                "nickname": {
                  "description": "Wallet's short name.",
                  "type": "string"
                },
                "password": {
                  "description": "Private key password.",
                  "type": "string"
                }
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "New wallet created.",
            "schema": {
              "$ref": "#/definitions/Wallet"
            }
          },
          "400": {
            "description": "Bad request.",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "422": {
            "description": "Unprocessable Entity - syntax is correct, but the server was unable to process the contained instructions.",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "500": {
            "description": "Internal Server Error - The server has encountered a situation it does not know how to handle.",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/mgmt/wallet/{nickname}": {
      "delete": {
        "produces": [
          "application/json"
        ],
        "operationId": "mgmtWalletDelete",
        "parameters": [
          {
            "type": "string",
            "description": "Wallet's short name.",
            "name": "nickname",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "204": {
            "description": "Wallet removed.",
            "schema": {
              "$ref": "#/definitions/Wallet"
            }
          },
          "400": {
            "description": "Bad request.",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "404": {
            "description": "Bad request.",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "422": {
            "description": "Unprocessable Entity - syntax is correct, but the server was unable to process the contained instructions.",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "500": {
            "description": "Internal Server Error - The server has encountered a situation it does not know how to handle.",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/website/{address}/{resource}": {
      "get": {
        "produces": [
          "media type"
        ],
        "operationId": "websiteGet",
        "parameters": [
          {
            "type": "string",
            "description": "Address containing the website.",
            "name": "address",
            "in": "path",
            "required": true
          },
          {
            "type": "string",
            "default": "index.html",
            "description": "Website resource.",
            "name": "resource",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Resource retrieved."
          },
          "400": {
            "description": "Bad request.",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "404": {
            "description": "Resource not found.",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "500": {
            "description": "Internal Server Error - The server has encountered a situation it does not know how to handle.",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "CmdExecuteFunctionParamsBodyCoins": {
      "description": "Coins to be send from caller to smart contract address.",
      "type": "object",
      "properties": {
        "parallel": {
          "description": "Number of parallel coins to transfer from the caller to the smart contract address.",
          "type": "number",
          "default": 0
        },
        "sequential": {
          "description": "Number of sequential coins to transfer from the caller to the smart contract address.",
          "type": "number",
          "default": 0
        }
      }
    },
    "CmdExecuteFunctionParamsBodyGaz": {
      "description": "Gaz attibutes. Gaz is a virtual resource consumed by node while running smart contract.",
      "type": "object",
      "properties": {
        "limit": {
          "description": "Maximum number of gaz unit that a node will be able consume.",
          "type": "integer",
          "default": 700000000
        },
        "price": {
          "description": "Price of a gaz unit.",
          "type": "number",
          "default": 0
        }
      }
    },
    "Error": {
      "description": "Error object.",
      "type": "object",
      "required": [
        "code",
        "message"
      ],
      "properties": {
        "code": {
          "description": "error code.",
          "type": "string"
        },
        "message": {
          "description": "error message.",
          "type": "string"
        }
      }
    },
    "NodeItems0": {
      "type": "object",
      "properties": {
        "cpu": {
          "type": "array",
          "items": {
            "type": "number"
          }
        },
        "memory": {
          "type": "array",
          "items": {
            "type": "number"
          }
        },
        "network": {
          "type": "array",
          "items": {
            "type": "number"
          }
        },
        "storage": {
          "type": "array",
          "items": {
            "type": "number"
          }
        }
      }
    },
    "StackingItems0": {
      "type": "object",
      "properties": {
        "address": {
          "type": "string"
        },
        "gains": {
          "type": "number"
        },
        "rolls": {
          "type": "integer"
        },
        "slashing": {
          "type": "integer"
        }
      }
    },
    "Wallet": {
      "description": "Wallet object (V0).",
      "type": "object",
      "required": [
        "nickname",
        "address",
        "keyPairs"
      ],
      "properties": {
        "address": {
          "description": "wallet's address.",
          "type": "string"
        },
        "keyPairs": {
          "description": "wallet's key pairs.",
          "type": "array",
          "items": {
            "$ref": "#/definitions/WalletKeyPairsItems0"
          }
        },
        "nickname": {
          "description": "wallet's nickname.",
          "type": "string"
        }
      }
    },
    "WalletItems0": {
      "type": "object",
      "properties": {
        "balance": {
          "type": "number"
        },
        "coin": {
          "type": "string"
        }
      }
    },
    "WalletKeyPairsItems0": {
      "type": "object",
      "required": [
        "privateKey",
        "publicKey",
        "salt",
        "nonce"
      ],
      "properties": {
        "nonce": {
          "description": "Nonce used by the AES-GCM algorithm used to protect the key pair's private key.",
          "type": "string",
          "format": "base58check"
        },
        "privateKey": {
          "description": "Key pair's private key.",
          "type": "string",
          "format": "base58check"
        },
        "publicKey": {
          "description": "Key pair's public key.",
          "type": "string",
          "format": "base58check"
        },
        "salt": {
          "description": "Salt used by the PBKDF that generates the secret key used to protect the key pair's private key.",
          "type": "string",
          "format": "base58check"
        }
      }
    }
  }
}`))
}
