package blackjack

const (
    stand = "S"
    hit = "H"
    split = "P"
    win = "W"
)

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
    switch card {
    case "ace":
        return 11
    case "one":
        return 1
    case "two":
        return 2
    case "three":
        return 3
    case "four":
        return 4
    case "five":
        return 5
    case "six":
        return 6
    case "seven":
        return 7
    case "eight":
        return 8
    case "nine":
        return 9
    case "ten", "jack", "queen", "king":
        return 10
    default:
        return 0
    }
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
    dealerCardNum := ParseCard(dealerCard)
    cardSum := ParseCard(card1) + ParseCard(card2)
    switch {
    case cardSum == 22:
        return split
    case cardSum == 21:
        if dealerCardNum < 10 {
            return win
        } else {
            return stand
        }
    case cardSum >= 17:
        return stand
    case cardSum >= 12 && dealerCardNum < 7:
        return stand
    default:
        return hit
    }
}
