
#!/usr/bin/env python

user_input=12345
_rev=0
while(user_input>0):
  dig=user_input%10
  _rev=_rev*10+dig
  user_input=user_input//10
print("The reversed number is :",_rev)
