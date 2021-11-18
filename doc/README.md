# Introduction

go-rest was built for the simple basic api implementation for the developers. For the different endpoints, easily start the development for your frontend application using this api. We hope you enjoy these docs, and please don't hesitate to file an issue if you see anything missing.

# Use Cases

There are many reasons to use the go-rest API. The most common use case is to create crud application for many senarios, so that you can build custom application whatever you want
using this api.

# Authorization

All API requests doesn't require the use of a generated API key. You don'T need API key, or generate a new one, just send your request. All API requests welcome.
Be able to be start the development easily, we designed the open api.

# Responses

Many API endpoints return the JSON representation of the resources created or edited.

```json
[{"id":"1","name":"Kebap","isSpicy":"Yes"},{"id":"2","name":"Pide","isSpicy":"No"}]
```



# Status Codes

go-rest returns the following status codes in its API:

| Status Code | Description |
| :--- | :--- |
| 200 | `OK` |
| 201 | `CREATED` |
| 400 | `BAD REQUEST` |
| 404 | `NOT FOUND` |
| 500 | `INTERNAL SERVER ERROR` |
