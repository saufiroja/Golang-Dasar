## Type Data Integer (1)

| Type Data | Nilai Minimum        | Nilai Maksimum      |
| --------- | -------------------- | ------------------- |
| int8      | -128                 | 127                 |
| int16     | -32768               | 32767               |
| int32     | -2147483648          | 2147483647          |
| int64     | -9223372036854775808 | 9223372036854775807 |

## Type Data Integer (2)

| Type Data | Nilai Minimum | Nilai Maksimum       |
| --------- | ------------- | -------------------- |
| uint8     | 0             | 225                  |
| uint16    | 0             | 65535                |
| uint32    | 0             | 4294967295           |
| uint64    | 0             | 18446744073709551615 |

## Type Data Floating Point

| Type Data  | Nilai Minimum                                         | Nilai Maksimum        |
| ---------- | ----------------------------------------------------- | --------------------- |
| float32    | 1.18 x 10 pangkat -38                                 | 3.4 x 10 pangkat 38   |
| float64    | 2.23 x 10 pangkat -308                                | 1.80 x 10 pangkat 308 |
| complex64  | complex numbers with float32 real and imaginary parts |
| complex128 | complex numbers with float64 real and imaginary parts |

## Alias

| Type Data | Alias untuk    |
| --------- | -------------- |
| byte      | uint8          |
| rune      | int32          |
| int       | Minimal int32  |
| uint      | Minimal uint32 |

## Operasi && (Dan)

| Nilai 1 | Operator | Nilai 2 | Hasil |
| ------- | -------- | ------- | ----- |
| true    | &&       | true    | true  |
| true    | &&       | false   | false |
| false   | &&       | true    | false |
| false   | &&       | false   | false |

## Operasi || (Atau)

| Nilai 1 | Operator | Nilai 2 | Hasil |
| ------- | -------- | ------- | ----- |
| true    | atau     | true    | true  |
| true    | atau     | false   | true  |
| false   | atau     | true    | true  |
| false   | atau     | false   | false |

## Operasi !

| Operator | Nilai 2 | Hasil |
| -------- | ------- | ----- |
| !        | true    | false |
| !        | false   | true  |

## Membuat slice dari array

| Membuat slice   | keterangan                                                             |
| --------------- | ---------------------------------------------------------------------- |
| array[low:high] | Membuat slice dari array dimulai index low sampai index sebelum akhir  |
| array[low:]     | Membuat slice dari array dimulai index low sampai index akhir di array |
| array[:high]    | Membuat slice dari array dimulai index 0 sampai index sebelum akhir    |
| array[:]        | Membuat slice dari array dimulai index 0 sampai index akhir di array   |

## Operator Pointer

- (&dan) adalah untuk membuat sebuah variable dengan nilai pointer ke varibale yang lain, kita bisa menggunakan operator (&dan) diikuti dengan nama variable nya
- (*bintang) adalah jika kita ingin mengubah seluruh varibale yang mengacu pada data tersebut, kita bisa menggunakan operator (*bintang)
