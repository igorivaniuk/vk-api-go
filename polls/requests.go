package polls

const (
	EnumNominative    = "nom"
	EnumGenitive      = "gen"
	EnumDative        = "dat"
	EnumAccusative    = "acc"
	EnumInstrumental  = "ins"
	EnumPrepositional = "abl"
)

type GetByIdRequest struct {
	OwnerId int  `param:"owner_id"`
	IsBoard bool `param:"is_board"`
	PollId  int  `param:"poll_id"`
}

type AddVoteRequest struct {
	OwnerId  int  `param:"owner_id"`
	PollId   int  `param:"poll_id"`
	AnswerId int  `param:"answer_id"`
	IsBoard  bool `param:"is_board"`
}

type DeleteVoteRequest struct {
	OwnerId  int  `param:"owner_id"`
	PollId   int  `param:"poll_id"`
	AnswerId int  `param:"answer_id"`
	IsBoard  bool `param:"is_board"`
}

type GetVotersRequest struct {
	OwnerId     int    `param:"owner_id"`
	PollId      int    `param:"poll_id"`
	AnswerIds   string `param:"answer_ids"`
	IsBoard     bool   `param:"is_board"`
	FriendsOnly bool   `param:"friends_only"`
	Offset      int    `param:"offset"`
	Count       int    `param:"count"`
	Fields      string `param:"fields"`
	NameCase    string `param:"name_case"`
}

type CreateRequest struct {
	Question    string `param:"question"`
	IsAnonymous bool   `param:"is_anonymous"`
	OwnerId     int    `param:"owner_id"`
	AddAnswers  string `param:"add_answers"`
}

type EditRequest struct {
	OwnerId       int    `param:"owner_id"`
	PollId        int    `param:"poll_id"`
	Question      string `param:"question"`
	AddAnswers    string `param:"add_answers"`
	EditAnswers   string `param:"edit_answers"`
	DeleteAnswers string `param:"delete_answers"`
}
