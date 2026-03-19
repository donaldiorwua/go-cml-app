package processor

import (
	"strings"
	"testing"
)

func run(input string) string {
	return Process(input)
}

func TestHexConversion(t *testing.T) {
	input := "1E (hex) files were added"
	expected := "30 files were added"

	result := run(input)
	if result != expected {
		t.Errorf("hex failed: got %q, want %q", result, expected)
	}
}

func TestBinConversion(t *testing.T) {
	input := "It has been 10 (bin) years"
	expected := "It has been 2 years"

	result := run(input)
	if result != expected {
		t.Errorf("bin failed: got %q, want %q", result, expected)
	}
}

func TestUp(t *testing.T) {
	input := "Ready, set, go (up) !"
	expected := "Ready, set, GO!"

	result := run(input)
	if result != expected {
		t.Errorf("up failed: got %q, want %q", result, expected)
	}
}

func TestLow(t *testing.T) {
	input := "I should stop SHOUTING (low)"
	expected := "I should stop shouting"

	result := run(input)
	if result != expected {
		t.Errorf("low failed: got %q, want %q", result, expected)
	}
}

func TestCap(t *testing.T) {
	input := "Welcome to the brooklyn bridge (cap)"
	expected := "Welcome to the Brooklyn Bridge"

	result := run(input)
	if result != expected {
		t.Errorf("cap failed: got %q, want %q", result, expected)
	}
}

func TestUpN(t *testing.T) {
	input := "This is so exciting (up, 2)"
	expected := "This is SO EXCITING"

	result := run(input)
	if result != expected {
		t.Errorf("upN failed: got %q, want %q", result, expected)
	}
}

func TestLowN(t *testing.T) {
	input := "STOP SHOUTING NOW (low, 2)"
	expected := "stop shouting NOW"

	result := run(input)
	if result != expected {
		t.Errorf("lowN failed: got %q, want %q", result, expected)
	}
}

func TestCapN(t *testing.T) {
	input := "this is so exciting (cap, 2)"
	expected := "this is So Exciting"

	result := run(input)
	if result != expected {
		t.Errorf("capN failed: got %q, want %q", result, expected)
	}
}

func TestPunctuationSpacing(t *testing.T) {
	input := "I was sitting over there ,and then BAMM !!"
	expected := "I was sitting over there, and then BAMM!!"

	result := run(input)
	if result != expected {
		t.Errorf("punctuation failed: got %q, want %q", result, expected)
	}
}

func TestPunctuationGroups(t *testing.T) {
	input := "I was thinking ... You were right"
	expected := "I was thinking... You were right"

	result := run(input)
	if result != expected {
		t.Errorf("punctuation group failed: got %q, want %q", result, expected)
	}
}

func TestQuotesSingleWord(t *testing.T) {
	input := "I am exactly how they describe me: ' awesome '"
	expected := "I am exactly how they describe me: 'awesome'"

	result := run(input)
	if result != expected {
		t.Errorf("quotes single failed: got %q, want %q", result, expected)
	}
}

func TestQuotesMultipleWords(t *testing.T) {
	input := "As Elton John said: ' I am the most well-known homosexual in the world '"
	expected := "As Elton John said: 'I am the most well-known homosexual in the world'"

	result := run(input)
	if result != expected {
		t.Errorf("quotes multiple failed: got %q, want %q", result, expected)
	}
}

func TestGrammarAtoAn(t *testing.T) {
	input := "There it was. A amazing rock!"
	expected := "There it was. An amazing rock!"

	result := run(input)
	if result != expected {
		t.Errorf("grammar failed: got %q, want %q", result, expected)
	}
}

func TestFullExample(t *testing.T) {
	input := "it (cap) was the best of times, it was the worst of times (up) , it was the age of wisdom, it was the age of foolishness (cap, 6) , it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, IT WAS THE (low, 3) winter of despair."

	expected := "It was the best of times, it was the worst of TIMES, it was the age of wisdom, It Was The Age Of Foolishness, it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, it was the winter of despair."

	result := run(input)

	if strings.TrimSpace(result) != strings.TrimSpace(expected) {
		t.Errorf("full test failed:\n got: %q\nwant: %q", result, expected)
	}
}
