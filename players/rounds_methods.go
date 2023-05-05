package players

import "errors"

func (p *drop) fold(uid int) {
	for i, v := range p.players {
		//fold
		if v.Uid == uid {
			p.players[i].Status = false
			p.Fold = append(p.Fold, v.Uid)
			break
		}
	}
}
func (p *drop) call(uid int) error {
	for i, v := range p.players {
		if v.Uid == uid {
			if v.Bank+p.bets[v.Uid] < p.max_bet {
				return errors.New("not enough money")
			}
			p.players[i].Bank -= p.max_bet - p.bets[v.Uid]
			p.bets[v.Uid] = p.max_bet
			if p.players[i].Bank == 0 {
				p.players[i].Status = false
			}
			break
		}
	}
	return nil
}
func (p *drop) check(uid int) error {
	for _, v := range p.players {
		if v.Uid == uid {
			//success check
			if p.bets[v.Uid] == p.max_bet {
				return nil
			}
			break
		}
	}
	return errors.New("not enough money")
}

func (p *drop) raise(uid int, value int) error {
	for i, v := range p.players {
		if v.Uid == uid {

			//success raise
			if v.Bank >= value && p.bets[v.Uid]+value > p.max_bet {
				p.bets[v.Uid] += value
				p.max_bet = p.bets[v.Uid]
				p.link_active_player = i + len(p.players) - 1
				p.players[i].Bank -= value
				if p.players[i].Bank == 0 {
					p.players[i].Status = false
				}
				return nil
			}
			break
		}
	}
	return errors.New("not enough money")
}

func (p *drop) all_in(uid int) {
	for i, v := range p.players {
		if v.Uid == uid {

			//all_in
			p.bets[v.Uid] += v.Bank
			if p.bets[v.Uid] > p.max_bet {
				p.max_bet = p.bets[v.Uid]
				p.link_active_player = i + len(p.players) - 1
			}
			p.players[i].Bank = 0
			p.players[i].Status = false
			break
		}
	}
}

// раздача блайндов
func (p *drop) blinds(value int) {
	j := p.link_diller
	max_players := len(p.players)
	for {
		if p.players[j%max_players].Status != false {
			err := p.raise(p.players[j%max_players].Uid, value)
			if err != nil {
				p.all_in(p.players[j%max_players].Uid)
			}
			j += 1
			break
		}
		j += 1
	}
	for {
		if p.players[j%max_players].Status != false {
			err := p.raise(p.players[j%max_players].Uid, 2*value)
			if err != nil {
				p.all_in(p.players[j%max_players].Uid)
			}
			break
		}
		j += 1
	}
}

func (p *drop) next_player(j int) int {
	max := len(p.players)
	for !p.players[j%max].Status {
		j += 1
	}
	return j
}
