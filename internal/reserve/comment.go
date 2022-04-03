package reserve

import "encoding/json"

type Comment string

func (c *Comment) UnmarshalJSON(data []byte) error {
	var value string
	err := json.Unmarshal(data, &value)
	if err != nil {
		return err
	}

	*c = Comment(value)
	return nil
}
