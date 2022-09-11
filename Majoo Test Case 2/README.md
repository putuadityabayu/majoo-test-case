# Majoo Test Cases 2

Terdapat block code sbb:

Terdapat model area dengan struct dibawah ini
```go
type (
  Area struct {
    ID int64 `gorm:"column:id;primaryKey;"`
    AreaValue int64 `gorm:"column:area_value"`
    AreaType string `gorm:"column:type"`
  }
)
```

Handler memiliki function seperti dibawah ini:

```go
func (_r *AreaRepository) InsertArea(param1 int32, param2 int64, type []string, ar *Model.Area)
  (err error) {
  inst := _r.DB.Model(ar)
  Var area int
  area = 0
  switch type {
  case ‘persegi panjang’:
    var area := param1 * param2
    ar.AreaValue = area
    ar.AreaType = ‘persegi panjang’
    err = _r.DB.create(&ar).Error
    if err != nil {
      return err
    }
  case ‘persegi’:
    var area = param1 * param2
    ar.AreaValue = area
    ar.AreaType = ‘persegi’
    err = _r.DB.create(&ar).Error
    if err != nil {
      return err
    }
  case segitiga:
    area = 0.5 * (param1 * param2)
    ar.AreaValue = area
    ar.AreaType = ‘segitiga’
    err = _r.DB.create(&ar).Error
    if err != nil {
      return err
    }
  default:
    ar.AreaValue = 0
    ar.AreaType = ‘undefined data’
    err = _r.DB.create(&ar).Error
    if err != nil {
      return err
    }
  }
}
```

Dan dipanggil dengan cara berikut:

```go
err = _u.repository.InsertArea(10, 10, ‘persegi’)
if err != nil {
  log.Error().Msg(err.Error())
  err = errors.New(en.ERROR_DATABASE)
  return err
}
```

## Dari keterangan diatas, apa saja yang harus diperbaiki dan berikan contoh konkritnya!

---

## JAWABAN

* Penjelasan jawaban dapat dilihat pada file [Majoo Test Case 2.pdf](MAJOO%20TEST%20CASE%202.pdf)
* Contoh konkrit perbaikan dapat dilihat pada file [main.go](main.go)