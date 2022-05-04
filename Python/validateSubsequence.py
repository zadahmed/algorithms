# Algorithm to validate if subsequence within a sequence 

# Time complexity is O(N) / Space is O(1)

def validateSequence(array, subsequence):
    arrIndex = 0
    seqIndex = 0
    while arrIndex < len(array) and seqIndex < len(subsequence):
        if array[arrIndex] == subsequence[seqIndex]:
            seqIndex += 1
        arrIndex += 1
    return seqIndex == len(subsequence)                

