curl --location 'localhost:8080/v1/queue/get' \
--header 'Content-Type: application/json' \
--data-raw '{
    "uuid": "some-connection-uuid",
    "maxsize": 10,
    "users": [
        {
            "userid": "userid-user-1",
            "firstname": "John",
            "lastname": "Doe",
            "email": "tt@tt.com"
        },
        {
            "userid": "userid-user-2",
            "firstname": "John",
            "lastname": "Doe",
            "email": "tt@tt.com"
        }
    ]
}'