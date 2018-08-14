# Email Read Receipts

## Commands

### Bring up app
```bash
source err.sh
err gub run
```

### Insert Name into DB
```bash
curl --header "Content-Type: application/json"   --request POST   --data '{"name":"321","key":"123"}' localhost:5000/name
```

### Fetch Img
```bash
curl localhost:5000/img/XXXX.png
```