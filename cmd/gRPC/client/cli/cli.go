package cli

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/ydsxiong/summingservice/gRPC/domain"

	"github.com/ydsxiong/summingservice/gRPC/service"
)

const ServiceMenu = `Menu
===========================================
Choose an option (1, or 2) first, then enter their corresponding values 
1) Type a name or label for the sum action, followed by a list of numbers to be summed up together, e.g. test1 [1,2,3]
2) Type a keyword only (of those sum action names/labels previous entered) to query a list of summaries. e.g. test
Type 'quit' to exit the application.
`
const InValidInput = "Invalid input!"
const InternalError = "Unexpected error occurred!"
const ThanksForPlaying = "Thanks for trying out this new sum service!"

type CLI struct {
	scanner *bufio.Scanner
	output  io.Writer
	client  service.SumServiceClient
	ctx     context.Context
}

func NewSumServiceCLI(input io.Reader, output io.Writer, client service.SumServiceClient, ctx context.Context) *CLI {
	return &CLI{bufio.NewScanner(input), output, client, ctx}
}

func (cl *CLI) scanline() string {
	cl.scanner.Scan()
	return cl.scanner.Text()
}

/**
This mehtod is to allow user to continueously use the service until they typed in a 'quit' command
*/
func (cl *CLI) Run() {
	for {
		if toContinue := cl.HandleInputOptions(); !toContinue {
			break
		}
	}
}

/**
This mehtod is to allow user to execute the command the service once and then terminate the app.
*/
func (cl *CLI) HandleInputOptions() (toContinue bool) {
	toContinue = true
	cl.println(ServiceMenu)
	line := cl.scanline()
	if "quit" == line {
		cl.println(ThanksForPlaying)
		return false
	}
	option, err := strconv.Atoi(line)
	if err != nil {
		cl.println(InValidInput)
	} else {
		if option == 1 {
			toContinue = cl.handleOption1()
		} else if option == 2 {
			toContinue = cl.handleOption2()
		}
	}
	return
}

func (cl *CLI) handleOption1() (toContinue bool) {
	toContinue = true
	input := cl.scanline()
	strs := strings.Split(input, " ")
	var numbers []int64
	var result string
	if len(strs) < 2 {
		result = InValidInput
	} else {
		err := json.Unmarshal([]byte(strs[1]), &numbers)
		if err != nil {
			result = InValidInput
		} else {
			res, err := cl.client.Sum(cl.ctx, &domain.SumRequest{Name: strs[0], Numbers: numbers})
			if err != nil {
				result = InternalError
				toContinue = false
			} else if res == nil {
				result = InValidInput
			} else {
				result = strconv.FormatInt(res.Sum, 10)
			}
		}
	}
	cl.println(result)
	return
}

func (cl *CLI) handleOption2() (toContinue bool) {
	toContinue = true
	keyword := cl.scanline()
	stream, err := cl.client.GetAllSums(cl.ctx, &domain.SumFilter{Keyword: keyword})
	if err != nil {
		if err == io.EOF {
			return
		}
		cl.println(InternalError)
		toContinue = false
	} else {
		for {
			result, err := stream.Recv()
			if err != nil {
				if err == io.EOF {
					break
				}
				cl.println(InternalError)
				toContinue = false
			} else if strings.Contains(result.Input.Name, keyword) {
				cl.println(result)
			}
		}
	}
	return
}

func (cl *CLI) println(outcome interface{}) {
	fmt.Fprintln(cl.output, outcome)
}
