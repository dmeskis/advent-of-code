with open("input") as f:
    lines = f.readlines()   

'''I check if there is a number right above or below the special character and if there is, I do not need to check
diagonals for that direction anymore as if there cant be both a number directly above as well as in one of 
the upper diagonals, as it would form one big number. 
'''
def lookright(temp,col,row):
    if not lines[row][col].isdecimal():
        return temp
    else:
        temp.append(lines[row][col])
        return lookright(temp,col+1,row)
    
def lookleft(temp,col,row):
    if not lines[row][col].isdecimal():
        temp.reverse()
        return temp
    else:
        temp.append(lines[row][col])
        return lookleft(temp,col-1,row)

def is_valid_ind(row=0,col=0):
    return (row < len(lines) and col < len(lines[0])-1) and (row >=0 and col >= 0)
    

_IGNORE = '1234567890.'



coll = []
for row in range(len(lines)):
   for col in range(len(lines[row])-1):
       if lines[row][col] not in _IGNORE:
           directly_above = False
           directly_below = False

           #directly above
           if row-1 >= 0:
                if lines[row-1][col].isdecimal():
                    left = lookleft([],col,row-1)
                    right = lookright([],col+1,row-1)
                    joint  = left+right
                    directly_above = True
                    coll.append(joint)


           #directly below
           if is_valid_ind(row+1):
                if lines[row+1][col].isdecimal():
                    left = lookleft([],col,row+1)
                    right = lookright([],col+1,row+1)
                    joint  = left+right
                    directly_below = True
                    coll.append(joint)

           
           
           #directly to the left
           if is_valid_ind(row,col-1) :
               if lines[row][col-1].isdecimal():
                   left = lookleft([],col-1,row)
                   coll.append(left)
           
           #directly to the right
           if is_valid_ind(row,col+1):
               if lines[row][col+1].isdecimal():
                   right = lookright([],col+1,row)
                   coll.append(right)

            

                ##DIAGONALS##


           if not directly_above:
            #top right
            if is_valid_ind(row-1,col+1):
                if lines[row-1][col+1].isdecimal():
                    right = lookright([],col+1,row-1)
                    coll.append(right)

            #top left
            if is_valid_ind(row-1,col-1):
                    if lines[row-1][col-1].isdecimal():
                        left = lookleft([],col-1,row-1)
                        coll.append(left)
         


           if not directly_below:
                #bottom right       
                if is_valid_ind(row+1,col+1):
                    if lines[row+1][col+1].isdecimal():
                        right = lookright([],col+1,row+1)
                        coll.append(right)
                #bottom left
                if is_valid_ind(row+1,col-1):
                    if lines[row+1][col-1].isdecimal():
                        left = lookleft([],col-1,row+1)
                        coll.append(left)
                    

sum_acc    =0   
for item in coll:
    int_acc = ''
    for char in item:
        int_acc+=char
    sum_acc+=int(int_acc)

print(sum_acc)
