package day02

import (
	"fmt"
)

type Move string

const (
	Rock     Move = "Rock"
	Paper    Move = "Paper"
	Scissors Move = "Scissors"
)

type ScoreKeeper struct {
	CurrentScore  int
	WinScore      int
	DrawScore     int
	LoseScore     int
	RockScore     int
	PaperScore    int
	ScissorsScore int
}

func NewScoreKeeper() *ScoreKeeper {
	return &ScoreKeeper{
		CurrentScore:  0,
		WinScore:      6,
		DrawScore:     3,
		LoseScore:     0,
		RockScore:     1,
		PaperScore:    2,
		ScissorsScore: 3,
	}
}

func (k *ScoreKeeper) Play(opponentMove string, yourMove string) error {
	opponent, err := k.MapOpponentMove(opponentMove)
	if err != nil {
		return err
	}

	your, err := k.MapYourMove(yourMove)
	if err != nil {
		return err
	}

	totalScore := 0

	winScore, err := k.CalculateWin(opponent, your)
	if err != nil {
		return err
	}

	yourMoveScore, err := k.CalculateYourMoveScore(your)
	if err != nil {
		return err
	}

	totalScore += winScore
	totalScore += yourMoveScore

	k.CurrentScore += totalScore
	return nil
}

// PlayPart2 is the same as Play but maps your move a little differently
func (k *ScoreKeeper) PlayPart2(opponentMove string, winCondition string) error {
	opponent, err := k.MapOpponentMove(opponentMove)
	if err != nil {
		return err
	}

	your, err := k.MapYourMovePart2(opponent, winCondition)
	if err != nil {
		return err
	}

	totalScore := 0

	winScore, err := k.CalculateWin(opponent, your)
	if err != nil {
		return err
	}

	yourMoveScore, err := k.CalculateYourMoveScore(your)
	if err != nil {
		return err
	}

	totalScore += winScore
	totalScore += yourMoveScore

	k.CurrentScore += totalScore
	return nil
}

func (k *ScoreKeeper) MapOpponentMove(opponentMove string) (Move, error) {
	switch opponentMove {
	case "A":
		return Rock, nil
	case "B":
		return Paper, nil
	case "C":
		return Scissors, nil
	default:
		return "", fmt.Errorf("invalid opponent move: %s", opponentMove)
	}
}

func (k *ScoreKeeper) MapYourMove(yourMove string) (Move, error) {
	switch yourMove {
	case "X":
		return Rock, nil
	case "Y":
		return Paper, nil
	case "Z":
		return Scissors, nil
	default:
		return "", fmt.Errorf("invalid your move: %s", yourMove)
	}
}

func (k *ScoreKeeper) MapYourMovePart2(opponent Move, winCondition string) (Move, error) {
	if opponent == Rock {
		switch winCondition {
		case "X":
			return Scissors, nil
		case "Y":
			return Rock, nil
		case "Z":
			return Paper, nil
		}
	}

	if opponent == Paper {
		switch winCondition {
		case "X":
			return Rock, nil
		case "Y":
			return Paper, nil
		case "Z":
			return Scissors, nil
		}
	}

	if opponent == Scissors {
		switch winCondition {
		case "X":
			return Paper, nil
		case "Y":
			return Scissors, nil
		case "Z":
			return Rock, nil
		}
	}

	return "", fmt.Errorf("could not determine your move: opponent %s, winCon %s", opponent, winCondition)
}

func (k *ScoreKeeper) CalculateWin(opponent Move, your Move) (int, error) {
	if opponent == your {
		return k.DrawScore, nil
	}

	if your == Rock {
		if opponent == Paper {
			return k.LoseScore, nil
		}
		if opponent == Scissors {
			return k.WinScore, nil
		}
	}

	if your == Paper {
		if opponent == Scissors {
			return k.LoseScore, nil
		}
		if opponent == Rock {
			return k.WinScore, nil
		}
	}

	if your == Scissors {
		if opponent == Rock {
			return k.LoseScore, nil
		}
		if opponent == Paper {
			return k.WinScore, nil
		}
	}

	return 0, fmt.Errorf("could not calculate winning score: opponent %v, your %v", opponent, your)
}

func (k *ScoreKeeper) CalculateYourMoveScore(your Move) (int, error) {
	switch your {
	case Rock:
		return k.RockScore, nil
	case Paper:
		return k.PaperScore, nil
	case Scissors:
		return k.ScissorsScore, nil
	default:
		return 0, fmt.Errorf("could not calculate your move score: %v", your)
	}
}
