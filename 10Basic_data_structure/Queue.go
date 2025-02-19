package Basic_data_structure

import "errors"

type Queue struct {
	List []int
}

func (q *Queue) Enqueue(v int) {
	q.List = append(q.List, v)
}
func (q *Queue) Dequeue() (int, error) {
	if len(q.List) == 0 {
		return 0, errors.New("queue is empty")
	} else {
		v := q.List[0]
		q.List = q.List[1:]
		return v, nil
	}
}
