# CloudQuery cloudquery-struct-nullable-fields Source Plugin

Repo created to demonstrate a potential bug in `transformers.TransformWithStruct` function in the cloudquery SDK.  
When providing native fields (non-pointer), the expected outcome would be non-nullable fields in the destination.  
When providing nullable fields (poitners), the expected outcome would be nullable fields in the destination.

## Screenshot

![Image](https://github.com/user-attachments/assets/8e925655-2f25-4088-b995-2ea4ec2c52e3)

## Reproducing steps

1. Run `go build . && cloudquery sync .`
2. Check nullability of fields with a SQLite db viewer