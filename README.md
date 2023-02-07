# **Project: retail-management**
### Quick Run Project:
- Setup: make setup
- Run app: make run

### **Technology**:
- Using Go 1.19
- Mongodb cloud

### Database
- Product: product of store
- Branch: infor about name, address, list of requestId inventory
- Inventory: infor about name, list of product
- Invoice: infor about invoice after sell product

### Project structure
- Workflow: Request => transport => business => storage => Database


### API ENDPOINTS(localhost:12345)
    1. Get List Info:
        Path: http:/localhost:12345/get/info
        Method: GET
        Success: 200
          {
            "result": [
                {
                    "Branch": {
                        "ID": "63df3e04b2d08d6dbd93f45b",
                        "Name": "PhuNhuan",
                        "Address": "",
                        "Inventory": [
                            "I1"
                        ]
                    },
                    "InventoryDetail": [
                        {
                            "ID": "63df3ea0b2d08d6dbd93f45d",
                            "Name": "ABC",
                            "Products": [
                                {
                                    "ID": "000000000000000000000000",
                                    "Name": "hep",
                                    "Amount": 11111,
                                    "OriginalAmount": 0,
                                    "Quantity": 470,
                                    "RequestId": "P1"
                                },
                                {
                                    "ID": "63df3f69b2d08d6dbd93f461",
                                    "Name": "Wine",
                                    "Amount": 75000,
                                    "OriginalAmount": 0,
                                    "Quantity": 130,
                                    "RequestId": "P3"
                                }
                            ],
                            "RequestId": "I1"
                        }
                    ]
                },
                {
                    "Branch": {
                        "ID": "63df3e70b2d08d6dbd93f45c",
                        "Name": "TanPhu",
                        "Address": "",
                        "Inventory": [
                            "I1",
                            "I2"
                        ]
                    },
                    "InventoryDetail": [
                        {
                            "ID": "63df3ea0b2d08d6dbd93f45d",
                            "Name": "ABC",
                            "Products": [
                                {
                                    "ID": "000000000000000000000000",
                                    "Name": "hep",
                                    "Amount": 11111,
                                    "OriginalAmount": 0,
                                    "Quantity": 470,
                                    "RequestId": "P1"
                                },
                                {
                                    "ID": "63df3f69b2d08d6dbd93f461",
                                    "Name": "Wine",
                                    "Amount": 75000,
                                    "OriginalAmount": 0,
                                    "Quantity": 130,
                                    "RequestId": "P3"
                                }
                            ],
                            "RequestId": "I1"
                        },
                        {
                            "ID": "63dfaad4b2d08d6dbd93f46c",
                            "Name": "BCD",
                            "Products": [
                                {
                                    "ID": "000000000000000000000000",
                                    "Name": "Noodle",
                                    "Amount": 5000,
                                    "OriginalAmount": 0,
                                    "Quantity": 2989,
                                    "RequestId": "P1"
                                },
                                {
                                    "ID": "000000000000000000000000",
                                    "Name": "Wine",
                                    "Amount": 75000,
                                    "OriginalAmount": 0,
                                    "Quantity": 4994,
                                    "RequestId": "P2"
                                }
                            ],
                            "RequestId": "I2"
                        }
                    ]
                }
            ]
        }

    2. Sell Product:
        Path: http:/localhost:12345/sell/product
        Method: POST
        Payload:
            {
              "branchName": "TanPhu",
              "products": [
                  {
                      "requestId": "P1",
                      "quantity": 1
                  },
                  {
                      "requestId": "P2",
                      "quantity": 2
                  }
              ]
          }
        Success: 200
          {
            "result": {
                "ID": "000000000000000000000000",
                "CreatedDate": "2023-02-06T23:17:30.252469+07:00",
                "Products": [
                    {
                        "ID": "63df3eebb2d08d6dbd93f45f",
                        "Name": "Noodle",
                        "Amount": 5000,
                        "OriginalAmount": 0,
                        "Quantity": 1,
                        "RequestId": "P1"
                    },
                    {
                        "ID": "63df3f44b2d08d6dbd93f460",
                        "Name": "Beer",
                        "Amount": 13000,
                        "OriginalAmount": 0,
                        "Quantity": 2,
                        "RequestId": "P2"
                    }
                ],
                "TotalAmount": 31000
            }
        }
            
    3. Import product
        Path: http:/localhost:12345/import/product
        Method: POST
        Payload:
            {
              "inventoryId": "I3",
              "products": [
                  {
                      "requestId": "P2",
                      "quantity": 11
                  }
              ]
          }
        Success: 201
            {
              "result": "Success"
            }
    4. Move between two inventory
        Path: http:/localhost:12345/move/product
        Method: Post
        Payload:
            {
              "inputInventory": {
                  "requestId": "I2",
                  "products": [
                      {
                          "requestId": "P1",
                          "quantity": 10
                      },
                      {
                          "requestId": "P2",
                          "quantity": 11
                      }
                  ]
              },
              "outputInventory": "I1"
          }
        Success: 200
            {
                "result": "Success"
            }
  
