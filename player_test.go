package gogo

import (
	"reflect"
	"testing"
)

func TestPlayer(t *testing.T) {
	t.Run("other palyer in perspective of BLACK is WHITE", func(t *testing.T) {
		player, _ := NewPlayer(BLACK)
		got := player.other()
		want, _ := NewPlayer(WHITE)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %q, want %q", got, want)
		}
	})
	t.Run("other palyer in perspective of WHITE is BLACK", func(t *testing.T) {
		player, _ := NewPlayer(WHITE)
		got := player.other()
		want, _ := NewPlayer(BLACK)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %q, want %q", got, want)
		}
	})

}
