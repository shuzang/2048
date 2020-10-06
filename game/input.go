package game

import (
	"errors"

	"github.com/eiannone/keyboard"
)

// GetCahrKeystroke returns a key without pressing enter/return key
// supported keys are [w a s d] and Arrow keys
// below magical numbers are their key codes
func GetCharKeystroke() (Dir, error) {
	char, key, err := keyboard.GetSingleKey()
	if err != nil {
		return ERROR_KEY, err
	}
	//fmt.Printf("You pressed: %c, key %X\r\n", char, key)
	if int(char) == 0 {
		switch key {
		case keyboard.KeyArrowUp:
			return UP, nil
		case keyboard.KeyArrowDown:
			return DOWN, nil
		case keyboard.KeyArrowLeft:
			return LEFT, nil
		case keyboard.KeyArrowRight:
			return RIGHT, nil
		case keyboard.KeyEsc:
			return QUIT, nil
		default:
			return ERROR_KEY, errors.New("GameOverError")
		}
	} else {
		switch char {
		case 119:
			return UP, nil
		case 97:
			return LEFT, nil
		case 115:
			return DOWN, nil
		case 100:
			return RIGHT, nil
		default:
			return ERROR_KEY, errors.New("GameOverError")
		}
	}
}
