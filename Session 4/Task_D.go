package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type party struct {
	votes   int
	qseats  int
	ostatok float64
}

type partyostatok struct {
	name    string
	ostatok float64
}

type partyvotes struct {
	name  string
	votes int
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	ansmap := make(map[string]*party)
	var anslist []string
	var totalvotes int

	for {
		line, _ := reader.ReadString('\n')
		if line == "" || line == "\n" {
			break
		}
		linelist := strings.Fields(line)
		name := strings.Join(linelist[:len(linelist)-1], " ")
		votes, _ := strconv.Atoi(linelist[len(linelist)-1])
		anslist = append(anslist, name)
		totalvotes += votes
		ansmap[name] = &party{votes: votes, qseats: 0, ostatok: 0}
	}

	var currmest int
	var ostatoklist []partyostatok
	var voteslist []partyvotes

	for key := range ansmap {
		qmestfloat := float64(ansmap[key].votes) / float64(totalvotes) * 450
		qmest := int(float64(ansmap[key].votes) / float64(totalvotes) * 450)
		ansmap[key].qseats = qmest
		ansmap[key].ostatok = (qmestfloat - float64(qmest)) * float64(ansmap[key].votes) / float64(qmestfloat)
		currmest += qmest
		ostatoklist = append(ostatoklist, partyostatok{name: key, ostatok: ansmap[key].ostatok})
		voteslist = append(voteslist, partyvotes{name: key, votes: ansmap[key].votes})
	}

	if currmest < 450 {
		sort.SliceStable(ostatoklist, func(i, j int) bool {
			return ostatoklist[i].ostatok > ostatoklist[j].ostatok
		})

		for _, value := range ostatoklist {
			if currmest < 450 && value.ostatok != 0 {
				ansmap[value.name].qseats += 1
				currmest += 1
			} else {
				break
			}
		}
	}

	if currmest < 450 {
		sort.SliceStable(voteslist, func(i, j int) bool {
			return voteslist[i].votes > voteslist[j].votes
		})

		for currmest < 450 {
			ansmap[voteslist[0].name].qseats += 1
			currmest += 1
		}
	}

	for _, value := range anslist {
		fmt.Println(value, ansmap[value].qseats)
	}
}
