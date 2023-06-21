# postgresStore

Golang library to store objects (byte slices) in Postgres / Timescale.

```golang
import "github.com/mathisve/postgresStore"
```

## Example
```golang
err = c.UploadObject(postgresStore.Object{
	ObjectName: filename,
	Bytes:      data,
})
if err != nil {
	log.Println(err)
}
```