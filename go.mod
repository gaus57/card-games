module github.com/gaus57/card-games

go 1.12

require (
    github.com/gorilla/websocket v1.4.0
    card-games/decks/classic36 v0.0.0
    card-games/games/pharaoh v0.0.0
)

replace (
    card-games/decks/classic36 => ./decks/classic36
    card-games/games/pharaoh => ./games/pharaoh
)
