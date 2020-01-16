package piscine

import "github.com/01-edu/z01"

func EightQueens() {
	for i := '1'; i <= '8'; i++ {
		for j := '1'; j <= '8'; j++ {
			if j != i && j != i+1 && j != i-1 {
				for k := '1'; k <= '8'; k++ {
					if k != j && k != j+1 && k != j-1 && k != i && k != i+2 && k != i-2 {
						for l := '1'; l <= '8'; l++ {
							if l != k && l != k+1 && l != k-1 && l != j && l != j+2 && l != j-2 && l != i && l != i+3 && l != i-3 {
								for m := '1'; m <= '8'; m++ {
									if m != l && m != l+1 && m != l-1 && m != k && m != k+2 && m != k-2 && m != j && m != j+3 && m != j-3 && m != i && m != i+4 && m != i-4 {
										for n := '1'; n <= '8'; n++ {
											if n != m && n != m+1 && n != m-1 && n != l && n != l+2 && n != l-2 && n != k && n != k+3 && n != k-3 && n != j && n != j+4 && n != j-4 && n != i && n != i+5 && n != i-5 {
												for p := '1'; p <= '8'; p++ {
													if p != n && p != n+1 && p != n-1 && p != m && p != m+2 && p != m-2 && p != l && p != l+3 && p != l-3 && p != k && p != k+4 && p != k-4 && p != j && p != j+5 && p != j-5 && p != i && p != i+6 && p != i-6 {
														for q := '1'; q <= '8'; q++ {
															if q != p && q != p+1 && q != p-1 && q != n && q != n+2 && q != n-2 && q != m && q != m+3 && q != m-3 && q != l && q != l+4 && q != l-4 && q != k && q != k+5 && q != k-5 && q != j && q != j+6 && q != j-6 && q != i && q != i+7 && q != i-7 {
																z01.PrintRune(i)
																z01.PrintRune(j)
																z01.PrintRune(k)
																z01.PrintRune(l)
																z01.PrintRune(m)
																z01.PrintRune(n)
																z01.PrintRune(p)
																z01.PrintRune(q)
																z01.PrintRune('\n')
															}
														}
													}
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
}
