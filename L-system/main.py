import turtle

SYSTEM_RULES = {} 


def derivation(axiom, steps):
    derived = [axiom] 
    for _ in range(steps):
        next_seq = derived[-1]
        next_axiom = [rule(char) for char in next_seq]
        derived.append(''.join(next_axiom))
    return derived


def rule(sequence):
    if sequence in SYSTEM_RULES:
        return SYSTEM_RULES[sequence]
    return sequence


def draw_l_system(turtle, SYSTEM_RULES, seg_length, angle):
    stack = []
    for command in SYSTEM_RULES:
        turtle.pd()
        if command in ["F", "G", "R", "L"]:
            turtle.forward(seg_length)
        elif command == "f":
            turtle.pu() 
            turtle.forward(seg_length)
        elif command == "+":
            turtle.right(angle)
        elif command == "-":
            turtle.left(angle)
        elif command == "[":
            stack.append((turtle.position(), turtle.heading()))
        elif command == "]":
            turtle.pu() 
            position, heading = stack.pop()
            turtle.goto(position)
            turtle.setheading(heading)

def set_turtle(alpha_zero):
    r_turtle = turtle.Turtle()  # recursive turtle
    r_turtle.screen.title("L-System")
    r_turtle.speed(0)  # adjust as needed (0 = fastest)
    r_turtle.setheading(alpha_zero)  # initial heading
    return r_turtle


def main():
    rule_num = 1
    while True:
        rule = input("Enter P[%d]: " % rule_num)
        if rule == '0':
            break
        key, value = rule.split("->")
        SYSTEM_RULES[key] = value
        rule_num += 1

    axiom = input("Enter w(0): ")
    iterations = int(input("Enter number of iterations: "))

    model = derivation(axiom, iterations)  # axiom (initial string), nth iterations

    segment_length = int(input("Enter step size (segment length): "))
    alpha_zero = float(input("Enter initial alpha: "))
    angle = float(input("Enter angle increment (i): "))

    # Set turtle parameters and draw L-System
    r_turtle = set_turtle(alpha_zero)  
    turtle_screen = turtle.Screen()  
    turtle_screen.screensize(1920, 1080)
    draw_l_system(r_turtle, model[-1], segment_length, angle)  # draw model
    turtle_screen.exitonclick()


if __name__ == "__main__":
    main()