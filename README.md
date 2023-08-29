# Cara menjalankan program
-  Run ``go run main.go``
-  Pilih program
# Kompleksitas
Inisialisasi dan Persiapan:

Pembuatan slice completedIndexes berukuran sebanding dengan panjang input string str. Ini memiliki kompleksitas O(n), di mana n adalah panjang string str.
Iterasi Utama:

Ada dua loop bersarang dalam kode ini. Loop pertama iterasi melalui karakter-karakter dalam string str. Ini memiliki kompleksitas O(n), di mana n adalah panjang string str.
Pencarian Bracket:

Di dalam loop utama, Anda memiliki loop kedua (nested loop) yang mencari tanda kurung yang sesuai dengan tanda kurung saat ini. Loop kedua ini juga memiliki kompleksitas O(n), di mana n adalah panjang sisa string yang belum diperiksa.
Operasi Lainnya:

Setiap operasi yang dilakukan dalam loop kedua hanya melibatkan operasi operasi dasar seperti perbandingan, penjumlahan, dan operasi slice. Ini tidak mempengaruhi kompleksitas secara signifikan.
Ketika meninjau keseluruhan kode dan melihat bahwa ada dua loop bersarang dengan kompleksitas O(n) masing-masing, Anda bisa mengalikan kompleksitas keduanya. Dengan begitu, kompleksitas keseluruhan program ini dapat dianggap sebagai O(n^2), di mana n adalah panjang string input.