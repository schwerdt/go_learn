package main

import "fmt"
import "math/rand"
import "bufio"
import "os"
import "time"

type Game struct {
  deck []int
  dealer_hand []int
  player_hand []int
}

type Card struct {
  CardValue int
  CardName string
  AceFlag
}

func main() {
  game := NewGame()
  game.InitialDeal()

  fmt.Println("player: \n", PrintGameHand(game.player_hand))
  fmt.Println("dealer: \n", PrintGameHand(game.dealer_hand))

  reader := bufio.NewReader(os.Stdin)
  fmt.Println("Would you like another card? (yes/no)")
  text, _ := reader.ReadString('\n')
  for (text == "yes\n")  {
    game.player_hand = append(game.player_hand, game.Deal())
    fmt.Println("player:", PrintGameHand(game.player_hand))
    fmt.Println("Would you like another card? (yes/no)?")
    text, _ = reader.ReadString('\n')
  }
  fmt.Println("score", CalcScore(game.player_hand))
  fmt.Println("hand", PrintFullHand(game.player_hand))
  if CalcScore(game.player_hand) > CalcScore(game.dealer_hand) && CalcScore(game.player_hand) <= 21 {
    fmt.Println("Player won")
   } else {
    fmt.Println("Dealer won") }

  fmt.Println("player hand:", PrintFullHand(game.player_hand))
  fmt.Println("dealer hand:", PrintFullHand(game.dealer_hand))

}

func NewGame() *Game {
  game := Game {
    deck: make([]int, 52),
    dealer_hand: make([]int, 0),
    player_hand: make([]int, 0),
  }
  for i := 0; i < 52; i++ {
    game.deck[i] = i
  }
  return &game
}

func (g *Game) Deal() int {
  // Randomly select an index from remaining cards, remove it from the array
  rand.Seed(time.Now().UTC().UnixNano())
  index := rand.Intn(len(g.deck))
  card := g.deck[index]
  g.deck = append(g.deck[:index], g.deck[index+1:]...)
  return card
}

func (g *Game) InitialDeal() {
  NumInitialCards := 2
  for i:= 0; i < NumInitialCards; i++ {
    g.dealer_hand = append(g.dealer_hand, g.Deal())
    g.player_hand = append(g.player_hand, g.Deal())
  }
}


func CardSuit(card int) string {
  switch(card/13) {
    case 0:
      return "hearts"
    case 1:
      return "spades"
    case 2:
      return "diamonds"
    case 3:
      return "clubs"
  }
  return "Something is wrong"
}

func CardName(card int) string {
  card_map := map[int]string {
    12: "king",
    11: "queen",
    10: "jack",
    9: "10",
    8: "9",
    7: "8",
    6: "7",
    5: "6",
    4: "5",
    3: "4",
    2: "3",
    1: "2",
    0: "ace",
  }
  return card_map[(card%13)]
}

func PrintFullHand(hand []int) string {
  full_hand := ""
  for _, val := range hand {
    full_hand += (CardName(val) + " of " + CardSuit(val) + " ")
  }
  return full_hand
}
 func PrintGameHand(hand []int) string {
   full_hand := "face down card\n"
   for i := 1; i< len(hand); i++ {
     full_hand += " " + (CardName(hand[i]) + " of " + CardSuit(hand[i]) + " ")
  }
  return full_hand
}

func CalcScore(hand []int) int {
  total := 0
  ace_flag := false
   for _, card :=range hand {
     if card%13 > 9 {
       total += 10
     } else if card%13 > 0 {
       total += card%13 + 1
     } else {
       total += 1
       ace_flag = true
     }
   }
  if (total + 10) <= 21 && ace_flag == true {
      total += 10
  }
    return total
}
