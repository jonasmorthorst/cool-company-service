# Cool company service

A small service to list companies.

The project consists of two sub-project:
- A REST-Api written in Golang - located at /api
- A Vue.JS client - located at /client

Hosted on Heroku
- https://cool-company-service-client.herokuapp.com
- https://cool-company-service-api.herokuapp.com
## Running the projects
The projects can easily be run locally with docker compose. The following command will spawn 4 docker images; the api, the client, a mysql db and an adminer.

```
docker-compose up
```

Make sure to have docker installed and running and docker-compose installed.

The following docker images will then be running - all on localhost:

| Image | Port |
| ------------- | ------------- |
| API  | 8081  |
| Client  | 8082  |
| MySQL  | 3308  |
| Adminer  | 8083  |

The MySQL DB will have the following credentials:

| Desc | Value |
| :------------- | :------------- |
| MYSQL_USER  | root  |
| MYSQL_PASSWORD  | root  |
| MYSQL_DB  | dev_db  |
| MYSQL_HOST  | mysql-dev  |
| MYSQL_PORT  | 3306  |

The API has the following endpoints:

| Method | Url | Description |
| :--- | :---- | :------|
| GET  | /api/v1/companies  | Gets all companies|
| GET  | /api/v1/companies/:id  | Gets a specific company |
| POST  | /api/v1/companies  | Create a company |
| PATCH  | /api/v1/companies/:id  | Update a company |
| DELETE  | /api/v1/companies/:id  | Delete a company |
| POST  | /api/v1/companies/:id/owners  | Add an owner to a company |
| DELETE  | /api/v1/companies/:id/owners/:ownerId  | Remove an owner from a company |

When creating/updating a company, you send following properties

| Property | Required | Type | 
| :------------- | :------------- |:---|
| name  | x  | string |
| companyID  | x  | string |
| email  |   | string |
| phone  |   | string |
| address  | x  | string |
| zipCode  | x  | string |
| city  | x  | string |
| country  | x  | string |

When adding an owner, you send following properties

| Property | Required | Type | 
| :------------- | :------------- |:---|
| name  | x  | string |
| title  | x  | string |

### Examples
Creating a company
```
curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"name":"Test company","companyId":"123","email":"test@test.com","phone":"12345678","address":"Test address","zipCode":"1234","city":"Test city","country":"Denmark"}' \
  https://cool-company-service-api.herokuapp.com/api/v1/companies
```

Adding an owner
```
curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"name":"Test owner","title":"Boss"}' \
  https://cool-company-service-api.herokuapp.com/api/v1/companies/1/owners
```


