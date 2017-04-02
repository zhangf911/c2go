package ast

type NotNullAttr struct {
	Address  string
	Position string
}

func parseNotNullAttr(line string) NotNullAttr {
	groups := groupsFromRegex(
		"<(?P<position>.*)> 1",
		line,
	)

	return NotNullAttr{
		Address: groups["address"],
		Position: groups["position"],
	}
}
