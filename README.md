## Dream 11 Backend API


## Overview  

This project provides backend APIs for managing users, contests, teams, and players for a fantasy sports application similar to Dream 11. It is built using the Gin web framework in Go (Golang) and integrates with PostgreSQL for data storage.

# Features
- **User Management:**   :Includes user signup and login functionality.
- **Contest Management** :Allows users to join contests and manage contest details.
- **Team Management**    :Enables users to create teams of players for contests.
- **Player Management**  :Provides APIs for creating new players.
- **Wallet Management**  :Supports loading money into user wallets for contest entry fees.


## Technology Stack

- **Golang**: The backend is written in Go (Golang), a statically typed, compiled language.

- **Gin**: The Gin web framework is used to create RESTful APIs and handle HTTP requests.

- **GORM**: ORM library for database interactions.

- **Postgres**: PostgreSQL is an advanced, enterprise class open source relational database that supports both SQL and JSON  querying. 
                It is a highly stable database management system, which has contributed to its high levels of resilience,and correctness. 
   

## Setup

The application will be accessible at `http://localhost:8080`.

##  Usage
 
- **Signup and Login**           : Use /signup and /login endpoints to register and authenticate users.
- **Contest and Team Management**: Create contests, join contests, and create teams using respective endpoints.
- **Player Management**          : Create new players using the /player endpoint.
- **Wallet Management**          : Load money into user wallets using the /loadmoney endpoint.

# Models

## User

- **ID**: `uint`
- **CreatedAt**: `time.Time`
- **UpdatedAt**: `time.Time`
- **DeletedAt**: `gorm.DeletedAt`
- **Name**: `string`
- **Email**: `string`

## Wallet

- **ID**: `uint`
- **CreatedAt**: `time.Time`
- **UpdatedAt**: `time.Time`
- **DeletedAt**: `gorm.DeletedAt`
- **UserID**: `uint`
- **Balance**: `float64`

## Contest

- **ID**: `uint`
- **CreatedAt**: `time.Time`
- **UpdatedAt**: `time.Time`
- **DeletedAt**: `gorm.DeletedAt`
- **Name**: `string`
- **EntryFee**: `float64`

## Player

- **ID**: `uint`
- **CreatedAt**: `time.Time`
- **UpdatedAt**: `time.Time`
- **DeletedAt**: `gorm.DeletedAt`
- **Name**: `string`
- **Team**: `string`
- **CreditScore**: `float64`

## UserTeam

- **ID**: `uint`
- **CreatedAt**: `time.Time`
- **UpdatedAt**: `time.Time`
- **DeletedAt**: `gorm.DeletedAt`
- **UserID**: `uint`
- **ContestID**: `uint`
- **PlayerIDs**: `string`