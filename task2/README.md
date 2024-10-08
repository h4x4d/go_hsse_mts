# Second Go Homework
### Test coverage: 78% (all except error handling)


## Server:

### GET: /version
Gives:
```json
{
  "version": "x.x.x"
}
```
(0.1.0 now)

### POST: /decode
With body:
```json
{
  "inputString": $base64String
}
```
Gives:
```json
{
  "outputString": $decodedString
}
```
### GET: /hard-op
Gives:
```json
{
  "status": "success" or "fail",
  "sleepTime": 10 <= int < 20
}
```
with http code 200 or 500 chosen randomly

## Client
Gives same opportunities with Go functions, but stop /hard-op when it lasts more than 15 seconds.