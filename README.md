# aes-encryption


Have two endpoints one that will take a string the can be encrypted and an id with which it can be accessed,
Sample curl call is as below:

To run the application do : make run


To store:
Id has to be more of length 16, 24, 32
```
curl -X POST \
  http://localhost:8080/store/ \
  -H 'cache-control: no-cache' \
  -H 'content-type: application/json' \
  -d '{
	"data": "In some cases, it may be desirable to store binary data such as media files in BLOB or TEXT columns. You may find MySQL'\''s string handling functions useful for working with such data",
	"id": "eeee22234dsfsdfd"
}'

```

To retrieve:
You will get key back in response of the store endpoint
```
curl -X GET \
  http://localhost:8080/get \
  -H 'cache-control: no-cache' \
  -H 'id: eeee22234dsfsdfd' \
  -H 'key: hTHctcuAxhxKQFDa' \
 ```