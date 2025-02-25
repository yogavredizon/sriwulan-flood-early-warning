# sriwulan-flood-early-warning
Aplikasi ini digunakan untuk memprediksi banjir rob yang terjadi di Desa Sriwulan, Kabupaten Demak, Jawa Tengah. Semua data diambil dari website resmi [BMKG](https://www.bmkg.go.id/).

> Data akan diambil melalui web scrapping, jika BMKG tidak menyediakan API cuaca dan juga pasang surut air laut.

*List URL untuk web scrapping* :

| Keterangan | URL |
| ----------- | ----------- |
|Cuaca|https://www.bmkg.go.id/cuaca/prakiraan-cuaca/33.21.04.2011|
|Tinggi Gelombang|https://www.bmkg.go.id/cuaca/maritim/P.K.04|

> *33.21.04.2011* merupakan kode wilayah desa sriwulan

> *P.K.04* merupakan kode perairan semarang - demak

## API Design
| Method | URL |
| ----------- | ----------- |
| GET | /sriwulan/api/weather |
| GET | /sriwulan/api/weather/{date} |
| GET | /sriwulan/api/tides |
| GET | /sriwulan/api/tides/{dates} |

1. Mendapatkan daftar cuaca
> `GET /sriwulan/api/weather'

Response : 

```json
{
  "code" : 200,
  "message" : "Berhasil mengambil data cuaca",
  "data" : [
              {
                "date" :"2025-02-25;19:19:1",
                "degree" : 25,
                "status" : "Berawan",
                "hummidity" : 93,
                "Wind" : 3,
                "Visibillity" : 11.27",
              },
                            {
                "date" :"2025-02-25;20:10:10",
                "degree" : 25,
                "status" : "Berawan",
                "hummidity" : 93,
                "Wind" : 3,
                "Visibillity" : 11.27",
              },
  ]
}```
 
