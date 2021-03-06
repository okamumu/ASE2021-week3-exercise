package main

var r2adata map[string]int
var rlabel1 []string
var rlabel10 []string
var rlabel100 []string
var rlabel1000 []string

func init() {
	r2adata = make(map[string]int)
	r2adata["I"] = 1
	r2adata["II"] = 2
	r2adata["III"] = 3
	r2adata["IV"] = 4
	r2adata["V"] = 5
	r2adata["VI"] = 6
	r2adata["VII"] = 7
	r2adata["VIII"] = 8
	r2adata["IX"] = 9

	r2adata["X"] = 10
	r2adata["XX"] = 20
	r2adata["XXX"] = 30
	r2adata["XL"] = 40
	r2adata["L"] = 50
	r2adata["LX"] = 60
	r2adata["LXX"] = 70
	r2adata["LXXX"] = 80
	r2adata["XC"] = 90

	r2adata["C"] = 100
	r2adata["CC"] = 200
	r2adata["CCC"] = 300
	r2adata["CD"] = 400
	r2adata["D"] = 500
	r2adata["DC"] = 600
	r2adata["DCC"] = 700
	r2adata["DCCC"] = 800
	r2adata["CM"] = 900

	r2adata["M"] = 1000
	r2adata["MM"] = 2000
	r2adata["MMMM"] = 3000
	r2adata["MMMMM"] = 4000

	rlabel1 = []string{"VIII", "III", "VII", "II", "IV", "VI", "IX", "I", "V"}
	rlabel10 = []string{"LXXX", "XXX", "LXX", "XX", "XL", "LX", "XC", "X", "L"}
	rlabel100 = []string{"DCCC", "CCC", "DCC", "CC", "CD", "DC", "CM", "C", "D"}
	rlabel1000 = []string{"MMMM", "MMM", "MM", "M"}
}

func r2a(x string) int {
	y := 0
	for _, u := range rlabel1000 {
		if len(x) >= len(u) && x[0:len(u)] == u {
			y += r2adata[u]
			x = x[len(u):len(x)]
			break
		}
	}
	for _, u := range rlabel100 {
		if len(x) >= len(u) && x[0:len(u)] == u {
			y += r2adata[u]
			x = x[len(u):len(x)]
			break
		}
	}
	for _, u := range rlabel10 {
		if len(x) >= len(u) && x[0:len(u)] == u {
			y += r2adata[u]
			x = x[len(u):len(x)]
			break
		}
	}
	for _, u := range rlabel1 {
		if len(x) >= len(u) && x[0:len(u)] == u {
			y += r2adata[u]
			x = x[len(u):len(x)]
			break
		}
	}
	if len(x) != 0 {
		return 0
	} else {
		return y
	}
}
