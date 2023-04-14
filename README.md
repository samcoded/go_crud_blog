# SIMPLE BLOG API

Web Restful API for blog with create, read, update and delete functionality

## TOOLS

-   GO
-   Gin (Web server)
-   Gorm ORM with postgresql (Database)

## Dependencies

```

go mod tidy

go run main.go
```

## API Routes

```[GET] /posts/ - Returns all posts

    [GET] /posts/{id} - Return post with ID

    [POST] /posts/ - Create and return post
    @body: {title, content}

    [PUT] /posts/{ID} - Update post with ID and return post
    @body: {title, content}

    [DELETE] /posts/{id} - Delete post with ID
```

## Live Demo

https://go-crud-blog.herokuapp.com/
