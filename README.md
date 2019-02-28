



<div align="center">
<h1>Roquette</h1>
<p>Simple CLI to stress your database.</p>
<a href="https://travis-ci.org/estebgonza/roquette" title="Roquette build">
  <img src="https://travis-ci.org/estebgonza/roquette.svg?branch=master" alt="Build Status"/>
</a>
</div>

## Getting started
Roquette is a Go program available on multiple platforms. Currently, you should build Roquette to use it.

### Prerequisites
- Go
- Git

### Installation
```bash 
go get -u github.com/estebgonza/roquette
```

### Usage
Roquette need two files to stress a database.
#### database.json
```json
{
    "driver": "driverName",
    "connection": {
        "host": "localhost",
        "port": 10000,
        "user": "yourUser",
        "pass": "yourPassword"
    }
}
```
#### plan.json
```json
{
    "name": "My plan",
    "concurrent-level": 1,
    "queries": [{
            "sql": "SHOW TABLES",
            "repeat": 5
        },
        {
            "sql": "SELECT * FROM my_table LIMIT 10",
            "repeat": 2
        }
    ]
}
```

### Run your plan!
```bash
roquette -run
```
