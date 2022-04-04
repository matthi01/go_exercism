package strand

/*
	DNA -> 	RNA
	G 	-> 	C
	C 	->	G
	T 	-> 	A
	A 	-> 	U
*/

func ToRNA(dna string) string {
	var rna []rune
	for _, nucleotide := range dna {
		switch nucleotide {
		case 'G':
			rna = append(rna, 'C')
		case 'C':
			rna = append(rna, 'G')
		case 'T':
			rna = append(rna, 'A')
		case 'A':
			rna = append(rna, 'U')
		default:
			return ""
		}
	}
	return string(rna)
}
