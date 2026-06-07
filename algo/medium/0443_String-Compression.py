class Solution:
    def compress(self, chars: List[str]) -> int:
        length = len(chars)
        if length == 1:
            return 1
        prev = chars[0]
        counter = 1
        idxToInsert = 0
        for i in range(1, length+1):
            if i < length:
                cur = chars[i]
            if i == length or cur != prev:
                chars[idxToInsert] = prev
                idxToInsert += 1
                digits = self.numToList(counter)
                for j, digit in enumerate(digits):
                    chars[idxToInsert+j] = digit
                idxToInsert += len(digits)
                prev = cur
                counter = 1
            else:
                counter += 1
        return idxToInsert

    def numToList(self, num: int) -> List[str]:
        if num == 1:
            return []
        digits = []
        while num != 0:
            digit = num % 10
            digits.append(str(digit))
            num //= 10
        digits.reverse()
        return digits
