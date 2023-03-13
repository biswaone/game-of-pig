package strategyparse

import (
	"errors"
	"strconv"
	"strings"
)

//3 stories
//same binary handle all the 3 stories based on CLI arguments
//parseArgs function will trigger different stories depending on
//the arguments passed

type GameStrategy struct {
	player1  int
	player2  int
	strategy int // To indicate which strategy to use
	// 1- both users fixed strategy
	// 2- one user fixed strategy and other variable strategy
	// 3- both variable strategy
}

func ValidNumber(s []string) (bool, error) {
	var err error
	_, err = strconv.Atoi(s[0])
	if err != nil {
		return false, err
	}
	_, err = strconv.Atoi(s[1])
	if err != nil {
		return false, err
	}
	return true, nil
}

func ParseAndValidateArgs(args []string) (GameStrategy, error) {
	s := GameStrategy{}
	var err error

	if len(args) != 2 {
		return s, errors.New("invalid number of arguments")
	}

	//check for -ve numbers
	if strings.HasPrefix(args[0], "-") || strings.HasPrefix(args[1], "-") {
		isValid, err := ValidNumber(args)
		if !isValid {
			return s, err
		}
		return s, errors.New("must specify a number greater than 0")
	}

	//story 1
	if !strings.Contains(args[0], "-") && !strings.Contains(args[1], "-") {

		//check for valid strings then assign it as strategy 1
		isValid, err := ValidNumber(args)
		if !isValid {
			return s, err
		}
		// validated the strings as a number then assigning
		s.player1, _ = strconv.Atoi(args[0])
		s.player2, _ = strconv.Atoi(args[1])

		s.strategy = 1

		//check if the number is between specified range
		if s.player1 > 100 || s.player2 > 100 {
			return s, errors.New("must specify a number less than equal to 100")
		}

	}

	//story 2
	if !strings.Contains(args[0], "-") && args[1] == "1-100" {
		s.player1, err = strconv.Atoi(args[0])
		if err != nil {
			s.player1 = 0 //Return an zero initialized struct in case of error
			return s, err
		}

		if s.player1 > 100 {
			s.player1 = 0
			return s, errors.New("must specify a number less than equal to 100")
		}

		s.strategy = 2

	}

	// story 3
	if args[0] == "1-100" && args[1] == "1-100" {
		s.strategy = 3
	}

	//if any of the condition doesn't match
	return s, errors.New("invalid syntax")
}
