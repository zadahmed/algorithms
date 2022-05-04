# Create an algorithm that finds pairs of numbers in an array that would sum to a target number provided


# Non Optimal Solution
# o(n^2) time / o(1) sum complexity 
def twoNumberSum(targetArray,targetSum):
    for i in range(len(targetArray) - 1):
        firstNum = targetArray[i]
        for j in range(i+1,len(targetArray)):
            secondNum = targetArray[j]
            if firstNum + secondNum is targetSum:
                return [firstNum, secondNum]
    return []


#Optimal Solution using Hashtable
# o(n) time / o(n) space complexity
def twoNumberSumHashTable(array, targetSum):
    hash_table = {}
    for num in array:
        if targetSum - num in hash_table:
            return [targetSum-num,num]
        else:
            hash_table[num] = True
    return []

x = twoNumberSumHashTable([2, 4, 5, 4,-3,9,4,-2], 8)

