# ERD
The ERD of this application. 
![ERD](/images/erd.png)
1. Products table as master product.
2. Stocks table that has relationship with products to record product stocks.
3. Inbounds table to record incoming product
4. Outbounds table to record outgoing product from an order.

# Use Case
1. Product : 
    - Add Product (to add product to database)
    - Find Product By ID (to find product from database by product ID)
    - Get All Products (to get all existing product in database) 

2. Stock :
    - Add Stock (to add product stock)
    - Get Latest Product Stock (to retreive latest record of product stock, sorted by expiry date, if the product is not expireable, it will be sorted by inbound date)
    - Get Product Total Stock (to get total stock of specific product)
    - Update Product Stock (to update product stock)

3. Inbound :
    - Add Inbound (to add inbound record to the database)

4. Outbound :
    - Add Outbound (to add outbound record to the database)

# Tech Stack
1. [Golang](https://golang.org) ft. [Gin Gonic](https://github.com/gin-gonic/gin)
2. [MySQL](https://www.mysql.com/)

# API Docs (Postman)
https://documenter.getpostman.com/view/4340319/TW6zFSMC