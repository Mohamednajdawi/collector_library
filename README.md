# Collector Library App

A professional, minimalist collector library for Amiibo, featuring a Go backend (Clean Architecture) and a React+TypeScript frontend.

## Prerequisites

- Go 1.21+
- Node.js 18+
- Supabase Project (PostgreSQL)

## Setup Instructions

### 1. Database Setup

1.  Log in to your Supabase Dashboard.
2.  Go to the SQL Editor.
3.  Run the contents of `backend/db/schema.sql`.

### 2. Backend Setup

1.  Navigate to the `backend` directory:
    ```bash
    cd backend
    ```
2.  Create a `.env` file (copy from example or create new):
    ```bash
    echo "DATABASE_URL=postgres://user:password@host:port/dbname" > .env
    ```
    *Replace with your actual Supabase connection string (Transaction Mode).*

3.  **Ingest Data**:
    Run the seeder script to populate your database with the scraped Amiibo data.
    ```bash
    go run cmd/seeder/main.go
    ```

4.  **Start Server**:
    ```bash
    go run cmd/server/main.go
    ```
    The server will start on `http://localhost:8080`.

### 3. Frontend Setup

1.  Navigate to the `frontend` directory:
    ```bash
    cd frontend
    ```
2.  Install dependencies:
    ```bash
    npm install
    ```
3.  Start the development server:
    ```bash
    npm run dev
    ```
4.  Open your browser to the URL shown (usually `http://localhost:5173`).

## Features

- **Data Scraper**: Custom scraping logic to get high-quality Amiibo data.
- **Clean Architecture**: Go backend organized by Domain, Ports, and Adapters.
- **Premium UI**: Glassmorphism design system using vanilla CSS.
