package main

func main() {
	Hey()
}

var b1, b2, b3, b4 bool

func Hey() { // 22

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

// elseを加算すると仮定
func Bar() { // 21
	if b1 { // +1
		if b2 { // +2
			if b3 { // +3
			} else if b4 { // +1
			} else { // +1
			}
		}
	} else { // +1
		if b1 { // +2
			if b2 { // +3
			} else { // +1
				if b3 { // +4
				} else if b4 { // +1
				} else { // +1
				}
			}
		}
	}
}
