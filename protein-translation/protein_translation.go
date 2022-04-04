package protein

import (
	"errors"
)

var (
	ErrInvalidBase error = errors.New("invalid base encountered")
	ErrStop        error = errors.New("stop sequence encountered")
)

func FromRNA(rna string) ([]string, error) {
	var codons []string
	var codon string
	var counter int = 0
	for i := 0; i < len(rna); i++ {
		codon += string(rna[i])
		counter++
		if counter == 3 {
			codons = append(codons, codon)
			codon = ""
			counter = 0
		}
	}

	var proteins []string
	for _, codon := range codons {
		protein, err := FromCodon(codon)
		if err != nil {
			if err == ErrStop {
				return proteins, nil
			}
			return nil, err
		}
		proteins = append(proteins, protein)
	}
	return proteins, nil
}

/*
	Codon								Protein
	AUG									Methionine
	UUU, UUC						Phenylalanine
	UUA, UUG						Leucine
	UCU, UCC, UCA, UCG	Serine
	UAU, UAC						Tyrosine
	UGU, UGC						Cysteine
	UGG									Tryptophan
	UAA, UAG, UGA				STOP
*/
func FromCodon(codon string) (string, error) {
	switch codon {
	case "AUG":
		return "Methionine", nil
	case "UUU", "UUC":
		return "Phenylalanine", nil
	case "UUA", "UUG":
		return "Leucine", nil
	case "UCU", "UCC", "UCA", "UCG":
		return "Serine", nil
	case "UAU", "UAC":
		return "Tyrosine", nil
	case "UGU", "UGC":
		return "Cysteine", nil
	case "UGG":
		return "Tryptophan", nil
	case "UAA", "UAG", "UGA":
		return "", ErrStop
	default:
		return "", ErrInvalidBase
	}
}
