# Cli tools built in Go

### Csv to Json

```
cd ./csvToJson
```
```
go run . [pathToCsvFile]
```
---
### Net tool
```
cd ./netTool
```

List all available params
```
go run . --help
```
Usage example
```
go run . ns --host google.com
```
Options
| Plugin | README |
| ------ | ------ |
| ns | Look name servers for particular Host |
| ip | Look IP's for particular Host |
| cname | Look CNAME's for particular Host |
| mx | Look MX (mail exchange) records for particular Host |

### Pull dependencies
```
go mod tidy
```
to add missing requirements and to drop unused requirements.
#### Build binary 
```
goo build . (in each directory)
```
- [Go Modules references](https://go.dev/ref/mod)
## 🔗 Links
[![linkedin](https://img.shields.io/badge/linkedin-0A66C2?style=for-the-badge&logo=linkedin&logoColor=white)](https://www.linkedin.com/in/dominik-polzer-hi-o/)

[![twitter](https://img.shields.io/badge/twitter-1DA1F2?style=for-the-badge&logo=twitter&logoColor=white)](https://twitter.com/dompolzer/)
