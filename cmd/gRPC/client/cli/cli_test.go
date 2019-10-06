package cli_test

import (
	"bytes"
	"context"
	"io"
	"strings"
	"testing"

	"github.com/ydsxiong/summingservice/cmd/gRPC/client/cli"
	"github.com/ydsxiong/summingservice/gRPC/domain"
	"github.com/ydsxiong/summingservice/gRPC/service"
	"google.golang.org/grpc"
)

type mockSumQueryResponse struct {
	result []*domain.SumResponse
	count  int
	err    error
	grpc.ClientStream
}

// mock out data streaming behaviour
func (sq *mockSumQueryResponse) Recv() (*domain.SumResponse, error) {
	if sq.count < len(sq.result) {
		res := sq.result[sq.count]
		sq.count++
		return res, sq.err
	}
	return nil, io.EOF
}

type mockServiceClient struct {
	name      string
	numbers   []int64
	sum       int64
	err       error
	qryResult *mockSumQueryResponse
}

func (c *mockServiceClient) Sum(ctx context.Context, in *domain.SumRequest, opts ...grpc.CallOption) (*domain.SumResponse, error) {
	if in.Name == c.name && areEqual(c.numbers, in.Numbers) {
		return &domain.SumResponse{Sum: c.sum}, c.err
	}
	return nil, nil
}

func (c *mockServiceClient) GetAllSums(ctx context.Context, in *domain.SumFilter, opts ...grpc.CallOption) (service.SumService_GetAllSumsClient, error) {
	return c.qryResult, nil
}

func TestCLI(t *testing.T) {

	testcases := []struct {
		option   string
		cmdInput string
		name     string
		numbers  []int64
		sum      int64
		qryrs    *mockSumQueryResponse
		output   []string
		status   bool
	}{
		{
			"1", "sum1 [1,3,5]", "sum1", []int64{1, 3, 5}, 9, nil, []string{cli.ServiceMenu, "9"}, true,
		},
		{
			"1", "sum2 [1,a,5]", "sum2", []int64{}, 0, nil, []string{cli.ServiceMenu, cli.InValidInput}, true,
		},
		{
			"1", "sum3 []", "sum3", []int64{}, 0, nil, []string{cli.ServiceMenu, "0"}, true,
		},
		{
			"quit", "", "quit", []int64{}, 0, nil, []string{cli.ServiceMenu, cli.ThanksForPlaying}, false,
		},
		{
			"random", "", "random input", []int64{}, 0, nil, []string{cli.ServiceMenu, cli.InValidInput}, true,
		},
		{
			"2", "sum", "query1", []int64{}, 0, createQueryResponseWithStreamData(),
			[]string{cli.ServiceMenu, "input:<name:\"sum3\" numbers:1 > sum:1 ", "input:<name:\"sum1\" numbers:1 numbers:3 > sum:4 "}, true,
		},
		{
			"2", "sum1", "query2", []int64{}, 0, createQueryResponseWithStreamData(),
			[]string{cli.ServiceMenu, "input:<name:\"sum1\" numbers:1 numbers:3 > sum:4 "}, true,
		},
		{
			"2", "random", "query3", []int64{}, 0, createQueryResponseWithStreamData(),
			[]string{cli.ServiceMenu}, true,
		},
	}
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()
			out := &bytes.Buffer{}
			in := userTypesIn(tc.option, tc.cmdInput)
			cmdl := cli.NewSumServiceCLI(in, out, &mockServiceClient{tc.name, tc.numbers, tc.sum, nil, tc.qryrs}, ctx)
			status := cmdl.HandleInputOptions()
			if status != tc.status {
				t.Errorf("expected status %t, but got %t", tc.status, status)
			}
			assertMessagesSentToUser(t, out, tc.output)
		})
	}
}

func userTypesIn(messages ...string) io.Reader {
	return strings.NewReader(strings.Join(messages, "\n"))
}

func assertMessagesSentToUser(t *testing.T, stdout *bytes.Buffer, messages []string) {
	t.Helper()
	want := strings.Join(messages, "\n") + "\n"
	got := stdout.String()
	if got != want {
		t.Errorf("got %q sent to stdout but expected %+v", got, messages)
	}
}

func areEqual(a, b []int64) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func createQueryResponseWithStreamData() *mockSumQueryResponse {
	return &mockSumQueryResponse{result: []*domain.SumResponse{
		&domain.SumResponse{Sum: 1, Input: &domain.SumRequest{Name: "sum3", Numbers: []int64{1}}},
		&domain.SumResponse{Sum: 4, Input: &domain.SumRequest{Name: "sum1", Numbers: []int64{1, 3}}}}}
}
