package game

import (
	"errors"

	"github.com/eiannone/keyboard"
)

// getCahrKeystroke returns a key without pressing enter/return key
// supported keys are [w a s d] and Arrow keys
// below magical numbers are their key codes
func getCharKeystroke() (Dir, error) {
	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	defer func() {
		_ = keyboard.Close()
	}()
	char, key, err := keyboard.GetKey()
	ans := int(char)
	if ans == 0 {
		ans = int(key)
	}
	if err != nil {
		return NO_DIR, err
	}
	switch ans {
	case 119, 65517:
		return UP, nil
	case 97, 65515:
		return LEFT, nil
	case 115, 65516:
		return DOWN, nil
	case 100, 65514:
		return RIGHT, nil
	case 3:
		return NO_DIR, errors.New("GameOverError")
	}
	return NO_DIR, nil
}
