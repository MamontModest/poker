package players

import "errors"

func (p *drop) fold(uid int) {
	for i, v := range p.players {
		//fold
		if v.Uid == uid {
			p.players[i].Status = false
			p.Fold = append(p.Fold, v)
			break
		}
	}
}
func (p *drop) call(uid int) error {
	for i, v := range p.players {
		if v.Uid == uid {
			if v.Bank+p.bets[v] < p.max_bet {
				return errors.New("not enough money")
			}
			p.players[i].Bank -= p.max_bet - p.bets[v]
			p.bets[v] = p.max_bet
			p.link_player += len(p.players) - 1
			break
		}
	}
	return nil
}
func (p *drop) check(uid int) error {
	for _, v := range p.players {
		if v.Uid == uid {
			//success check
			if p.bets[v] == p.max_bet {
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
			if v.Bank >= value && p.bets[v]+value > p.max_bet {
				p.bets[v] += value
				p.max_bet = p.bets[v]
				p.link_player = (i + len(p.players) - 1) % len(p.players)
				p.players[i].Bank -= value
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
			p.bets[v] += v.Bank
			if p.bets[v] > p.max_bet {
				p.max_bet = p.bets[v]
				p.link_player += len(p.players) - 1
			}
			p.players[i].Bank = 0
			break
		}
	}
}

// раздача блайндов
func (p *drop) blinds(value int) {

}
