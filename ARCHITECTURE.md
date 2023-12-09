# Architecture

- [Architecture](#architecture)
  - [Entities](#entities)
    - [Investor](#investor)
    - [Asset](#asset)
    - [Transaction](#transaction)
    - [Order](#order)
    - [Order Queue](#order-queue)
    - [Book](#book)
  - [Transformers](#transformers)
  - [Concurrency and channels](#concurrency-and-channels)
  - [Testing](#testing)
  - [Error handling](#error-handling)
  - [Logging](#logging)
  - [Infrastructure](#infrastructure)

## Entities

The entities are the core for application, in this are the declarations, methods and core business. The logic about all of these are as following:

### Investor

The `investor` is the primary entity of exchange. This entity describes a client that want to make a trade. In this entity we find your name and asset position, as well as business logic to add and update asset positions

### Asset

The `asset` represents the good to be transacted, this have a name and volume in the market.

### Transaction

The `transaction` represents a market transaction between two interest parties that have movements shares by a determined price and total amount. This entity has a date time as timestamp to identify when occurred.

In addition, in this entity we could find some important business logics to determine orders pending shares (when a order were partially filled), calculate amount of a transaction and interact with a transaction status.

### Order

The `order` is the object that represents the flow of a market share. In this we could found the investor that interacts with it, the asset to be transacted, the shares quantity, pending shares, price of order, the type of order, status and transactions involved to close it.

### Order Queue

The `order_queue` is a auxiliary entity to balance the orders incoming. This entity assign important methods to initiate a heap data structure that will be implemented to organise the `book` of transactions.

### Book

The `book` is the core of transactions made by `investors`. In this entity we define the orders, transactions and channels to balance the orders published by clients.

## Transformers

To provide inputs and outputs for broker and provide an interface with services that will accomplish with, we use the [Data mapper pattern](https://en.wikipedia.org/wiki/Data_mapper_pattern). The mappers are disposed in the `app/market/transformer` and are used to map all inputs and outputs to be processed before sent or recovered by broker.

## Concurrency and channels

The Golang was chosen for this project to provide a build-in consistent and easy way to work with concurrency, multithreading and channels.

## Testing

This project uses the Golang build-in and [Testify](https://github.com/stretchr/testify) packages to provide a solution to write and execute testing to cover application necessities.

## Error handling

> Under development

## Logging

> Under development

## Infrastructure

This project uses the Apache Kafka as broker to consume and publish transactions.
