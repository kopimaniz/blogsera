## handle null from db

di package sql atau di msql driver ada type Null yang mengimpelement Scan

## value vs pointer

- jika struct sebaiknya selalu gunakan pointer bagaik di paramater atau return, dengan memberi return pointer maka kita bisa memberi nilai nil pada return

## how to log

## peran layer

### repository

- repository menghubungkan data dari layer diatasnya ke database
- tidak ada validasi data di layer ini

### service

- layer service tidak tergantung dengan layer dibawahnya, misal di sql ada `Err.NoRows` , layer service tidak boleh membaca itu, jika perlu handling error maka buatkan common error yang general, tidak terikat dengan suatu library
- layer service tidak menerima data dari layer diatasnya, misal layer diatas service adalah http Handler, maka layer service tidak perlu menerim input `http.ResponseWriter` atau `http.Request`

## handler

- handler bertugas memparsing data dari luar, dan mengarahkan sesuai dengan service yang digunakan
- dihandler tidak ada pengolahan data dan validasi

## bagaimana penamaan file yang baik

## bagaimana membaut time.Time mengikuti omitempty pada struct

gunakan pointer *time.Time sql scan tidak peduli dengan data pointer, jika data adalah pointer maka akan otomatis mengisi value dari pointer
