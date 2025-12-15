package day10

import (
	"advent-of-code/go_utils"
	"fmt"
	"log"
	"math"
	"regexp"
	"slices"
	"sort"
	"strconv"
	"strings"
)

// https://adventofcode.com/2025/day/10

func Run(part *string) {

	if part == nil {
		part1("day10/input.txt")
		Part2PyConv()
		return
	}

	switch *part {
	case "1":
		part1("day10/input.txt")
	case "2":
		Part2PyConv()
	default:
		fmt.Println("Invalid Input")
	}
}

type machine struct {
	indicatorGoal       []bool
	buttons             [][]int
	joltageRequirements []int
}

func parseInput(path string) []machine {
	lines, err := go_utils.ReadIntoStrArr(path)

	if err != nil {
		log.Fatalf("Error parsing input: %s", err)
	}

	var machines []machine

	re := regexp.MustCompile(`\[(.*?)\]\s(.*)\s{([0-9,]+)}`)
	for _, line := range lines {
		var m machine
		matches := re.FindStringSubmatch(line)
		for _, r := range matches[1] {
			m.indicatorGoal = append(m.indicatorGoal, r == '#')
		}

		wiringGroups := strings.Fields(matches[2])
		m.buttons = make([][]int, len(wiringGroups))

		for i, wiringGroup := range wiringGroups {
			buttonsChanged := strings.Split(wiringGroup[1:len(wiringGroup)-1], ",")
			for _, btnStr := range buttonsChanged {
				btnInt, err := strconv.Atoi(btnStr)
				if err != nil {
					log.Fatalf("Error converting button wiring %s: %s", btnStr, err)
				}

				m.buttons[i] = append(m.buttons[i], btnInt)
			}
		}

		for _, joltage := range strings.Split(matches[3], ",") {
			jInt, err := strconv.Atoi(joltage)
			if err != nil {
				log.Fatalf("Error converting joltage %s: %s", joltage, err)
			}
			m.joltageRequirements = append(m.joltageRequirements, jInt)
		}

		machines = append(machines, m)
	}

	return machines
}

func printIndicators(indicators []bool) string {
	out := ""
	for _, val := range indicators {
		if val {
			out += "#"
		} else {
			out += "."
		}
	}
	return out
}

func printJoltages(joltages []int) string {
	return fmt.Sprintf("%v", joltages)
}

func printBtnGrouping(grouping []int) string {
	out := ""

	for _, val := range grouping {
		if len(out) != 0 {
			out += ","
		}
		out += strconv.Itoa(val)
	}

	return out
}

func (m *machine) doIndicatorsMatch(testIndicators []bool) bool {
	for i, val := range testIndicators {
		if val != m.indicatorGoal[i] {
			return false
		}
	}

	return true
}

func (m *machine) areJoltagesCorrect(joltages []int) int {
	for i, j := range m.joltageRequirements {
		currJoltage := joltages[i]

		if currJoltage > j {
			return -1
		}

		if currJoltage < j {
			return 0
		}
	}

	return 1
}

func blinkIndicators(curr []bool, flipIs []int) []bool {
	cp := make([]bool, len(curr))
	copy(cp, curr)
	for _, i := range flipIs {
		cp[i] = !cp[i]
	}

	return cp
}

func addJoltages(curr []int, addIs []int) []int {
	cp := make([]int, len(curr))
	copy(cp, curr)
	for _, i := range addIs {
		cp[i] += 1
	}

	return cp
}

func getVisitedMapKey(btnsClicked []int, currPattern []bool) string {
	return fmt.Sprintf("%s|%s", printBtnGrouping(btnsClicked), printIndicators(currPattern))
}

func getFastestMatchRecursive(m machine, currPattern []bool, blinkCount int, visited map[string]int, fastest int) int {
	// Disregard if the fastest is done before this
	if blinkCount > fastest {
		return math.MaxInt
	}

	if m.doIndicatorsMatch(currPattern) {
		return blinkCount
	}

	fastestMatch := fastest
	for _, btnGroup := range m.buttons {
		key := getVisitedMapKey(btnGroup, currPattern)

		if knownCount, ok := visited[key]; ok {
			if blinkCount < fastestMatch && knownCount < fastest {
				fastestMatch = (blinkCount + knownCount)
			}
			continue
		}

		blinked := blinkIndicators(currPattern, btnGroup)
		keyBlinked := getVisitedMapKey(btnGroup, blinked)

		visited[keyBlinked] = math.MaxInt
		fastestIter := getFastestMatchRecursive(m, blinked, blinkCount+1, visited, fastestMatch)
		visited[keyBlinked] = fastestIter

		if fastestIter < fastestMatch {
			fastestMatch = fastestIter
		}
	}

	return fastestMatch
}

type state struct {
	pattern []bool
	count   int
}

func solveP1(m machine) int {
	startState := state{
		pattern: make([]bool, len(m.indicatorGoal)),
		count:   0,
	}

	queue := []state{startState}
	visited := map[string]bool{printIndicators(startState.pattern): true}

	for len(queue) > 0 {
		s := queue[0]
		queue = queue[1:]

		if m.doIndicatorsMatch(s.pattern) {
			return s.count
		}

		for _, btn := range m.buttons {
			next := blinkIndicators(s.pattern, btn)
			key := printIndicators(next)

			if !visited[key] {
				visited[key] = true
				queue = append(queue, state{next, s.count + 1})
			}
		}
	}

	return math.MaxInt
}

func part1(path string) int {
	fmt.Println("Day 10, Part 1: START")
	timer := go_utils.Timer{}
	result := 0

	machines := parseInput(path)

	timer.Start()

	for _, m := range machines {
		minSteps := solveP1(m)
		if minSteps < 0 {
			log.Fatalf("Have negative steps! %d", minSteps)
		}

		result += minSteps
	}

	timer.End()

	fmt.Printf("day 10, part 1 result: %d | %s\n", result, timer.TimeLapsed())
	return result
}

type stateP2 struct {
	currJoltageCount []int
	count            int
}

// This definitely takes way too long
func solveP2Naive(m machine) int {
	startState := stateP2{
		currJoltageCount: make([]int, len(m.joltageRequirements)),
		count:            0,
	}

	queue := []stateP2{startState}
	visited := map[string]bool{printJoltages(startState.currJoltageCount): true}

	for len(queue) > 0 {
		s := queue[0]
		queue = queue[1:]

		joltageMatch := m.areJoltagesCorrect(s.currJoltageCount)

		if joltageMatch == -1 {
			// Yeah this is too much continue
			continue
		}
		if joltageMatch == 1 {
			// We have a match
			return s.count
		}

		for _, btn := range m.buttons {
			next := addJoltages(s.currJoltageCount, btn)
			key := printJoltages(next)

			if !visited[key] {
				visited[key] = true
				queue = append(queue, stateP2{next, s.count + 1})
			}
		}
	}

	return math.MaxInt
}

func areAllAtZero(joltages []int) bool {
	for _, j := range joltages {
		if j > 0 {
			return false
		}
	}

	return true
}

func areAllEven(joltages []int) bool {
	for j := range joltages {
		if j%2 != 0 {
			return false
		}
	}

	return true
}

func sumPresses(buttons [][]int, btnsPressed []int) ([]int, [][]int) {
	joltagesEffected := make([]int, len(buttons[0]))
	outBtns := [][]int{}

	for i := range joltagesEffected {
		joltagesEffected[i] = 0
	}

	for _, btnI := range btnsPressed {
		outBtns = append(outBtns, buttons[btnI])
		for i, j := range buttons[btnI] {
			joltagesEffected[i] += j
		}
	}

	return joltagesEffected, outBtns
}

type pattern struct {
	cost    int
	pattern []int
	buttons [][]int
}

func generatePatterns(buttons [][]int) []pattern {
	pMap := make(map[string]pattern)

	numBtns := len(buttons)

	for patternLen := 0; patternLen <= numBtns; patternLen++ {
		combos := go_utils.Combinations(go_utils.Range(numBtns), patternLen)
		for _, btnIs := range combos {
			pat, btns := sumPresses(buttons, btnIs)
			key := fmt.Sprint(pat)
			if _, ok := pMap[key]; !ok {
				pMap[key] = pattern{
					cost:    patternLen,
					pattern: pat,
					buttons: btns,
				}
			}
		}
	}

	patternsList := make([]pattern, 0, len(pMap))
	for _, p := range pMap {
		patternsList = append(patternsList, p)
	}
	sort.Slice(patternsList, func(i, j int) bool { return patternsList[i].cost < patternsList[j].cost })

	fmt.Println("PATTERNS\n", patternsList)

	return patternsList
}

func buildBtnMatrix(btns [][]int, jCount int) [][]int {
	m := make([][]int, len(btns))
	for bI, btn := range btns {
		for i := 0; i < jCount; i++ {
			if slices.Contains(btn, i) {
				m[bI] = append(m[bI], 1)
			} else {
				m[bI] = append(m[bI], 0)
			}
		}
	}

	return m
}

func solveP2(remaining []int, patterns []pattern, memo map[string]int, pathMemo map[string][][]int) (int, [][]int) {
	if slices.Min(remaining) < 0 {
		log.Fatal("Remaining has a negative!", remaining)
	}
	if areAllAtZero(remaining) {
		return 0, [][]int{}
	}

	key := fmt.Sprint(remaining)
	fmt.Println("Key", key)
	if val, ok := memo[key]; ok {
		return val, pathMemo[key]
	}

	res := math.MaxInt
	var minPath [][]int

	for _, p := range patterns {
		fits := true
		newRemaining := make([]int, len(remaining))

		if len(p.pattern) != len(remaining) {
			fmt.Println("PATTERN MISMATCH")
		}
		for i := 0; i < len(remaining); i++ {
			if p.pattern[i] <= remaining[i] && p.pattern[i]%2 == remaining[i]%2 {
				newRemaining[i] = (remaining[i] - p.pattern[i]) / 2
			} else {
				fits = false
				break
			}

		}

		if fits {
			cost, path := solveP2(newRemaining, patterns, memo, pathMemo)
			calculated := p.cost + 2*cost
			if calculated < res {
				res = calculated
				minPath = append(p.buttons, path...)
			}
			// res = min(res, p.cost+(2*solveP2(newRemaining, patterns, memo)))
		}
	}

	memo[key] = res
	pathMemo[key] = minPath
	return res, minPath
}

// Incorrect
// 14439 -> Too low!
func part2(path string) int {
	fmt.Println("Day 10, Part 2: START")
	result := 0

	machines := parseInput(path)
	timer := go_utils.Timer{}

	timer.Start()

	for _, m := range machines {
		matrix := buildBtnMatrix(m.buttons, len(m.joltageRequirements))
		patterns := generatePatterns(matrix)
		memo := make(map[string]int)
		pathMemo := make(map[string][][]int)
		minSteps, minPath := solveP2(m.joltageRequirements, patterns, memo, pathMemo)
		fmt.Println("Solved: requirements", m.joltageRequirements, "Steps", minSteps, "Path", minPath)
		if minSteps < 0 {
			log.Fatalf("Have negative steps! %d", minSteps)
		}

		result += minSteps
	}

	timer.End()

	fmt.Printf("day 10, part 2 result: %d | %s\n", result, timer.TimeLapsed())

	return result
}
