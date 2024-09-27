`# Payment Gateway API

## Overview

This project implements a simple payment gateway API that allows merchants to process payments and retrieve payment details. It includes a bank simulator to test the payment flow.

## Table of Contents

- [How to Run the Solution](#how-to-run-the-solution)
- [Prerequisites](#prerequisites)
- [Installation and Running](#installation-and-running)
- [API Endpoints](#api-endpoints)
- [Process a Payment](#process-a-payment)
- [Retrieve Payment Details](#retrieve-payment-details)
- [Assumptions Made](#assumptions-made)
- [Areas for Improvement](#areas-for-improvement)
- [Running Tests](#running-tests)
- [Contact Information](#contact-information)

## How to Run the Solution

### Prerequisites

- **Go** installed (version 1.18 or higher). Download from [golang.org](https://golang.org/dl/).

### Installation and Running

1. **Clone the repository:**

   ```bash
   git clone https://github.com/MikkelAllison/payment-gateway.git
   cd payment-gateway
   ```

2. **Install dependencies:**

    ```bash
    go mod download
    ```
3. **Run the application:**

    ```bash
    go run cmd/main.go
    ```
   
4. **The API is now available at `http://localhost:8080`.**

API Endpoints
-------------

### Process a Payment

-   **URL:** `/payments`

-   **Method:** `POST`

-   **Headers:**
    -   `Content-Type: application/json`
    -   `Merchant-ID: your_merchant_id`
    
- **Request Body:**

    ```json
    {
    "card_number": "4242424242424242",
    "expiry_month": 12,
    "expiry_year": 2025,
    "amount": 1000,
    "currency": "GBP",
    "cvv": "123"
    }
    ```
    
-   **Response:**

    -   **Status Code:** `200 OK`

    -   **Body:**

    ```json
    {
      "id": "txn_xxxxxxxxxx",
      "status": "approved"
    }
    ```
    

**Example using cURL:**

```bash
curl -X POST http://localhost:8080/payments \
-H "Content-Type: application/json" \
-H "Merchant-ID: merchant_123" \
-d '{
  "card_number": "4242424242424242",
  "expiry_month": 12,
  "expiry_year": 2025,
  "amount": 1000,
  "currency": "GBP",
  "cvv": "123"
}'
```

### Retrieve Payment Details

-   **URL:** `/payments/{id}`

-   **Method:** `GET`

-   **Headers:**

    -   `Content-Type: application/json`
-   **Response:**

    -   **Status Code:** `200 OK`

    -   **Body:**

        ```json
        {
        "id": "txn_xxxxxxxxxx",
        "card_number": "**** **** **** 4242",
        "expiry_month": 12,
        "expiry_year": 2025,
        "amount": 1000,
        "currency": "GBP",
        "status": "approved"
        }
        ```
**Example using cURL:**

```bash
curl -X GET http://localhost:8080/payments/txn_xxxxxxxxxx \
-H "Content-Type: application/json" \
-H "Merchant-ID: merchant_123"
```

Replace `txn_xxxxxxxxxx` with the actual transaction ID received from the payment processing response.

Assumptions Made
----------------

-   **In-Memory Storage:** Payments are stored in-memory using a Go map. Data will be lost when the application stops.
-   **Simplified Validation:**
    -   Card number length is checked to be 16 digits but not validated using a more complex method (e.g. the Luhn algorithm).
    -   The CVV is collected during payment processing but is not stored or returned in any API responses to enhance security.
-   **Security:**
    -   **Sensitive Data Masking:** Card numbers are masked when retrieved.
    -   **No Encryption:** Sensitive data is not encrypted in memory.
-   **Authentication:** Not implemented. Merchant identification is simulated via the `Merchant-ID` header.
-   **Bank Simulator:** The bank simulator always approves payments for testing purposes.

Areas for Improvement
---------------------

-   **Storage:** Implement a database (e.g., PostgreSQL, MongoDB) to store payments across application restarts.
-   **Enhanced Validation:**
    -   Use the Luhn algorithm to validate card numbers.
    -   Validate CVV length and format.
    -   Validate expiration dates against the current date.
-   **Security Enhancements:**
    -   Implement proper authentication and authorization mechanisms.
-   **Error Handling:**
    -   Provide more detailed error messages without exposing sensitive information.
-   **Testing:**
    -   Increase test coverage, including unit, integration, and end-to-end tests.
    -   Implement tests for further edge cases and error scenarios.



Running Tests
-------------

To run the tests for the application, execute:

```bash
go test ./...
```

Contact Information
-------------------

For any questions or support, please contact:

-   **Name:** Mikkel Allison
-   **Email:** mikkel_a01@hotmail.com