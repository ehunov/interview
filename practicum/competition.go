package practicum

import "sort"

type Work struct {
	Score int
}

type Works map[int]*Work

type CompetitionInterface interface {
	WithWorks(works Works)
	Like(participantId int)
	Dislike(participantId int)
	GetBestWorks(count int) Works
}

type CompetitionWithSortingOnGetter struct {
	works Works
}

func (receiver *CompetitionWithSortingOnGetter) WithWorks(works Works) {
	receiver.works = works
}

func (receiver *CompetitionWithSortingOnGetter) Like(participantId int) {
	receiver.works[participantId].Score++
}

func (receiver *CompetitionWithSortingOnGetter) Dislike(participantId int) {
	receiver.works[participantId].Score--
}

func (receiver *CompetitionWithSortingOnGetter) GetBestWorks(count int) Works {
	type Pair struct {
		ParticipantId int
		Work          *Work
	}

	i, pairs := 0, make([]Pair, len(receiver.works))

	for participantId, work := range receiver.works {
		pairs[i] = Pair{
			participantId,
			work,
		}
		i++
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].Work.Score > pairs[j].Work.Score
	})

	result := make(Works, count)

	for _, pair := range pairs[0:count] {
		result[pair.ParticipantId] = pair.Work
	}

	return result
}
