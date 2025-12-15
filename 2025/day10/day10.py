import collections, functools, itertools


def powerset(s): 
    for r in range(len(s)+1): yield from itertools.combinations(s,r)


def solve(line):
    lights, *buttons, demands = [g[1:-1] for g in line.split()]

    lights  = tuple(l=='#' for l in lights)
    buttons = tuple(set(eval(b+',')) for b in buttons)
    demands = tuple(eval(demands))

    options, output = collections.defaultdict(list), dict()
    for pressed in powerset(range(len(buttons))):
        supply = [len([1 for b in pressed if j in buttons[b]])
                         for j in range(len(demands))]
        parity = tuple(j%2 for j in supply)

        options[parity] += [pressed]
        output[pressed] = supply


    @functools.cache
    def opt(demands):
        if min(demands)  < 0: return 1e99
        if sum(demands) == 0: return 0

        answer = 1e99
        parity = tuple(j%2 for j in demands)

        for pressed in options[parity]:
            remain = tuple((j-s)//2 for j,s in zip(demands, output[pressed]))
            answer = min(answer, len(pressed) + 2*opt(remain))

        return answer

    return min(map(len, options[lights])), opt(demands)


print(*map(sum, zip(*map(solve, open('input.txt')))))
