package tgwrap

// Poll is a Telegram API object.
type Poll struct {
	//id	String	Unique poll identifier
	//question	String	Poll question, 1-255 characters
	//options	Array of PollOption	List of poll options
	//total_voter_count	Integer	Total number of users that voted in the poll
	//is_closed	Boolean	True, if the poll is closed
	//is_anonymous	Boolean	True, if the poll is anonymous
	//type	String	Poll type, currently can be “regular” or “quiz”
	//allows_multiple_answers	Boolean	True, if the poll allows multiple answers
	//correct_option_id	Integer	Optional. 0-based identifier of the correct answer option. Available only for polls in the quiz mode, which are closed, or was sent (not forwarded) by the bot or to the private chat with the bot.
}

// PollAnswer represents an answer of a user in a non-anonymous poll.
type PollAnswer struct {
	//poll_id	String	Unique poll identifier
	//user	User	The user, who changed the answer to the poll
	//option_ids	Array of Integer	0-based identifiers of answer options, chosen by the user. May be empty if the user retracted their vote.

}

//PollOption contains information about one answer option in a poll.
type PollOption struct {

	//text	String	Option text, 1-100 characters
	//voter_count	Integer	Number of users that voted for this option
}
