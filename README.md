# sriwulan-flood-early-warning
Aplikasi ini digunakan untuk memprediksi banjir rob yang terjadi di Desa Sriwulan, Kabupaten Demak, Jawa Tengah. Semua data diambil dari website resmi BMKG [title](https://www.bmkg.go.id/).

## API Design
> Data akan diambil melalui web scrapping, jika BMKG tidak menyediakan API cuaca dan juga pasang surut air laut.
1. List URL untuk web scrapping :
   
| Syntax | Description |
| ----------- | ----------- |
|Cuaca|https://www.bmkg.go.id/cuaca/prakiraan-cuaca/33.21.04.2011|
|Tinggi Gelombang|https://www.bmkg.go.id/cuaca/maritim/P.K.04|
