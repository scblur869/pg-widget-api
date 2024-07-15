curl --location 'localhost:4000/api/v1/addNew' \
--header 'Content-Type: application/json' \
--data '{
    "name": "stripe01",
    "category": "american",
    "color" : "red"
}'

curl --location 'localhost:4000/api/v1/addNew' \
--header 'Content-Type: application/json' \
--data '{
    "name": "stripe02",
    "category": "american",
    "color" : "white"
}'

curl --location 'localhost:4000/api/v1/addNew' \
--header 'Content-Type: application/json' \
--data '{
    "name": "stars01",
    "category": "american",
    "color" : "blue"
}'

curl --location 'localhost:4000/api/v1/addNew' \
--header 'Content-Type: application/json' \
--data '{
    "name": "stars02",
    "category": "american",
    "color" : "white"
}'

curl --location 'localhost:4000/api/v1/addNew' \
--header 'Content-Type: application/json' \
--data '{
    "name": "pole01",
    "category": "american",
    "color" : "silver"
}'

curl --location 'localhost:4000/api/v1/getAll'
