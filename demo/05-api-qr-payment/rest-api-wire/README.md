
comment mockgen
```
//go:generate mockgen -source=./<filename>.go -destination=./mock_<filename>/mock_<filename>.go -package=mock_<filename>
```

command generate
```bash
go generate ./... 
```