package practicum

import (
	"reflect"
	"testing"
)

func TestCompetition(t *testing.T) {
	testCompetition(t, &CompetitionWithSortingOnGetter{})
}

func testCompetition(t *testing.T, competition CompetitionInterface) {
	competition.WithWorks(Works{
		1:  &Work{10},
		2:  &Work{10},
		3:  &Work{1},
		4:  &Work{2},
		5:  &Work{6},
		6:  &Work{9},
		7:  &Work{12},
		8:  &Work{1},
		9:  &Work{16},
		10: &Work{2},
	})
	competition.Like(1)
	competition.Like(1)
	competition.Like(1)
	competition.Dislike(1)
	competition.Like(2)
	competition.Dislike(2)
	competition.Like(9)

	actualBestWorks := competition.GetBestWorks(5)
	expectedBestWorks := Works{
		9: &Work{17},
		1: &Work{12},
		7: &Work{12},
		2: &Work{10},
		6: &Work{9},
	}

	if !reflect.DeepEqual(actualBestWorks, expectedBestWorks) {
		t.Errorf("GetBestWorks() = %v, want %v", actualBestWorks, expectedBestWorks)
	}
}
