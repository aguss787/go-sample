# hoho haha

## Migration

```bash
migrate create -dir migrations -ext sql addHubberTable
migrate -source file://migrations -database "postgresql://glints:glints@localhost:5432/glints?sslmode=disable" up 1
migrate -source file://migrations -database "postgresql://glints:glints@localhost:5432/glints?sslmode=disable" down 1
```

## Sample graphql

playground at: localhost:8080

```graphql
{
  hubbers {
    code
    id
    name
  }
}
```