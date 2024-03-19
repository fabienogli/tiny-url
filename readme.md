# How to use
## Requirements
* docker
* docker-compose
* http client

## How to install
### With Makefile
Run the following command
```bash
make up
````
### With Docker-compose
Run the following command
```bash
docker-compose up -d 
````

## Follow executions
### With Makefile
Run the following command
```bash
make logs
```

### With docker-compose
```bash
docker-compose logs -f
```

### Routes
| Method | route      | body                                | Answer                                                                                |
|--------|------------|-------------------------------------|---------------------------------------------------------------------------------------|
| POST   | /          | {"url":"https://my-url-to-shorten"} | 201: {"url":"https://my-url-to-shorten","slug":"Mjg5NDM0"} 500: internal server error |
| GET    | /<my-slug> |                                     | 301:  redirect to my saved url 404: slug unknown                                      |

### How to insert url inside
#### With curl
```bash
curl -X POST -i localhost:8000 -d '{"url": "https://superuser.com/questions/272265/getting-curl-to-output-http-status-code"}'
``` 
# Ressources used

https://hpmahesh.medium.com/creating-a-simple-tiny-url-generator-using-golang-postgresql-and-redis-df8a29f2deab

https://anandjoshi.me/articles/2016-10/URL-Shortening

https://www.eddywm.com/lets-build-a-url-shortener-in-go-part-3-short-link-generation/

https://stackoverflow.com/questions/1562367/how-do-short-urls-services-work

docker  
https://www.mitrais.com/news-updates/how-to-dockerize-a-restful-api-with-golang-and-postgres/