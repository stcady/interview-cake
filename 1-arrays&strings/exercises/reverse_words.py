def reverse(list_of_chars, left_index, right_index):

    left_index  = 0
    right_index = len(list_of_chars) - 1

    while left_index < right_index:
        # Swap characters
        list_of_chars[left_index], list_of_chars[right_index] = \
            list_of_chars[right_index], list_of_chars[left_index]
        # Move towards middle
        left_index  += 1
        right_index -= 1

def reverse_words(array):
    reverse(array, 0, len(array)-1)
    i=0
    for j in range(len(array)):
        if (j == len(array)) or (array[j] == ' '):
            reverse(array, i, j-1)
            i = j+1


if __name__ == "__main__":
    message = [ 'c', 'a', 'k', 'e', ' ',
            'p', 'o', 'u', 'n', 'd', ' ',
            's', 't', 'e', 'a', 'l' ]
    reverse_words(message)
    print(''.join(message))