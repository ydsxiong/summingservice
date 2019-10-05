package impl_test

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/ydsxiong/summingservice/datastore"
	"github.com/ydsxiong/summingservice/gRPC/domain"
	"google.golang.org/grpc"

	"github.com/ydsxiong/summingservice/gRPC/impl"
)

type mockStore struct {
	num int
}

func (s *mockStore) Save(msg datastore.DataMessage) error {
	s.num++
	return nil
}

func (s *mockStore) FetchAll() ([]datastore.DataMessage, error) {
	return nil, nil
}

func TestServiceSum(t *testing.T) {

	testcases := []struct {
		name    string
		numbers []int64
		sum     int64
	}{
		{
			"test1",
			[]int64{1},
			1,
		},
		{
			"test2",
			[]int64{1, 3},
			4,
		},
		{
			"test3",
			[]int64{1, 3, 5, 7, 9},
			25,
		},
	}

	for _, tc := range testcases {
		store := new(mockStore)
		server := impl.NewSumServer(store)
		t.Run(tc.name, func(t *testing.T) {
			request := domain.SumRequest{Name: tc.name, Numbers: tc.numbers}
			res, err := server.Sum(context.Background(), &request)
			if err != nil {
				t.Errorf("Unexpected error occurred: %v", err)
			} else if res == nil {
				t.Errorf("Unexpected null response received")
			} else if res.Sum != tc.sum {
				t.Errorf("got %d, wanted %d", res.Sum, tc.sum)
			} else if store.num != 1 {
				t.Errorf("expected one message to be saved, but found %d", store.num)
			}
		})
	}
}

var errorSendingResponse = errors.New("Message sending interrupted...")

type mockServerHandler struct {
	msgs []datastore.DataMessage
	grpc.ServerStream
	errorOccuured bool
}

func (h *mockServerHandler) Send(res *domain.SumResponse) error {
	if h.errorOccuured {
		return errorSendingResponse
	}
	h.msgs = append(h.msgs, datastore.DataMessage{Name: res.Input.Name, Input: res.Input.Numbers, Output: res.Sum})
	return nil
}
func (h *mockServerHandler) NumOfMessagesSent() int {
	return len(h.msgs)
}

func TestServiceGetSums(t *testing.T) {

	store := datastore.NewInMemoryDataStore()
	store.Save(datastore.DataMessage{"sum1", []int64{1}, 1})
	store.Save(datastore.DataMessage{"sum2", []int64{1, 3, 5}, 9})
	server := impl.NewSumServer(store)

	testcases := []struct {
		name   string
		filter *domain.SumFilter
		msg    []datastore.DataMessage
		err    error
	}{
		{
			"test1",
			nil,
			[]datastore.DataMessage{datastore.DataMessage{"sum1", []int64{1}, 1}, datastore.DataMessage{"sum2", []int64{1, 3, 5}, 9}},
			nil,
		},
		{
			"test2",
			&domain.SumFilter{Keyword: "sum"},
			[]datastore.DataMessage{datastore.DataMessage{"sum1", []int64{1}, 1}, datastore.DataMessage{"sum2", []int64{1, 3, 5}, 9}},
			nil,
		},
		{
			"test3",
			&domain.SumFilter{Keyword: "sum1"},
			[]datastore.DataMessage{datastore.DataMessage{"sum1", []int64{1}, 1}},
			nil,
		},
		{
			"test4",
			&domain.SumFilter{Keyword: "foo"},
			[]datastore.DataMessage{},
			nil,
		},
		{
			"test5",
			nil,
			[]datastore.DataMessage{},
			errorSendingResponse,
		},
	}

	for _, tc := range testcases {
		var serverHandler = &mockServerHandler{}
		if tc.name == "test5" {
			serverHandler.errorOccuured = true
		}
		t.Run(tc.name, func(t *testing.T) {
			err := server.GetAllSums(tc.filter, serverHandler)
			if tc.err == nil && err != nil {
				t.Errorf("Unexpected error occurred: %v", err)
			} else if tc.err != nil && err == nil {
				t.Errorf("Expected an error, but didn't get any")
			} else if len(serverHandler.msgs) != len(tc.msg) {
				t.Errorf("got number of message %d, wanted %d", len(serverHandler.msgs), len(tc.msg))
			} else {
				assertEqual(t, serverHandler.msgs, tc.msg)
			}
		})
	}
}

func assertEqual(t *testing.T, got []datastore.DataMessage, wanted []datastore.DataMessage) {
	t.Helper()
	if len(got) != len(wanted) {
		t.Errorf("got number of message %d, wanted %d", len(got), len(wanted))
	} else {
		for i, msg := range got {
			if !reflect.DeepEqual(msg, wanted[i]) {
				t.Errorf("got %v, wanted %v", got, wanted)
			}
		}
	}

}
