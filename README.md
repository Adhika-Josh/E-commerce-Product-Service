# E-commerce Product Service

The E-commerce Product Service is a microservice built using Go and the Gin framework. It provides APIs for managing products and administrators in an e-commerce system.

## Table of Contents

- [Installation](#installation)
- [Usage](#usage)
- [API Endpoints](#api-endpoints)
- [Testing](#testing)
- [Contributing](#contributing)

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/your-username/e-commerce-product-service.git
   ```

2. Navigate to the project directory:

   ```bash
   cd e-commerce-product-service
   ```

3. Install dependencies:

   ```bash
   go mod tidy
   ```

## Usage

1. Start the server:

   ```bash
   go run main.go
   ```

2. The server will start at `http://localhost:8080` by default.

## API Endpoints

### Add Product

- **Method**: `POST`
- **URL**: `/product-service/v1/add-product`
- **Description**: Adds a new product to the system.
- **Request Body**:

  ```json
  {
      "item_name": "Product Name",
      "item_category": "Product Category",
      "item_quantity": "Product Quantity",
      "item_price": "Product Price"
  }
  ```

- **Response**: 

  - Status: `200 OK`
  - Content: 

    ```json
    {
        "status": 200,
        "message": "Product added successfully"
    }
    ```

### Get All Products

- **Method**: `GET`
- **URL**: `/product-service/v1/get-all-products`
- **Description**: Retrieves a list of all products in the system.
- **Response**: 

  - Status: `200 OK`
  - Content: 

    ```json
    [
        {
            "item_id": 1,
            "item_name": "Product Name",
            "item_category": "Product Category",
            "item_quantity": "Product Quantity",
            "item_price": "Product Price"
        },
        {
            "item_id": 2,
            "item_name": "Product Name",
            "item_category": "Product Category",
            "item_quantity": "Product Quantity",
            "item_price": "Product Price"
        }
    ]
    ```

### Update Product

- **Method**: `PATCH`
- **URL**: `/product-service/v1/update-product`
- **Description**: Updates the details of an existing product.
- **Request Body**:

  ```json
  {
      "item_id": 1,
      "item_name": "New Product Name",
      "item_category": "New Product Category",
      "item_quantity": "New Product Quantity",
      "item_price": "New Product Price"
  }
  ```

- **Response**: 

  - Status: `200 OK`
  - Content: 

    ```json
    {
        "status": 200,
        "message": "Product updated successfully"
    }
    ```

### Create Admin

- **Method**: `POST`
- **URL**: `/product-service/v1/create-admin`
- **Description**: Creates a new administrator account.
- **Request Body**:

  ```json
  {
      "name": "Admin Name",
      "username": "Admin Username",
      "password": "Admin Password"
  }
  ```

- **Response**: 

  - Status: `200 OK`
  - Content: 

    ```json
    {
        "status": 200,
        "message": "Admin account created successfully"
    }
    ```

### Login Admin

- **Method**: `POST`
- **URL**: `/product-service/v1/login-admin`
- **Description**: Logs in an administrator.
- **Request Body**:

  ```json
  {
      "username": "Admin Username",
      "password": "Admin Password"
  }
  ```

- **Response**: 

  - Status: `200 OK`
  - Content: 

    ```json
    {
        "status": 200,
        "message": "Login successful",
        "admin_pid": "Admin PID"
    }
    ```

## Testing

To run tests, execute the following command:

```bash
go test ./...
```

## Contributing

Contributions are welcome! If you'd like to contribute to this project, please follow these steps:

1. Fork the repository.
2. Create your feature branch (`git checkout -b feature/your-feature-name`).
3. Commit your changes (`git commit -am 'Add some feature'`).
4. Push to the branch (`git push origin feature/your-feature-name`).
5. Create a new Pull Request.

