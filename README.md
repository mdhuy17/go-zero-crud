step 0 
```
etcd
```

step 1
```
docker compose up -d
```

step 2
```
cd rpc/service/company_service/
go run companyservice.go -f etc/companyservice.yaml
```

step 3
```
cd rpc/service/user_service/
go run userservice.go -f etc/userservice.yaml
```

step 4
```
cd back to root
go run gateway.go -f etc/gateway-api.yaml

```
