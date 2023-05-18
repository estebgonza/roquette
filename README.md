



<div align="center">
  <h1>roquette</h1>
  <p>Simple CLI to stress your database.</p>
  <a href="https://goreportcard.com/report/github.com/estebgonza/roquette" title="Roquette SQL Stresser Report">
    <img src="https://goreportcard.com/badge/github.com/estebgonza/roquette" alt="Roquette SQL Stresser Report"/>
  </a>
  <a href="https://coveralls.io/github/estebgonza/roquette?branch=master" title="Roquette SQL Stresser Coverage">
    <img src="https://coveralls.io/repos/github/estebgonza/roquette/badge.svg?branch=master" alt="Roquette SQL Stresser Coverage"/>
  </a>
</div>

### Installation
```bash 
go get -u github.com/estebgonza/roquette
```

### Supported Databases
- Hive
- Postgres

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
roquette run
```
