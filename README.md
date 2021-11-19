# Products-Api Documentation

products-api adalah api untuk menampilkan 

products-api dibuat menggunakan desain clean architecture karena dengan menggunakan desain clean architecture lebih mudah dalam menangani error atau bug dan juga maintenance.

## General info

products-api menggunakan database MySQL dan menggunakan bahasa pemrograman Go.

## APIs

```
base URL : localhost:8080

POST /product
'Accept: application/json' http://localhost:8080/product

Input example:
{
  "name": "gayung",
  "price": 60000,
  "description": "gayung ajaib",
  "quantity": 10
}

Response 200:
{
    "code": 200,
    "message": "success",
    "data": [
        {
            "id": 1,
            "created_at": "2021-11-19T14:26:34.465+07:00",
            "name": "gayung",
            "price": "8000",
            "description": "gayung ajaib",
            "quantity": "8"
        },
        {
            "id": 2,
            "created_at": "2021-11-19T14:42:10.215+07:00",
            "name": "gayung",
            "price": "8000",
            "description": "gayung ajaib",
            "quantity": "8"
        },
        {
            "id": 3,
            "created_at": "2021-11-19T14:42:43.039+07:00",
            "name": "baju",
            "price": "2000",
            "description": "baju ajaib",
            "quantity": "8"
        },
        {
            "id": 4,
            "created_at": "2021-11-19T14:43:35.863+07:00",
            "name": "anduk",
            "price": "2000",
            "description": "anduk ajaib",
            "quantity": "8"
        }
    ]
}

Response 500:
{
    "code": 500,
    "message": "internal server error",
    "error": "dial tcp [::1]:3306: connectex: No connection could be made because the target machine actively refused it."
}

GET api/products
'Accept: application/json' http://localhost:8080/api/products

query:
- sortBy
- sortType

sort requirement(add in query params):
- default/produk terbaru: sortBy = created_at, sortType = DESC
- produk termurah: sortBy = price, sortType = ASC
- produk termahal: sortBy = price, sortType = DESC
- sort by product name (A-Z): sortBy = name, sortType = ASC
- sort by product name (Z-A): sortBy = name, sortType = DESC

example hit API:
localhost:8080/product?sortBy=name&sortType=desc

Response example:
{
    "code": 200,
    "message": "success",
    "data": [
        {
            "id": 1,
            "created_at": "2021-11-19T14:26:34.465+07:00",
            "name": "gayung",
            "price": "8000",
            "description": "gayung ajaib",
            "quantity": "8"
        },
        {
            "id": 2,
            "created_at": "2021-11-19T14:42:10.215+07:00",
            "name": "gayung",
            "price": "8000",
            "description": "gayung ajaib",
            "quantity": "8"
        },
        {
            "id": 3,
            "created_at": "2021-11-19T14:42:43.039+07:00",
            "name": "baju",
            "price": "2000",
            "description": "baju ajaib",
            "quantity": "8"
        },
        {
            "id": 4,
            "created_at": "2021-11-19T14:43:35.863+07:00",
            "name": "anduk",
            "price": "2000",
            "description": "anduk ajaib",
            "quantity": "8"
        }
    ]
}
```
