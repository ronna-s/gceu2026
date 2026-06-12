package fileservice

type Part struct {
	p []byte
}

type Client struct {
	service *service
}

func NewClient() *Client {
	return &Client{service: DefaultService}
}

func (c Client) GetPart(p *Part) ([]byte, int) {
	return c.service.GetPart(p)
}

func (c *Client) GetParts() []*Part {
	return c.service.GetParts()
}

func Map[T1, T2 any](s1 []T1, m func(T1) T2) []T2 {
	s2 := make([]T2, len(s1))
	for i, e := range s1 {
		s2[i] = m(e)
	}
	return s2
}
