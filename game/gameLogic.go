package game

import (
	"strings"
)

type Hangman struct {
	Word         string   // Le mot a deviner
	Guessed      []string // Lettres correctement devinées
	Attempts     int      // Essais restants
	WrongGuesses []string // Lettres pas correctement devinées
}

// Initalisation d'une nouvelle partie avec un mot a deviner et un nombre max d'essais
func NewGame(word string, attempts int) *Hangman {
	return &Hangman{
		Word:     word,
		Guessed:  make([]string, len(word)),
		Attempts: attempts,
	}
}

// Fonction qui permet d'essayer de deviner une lettre du mot
func (h *Hangman) Guess(letter string) bool {
	letter = strings.ToLower(letter)

	if strings.Contains(h.Word, letter) {
		// Si la lettre est bonne, on la revèle
		for i, l := range h.Word {
			if string(l) == letter {
				h.Guessed[i] = letter
			}
		}
		return true
	}

	// Si la lettre n'est pas dans le mot, on réduis le nombre d'essais
	h.WrongGuesses = append(h.WrongGuesses, letter)
	h.Attempts--
	return false
}

// Fonction qui vérifie que la partie soit terminée ( gagné ou perdu )
func (h *Hangman) IsGameOver() bool {
	return h.Attempts == 0 || h.IsWordGuessed()
}

// Vérifie si le mot complet est deviné
func (h *Hangman) IsWordGuessed() bool {
	for _, letter := range h.Guessed {
		if letter == "" {
			return false
		}
	}
	return true
}
