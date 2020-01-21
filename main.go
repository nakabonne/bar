package main

func main() {
	Hey()
}

func Hey() { // 22
	var b1, b2, b3, b4 bool

	if b1 { // +1
		if b2 { // +2
		} else {
			if b3 { // +3
			}
		}
	}

	if b1 { // +1
		if b2 { // +2
		} else if b3 {
			if b4 { // +3
			}
		}
	}

	if b1 { // +1
		if b2 { // +2
			if b3 { // +3
				if b4 { // +4
				}
			}
		}
	}
}

func Foo() { // 23
	var b1, b2, b3, b4 bool

	if b1 { // +1
		if b2 { // +2
		} else {
			if b3 { // +3
			}
		}
	}

	if b1 { // +1
		if b2 { // +2
		} else if b3 {
			if b4 { // +3
			}
		}
	}

	if b1 { // +1
		if b2 { // +2
			if b3 { // +3
				if b4 { // +4
				}
			}
		}
	}

	if b1 { // +1
	} else {
	}
}
