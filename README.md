# Go PIN validation sample
Have 2 versions:
- Config by code
- Config by yml file

## 1. Config by code:
**Run:**
```bash
cd ./config-by-code
go run main.go
```

## 2. Config by yml file:
**Description:** The business logic validation is loaded from the `configs/validationRules.yml` file. Therefore, when we need update this logic, we just update this the yml file, we don't need to change the code

**Run:**
```bash
cd ./config-by-yml
go mod tidy
go run main.go
```