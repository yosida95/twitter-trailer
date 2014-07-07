package trailer

type Event struct {
	Target       User                   `json:"target"`
	Source       User                   `json:"source"`
	Event        string                 `json:"event"`
	TargetObject map[string]interface{} `json:"target_object"`
}

type DeleteEvent struct {
	Delete struct {
		Status struct {
			Id     uint64 `json:"id"`
			UserId uint64 `json:"user_id"`
		} `json:"status"`
	} `json:"delete"`
}
