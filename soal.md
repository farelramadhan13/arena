Arena, Tahap 1: Pembukaan Pertandingan

Cerita. Di arena ada satu monster bernama Lord. Begitu pertandingan dibuka, semua hero menyerangnya sekaligus. Tidak ada yang menunggu giliran, dan tiap hero punya ritme serangnya sendiri. Ada yang cepat tapi lemah, ada yang lambat tapi sakit.

Kebutuhan:

Minimal 3 hero. Tiap hero punya nama, damage, dan jeda antar serangan yang berbeda.
Setiap hero menyerang beberapa kali (misalnya 5). Tiap serangan mencetak satu baris yang memuat nama hero, damage, dan serangan ke berapa.
Semua hero mulai bersamaan, bukan satu per satu. Buktinya, output tercampur antar hero, dan hero bejeda cepat tampil lebih sering di awal.
Program mencetak kalimat pembuka sebelum pertandingan dan kalimat penutup setelah semua serangan selesai.
Cetak jumlah goroutine yang hidup sebelum dan sesudah pertandingan.
Kode dipisah ke minimal dua file: satu untuk titik masuk program, satu untuk hero. Hero harus berupa tipe sendiri (struct) dengan perilaku sendiri (method), bukan sekumpulan variabel lepas.

Batasan tahap ini: belum boleh channel, Mutex, WaitGroup, atau atomic. Untuk menunggu pertandingan selesai, pakai cara paling sederhana yang Anda tahu. Cara itu sengaja akan kita ganti di Tahap 4, dan kenapa harus diganti akan Anda sadari sendiri.

Kriteria lulus:

go run . berjalan tanpa error.
go vet ./... bersih, dan go run -race . tidak mengeluarkan peringatan.
Dijalankan tiga kali, urutan baris tidak selalu sama.
Kodenya sudah diformat dengan go fmt ./....

Pertanyaan konsep (jawab dengan kalimat, tanpa kode):

Kenapa urutan output berbeda tiap kali dijalankan?
Apa yang terjadi pada hero yang belum selesai menyerang kalau main sudah selesai? Kenapa?
Berapa jumlah goroutine sebelum dan sesudah pertandingan, dan kenapa angkanya demikian?
Kenapa menunggu dengan jeda waktu tetap itu cara yang rapuh? Beri satu skenario konkret di mana cara itu gagal.

Bonus (opsional): beri tiap serangan peluang 20% menjadi serangan kritis dengan damage berlipat, dan tandai di output.

Yang perlu Anda kirim
Isi semua file .go Anda
Output dari tiga kali go run .
Output go run -race . dan go vet ./...
Jawaban empat pertanyaan konsep di atas
