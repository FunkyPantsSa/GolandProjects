package Other

u := [...]int{11, 12, 13, 14, 15}
fmt.Println("array:", u) // [11, 12, 13, 14, 15]
s := u[1:3]
fmt.Printf("slice(len=%d, cap=%d): %v\n", len(s), cap(s), s) // [12, 13]
s = append(s, 24) // [12, 13,  24]

fmt.Println("after append 24, array:", u)
fmt.Printf("after append 24, slice(len=%d, cap=%d): %v\n", len(s), cap(s), s) // 3 4
s = append(s, 25) // [12, 13,  24, 25]

fmt.Println("after append 25, array:", u)
fmt.Printf("after append 25, slice(len=%d, cap=%d): %v\n", len(s), cap(s), s) // 4 4
s = append(s, 26)
fmt.Println("after append 26, array:", u) // [12, 13,  24, 25, 26]
fmt.Printf("after append 26, slice(len=%d, cap=%d): %v\n", len(s), cap(s), s) // 5 8

s[0] = 22
fmt.Println("after reassign 1st elem of slice, array:", u)
fmt.Printf("after reassign 1st elem of slice, slice(len=%d, cap=%d): %v\n", len(s), cap(s), s) //5 8   //// [22, 13,  24, 25, 26]
