import turtle

RULES = {} 

def draw(turtle, RULES, length, angle):
    commandPipeLine = []
    for command in RULES:
        turtle.pd()
        if command  == "F":
            turtle.forward(length)
        elif command == "f":
            turtle.pu() 
            turtle.forward(length)
        elif command == "+":
            turtle.right(angle)
        elif command == "-":
            turtle.left(angle)
        elif command == "[":
            commandPipeLine.append((turtle.position(), turtle.heading()))
        elif command == "]":
            turtle.pu() 
            position, heading = commandPipeLine.pop()
            turtle.goto(position)
            turtle.setheading(heading)

def ruleSequence(sequence):
    if sequence in RULES:
        return RULES[sequence]
    return sequence

def getInstances(axiom, steps):
    instance = [axiom] 
    for _ in range(steps):
        next_seq = instance[-1]  # get the last element of sequence
        next_axiom = [ruleSequence(char) for char in next_seq]
        instance.append(''.join(next_axiom))
    return instance            

def setUp(alphaInit):
    r_turtle = turtle.Turtle()  
    r_turtle.screen.title("L-System")
    r_turtle.speed(0)  
    r_turtle.setheading(alphaInit)
    return r_turtle

def main():
    rule_num = 1
    while True:
        rule = input("Enter P[%d]: " % rule_num)
        if rule == '0':
            break
        key, value = rule.split("->")
        RULES[key] = value
        rule_num += 1

    axiom = input("Enter w(0): ")
    iterations = int(input("Enter number of iterations: "))

    model = getInstances(axiom, iterations)

    segmentLength = int(input("Enter segment length: "))
    alphaInit = float(input("Enter initial alpha: "))
    angle = float(input("Enter angle: "))

    r_turtle = setUp(alphaInit)  
    turtle_screen = turtle.Screen()  
    turtle_screen.screensize(1920, 1080)
    draw(r_turtle, model[-1], segmentLength, angle)
    turtle_screen.exitonclick()

if __name__ == "__main__":
    main()