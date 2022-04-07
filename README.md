
# Golang path resolver challange app

Simple REST service which receives path tree data in json format, resolve and calculate max value of paths.
## Run locally

Clone the project

```bash
  git clone https://github.com/guneyin/go-path-resolver.git
```

Go project directory

```bash
  cd go-path-resolver
```

Build project

```bash
  go build
```

Run

```bash
  .\go-path-resolver
```

  
## API Usage

#### Request

```http
  POST http://127.0.0.1:3001/
```

#### Body

```json
{
    "tree": {
    "nodes": [
        {"id": "1", "left": "2", "right": "3", "value": 1},
        {"id": "3", "left": null, "right": null, "value": 3},
        {"id": "2", "left": null, "right": null, "value": 2}
    ],
    "root": "1"
    }
}
```

#### Response

```json
{
    "path_sum": 4,
    "node_count": 3,
    "execute_duration": 0
}
```

  
## Testing

Run the following command to run the tests

```bash
  go test
```

  
