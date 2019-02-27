package glicko

type Player struct {
    pre     *Rating
    post    *Rating
    matches []*match
}

func (player *Player) addMatch(match *match) {
    player.matches = append(player.matches, match)
}

func (player *Player) Rating() *Rating {
    return player.post
}

func NewPlayer(pre *Rating) *Player {
    return &Player{
        pre:  pre,
        post: NewRating(pre.r, pre.rd, pre.sigma),
    }
}

func NewDefaultPlayer() *Player {
    return NewPlayer(NewDefaultRating())
}
