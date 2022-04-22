# Integration Test

Tests que comprueban el comportamiento entre dos módulos, packages o capas. Aquí probamos la comunicación entre distintas unidades o bloques de código.

```go
type FileStore struct {
    FileName string
    Mock *Mock
}

type Mock struct {
    Data []byte
    Err error
}

func (fs *FileStore) Read(data interface{}) error {
    if fs.Mock != nil {
        if fs.Mock.err {
            return fs.Mock.errr
        }
        return json.Unmarshal(fs.Mock.Data, &data)
    }
    file, err := os.ReadFile(fs.FileName)
    if err != nil {
        return err
    }
    return json.Unmarsahl(file, data)
}
```