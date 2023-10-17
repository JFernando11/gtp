## GovTech Procurement

This project was made to fulfill recruitment task at GTP

## Setup

Before running the application, make sure to install Go Language and MySQL in your environment.

## Build

To build the Go application, run the following command:

```bash
go build
```

## Run

After building the application, you can run it using the following command:

```bash
go run main.go
```

## API Endpoints

- Create a new product: POST /products/
- Get all products: GET /products/
- Filter all products by SKU: GET /products/?sku={sku}
- Filter all products by title: GET /products/?title={title}
- Filter all products by category: GET /products/?category={category}
- Filter all products by etalase: GET /products/?etalase={etalase}
- Sort all products by date (newest): GET /products/?sort=newest
- Sort all products by date (oldest): GET /products/?sort=oldest
- Sort all products by average rating (highest): GET /products/?sort=highest-rated
- Sort all products by average rating (lowest): GET /products/?sort=lowest-rated
- Get a product by specific ID: GET /product/{id}
- Update a product by specific ID: PUT /product/{id}
- Delete a product by specific ID: DELETE /product/{id}
- Add a review for a product by specific ID: POST /reviews/{id}
- Get all reviews for a product by specific ID: GET /reviews/{id}
- Add an image for a product by specific ID: POST /images/{id}
- Get all images for a product by specific ID: GET /images/{id}
