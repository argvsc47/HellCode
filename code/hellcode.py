from sys import argv

rooms = [k for k in range(0,20)]

ptr = 0 #ends at 19
Break = 0
idx = 0

def begLoop(scr):
	global idx, Break
	while not Break:
		idx += 1
		Alloc(scr)
	Break = 0

def endLoop():
	global Break
	Break = 1

def Alloc(scr):
	global idx, ptr
	char = scr[idx]
	if char == "[": 
		if rooms[ptr] == 0: rooms[ptr] = 256
		else: rooms[ptr] -= 1
	if char == "]":
		if rooms[ptr] == 256: rooms[ptr] = 0
		else: rooms[ptr] += 1
	if char == "#": 
		if ptr == 19: ptr = 0
		else: ptr += 1
	if char == "!":
		if ptr == 0: ptr = 19
		else: ptr -= 1
	if char == "*": rooms[ptr] = 256
	if char == "@": rooms[ptr] = 0
	if char == "~": print(chr(rooms[ptr]))
	if char == "$": rooms[ptr] = ord(input())
	if char == "{": begLoop(scr)
	if char == "}": endLoop()
	if char == "%": exit()

def Parse(argc):
	global idx
	f = open(argc, "r")
	script = [l for l in f.read()]
	f.close()
	symList = [k for k in "[]*@!#~${}%"]
	while idx < len(script):
		char = script[idx]
		if not char in symList:
			pass
		else:
			Alloc(script)
		idx += 1

Parse(argv[1])
