# 💳 Wallet System (Go + gRPC + JWT + YooKassa SDK)

A backend system written in Go using a microservice architecture with gRPC communication.  
The project simulates a simple financial system with authentication via jwt, payments via YooKassa, and balance management.

---
## 🛠 Technology Stack

![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![Fiber](https://img.shields.io/badge/Fiber-v2.0-00ADD8?style=for-the-badge&logo=gofiber&logoColor=white)
![GORM](https://img.shields.io/badge/GORM-316192?style=for-the-badge&logo=gorm&logoColor=white)
![gRPC](https://img.shields.io/badge/gRPC-4285F4?style=for-the-badge&logo=grpc&logoColor=white)
![Protobuf](https://img.shields.io/badge/Protocol_Buffers-E3741C?style=for-the-badge&logo=google-protobuf&logoColor=white)
![PostgreSQL](https://img.shields.io/badge/PostgreSQL-316192?style=for-the-badge&logo=postgresql&logoColor=white)
![Docker](https://img.shields.io/badge/Docker-2496ED?style=for-the-badge&logo=docker&logoColor=white)


## 🧠 Architecture Overview

The system consists of **gRPC module + main backend API which both work in grpc client and server mode simultaneously for creating and confirming payment**:

* Main Backend (HTTP + Auth + API + Users with balance)
↓ (gRPC)
* Payment Service (YooKassa integration and balance + transactions) (gRPC)



---

## 🧱 Services

### 🟢 Main Backend
Responsible for:
- User registration & login
- JWT authentication (access + refresh tokens)
- HTTP API (Fiber)
- Creating payment requests
- User balance management

---

### 🟡 Payment Service (gRPC)
Responsible for:
- Creating payment sessions via YooKassa
- Handling payment webhooks
- Tracking payment status
- Communicating payment result to main backend

---


## 👤 User Flow
* Registration: \
User sends: POST /api/v1/new:
```
{
    "first_name":"John",
    "last_name":"Pork",
    "username": "Johny1821",
    "email":"john12@gmail.com",
    "password":"hello1212"
}

```
Reply:
```

{
    "message": "successfully registered"
}
```

* Login:\
User sends: GET /api/v1/login:
```
{
    "email":"john12@gmail.com",
    "password":"hello1212"
}

```

Reply:
```
{
    "id": "cf074995-c549-40d5-b7d7-da59318cf967",
    "first_name": "John",
    "last_name": "Pork",
    "username": "Johny1821",
    "balance": 0,
    "email": "john12@gmail.com"
}

and tokens in cookies:
refresh_token: eyJhbGciOiJIUzI1NiIs...5iYcU
access_token: eyJhbGciOiJIUzI1NiIs...-cFhC7U
```


## 💳 Payment Flow

1. User sends: POST /api/v1/auth/payment { "price": 129 } with cookies

2. Main backend sends request via grpc to payment module

3. Payment module sends request to yookassa and return response:

```
{
    "result": "https://yoomoney.ru/checkout/payments/v2/contract?orderId=318d8984-000f-5000-b000-14607dcab9c2"
}
```

4. After confirming the payment in your personal yookassa account webhook updates payment status in payment Service.  
Then Payment Service triggers main backend via gRPC

5. User sends: GET /api/v1/auth/get with cookies
Reply:
```
{
    "id": "cf074995-c549-40d5-b7d7-da59318cf967",
    "first_name": "John",
    "last_name": "Pork",
    "username": "Johny1821",
    "balance": 129,
    "email": "john12@gmail.com"
}
```

## 🚦 Getting Started

### Prerequisites
*   Docker & Docker Compose
*   Go 1.21+ (for local development)

### Configuration
1.  **Clone the repository:**
    ```bash
    git clone https://github.com/SH1roV12/BalanceApp.git
    ```
2.  **Environment Setup:** Define your .env file.

3. Run docker compose
```
docker compose up
```


4. Start ngrok 
```
ngrok http HTTP_YOOKASSA_PORT 
```
and paste the link in your personal account in yookassa 

5. Use routes:
```
POST	/api/v1/new	
POST	/api/v1/login	
GET	/api/v1/users	
GET	/api/v1/refresh	
POST	/api/v1/auth/payment	JWT Middleware(send with cookies)
GET	/api/v1/auth/get	JWT Middleware(send with cookies)
```