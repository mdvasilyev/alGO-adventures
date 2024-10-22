import sys


def is_opening_bracket(c):
    return c in '([{'


def closing_bracket(c):
    if c == '(':
        return ')'
    elif c == '[':
        return ']'
    else:
        return '}'


def check_brackets(s):
    st = []

    for val in s:
        if is_opening_bracket(val):
            st.append(val)
        else:
            if len(st) == 0 or val != closing_bracket(st[-1]):
                return False
            st.pop()

    return len(st) == 0


if __name__ == "__main__":
    s = sys.stdin.read().strip()

    if check_brackets(s):
        print("YES")
    else:
        print("NO")
